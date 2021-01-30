package LobbyWaitRoom

import (
	"Game/Components/Map"
	"Game/Components/States"
	"Game/Heroes/Users"
	"Game/Server"	
	"Game/Window"
	"fmt"
	

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

func GetUserFromList(u string, l []*Server.GameRequest)*Server.GameRequest{
	for _, value := range l{
		if value.PersonalInfo.Username == u{
			return value
		}
	}
	return nil
}

func (l *LobbyWaitRoom)ProcessNetworking(){
	//It is a func for making network for getting
	//the newest information about users in lobby.
	//To get it, not envolving on the fps it's run
	//in the goroutine

	if !l.currState.NetworkingStates.LobbyWaitRoom{
		l.currState.NetworkingStates.LobbyWaitRoom = true
		go func(){
			parser := Server.GameParser(new(Server.GameRequest))
			server := Server.Network(new(Server.N))
			server.Init(nil, l.userConfig, 1, nil, parser.Parse, "GetUsersInfoPrepLobby")
			server.Write()
			response := server.ReadGame(parser.Unparse)
			responseUser :=  GetUserFromList(l.userConfig.PersonalInfo.Username, response)
			
			if responseUser != nil{
				switch responseUser.Error{
				case "60":
					iswritten := func(u string)bool{
						for _, us := range l.winConf.WaitRoom.NewMembers{
							if u == us{
								return true
							}
						}
						return false
					}
					for _, user := range response{
						if len(l.winConf.WaitRoom.NewMembers) <= 4 && !iswritten(user.PersonalInfo.Username){
							l.winConf.WaitRoom.NewMembers = append(l.winConf.WaitRoom.NewMembers, user.PersonalInfo.Username)
						}
					}
				case "502":
					l.currState.MainStates.SetGame()
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

	server := Server.Network(new(Server.N))
	if (l.winConf.Win.MousePosition().X >= 361 && l.winConf.Win.MousePosition().X <= 596) && (l.winConf.Win.MousePosition().Y >= 73 && l.winConf.Win.MousePosition().Y <= 165) && l.winConf.Win.JustPressed(pixelgl.MouseButtonLeft){
		if l.winConf.WaitRoom.RoomType == "create"{
			parser := Server.GameParser(new(Server.GameRequest))
			server.Init(nil, l.userConfig, 1, nil, parser.Parse, "ClosePreparingLobby")
			server.Write()
			server.ReadGame(parser.Unparse)
			l.currState.MainStates.SetGame()
		}
	}
	if (l.winConf.Win.MousePosition().X >= 21 && l.winConf.Win.MousePosition().X <= 68) && (l.winConf.Win.MousePosition().Y >= 463 && l.winConf.Win.MousePosition().Y <= 507) && l.winConf.Win.JustPressed(pixelgl.MouseButtonLeft){
		if l.winConf.WaitRoom.RoomType == "create"{
			l.winConf.WaitRoom.NewMembers = []string{}
			l.currState.MainStates.SetCreateLobbyMenu()
				 
			parser := Server.GameParser(new(Server.GameRequest))
			server.Init(nil, l.userConfig, 1, nil, parser.Parse, "DeleteLobby")
			server.Write()
	 		server.ReadGame(parser.Unparse)
		}else{
			l.winConf.WaitRoom.NewMembers = []string{}
	 		l.currState.MainStates.SetJoinLobbyMenu()
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
	lobbyIdText := fmt.Sprintf("Lobby ID is: %s", l.userConfig.PersonalInfo.LobbyID)
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