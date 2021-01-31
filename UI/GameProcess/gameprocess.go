package GameProcess

import (
	Map "Game/Components/Map"
	"Game/Components/Sound"
	"Game/Components/States"
	"Game/Heroes/Animation"
	"Game/Heroes/Users"
	"Game/Server"
	"Game/UI/GameProcess/ConfigParsers"
	"Game/Window"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameProcess struct {
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf *Window.WindowConfig

	currState *States.States

	userConfig *Users.User

	mapComponents Map.MapConf
}

func (g *GameProcess) Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf) {
	g.winConf = winConf
	g.currState = currState
	g.userConfig = userConfig
	g.mapComponents = mapComponents
}

func GetUserFromList(u string, l []*Server.GameRequest) *Server.GameRequest {
	for _, value := range l {
		if value.PersonalInfo.Username == u {
			return value
		}
	}
	return nil
}

func (g *GameProcess) ProcessNetworking() {

	if !g.currState.NetworkingStates.GameProcess {
		g.currState.NetworkingStates.GameProcess = true
		go func() {
			parser := Server.GameParser(new(Server.GameRequest))
			server := Server.Network(new(Server.N))
			server.Init(nil, g.userConfig, 0, nil, parser.Parse, "GetUsersInfoReadyLobby")

			server.Write()
			response := server.ReadGame(parser.Unparse)
			responseUser := GetUserFromList(g.userConfig.PersonalInfo.Username, response)
			if response != nil{
				switch responseUser.Error {
				case "70":
					cp := ConfigParsers.ConfigParser(new(ConfigParsers.CP))
					cp.Init(g.winConf, g.userConfig)
					cp.ApplyConfig(responseUser)
					for _, value := range response {
						nu := cp.Unparse(value)
						cp.Commit(nu)
					}
				case "502":
					g.currState.MainStates.SetStartMenu()
				}
			}
			g.currState.NetworkingStates.GameProcess = false
		}()
	}
}

func (g *GameProcess) ProcessKeyboard() {

	currPosition := pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y))
	g.mapComponents.GetCollisions().GetDoorsCollisions().DoorTraker(currPosition)

	switch {
	case ((g.winConf.Win.Pressed(pixelgl.KeyW) ||
		g.winConf.Win.Pressed(pixelgl.KeyA) ||
		g.winConf.Win.Pressed(pixelgl.KeyS) ||
		g.winConf.Win.Pressed(pixelgl.KeyD)) && g.winConf.Win.JustPressed(pixelgl.KeySpace)) ||
		g.winConf.Win.JustPressed(pixelgl.KeySpace):
		ok, user := g.mapComponents.GetCollisions().GetHeroCollisions().IsHero(currPosition, g.winConf.GameProcess.OtherUsers).Near(30, 37)
		if ok {
			g.userConfig.Context.Additional = append(g.userConfig.Context.Additional, user, "1")
			parser := Server.GameParser(new(Server.GameRequest))
			server := Server.Network(new(Server.N))
			server.Init(nil, g.userConfig, 1, nil, parser.Parse, "UpdateUsersHealth")

			server.Write()
			server.ReadGame(parser.Unparse)

			g.userConfig.Context.Additional = []string{}
		}
		fallthrough
	default:
		switch {
		case g.winConf.Win.Pressed(pixelgl.KeyW):
			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y+2)), g.winConf.GameProcess.OtherUsers, "top")
			if coll {
				return
			}

			if g.userConfig.Pos.Y <= g.mapComponents.GetHeroBorder().Top() {
				g.userConfig.Pos.Y += 3
			}
			if g.winConf.Cam.CamPos.Y < g.mapComponents.GetCamBorder().Top()+50 {
				if g.userConfig.Pos.Y >= int(g.winConf.Win.Bounds().Center().Y) {
					g.winConf.Cam.CamPos.Y += 5
				}
			}
		case g.winConf.Win.Pressed(pixelgl.KeyA):
			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X-2), float64(g.userConfig.Pos.Y)), g.winConf.GameProcess.OtherUsers, "left")
			if coll {
				return
			}

			if g.userConfig.Pos.X >= g.mapComponents.GetHeroBorder().Left() {
				g.userConfig.Pos.X -= 3
			}
			if g.winConf.Cam.CamPos.X >= g.mapComponents.GetCamBorder().Left() {
				if g.userConfig.Pos.X <= int(g.winConf.Win.Bounds().Center().X) {
					g.winConf.Cam.CamPos.X -= 5
				}
			}
		case g.winConf.Win.Pressed(pixelgl.KeyS):
			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y-2)), g.winConf.GameProcess.OtherUsers, "bottom")
			if coll {
				return
			}

			if g.userConfig.Pos.Y >= g.mapComponents.GetHeroBorder().Bottom() {
				g.userConfig.Pos.Y -= 3
			}
			if g.winConf.Cam.CamPos.Y >= g.mapComponents.GetCamBorder().Bottom() {
				if g.userConfig.Pos.Y <= int(g.winConf.Win.Bounds().Center().Y) {
					g.winConf.Cam.CamPos.Y -= 5
				}
			}
		case g.winConf.Win.Pressed(pixelgl.KeyD):
			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X+2), float64(g.userConfig.Pos.Y)), g.winConf.GameProcess.OtherUsers, "right")
			if coll {
				return
			}

			if g.userConfig.Pos.X <= g.mapComponents.GetHeroBorder().Right() {
				g.userConfig.Pos.X += 3
			}
			if g.winConf.Cam.CamPos.X <= g.mapComponents.GetCamBorder().Right() {
				if g.userConfig.Pos.X >= int(g.winConf.Win.Bounds().Center().X) {
					g.winConf.Cam.CamPos.X += 5
				}
			}
		}
	}

}

func (g *GameProcess) ProcessTextInput() {
	//WARNING: it is not implemented!
}

func (g *GameProcess) ProcessMusic() {

	sound := Sound.Sound(new(Sound.S))
	sound.Init(g.currState)
	//sound.Play()
}

func (g *GameProcess) DrawAnnouncements() {
	//WARNING: it is not implemented!
}

func (g *GameProcess) DrawElements() {
	g.winConf.DrawGameBackground()

	//g.winConf.DrawGoldChest()

	g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	

	Animation.NewDefaultSwordAnimator(g.winConf, g.userConfig).Move()
	Animation.NewIconAnimator(g.winConf, g.userConfig).Move()

	for _, value := range g.winConf.GameProcess.OtherUsers {
		Animation.NewDefaultSwordAnimator(g.winConf, value).Move()
		Animation.NewIconAnimator(g.winConf, value).Move()
	}


	g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.Pos.X)*2.5)-31, (float64(g.userConfig.Pos.Y)*2.5)-30))

	g.winConf.DrawElementsPanel()


	g.mapComponents.GetCam().UpdateCam()

	var bias float64
	for i := 0; i <= g.userConfig.GameInfo.Health; i++ {
		g.winConf.DrawHPHeart(
			pixel.V(-40+bias, 1200),
		)
		bias += 100
	}

	if g.userConfig.GameInfo.Health < 1 {
		g.mapComponents.GetCam().SetDefaultCam()
		g.currState.MainStates.SetStartMenu()
	}
}

func (g *GameProcess) Run() {

	g.ProcessKeyboard()

	g.ProcessNetworking()

	g.DrawElements()

	g.ProcessMusic()
}
