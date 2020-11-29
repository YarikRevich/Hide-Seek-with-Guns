package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"Game/Utils"
	"Game/Window"
)


type States struct{
	StartMenu bool
	CreateLobby bool
	StayInWaitRoom bool
	Game bool
}


// func runConnectionWithServer(conn net.Conn, readWriteChan chan string){
// 	//Gets dialed connection and reads all the messages 

// 	go GetUpdates(conn, readWriteChan)
// }

//func SendNewConfig(parsedMessage string){}


func startComponentsPackage(){}


func choseActionGate(winConf WindowConfig, userConfig *Users.User){
	if StateS.StartMenu{
		createStartMenu(winConf, userConfig)
	}else if StateS.CreateLobby{
		createLobby(userConfig)
	}else if StateS.StayInWaitRoom{
		createLobbyAndStayInWaitRoom()
	}
}

func run(){
	//Creates window and does futher updations

	winConf := Window.CreateWindow()
	bgsprite := Window.DrawBackgroundImage(&winConf.win)
	
	//Loads all the available hero images
	availableHeroImages := Utils.GetAvailableHeroImages()
	//allSystemImages := Utils.LoadAllSystemImages()

	//Runs connection with server
	//readWriteChan := make(chan string)
	conn := GetConnection()
	//go runConnectionWithServer(conn, readWriteChan)

	userConfig := Users.User{
		Conn: conn, 
		HeroPicture: Utils.GetRandomName(availableHeroImages),
	}
	//CreateLobby(&userConfig)
	for !winConf.win.Closed(){
		choseActionGate(winConf, &userConfig)
		//runMainPipeLine(readWriteChan, &userConfig, &winConf, availableHeroImages)
		winConf.win.Update()
	}
}

func main(){
	pixelgl.Run(run)
}