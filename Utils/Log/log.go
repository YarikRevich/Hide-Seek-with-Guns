package Log

import (
	"os"
	"log"
	//"fmt"
	//"flag"
	"strings"
	"Game/Heroes/Users"
	"github.com/go-ping/ping"
)

type Logger interface{
	Init(userConfig *Users.User)
	GetPing()
	Show()
}

type Log struct{
	userConfig *Users.User
	stat *ping.Statistics
}

func (l *Log)Init(userConfig *Users.User){
	l.userConfig = userConfig
	l.stat = &ping.Statistics{}
}

func (l *Log)GetPing(){
	for{
		pinger, err := ping.NewPinger(strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0])
		if err != nil{
			panic(err)
		}
		pinger.Count = 2
		err = pinger.Run()
		if err != nil{
			panic(err)
		}
		l.stat = pinger.Statistics()
	}
}

func (l Log)Show(){
	// var statFlag bool

	// flag.BoolVar(&statFlag, "stat", false, "it turns on the debug regime")
	// flag.Parse()

	if len(os.Args) >= 2 && os.Args[1] == "stat"{
		log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\nServer: [%s]\nDelay: [%s]\nPackets: [PacketsLoss: %f, PacketsSent: %d, PacketsRecv: %d]\n-------------", 
			l.userConfig.Pos.X, l.userConfig.Pos.Y, l.userConfig.PersonalInfo.HeroPicture, l.userConfig.PersonalInfo.LobbyID, strings.Split(l.userConfig.Conn.RemoteAddr().String(), ":")[0], l.stat.AvgRtt, l.stat.PacketLoss, l.stat.PacketsSent, l.stat.PacketsRecv,
		)
	}
}