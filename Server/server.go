package Server


import (
	"net"
	"time"
)

func GetConnection(adress string)net.Conn{
	//Creates a connection for user

	conn, err := net.Dial("udp", adress)
	if err != nil{
		panic(err)
	}
	conn.SetReadDeadline(time.Now().Add(6000 * time.Millisecond))
	conn.Write([]byte("!_"))
	return conn
}	