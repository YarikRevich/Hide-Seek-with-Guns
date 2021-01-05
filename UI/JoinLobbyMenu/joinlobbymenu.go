package JoinLobbyMenu

import (
	"fmt"
	"time"
	"strings"
	"Game/Utils"
	"Game/Server"
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type JoinLobbyMenu struct{
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf        *Window.WindowConfig

	currState      *States.States
	
	userConfig     *Users.User

	mapComponents  Map.MapConf

}

func (j *JoinLobbyMenu)Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
	j.winConf       = winConf
	j.currState     = currState
	j.userConfig    = userConfig
	j.mapComponents = mapComponents
}

func (j *JoinLobbyMenu)ProcessNetworking(){


	if j.currState.SendStates.JoinRoom{
		
		server := Server.Network(new(Server.N))
		lobbyIdToJoin := strings.Join(j.winConf.TextAreas.JoinLobbyInput.WrittenText, "")
		j.userConfig.LobbyID = lobbyIdToJoin 
		server.Init(
			fmt.Sprintf(
				"AddToLobby///%s~/%s/%d/%d/%d/%d/0|0|0|0/%s", 
				lobbyIdToJoin,
				j.userConfig.Username,
				j.userConfig.X,
				j.userConfig.Y,
				j.userConfig.UpdationRun,
				j.userConfig.CurrentFrame,
				j.userConfig.HeroPicture,
			),
			j.userConfig.Conn,
			1,
		)
		server.Write()
		response := server.Read()
		if !Utils.MessageIsEmpty(response){
			if Utils.CheckErrorResp(response){
				j.winConf.WindowError.LobbyDoesNotExist = true
				j.winConf.WindowError.LobbyErrorStop = time.Now()
				j.winConf.WindowError.LobbyErrorText = "Such lobby doesn't exist!"
			}else{
				j.winConf.WaitRoom.RoomType = "join"
				j.currState.MainStates.SetWaitRoom()
			}
			j.currState.SendStates.JoinRoom = false
		}
	}
}

func (j *JoinLobbyMenu)ProcessKeyboard(){

	if j.winConf.WindowUpdation.JoinLobbyMenuFrame % 5 == 0 && j.winConf.WindowUpdation.JoinLobbyMenuFrame != 0{
		if (j.winConf.Win.MousePosition().X >= 21 && j.winConf.Win.MousePosition().X <= 68) && (j.winConf.Win.MousePosition().Y >= 468 && j.winConf.Win.MousePosition().Y <= 511) && j.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			j.winConf.TextAreas.JoinLobbyInput.WrittenText = []string{}
			j.currState.MainStates.SetStartMenu()
		}
	}

	if (j.winConf.Win.MousePosition().X >= 363 && j.winConf.Win.MousePosition().X <= 596) && (j.winConf.Win.MousePosition().Y >= 73 && j.winConf.Win.MousePosition().Y <= 165) && j.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		j.currState.SendStates.JoinRoom = true
	}
}

func (j *JoinLobbyMenu)ProcessTextInput(){
	j.winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Clear()
	if j.winConf.Win.Pressed(pixelgl.KeyBackspace){
		if j.winConf.WindowUpdation.JoinLobbyMenuFrame % 8 == 0{
			if len(j.winConf.TextAreas.JoinLobbyInput.WrittenText) > 0{
				j.winConf.TextAreas.JoinLobbyInput.WrittenText = Utils.RemoveIndex(j.winConf.TextAreas.JoinLobbyInput.WrittenText, len(j.winConf.TextAreas.JoinLobbyInput.WrittenText)-1)
			}
		}
	}
	if len(j.winConf.Win.Typed()) != 0 && len(j.winConf.TextAreas.JoinLobbyInput.WrittenText) < 10{
		j.winConf.TextAreas.JoinLobbyInput.WrittenText = append(j.winConf.TextAreas.JoinLobbyInput.WrittenText, j.winConf.Win.Typed())
	}
	for _, value := range j.winConf.TextAreas.JoinLobbyInput.WrittenText{
		fmt.Fprintf(j.winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea, value)
	}
	j.winConf.WindowUpdation.JoinLobbyMenuFrame++
	j.winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Draw(j.winConf.Win, pixel.IM.Scaled(j.winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea.Orig, 3))

}

func (j *JoinLobbyMenu)ProcessMusic(){
	//WARNING: it is not implemented!
}

func (j *JoinLobbyMenu)DrawAnnouncements(){

	j.winConf.TextAreas.JoinLobbyAnnouncement.Clear()
	j.winConf.TextAreas.JoinLobbyAnnouncement.Write([]byte("Write your lobby ID!"))
	j.winConf.TextAreas.JoinLobbyAnnouncement.Draw(j.winConf.Win, pixel.IM.Scaled(j.winConf.TextAreas.JoinLobbyAnnouncement.Orig, 3))

}

func (j *JoinLobbyMenu)DrawElements(){

	j.winConf.DrawJoinLobbyMenuBG()
	j.winConf.DrawErrorText()
}

func (j *JoinLobbyMenu)Run(){

	j.DrawElements()

	j.ProcessKeyboard()

	j.DrawAnnouncements()

	j.ProcessTextInput()

	j.ProcessNetworking()
}
