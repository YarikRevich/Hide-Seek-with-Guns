package Server


import (
	"net"
)

func GetConnection(adress string)net.Conn{
	//Creates a connection for user

	conn, err := net.Dial("udp", adress)
	if err != nil{
		panic(err)
	}
	conn.Write([]byte("!_"))
	return conn
}	

func GetUpdates(conn net.Conn, readWriteChan chan string){
	conn.Write([]byte("UpdateUser///1~10/20/0/0/0|0|0|0/testhero/user1::/20/30/0/0/0|0|0|0/testhero/user2"))
	for{
		buff := make([]byte, 2048)
		conn.Read(buff)
		readWriteChan <- string(buff)
	}
}