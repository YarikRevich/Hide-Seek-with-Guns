package Server


import (
	"net"
)

func GetConnection()net.Conn{
	conn, _ := net.Dial("udp", "127.0.0.1:9001")
	conn.Write([]byte("!_"))
	return conn
}


func GetUpdates(conn net.Conn, readWriteChan chan string){
	conn.Write([]byte("CreateLobby///1"))
	conn.Write([]byte("AddToLobby///1~"))
	conn.Write([]byte("ClosePreparingLobby///1~"))
	conn.Write([]byte("UpdateUser///1~10/20/0/0/0|0|0|0/testhero/user1::/20/30/0/0/0|0|0|0/testhero/user2"))
	for{
		buff := make([]byte, 2048)
		conn.Read(buff)
		readWriteChan <- string(buff)
	}
}