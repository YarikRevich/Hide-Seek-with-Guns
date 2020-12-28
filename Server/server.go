package Server


import (
	"net"
	"time"
)

func GetConnection(adress string)net.Conn{
	//Creates a connection for user

	var waittime time.Duration = 20000
	if len(adress) == 0{
		adress = "localhost:10100"
		waittime = 30000
	}
	conn, err := net.Dial("udp", adress)
	if err != nil{
		panic(err)
	}
	conn.SetWriteDeadline(time.Now().Add(waittime * time.Millisecond))
	conn.SetReadDeadline(time.Now().Add(waittime * time.Millisecond))
	conn.Write([]byte("!_"))
	return conn
}	