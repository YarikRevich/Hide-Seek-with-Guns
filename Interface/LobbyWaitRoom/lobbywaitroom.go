package LobbyWaitRoom

import (
	"fmt"
	"strings"
	"Game/Utils"
	"Game/Server"
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/States"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"Game/Interface/GameProcess/ConfigParsers"
)

func ClientIsWritten(client string, winConf *Window.WindowConfig)bool{
	for _, value := range winConf.WaitRoom.NewMembers{
		if strings.Contains(value, client){
			return true
		}
	}	
	return false
}

func ListenForChanges(winConf *Window.WindowConfig, userConfig *Users.User, currState *States.States){
	if winConf.WindowUpdation.WaitRoomFrame % 8 == 0 && winConf.WindowUpdation.WaitRoomFrame != 0{
		server := Server.Network(new(Server.N))
		switch winConf.WaitRoom.RoomType {
		case "create":
			if (winConf.Win.MousePosition().X >= 361 && winConf.Win.MousePosition().X <= 596) && (winConf.Win.MousePosition().Y >= 73 && winConf.Win.MousePosition().Y <= 165) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
				server.Init(fmt.Sprintf("ClosePreparingLobby///%s~", userConfig.LobbyID), userConfig.Conn)
				server.Write()
				currState.SetGame()
			}
			if (winConf.Win.MousePosition().X >= 21 && winConf.Win.MousePosition().X <= 68) && (winConf.Win.MousePosition().Y >= 463 && winConf.Win.MousePosition().Y <= 507) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
				currState.SetCreateLobbyMenu()
				server.Init(fmt.Sprintf("DeleteLobby///%s", userConfig.LobbyID), userConfig.Conn)
				server.Write()
			}
		case "join":
			if (winConf.Win.MousePosition().X >= 21 && winConf.Win.MousePosition().X <= 68) && (winConf.Win.MousePosition().Y >= 463 && winConf.Win.MousePosition().Y <= 507) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
				currState.SetCreateLobbyMenu()
			}	
		}
	}
	winConf.WindowUpdation.WaitRoomFrame++	
}

func CreateLobbyWaitRoom(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User){

	switch winConf.WaitRoom.RoomType{
	case "create":
		winConf.DrawWaitRoomMenuBG()
	case "join":
		winConf.DrawWaitRoomJoinBG()
	}

	//Writes announcement for the waiting room
	winConf.TextAreas.NewMembersAnnouncement.Clear()
	winConf.TextAreas.NewMembersAnnouncement.Write([]byte("Wait for members!"))
	winConf.TextAreas.NewMembersAnnouncement.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.NewMembersAnnouncement.Orig, 4))

	winConf.TextAreas.NewMembersTextArea.Clear()
	for _, value := range winConf.WaitRoom.NewMembers{
		winConf.TextAreas.NewMembersTextArea.Write([]byte(value + "\n"))
	}
	winConf.TextAreas.NewMembersTextArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.NewMembersTextArea.Orig, 2.5)) 

	winConf.TextAreas.CurrentLobbyIDArea.Clear()
	lobbyIdText := fmt.Sprintf("Lobby ID is: %s", userConfig.LobbyID)
	winConf.TextAreas.CurrentLobbyIDArea.Write([]byte(lobbyIdText))
	winConf.TextAreas.CurrentLobbyIDArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.CurrentLobbyIDArea.Orig, 2.5))

	ListenForChanges(winConf, userConfig, currState)

	server := Server.Network(new(Server.N))
	server.Init(fmt.Sprintf("GetUsersInfoLobby///%s", userConfig.LobbyID), userConfig.Conn)
	server.Write()
	response := server.Read()

	if !Utils.MessageIsEmpty(response){
		if Utils.CheckErrorResp(string(response)){
			currState.SetGame()
		}else{
			cleanedResp := Utils.CleanGottenResponse(string(response))
			if !Utils.IsOkResp(cleanedResp){
				unpursedUsers := ConfigParsers.UnparseUsers(cleanedResp)
				for _, value := range unpursedUsers{
					if !ClientIsWritten(value, winConf) && len(winConf.WaitRoom.NewMembers) <= 3{
						winConf.WaitRoom.NewMembers = append(winConf.WaitRoom.NewMembers, value)
					}
				}
			}
		}
	}
}