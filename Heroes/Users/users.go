package Users

import (
	"net"
)

type User struct{
	Username string
	Conn net.Conn
	X int
	Y int
	HeroPicture string
	CurrentFrame int
	CurrentFrameMatrix []string
	UpdationRun int
}
