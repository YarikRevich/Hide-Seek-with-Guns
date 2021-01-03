package Server

import (
	"fmt"
	"net"
	"time"
	"strings"
	"strconv"
)

const delay = 100 * time.Millisecond

var(
	currstate = 0
)


func GetConnection(adress string) net.Conn {
	//Creates a connection

	if len(adress) == 0 {
		adress = "localhost:10100"
	}
	addr, err := net.ResolveUDPAddr("udp", adress)
	if err != nil{
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	conn.SetReadDeadline(time.Now().Add(delay))
	conn.SetWriteDeadline(time.Now().Add(delay))
	conn.Write([]byte("!_"))
	return conn
}

type Network interface {
	//Special interface for networking along the game

	//Inits the important response and the connector to send the request through
	Init(request string, conn net.Conn)

	//Writes given response
	Write()

	//Reads upcoming bytes
	Read() []byte

	//Checks whether the request is old
	IsOld(buff []byte) bool

	//Formats response
	Format(request string)

	//Informats response
	Unformat(response string) int

	//Formats to work with
	FormatToWorkWith(buff []byte)[]byte
}

type N struct {
	//Struct for networking

	//Response saves the response as it is not banal :)
	request string

	//Saves the connection to send the response through
	conn     net.Conn
}

func (n *N) Init(request string, conn net.Conn) {
	//Inits the reponse and connection

	
	n.Format(request)
	n.conn = conn
}

func (n *N) Write() {
	//Writes the response placing deadline before

	for {
		n.conn.SetWriteDeadline(time.Now().Add(delay))
		num, err := n.conn.Write([]byte(n.request))
		if !((err != nil && num == 0) || num == 0) {
			break
		}
	}
}

func (n *N) Read() []byte {
	//Tries to read upcoming bytes

	buff := make([]byte, 40000)
	var timeout int = 0
	for {
		n.conn.SetReadDeadline(time.Now().Add(delay))
		num, err := n.conn.Read(buff)
		if !((err != nil && num == 0) || num == 0){
			if !n.IsOld(buff){
				return n.FormatToWorkWith(buff)
			}
		}
		timeout++
		if timeout == 5{
			n.Write()
			timeout = 0
		}
	}
}

func (n N) FormatToWorkWith(buff []byte)[]byte{
	var cleaned []byte
	for _, value := range buff{
		if value != 0{
			cleaned = append(cleaned, value)
		}
	}
	splitted := strings.Split(string(buff), "_")
	if len(splitted) == 1{
		return []byte(splitted[0])
	}
 	return []byte(splitted[1])
}

func (n *N) IsOld(buff []byte)bool{

	num := n.Unformat(string(buff))
	if num != currstate{
		return true
	}
	return false
}

func (n *N)Unformat(response string)int{
	reqnum := strings.Split(response, "_")
	num, err := strconv.Atoi(reqnum[0])
	if err != nil{
		panic(err)
	}
	return num
}

func (n *N) Format(request string){
	currstate++
	n.request = fmt.Sprintf("%d_%s", currstate, request)
}
