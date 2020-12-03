package LobbyWaitRoom

import (
	"strings"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"Game/Window"
	"Game/Heroes/Users"
	"fmt"
	"bytes"
)

func MessageIsEmpty(message []byte)bool{
	emptyMessage := make([]byte, 144)
	if bytes.Compare(message, emptyMessage) == 0{
		return true
	}
	return false
}

func ClientIsWritten(client string, waitRoom *Window.WaitRoom)bool{
	for _, value := range waitRoom.NewMembers{
		if strings.Contains(client, value){
			return true
		}
	}	
	return false
}

func SymbolIsAvailable(symbolToCheck string)bool{
	availableSymbols := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", ".", ":", "/",
	}
	for _, value := range availableSymbols{
		if value == symbolToCheck{
			return true
		}
	}
	return false
}

func CleanGottenResponse(resp string)string{
	cleanedResponse := ""
	for _, value := range resp{
		if SymbolIsAvailable(string(value)){
			cleanedResponse += string(value)
		}else{
			return cleanedResponse
		}
	}
	return cleanedResponse
}

func GetUpdates(userConfig *Users.User, waitRoom *Window.WaitRoom){
	for{
		requestText := fmt.Sprintf("GetMembersInLobby///%s", userConfig.LobbyID)
		userConfig.Conn.Write([]byte(requestText))
		buff := make([]byte, 144)

		userConfig.Conn.Read(buff)
		if !MessageIsEmpty(buff){
			cleanedResp := CleanGottenResponse(string(buff))
			splitedBuff := strings.Split(cleanedResp, "//")
			for _, value := range splitedBuff{
				if !ClientIsWritten(value, waitRoom) && len(waitRoom.NewMembers) <= 5{
					waitRoom.NewMembers = append(waitRoom.NewMembers, value)
				}
			}
		}
	}
}

func CreateLobbyWaitRoom(winConf Window.WindowConfig, currState *Users.States, userConfig *Users.User, waitRoom *Window.WaitRoom){
	if !waitRoom.GettingUpdates{
		go GetUpdates(userConfig, waitRoom)
		waitRoom.GettingUpdates = true
	}
	Window.DrawWaitRoomMenuBG(winConf)

	//Listens to actions 

	ListenForChanges(winConf, currState)
	//Writes announcement for the waiting room
	winConf.TextAreas.NewMembersAnnouncement.Clear()
	winConf.TextAreas.NewMembersAnnouncement.Write([]byte("Wait for members!"))
	winConf.TextAreas.NewMembersAnnouncement.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.NewMembersAnnouncement.Orig, 4))
	
	winConf.TextAreas.NewMembersTextArea.Clear()
	//fmt.Println(waitRoom.NewMembers)
	for _, value := range waitRoom.NewMembers{
		winConf.TextAreas.NewMembersTextArea.Write([]byte(value + "\n"))
	}
	winConf.TextAreas.NewMembersTextArea.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.NewMembersTextArea.Orig, 2.5))
}

func ListenForChanges(winConf Window.WindowConfig, currState *Users.States){
	if winConf.WindowUpdation.Frame % 10 == 0{
		if (winConf.Win.MousePosition().X >= 361 && winConf.Win.MousePosition().X <= 596) && (winConf.Win.MousePosition().Y >= 73 && winConf.Win.MousePosition().Y <= 165) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			currState.SetGame()
		}
	}
	winConf.WindowUpdation.Frame++
}