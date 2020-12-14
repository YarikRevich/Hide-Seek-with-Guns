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

func KeyBoardButtonListener(userConfig *Users.User, win *pixelgl.Window){
	if win.Pressed(pixelgl.KeyW){
		userConfig.Y += 7
	}else if win.Pressed(pixelgl.KeyA){
		userConfig.X -= 7
	}else if win.Pressed(pixelgl.KeyS){
		userConfig.Y -= 7
	}else if win.Pressed(pixelgl.KeyD){
		userConfig.X += 7	
	} 
}

func ReDraw(otherUsers *[]*Users.User, winConf *Window.WindowConfig){
	for _, value := range *otherUsers{
		Animation.MoveAndChangeAnim(value, winConf)
	}
}

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig){
	KeyBoardButtonListener(userConfig, winConf.Win)
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
}