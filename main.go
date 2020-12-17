package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"Game/Window"
	"Game/Interface/Menu"
	"Game/Interface/GameProcess"
	"Game/Interface/CreationLobbyMenu"
	"Game/Interface/LobbyWaitRoom"
	"Game/Interface/JoinLobbyMenu"
	"Game/Heroes/Users"
	"Game/Server"
	"Game/Utils"
	"Game/Utils/Log"
	"time"
)

var (
	frames = 0
	second = time.Tick(time.Second)
)

func choseActionGate(winConf *Window.WindowConfig, currState *Users.States, userConfig *Users.User){
	/* It is a main action gate which choses an
	   important menu to act and to draw. It can
	   chose such menues as:
	   - StartMenu 
	   - CreateLobbyMenu
	   - JoinLobbyMenu
	   - WaitRoom
	   - Game
	*/
	switch{
	case currState.StartMenu:
		Menu.ListenForActions(*winConf, currState)
	case currState.CreateLobbyMenu:
		CreationLobbyMenu.CreateLobbyMakingMenu(winConf, currState, userConfig)
	case currState.JoinLobbyMenu:
		JoinLobbyMenu.CreateJoinLobbyMenu(winConf, currState, userConfig)
	case currState.WaitRoom:
		LobbyWaitRoom.CreateLobbyWaitRoom(winConf, currState, userConfig)
	case currState.Game:
		GameProcess.CreateGame(userConfig, winConf)
	}
}

func run(){
	/* It is a main game starting func.
	   Firstly, it creates window with all the 
	   settings, then, draws starting background image,
	   and loads all the background images for all the 
	   menues. Due to put information configurates user
	   struction. Sets state-machine at the first state.
	   Runs 'choseActionGate' which choses important menu
	   to draw.
	*/

	fmt.Println("Write your username!")
	var username string
	fmt.Scan(&username)

	winConf := Window.CreateWindow()
	Window.DrawBackgroundImage(&winConf)
	Window.LoadCreationLobbyMenuBG(&winConf)
	Window.LoadJoinLobbyMenu(&winConf)
	Window.LoadWaitRoomMenuBG(&winConf)
	Window.LoadWaitRoomJoinBG(&winConf)
	Window.LoadGameBackground(&winConf)
	Window.DrawAllTextAreas(&winConf)
	Window.LoadAvailableHeroImages(&winConf)
	conn := Server.GetConnection()


	userConfig := Users.User{
		Username: username,
		Conn: conn,
		X: int(winConf.BGImages.Game.Bounds().Center().X),
		Y: int(winConf.BGImages.Game.Bounds().Center().Y),
		Game: &Users.Game{ReadWriteUpdate: make(chan string)},
		HeroPicture: Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
		CurrentFrameMatrix: []string{"0", "0", "0", "0"},
	}

	Window.SetCam(&winConf, userConfig)


	CurrState := Users.States{StartMenu: true}

	log := Log.Logger(&Log.Log{})
	log.Init(&userConfig)

	for !winConf.Win.Closed(){

		//Shows statistics about user if argument is placed
		log.Show()

		Window.UpdateBackground(&winConf)
		choseActionGate(&winConf, &CurrState, &userConfig)
		winConf.Win.Update()
		frames++
		select{
		case <- second:
			winConf.Win.SetTitle(fmt.Sprintf("Hide and seek|%d", frames))
			frames = 0
		default:
		}
	}
}

func main(){
	pixelgl.Run(run)
}