package LobbyWaitRoom

import (
	"fmt"
	"strings"
	"Game/Utils"
	"Game/Server"
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"Game/UI/GameProcess/ConfigParsers"
)

type LobbyWaitRoom struct{
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf        *Window.WindowConfig

	currState      *States.States
	
	userConfig     *Users.User

	mapComponents  Map.MapConf

}

func (l *LobbyWaitRoom)Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
	l.winConf       = winConf
	l.currState     = currState
	l.userConfig    = userConfig
	l.mapComponents = mapComponents
}

func (l *LobbyWaitRoom)ProcessNetworking(){
	//It is a func for making network for getting
	//the newest information about users in lobby.
	//To get it, not envolving on the fps it's run
	//in the goroutine

	if !l.currState.NetworkingStates.LobbyWaitRoom{
		l.currState.NetworkingStates.LobbyWaitRoom = true
		go func(){server := Server.Network(new(Server.N))
			server.Init(fmt.Sprintf("GetUsersInfoLobby///%s", l.userConfig.LobbyID), l.userConfig.Conn, 1)
			server.Write()
			response := server.Read()

			ready := Utils.CheckLobbyIsReady(response)
			if ready{
				l.currState.MainStates.SetGame()
			}else{
				if !Utils.CheckErrorResp(response){
					cleanedResp := Utils.CleanGottenResponse(response)
					if !Utils.IsOkResp(cleanedResp){
						unparsedUsers := ConfigParsers.UnparseUsers(cleanedResp)
						for _, value := range unparsedUsers{
							if len(l.winConf.WaitRoom.NewMembers) <= 3{
								if len(l.winConf.WaitRoom.NewMembers) != 0{
									if !strings.Contains(strings.Join(l.winConf.WaitRoom.NewMembers, " "), value){
										l.winConf.WaitRoom.NewMembers = append(l.winConf.WaitRoom.NewMembers, value)
									}
								
								}else{
								 	l.winConf.WaitRoom.NewMembers = append(l.winConf.WaitRoom.NewMembers, value)
								}
							}
						}
					}
				}
			}
			l.currState.NetworkingStates.LobbyWaitRoom = false
		}()
	}
}

func (l *LobbyWaitRoom)ProcessKeyboard(){

	if l.winConf.WaitRoom.RoomType == "create" && (l.winConf.Win.MousePosition().X >= 361 && l.winConf.Win.MousePosition().X <= 596) && (l.winConf.Win.MousePosition().Y >= 73 && l.winConf.Win.MousePosition().Y <= 165){
		l.winConf.DrawWaitRoomPressedButton()
	}

	if l.winConf.WindowUpdation.WaitRoomFrame % 6 == 0 && l.winConf.WindowUpdation.WaitRoomFrame != 0{

		server := Server.Network(new(Server.N))
		if (l.winConf.Win.MousePosition().X >= 361 && l.winConf.Win.MousePosition().X <= 596) && (l.winConf.Win.MousePosition().Y >= 73 && l.winConf.Win.MousePosition().Y <= 165) && l.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			if l.winConf.WaitRoom.RoomType == "create"{
				server.Init(fmt.Sprintf("ClosePreparingLobby///%s~", l.userConfig.LobbyID), l.userConfig.Conn, 1)
				server.Write()
				server.Read()
				l.currState.MainStates.SetGame()
			}
		}

		if (l.winConf.Win.MousePosition().X >= 21 && l.winConf.Win.MousePosition().X <= 68) && (l.winConf.Win.MousePosition().Y >= 463 && l.winConf.Win.MousePosition().Y <= 507) && l.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			if l.winConf.WaitRoom.RoomType == "create"{
				l.winConf.WaitRoom.NewMembers = []string{}
		 		l.currState.MainStates.SetCreateLobbyMenu()
		 		server.Init(fmt.Sprintf("DeleteLobby///%s", l.userConfig.LobbyID), l.userConfig.Conn, 1)
		 		server.Write()
		 		server.Read()
			}else{
				l.winConf.WaitRoom.NewMembers = []string{}
		 		l.currState.MainStates.SetJoinLobbyMenu()
			}
		}
	}
	l.winConf.WindowUpdation.WaitRoomFrame++
}

func (l *LobbyWaitRoom)ProcessTextInput(){
	//WARNING: it is not implemented!
}

func (l *LobbyWaitRoom)ProcessMusic(){
	//WARNING: it is not implemented!
}

func (l *LobbyWaitRoom)DrawAnnouncements(){

	l.winConf.TextAreas.NewMembersAnnouncement.Clear()
	l.winConf.TextAreas.NewMembersAnnouncement.Write([]byte("Wait for members!"))
	l.winConf.TextAreas.NewMembersAnnouncement.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.NewMembersAnnouncement.Orig, 4))

	l.winConf.TextAreas.NewMembersTextArea.Clear()
	for _, value := range l.winConf.WaitRoom.NewMembers{
		l.winConf.TextAreas.NewMembersTextArea.Write([]byte(value + "\n"))
	}
	l.winConf.TextAreas.NewMembersTextArea.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.NewMembersTextArea.Orig, 2.5)) 

	l.winConf.TextAreas.CurrentLobbyIDArea.Clear()
	lobbyIdText := fmt.Sprintf("Lobby ID is: %s", l.userConfig.LobbyID)
	l.winConf.TextAreas.CurrentLobbyIDArea.Write([]byte(lobbyIdText))
	l.winConf.TextAreas.CurrentLobbyIDArea.Draw(l.winConf.Win, pixel.IM.Scaled(l.winConf.TextAreas.CurrentLobbyIDArea.Orig, 2.5))

}

func (l *LobbyWaitRoom)DrawElements(){

	switch l.winConf.WaitRoom.RoomType{
	case "create":
		l.winConf.DrawWaitRoomMenuBG()
	case "join":
		l.winConf.DrawWaitRoomJoinBG()
	}

}

func (l *LobbyWaitRoom)Run(){

	l.DrawElements()

	l.ProcessKeyboard()

	l.DrawAnnouncements()

	l.ProcessNetworking()
}