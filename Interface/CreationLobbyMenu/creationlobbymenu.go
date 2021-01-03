package CreationLobbyMenu

import (
	"github.com/faiface/pixel/pixelgl"
	"Game/Window"
	"fmt"
	"Game/Server"
	"Game/Components/States"
	"github.com/faiface/pixel"
	"Game/Heroes/Users"
	"Game/Utils"
	"strings"
)

func RemoveIndex(s []string, index int)[]string{
	return append(s[:index], s[index+1:]...)
}

func ChangeLobbyIDInputArea(winConf *Window.WindowConfig){
	if winConf.Win.Pressed(pixelgl.KeyBackspace){
		if winConf.WindowUpdation.CreationMenuFrame % 8 == 0{
			if len(winConf.TextAreas.CreateLobbyInput.WrittenText) > 0{
				winConf.TextAreas.CreateLobbyInput.WrittenText = RemoveIndex(winConf.TextAreas.CreateLobbyInput.WrittenText, len(winConf.TextAreas.CreateLobbyInput.WrittenText)-1)
			}
		}
	}
	if len(winConf.Win.Typed()) != 0 && len(winConf.TextAreas.CreateLobbyInput.WrittenText) < 10{
		winConf.TextAreas.CreateLobbyInput.WrittenText = append(winConf.TextAreas.CreateLobbyInput.WrittenText, winConf.Win.Typed())
	}
	for _, value := range winConf.TextAreas.CreateLobbyInput.WrittenText{
		fmt.Fprintf(winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea, value)
	}
	winConf.WindowUpdation.CreationMenuFrame++
}

func CheckBackButton(winConf Window.WindowConfig, currState *States.States){
	if winConf.WindowUpdation.CreationMenuFrame % 8 == 0 && winConf.WindowUpdation.CreationMenuFrame != 0{
		if (winConf.Win.MousePosition().X >= 21 && winConf.Win.MousePosition().X <= 68) && (winConf.Win.MousePosition().Y >= 468 && winConf.Win.MousePosition().Y <= 511) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			winConf.TextAreas.CreateLobbyInput.WrittenText = []string{}
			currState.SetStartMenu()
		}
	}
} 

func CheckCreateButton(winConf Window.WindowConfig, currState *States.States, userConfig *Users.User){
	if (winConf.Win.MousePosition().X >= 342 && winConf.Win.MousePosition().X <= 612) && (winConf.Win.MousePosition().Y >= 75 && winConf.Win.MousePosition().Y <= 172) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		winConf.Senders.CreateRoom = true
	}
}

func CreateLobbyMakingMenu(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User){

	//Draws creation menu BG
	winConf.DrawCreationLobbyMenuBG()

	//Checks whether back button is pressed

	CheckBackButton(*winConf, currState)


	//Clears last text area for the announsing of the importance to write lobby ID
	winConf.TextAreas.WriteIDTextArea.Clear()
	fmt.Fprintf(winConf.TextAreas.WriteIDTextArea, "Write your lobby ID!")
	winConf.TextAreas.WriteIDTextArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.WriteIDTextArea.Orig, 4))

	//Puts pressed buttons to the input area
	winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Clear()
	ChangeLobbyIDInputArea(winConf)
	winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Orig, 3))

	//Checks whether create lobby button is pressed

	CheckCreateButton(*winConf, currState, userConfig)

	if winConf.Senders.CreateRoom{
		server := Server.Network(new(Server.N))

		writtenID := strings.Join(winConf.TextAreas.CreateLobbyInput.WrittenText, "")
		userConfig.LobbyID = writtenID
		server.Init(fmt.Sprintf("CreateLobby///%s", writtenID), userConfig.Conn)
		server.Write()
		server.Read()

		server.Init(
			fmt.Sprintf(
				"AddToLobby///%s~/%s/%d/%d/%d/%d/0|0|0|0/%s", 
				writtenID,
				userConfig.Username,
				userConfig.X,
				userConfig.Y,
				userConfig.UpdationRun,
				userConfig.CurrentFrame,
				userConfig.HeroPicture,
			), 
			userConfig.Conn,
		)
		server.Write()
		response := server.Read()

		if !Utils.MessageIsEmpty(response){
			winConf.WaitRoom.RoomType = "create"
			currState.SetWaitRoom()
			winConf.Senders.CreateRoom = false
			winConf.TextAreas.CreateLobbyInput.WrittenText = []string{}
		}
	}
}