package GameProcess

import (
	"fmt"
	"strings"
	"Game/Utils"
	"Game/Window"
	"Game/Server"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"Game/Components/Sound"
	"Game/Components/States"
	Map "Game/Components/Map"
	"Game/UI/GameProcess/ConfigParsers"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameProcess struct{
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf        *Window.WindowConfig

	currState      *States.States
	
	userConfig     *Users.User

	mapComponents  Map.MapConf

}

func (g *GameProcess)Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
	g.winConf       = winConf
	g.currState     = currState
	g.userConfig    = userConfig
	g.mapComponents = mapComponents
}

func (g *GameProcess)ProcessNetworking(){

	if !g.currState.NetworkingStates.GameProcess{

		g.currState.NetworkingStates.GameProcess = true
		go func(){server := Server.Network(new(Server.N))
			server.Init(fmt.Sprintf("GetUsersInfo///%s~", g.userConfig.LobbyID), g.userConfig.Conn, 0)
			server.Write()
			response := server.Read()

			server.Init(ConfigParsers.ParseConfig(g.userConfig), g.userConfig.Conn, 0)
			server.Write()

			if ConfigParsers.IsUsersInfo(string(response)){
				if cleaned := Utils.CleanGottenResponse([]byte(strings.Split(string(response), "~/")[1])); len(cleaned) != 0{
					g.winConf.GameProcess.OtherUsers = []*Users.User{}
					ConfigParsers.UnparseOthers(cleaned, *g.userConfig, g.winConf)
				}
			}
			g.currState.NetworkingStates.GameProcess = false
		}()
	}
}

func (g *GameProcess)ProcessKeyboard(){

	currPosition := pixel.V(float64(g.userConfig.X), float64(g.userConfig.Y))
	g.mapComponents.GetCollisions().GetDoorsCollisions().DoorTraker(currPosition)
	
	switch {
	case g.winConf.Win.Pressed(pixelgl.KeyW):
		coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.X), float64(g.userConfig.Y+2)), g.winConf.GameProcess.OtherUsers, "top")
		if coll{
			return
		}
		
		if g.userConfig.Y <= g.mapComponents.GetHeroBorder().Top(){
			g.userConfig.Y += 3
		}
		if g.winConf.Cam.CamPos.Y < g.mapComponents.GetCamBorder().Top(){
			if g.userConfig.Y >= int(g.winConf.Win.Bounds().Center().Y){
				g.winConf.Cam.CamPos.Y += 5
			}
		}
	case g.winConf.Win.Pressed(pixelgl.KeyA):
		coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.X-2), float64(g.userConfig.Y)), g.winConf.GameProcess.OtherUsers, "left")
		if coll{
			return
		}
		
		if g.userConfig.X >= g.mapComponents.GetHeroBorder().Left(){
			g.userConfig.X -= 3
		}
		if g.winConf.Cam.CamPos.X >= g.mapComponents.GetCamBorder().Left(){
			if g.userConfig.X <= int(g.winConf.Win.Bounds().Center().X){
				g.winConf.Cam.CamPos.X -= 5
			}
		}
	case g.winConf.Win.Pressed(pixelgl.KeyS):
		coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.X), float64(g.userConfig.Y-2)), g.winConf.GameProcess.OtherUsers, "bottom")
		if coll{
			return
		}

		if g.userConfig.Y >= g.mapComponents.GetHeroBorder().Bottom(){
			g.userConfig.Y -= 3
		}
		if g.winConf.Cam.CamPos.Y >= g.mapComponents.GetCamBorder().Bottom(){
			if g.userConfig.Y <= int(g.winConf.Win.Bounds().Center().Y){	
				g.winConf.Cam.CamPos.Y -= 5
			}
		}
	case g.winConf.Win.Pressed(pixelgl.KeyD):
		coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.X+2), float64(g.userConfig.Y)), g.winConf.GameProcess.OtherUsers, "right")
		if coll{
			return
		}

		if g.userConfig.X <= g.mapComponents.GetHeroBorder().Right(){
			g.userConfig.X += 3
		}
		if g.winConf.Cam.CamPos.X <= g.mapComponents.GetCamBorder().Right(){
			if g.userConfig.X >= int(g.winConf.Win.Bounds().Center().X){	
				g.winConf.Cam.CamPos.X += 5
			}
		}
	} 
}

func (g *GameProcess)ProcessTextInput(){
	//WARNING: it is not implemented!
}

func (g *GameProcess)ProcessMusic(){

	sound := Sound.Sound(new(Sound.S))
	sound.Init(g.currState)
	//sound.Play()
}

func (g *GameProcess)DrawAnnouncements(){
	//WARNING: it is not implemented!
}

func (g *GameProcess)DrawElements(){
	g.winConf.DrawGameBackground()

	//g.winConf.DrawGoldChest()

	g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	Animation.MoveAndChangeAnim(g.userConfig, g.winConf)

	for _, value := range g.winConf.GameProcess.OtherUsers{
		Animation.MoveAndChangeAnim(value, g.winConf)
	}

	g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.X)*2.5)-31, (float64(g.userConfig.Y)*2.5)-30))

	g.mapComponents.GetCam().UpdateCam()
}

func (g *GameProcess)Run(){

	g.ProcessNetworking()

	g.DrawElements()

	g.ProcessKeyboard()

	g.ProcessMusic()
}