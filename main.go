package main

import (
	"github.com/faiface/pixel/pixelgl"
	_ "github.com/faiface/pixel"
	_ "Game/Utils"
	"Game/Window"
	"Game/Interface/Menu"
	"Game/Interface/CreationLobbyMenu"
	"Game/Interface/LobbyWaitRoom"
	"Game/Heroes/Users"
	"Game/Server"
	"fmt"
)


// func runConnectionWithServer(conn net.Conn, readWriteChan chan string){
// 	//Gets dialed connection and reads all the messages 

// 	go GetUpdates(conn, readWriteChan)
// }

//func SendNewConfig(parsedMessage string){}


func startComponentsPackage(){}


func choseActionGate(winConf *Window.WindowConfig, currState *Users.States, userConfig *Users.User, waitRoom *Window.WaitRoom){
	//Choses action which refers to current state
	
	if currState.StartMenu{
		Menu.ListenForActions(*winConf, currState)
		return
	}
	if currState.CreateLobbyMenu{
		CreationLobbyMenu.CreateLobbyMakingMenu(winConf, currState, userConfig)
		return
	}
	if currState.JoinLobbyMenu{
		fmt.Println("joinlobby menu")
		return
	}
	if currState.WaitRoom{
		LobbyWaitRoom.CreateLobbyWaitRoom(*winConf, currState, userConfig, waitRoom)
		return
	}
	if currState.Game{
		fmt.Println("game!")
		return
	}
}

func run(){
	//Creates window and does futher updations

	winConf := Window.CreateWindow()
	Window.DrawBackgroundImage(&winConf)
	Window.LoadCreationLobbyMenuBG(&winConf)
	Window.LoadWaitRoomMenuBG(&winConf)
	Window.DrawAllTextAreas(&winConf)
	//Loads all the available hero images
	//availableHeroImages := Utils.GetAvailableHeroImages()
	//allSystemImages := Utils.LoadAllSystemImages()
	//Runs connection with server
	//readWriteChan := make(chan string)
	conn := Server.GetConnection()
	//go runConnectionWithServer(conn, readWriteChan)

	userConfig := Users.User{
		Conn: conn, 
		HeroPicture: "user1",//Utils.GetRandomName(availableHeroImages),
	}
	CurrState := Users.States{StartMenu: true}

	WaitRoom := Window.WaitRoom{}
	//CreateLobby(&userConfig)
	for !winConf.Win.Closed(){
		Window.UpdateBackground(&winConf)
		choseActionGate(&winConf, &CurrState, &userConfig, &WaitRoom)
		//runMainPipeLine(readWriteChan, &userConfig, &winConf, availableHeroImages)
		winConf.Win.Update()
	}
}

func main(){
	pixelgl.Run(run)
}