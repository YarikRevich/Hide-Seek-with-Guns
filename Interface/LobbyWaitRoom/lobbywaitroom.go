package LobbyWaitRoom

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)


func createLobby(userConfig *Users.User){
	userConfig.Conn.Write([]byte("CreateLobby///1"))
	userConfig.Conn.Write([]byte("AddToLobby///1~"))
}

func CreateLobbyWaitRoom(){}