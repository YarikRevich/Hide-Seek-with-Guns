package LobbyWaitRoom

import (
	"strings"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Utils"
	"fmt"
)

func ClientIsWritten(client string, winConf *Window.WindowConfig)bool{
	for _, value := range winConf.WaitRoom.NewMembers{
		if strings.Contains(value, client){
			return true
		}
	}	
	return false
}

func GetUpdates(userConfig *Users.User, winConf *Window.WindowConfig, currState *Users.States){
		requestText := fmt.Sprintf("GetMembersInLobby///%s", userConfig.LobbyID)
		userConfig.Conn.Write([]byte(requestText))
		buff := make([]byte, 144)
		//userConfig.Conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		userConfig.Conn.Read(buff)
		if !Utils.MessageIsEmpty(buff){
			if Utils.CheckErrorResp(string(buff)){
				currState.SetGame()
				return
			}
			cleanedResp := Utils.CleanGottenResponse(string(buff))
			splitedBuff := strings.Split(cleanedResp, "//")
			for _, value := range splitedBuff{
				if !ClientIsWritten(value, winConf) && len(winConf.WaitRoom.NewMembers) <= 3{
					winConf.WaitRoom.NewMembers = append(winConf.WaitRoom.NewMembers, value)
				}
			}
		}
}

func CreateLobbyWaitRoom(winConf *Window.WindowConfig, currState *Users.States, userConfig *Users.User){
	
	GetUpdates(userConfig, winConf, currState)

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

	//Listens to actions

	ListenForChanges(winConf, userConfig, currState)
}

func ListenForChanges(winConf *Window.WindowConfig, userConfig *Users.User, currState *Users.States){
	if winConf.WindowUpdation.WaitRoomFrame % 8 == 0 && winConf.WindowUpdation.WaitRoomFrame != 0{
		switch winConf.WaitRoom.RoomType {
		case "create":
			if (winConf.Win.MousePosition().X >= 361 && winConf.Win.MousePosition().X <= 596) && (winConf.Win.MousePosition().Y >= 73 && winConf.Win.MousePosition().Y <= 165) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
				formattedReq := fmt.Sprintf("ClosePreparingLobby///%s~", userConfig.LobbyID)
				userConfig.Conn.Write([]byte(formattedReq))
				currState.SetGame()
			}
		case "join":
			if (winConf.Win.MousePosition().X >= 21 && winConf.Win.MousePosition().X <= 68) && (winConf.Win.MousePosition().Y >= 463 && winConf.Win.MousePosition().Y <= 507) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
				currState.SetCreateLobbyMenu()
			}	
		}
	}
	winConf.WindowUpdation.WaitRoomFrame++	
}