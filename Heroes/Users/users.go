package Users

import (
	"net"
)

type User struct{
	Conn net.Conn
	Pos *Pos
	GameInfo *GameInfo
	PersonalInfo *PersonalInfo
	Animation *Animation
	Networking *Networking
	Context    *Context
}

type Pos struct{
	X int
	Y int
}

type GameInfo struct{
	Health       int
	//WeaponRadius int
}

type PersonalInfo struct{
	LobbyID string
	Username string
	HeroPicture string
}

type Animation struct{
	UpdationRun int
	CurrentFrame int
	CurrentFrameMatrix []float64	
}

type Networking struct{
	Index int
}

type Context struct{
	Additional []string
}



