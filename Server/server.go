package Server

import (
	"Game/Heroes/Users"
	"Game/Utils"
	"net"
	"os"
	"time"
)

const(
	//Containes the delay for read and write timeouts
	delay = 500 * time.Millisecond
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
	// err = conn.SetReadDeadline(time.Now().Add(delay))
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	// err = conn.SetWriteDeadline(time.Now().Add(delay))
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	return conn
}

type Network interface {
	//Special interface for networking along the game

	//Inits the important response and the connector to send the request through
	Init(c net.Conn, u *Users.User, r int, sreq func([]*StartRequest)[]byte, greq func([]*GameRequest)[]byte, t string)

	//Writes given response
	Write()

	//Reads upcoming bytes
	Read() []byte

	//Reads bytes being parsed via start request parser
	ReadStart(func([]byte)[]*StartRequest)[]*StartRequest
	
	//Reads bytes being parsed via game request parser
	ReadGame(func([]byte)([]*GameRequest, error))[]*GameRequest

	//Sets such called speed limit of checking
	SetRegime(r int)
}

type N struct {
	//Struct for networking

	conn net.Conn

	//Saves the connection to send the response through
	userconfig *Users.User     


	//Contains all the important settings fro two regimes
	//of work for networking. Firstly it takes tryLimits
	//which checks whether response for the request has't
	//come for this places times it sends a new request
	//(for game regime it is not important) and checkOld
	//checks all the indexes of responses not passing
	//old responses returning only the newest ones
	regime	 NC

	//Contains a type of read and write. To chose it you
	//have to place int or a constant
	startreq func([]*StartRequest)[]byte
	gamereq func([]*GameRequest)[]byte

	typ string
}

type NC struct{
	//It is a config for networking settings

	//Contains limits for response failuer
	tryLimit int

	tries int

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

func (n *N) Init(c net.Conn, u *Users.User, r int, sreq func([]*StartRequest)[]byte, greq func([]*GameRequest)[]byte, t string) {
	//Inits the reponse and connection

	n.SetRegime(r)
	n.userconfig = u
	n.conn = c
	n.startreq = sreq
	n.gamereq = greq
	n.typ = t
}

func (n *N) Write() {
	//Writes the response placing deadline before

	for {
		currindex++
		var request []byte
		switch{
		case n.startreq != nil:
			req := NewStartRequest(n.typ)
			req[0].Index = currindex
			request = n.startreq(req)
		case n.gamereq != nil:
			n.userconfig.Networking.Index = currindex
			req := NewGameRequest(n.typ, n.userconfig)
			request = n.gamereq(req)
		}
		var num int
		var err error
		switch{
		case n.conn != nil:
			n.conn.SetWriteDeadline(time.Now().Add(delay))
			num, err = n.conn.Write(request)
		case n.userconfig != nil:
			n.userconfig.Conn.SetWriteDeadline(time.Now().Add(delay))
			num, err = n.userconfig.Conn.Write(request)
		}
		if !((err != nil && num == 0) || num == 0) {
			break
		}
	}
}

func (n *N) Read() []byte {
	//Tries to read upcoming bytes

	buff := make([]byte, 30000)
	for {
		var num int
		var err error
		switch{
		case n.conn != nil:
			n.conn.SetReadDeadline(time.Now().Add(delay))
			num, err = n.conn.Read(buff)
		case n.userconfig != nil:
			n.userconfig.Conn.SetReadDeadline(time.Now().Add(delay))
			num, err = n.userconfig.Conn.Read(buff)
		}
		if !((err != nil && num == 0) || num == 0){
			return buff
		}
		if os.IsTimeout(err) && n.regime.tries == n.regime.tryLimit{
			n.Write()
			n.regime.tries = 0
		}else{
			n.regime.tries++
		}
	}
}

func(n *N) ReadStart(parser func([]byte)[]*StartRequest)[]*StartRequest{
	buff := n.Read()
	cl := Utils.Clean(buff)
	unformatted := parser(cl)
	if unformatted[0].Index == currindex && n.regime.checkOld{
		return unformatted
	}
	return unformatted
}

func(n *N) ReadGame(parser func([]byte)([]*GameRequest, error))[]*GameRequest{
	buff := n.Read()
	cl := Utils.Clean(buff)
	uf, _ := parser(cl)
	if uf[0].Networking.Index == currindex && n.regime.checkOld{
		return uf
	}
	return uf
}