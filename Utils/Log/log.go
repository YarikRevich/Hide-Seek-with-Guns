package Log

import (
	"os"
	"log"
	"Game/Heroes/Users"
)

type Logger interface{
	Init(userConfig *Users.User)
	Show()
}

type Log struct{
	userConfig *Users.User
}

func (l *Log)Init(userConfig *Users.User){
	l.userConfig = userConfig
}

func (l Log)Show(){
	if len(os.Args) >= 2 && os.Args[1] == "stat"{
		log.Printf("\n-------------\nX: [%d], Y: [%d]\nHeroImage: [%s]\nLobbyID: [%s]\n-------------", 
			l.userConfig.X, l.userConfig.Y, l.userConfig.HeroPicture, l.userConfig.LobbyID,
		)
	}
}