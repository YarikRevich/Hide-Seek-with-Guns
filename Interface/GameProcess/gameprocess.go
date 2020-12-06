package GameProcess

import (
	"fmt"
	_ "time"
	"Game/Window"
	//"Game/Interface/GameProcess/ConfigParsers"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"github.com/faiface/pixel/pixelgl"
)

func KeyBoardButtonListener(userConfig *Users.User, win *pixelgl.Window){
	if (win.MousePosition().X > 18 && win.MousePosition().X < 141) && (win.MousePosition().Y > 74 && win.MousePosition().Y < 102) && win.Pressed(pixelgl.MouseButtonLeft){}
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
		Animation.MoveAndChangeAnim(value, winConf.Win, winConf.Components.AvailableHeroImages)
	}
}

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig){
	KeyBoardButtonListener(userConfig, winConf.Win)
	Animation.MoveAndChangeAnim(userConfig, winConf.Win, winConf.Components.AvailableHeroImages)
}

func ListenToUsersInfo(userConfig *Users.User){
	
	buff := make([]byte, 144)
	userConfig.Conn.Read(buff)
	userConfig.Game.ReadWriteUpdate <- string(buff)
	
}

func CreateGame(userConfig *Users.User, winConf *Window.WindowConfig){

	formattedReq := fmt.Sprintf("GetUsersInfo///%s~", userConfig.LobbyID)
	userConfig.Conn.Write([]byte(formattedReq))

	go ListenToUsersInfo(userConfig)
	select{
	case response := <- userConfig.Game.ReadWriteUpdate:
		fmt.Println(response)
	}
	// select{
	// case response := <- userConfig.Game.ReadWriteUpdate:
	// 	Window.UpdateBackground(winConf)
	// 	ConfigParsers.UnparseCurrent(response, userConfig)
	// 	ChangePos(userConfig, winConf)
	// 	ConfigParsers.UnparseOthers(response, *userConfig, &otherUsers)
	// 	ReDraw(&otherUsers, winConf)
	// 	parsedMessage := ConfigParsers.ParseConfig(userConfig, otherUsers, response)
	// 	userConfig.Conn.Write([]byte(parsedMessage))
	// }
}