package Server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const(
	//Containes the delay for read and write timeouts

	delay = 12000 * time.Microsecond
)

var(
	//Containes the index of the current response to wait for

	currindex = 0
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
	Init(request string, conn net.Conn, r int)

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

	//Sets such called speed limit of checking
	SetRegime(r int)
}

type N struct {
	//Struct for networking

	//Response saves the response as it is not banal :)
	request  string

	//Saves the connection to send the response through
	conn     net.Conn

	//Contains all the important settings fro two regimes
	//of work for networking. Firstly it takes tryLimits
	//which checks whether response for the request has't
	//come for this places times it sends a new request
	//(for game regime it is not important) and checkOld
	//checks all the indexes of responses not passing
	//old responses returning only the newest ones
	regime	 NC
}

type NC struct{
	//It is a config for networking settings

	//Contains limits for response failuer
	tryLimit int

	//Sets whether the 'old-message indexing' is important
	checkOld bool
}

func (n *N) SetRegime(r int){
	//Sets configuration for the networking
	//due to places num of the corrisponding regime

	switch r{
	case 0:
		n.regime = NC{tryLimit: 1, checkOld: false}
	case 1:
		n.regime = NC{tryLimit: 4, checkOld: true}
	}
}

func (n *N) Init(request string, conn net.Conn, r int) {
	//Inits the reponse and connection

	n.SetRegime(r)
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

	buff := make([]byte, 30000)
	for {
		n.conn.SetReadDeadline(time.Now().Add(delay))
		num, err := n.conn.Read(buff)
		if !((err != nil && num == 0) || num == 0){
			if !n.IsOld(buff) && n.regime.checkOld{
				return n.FormatToWorkWith(buff)
			}
			return n.FormatToWorkWith(buff)

		}
		if os.IsTimeout(err){
			n.Write()
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
	//Checks whether response is old

	num := n.Unformat(string(buff))
	if num != currindex{
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
	currindex++
	n.request = fmt.Sprintf("%d_%s", currindex, request)
}
