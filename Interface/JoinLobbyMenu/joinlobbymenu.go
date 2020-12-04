package JoinLobbyMenu

import (
	"Game/Window"
	"Game/Heroes/Users"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func RemoveIndex(s []string, index int)[]string{
	return append(s[:index], s[index+1:]...)
}

func ChangeLobbyIDInputArea(winConf *Window.WindowConfig){

	if winConf.Win.Pressed(pixelgl.KeyBackspace){
		if winConf.WindowUpdation.JoinLobbyMenuFrame % 8 == 0{
			if len(winConf.TextAreas.JoinLobbyInput.WrittenText) > 0{
				winConf.TextAreas.JoinLobbyInput.WrittenText = RemoveIndex(winConf.TextAreas.JoinLobbyInput.WrittenText, len(winConf.TextAreas.JoinLobbyInput.WrittenText)-1)
			}
		}
	}
	if len(winConf.Win.Typed()) != 0 && len(winConf.TextAreas.JoinLobbyInput.WrittenText) < 10{
		winConf.TextAreas.JoinLobbyInput.WrittenText = append(winConf.TextAreas.JoinLobbyInput.WrittenText, winConf.Win.Typed())
	}
	for _, value := range winConf.TextAreas.JoinLobbyInput.WrittenText{
		fmt.Fprintf(winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea, value)
	}
	winConf.WindowUpdation.JoinLobbyMenuFrame++
}

func CheckBackButton(winConf *Window.WindowConfig, currState *Users.States){
	if winConf.WindowUpdation.JoinLobbyMenuFrame % 8 == 0 && winConf.WindowUpdation.JoinLobbyMenuFrame != 0{
		if (winConf.Win.MousePosition().X >= 21 && winConf.Win.MousePosition().X <= 68) && (winConf.Win.MousePosition().Y >= 468 && winConf.Win.MousePosition().Y <= 511) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			currState.SetStartMenu()
		}
	}
	winConf.WindowUpdation.JoinLobbyMenuFrame++
} 

func CheckJoinButton(winConf Window.WindowConfig, currState *Users.States){
	if (winConf.Win.MousePosition().X >= 363 && winConf.Win.MousePosition().X <= 596) && (winConf.Win.MousePosition().Y >= 73 && winConf.Win.MousePosition().Y <= 165) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		fmt.Println("JoinButton pressed!")
	}
}


func CreateJoinLobbyMenu(winConf *Window.WindowConfig, currState *Users.States, userConfig *Users.User){
	Window.DrawJoinLobbyMenuBG(*winConf)

	CheckBackButton(winConf, currState)

	//Draws announcement
	winConf.TextAreas.JoinLobbyAnnouncement.Clear()
	winConf.TextAreas.JoinLobbyAnnouncement.Write([]byte("Write your lobby ID!"))
	winConf.TextAreas.JoinLobbyAnnouncement.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.JoinLobbyAnnouncement.Orig, 3))

	//Draws input lobby ID
	winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Clear()
	ChangeLobbyIDInputArea(winConf)
	winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Orig, 3))

	CheckJoinButton(*winConf, currState)
}
