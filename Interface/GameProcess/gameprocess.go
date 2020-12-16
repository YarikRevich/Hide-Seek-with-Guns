package GameProcess

import (
	"strings"
	"Game/Utils"
	"fmt"
	_ "time"
	"Game/Window"
	"Game/Interface/GameProcess/ConfigParsers"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"github.com/faiface/pixel/pixelgl"
)

func KeyBoardButtonListener(userConfig *Users.User, winConf *Window.WindowConfig){

	if winConf.Win.Pressed(pixelgl.KeyW){
		if userConfig.Y <= int(winConf.BGImages.Game.Bounds().Max.Y-100){
			userConfig.Y += 5
		}
		if (winConf.Cam.CamPos.Y*2) < winConf.BGImages.Game.Bounds().Max.Y{
			if userConfig.Y >= int(winConf.Win.Bounds().Center().Y){
				winConf.Cam.CamPos.Y += 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyA){
		if userConfig.X >= -132{
			userConfig.X -= 5
		}
		if winConf.Cam.CamPos.X >= (winConf.BGImages.Game.Bounds().Center().X/2){
			if userConfig.X <= int(winConf.Win.Bounds().Center().X){
				winConf.Cam.CamPos.X -= 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyS){
		if userConfig.Y >= -61{
			userConfig.Y -= 5
		}
		if winConf.Cam.CamPos.Y >= (winConf.BGImages.Game.Bounds().Center().Y/2){
			if userConfig.Y <= int(winConf.Win.Bounds().Center().Y){	
				winConf.Cam.CamPos.Y -= 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyD){
		if userConfig.X <= int(winConf.BGImages.Game.Bounds().Max.X-220){
			userConfig.X += 5
		}
		if (winConf.Cam.CamPos.X*2) != winConf.BGImages.Game.Bounds().Max.X{
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

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig){
	KeyBoardButtonListener(userConfig, winConf)
	Animation.MoveAndChangeAnim(userConfig, winConf)
}

func ListenToUsersInfo(userConfig *Users.User)string{
	
	buff := make([]byte, 4096)
	userConfig.Conn.Read(buff)
	return string(buff)
}

func CreateGame(userConfig *Users.User, winConf *Window.WindowConfig){

	formattedReq := fmt.Sprintf("GetUsersInfo///%s~", userConfig.LobbyID)
	userConfig.Conn.Write([]byte(formattedReq))
	response := ListenToUsersInfo(userConfig)
	Window.DrawGameBackground(*winConf)

	//Draws main hero
	ChangePos(userConfig, winConf)
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
	Window.UpdateCam(winConf)
}