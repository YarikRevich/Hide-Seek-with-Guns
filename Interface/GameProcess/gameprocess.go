package GameProcess

import (
	"github.com/faiface/pixel"
	"strings"
	"Game/Utils"
	"fmt"
	_ "time"
	"Game/Window"
	"Game/Interface/GameProcess/ConfigParsers"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"github.com/faiface/pixel/pixelgl"
	"Game/Interface/GameProcess/Map"
)

func KeyBoardButtonListener(userConfig *Users.User, winConf *Window.WindowConfig, camBorder Map.CamBorder){

	heroBorder := Map.HeroBorder(&Map.HB{})

	collisions := Map.Collisions(&Map.C{})
	collisions.Init()


	switch {
	case winConf.Win.Pressed(pixelgl.KeyW):
		if collisions.IsCollision(pixel.V(float64(userConfig.X), float64(userConfig.Y+2))){
			return
		}
		if userConfig.Y <= heroBorder.Top(){
			userConfig.Y += 3
		}
		if winConf.Cam.CamPos.Y < camBorder.Top(){
			if userConfig.Y >= int(winConf.Win.Bounds().Center().Y){
				winConf.Cam.CamPos.Y += 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyA):
		if collisions.IsCollision(pixel.V(float64(userConfig.X-2), float64(userConfig.Y))){
			return
		}
		if userConfig.X >= heroBorder.Left(){
			userConfig.X -= 3
		}
		if winConf.Cam.CamPos.X >= camBorder.Left(){
			if userConfig.X <= int(winConf.Win.Bounds().Center().X){
				winConf.Cam.CamPos.X -= 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyS):
		if collisions.IsCollision(pixel.V(float64(userConfig.X), float64(userConfig.Y-2))){
			return
		}
		if userConfig.Y >= heroBorder.Bottom(){
			userConfig.Y -= 3
		}
		if winConf.Cam.CamPos.Y >= camBorder.Bottom(){
			if userConfig.Y <= int(winConf.Win.Bounds().Center().Y){	
				winConf.Cam.CamPos.Y -= 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyD):
		if collisions.IsCollision(pixel.V(float64(userConfig.X+2), float64(userConfig.Y))){
			return
		}
		if userConfig.X <= heroBorder.Right(){
			userConfig.X += 3
		}
		if winConf.Cam.CamPos.X <= camBorder.Right(){
			if userConfig.X >= int(winConf.Win.Bounds().Center().X){	
				winConf.Cam.CamPos.X += 5
			}
		}
	} 
}

func ReDraw(otherUsers *[]*Users.User, winConf *Window.WindowConfig){
	for _, value := range *otherUsers{
		Animation.MoveAndChangeAnim(value, winConf)
	}
}

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig, camBorder Map.CamBorder){
	KeyBoardButtonListener(userConfig, winConf, camBorder)
	Animation.MoveAndChangeAnim(userConfig, winConf)
}

func ListenToUsersInfo(userConfig *Users.User)string{
	
	buff := make([]byte, 4096)
	userConfig.Conn.Read(buff)
	return string(buff)
}

func CreateGame(userConfig *Users.User, winConf *Window.WindowConfig, camBorder Map.CamBorder){

	formattedReq := fmt.Sprintf("GetUsersInfo///%s~", userConfig.LobbyID)
	userConfig.Conn.Write([]byte(formattedReq))
	response := ListenToUsersInfo(userConfig)
	winConf.DrawGameBackground()


	//Draws main hero
	ChangePos(userConfig, winConf, camBorder)
	parsedMessage := ConfigParsers.ParseConfig(userConfig)
	userConfig.Conn.Write([]byte(parsedMessage))

	if ConfigParsers.IsUsersInfo(response){
		if cleaned := Utils.CleanGottenResponse(strings.Split(response, "~/")[1]); len(cleaned) != 0{
			
			//Draws other heroes
			otherUsers := []*Users.User{}
			ConfigParsers.UnparseOthers(cleaned, *userConfig, &otherUsers)
			ReDraw(&otherUsers, winConf)
		}
	}
	winConf.UpdateCam()
}