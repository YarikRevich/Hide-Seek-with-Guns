package Users

import (
	"net"
)

type User struct{
	X int
	Y int
	Conn net.Conn
	Game *Game
	Username string
	LobbyID string
	HeroPicture string
	UpdationRun int
	CurrentFrame int
	CurrentFrameMatrix []string
}

type Game struct{
	GameStarted bool
	ReadWriteUpdate chan string
}

