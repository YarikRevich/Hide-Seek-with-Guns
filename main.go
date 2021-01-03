package main

import (
	"fmt"
	"time"
	"Game/Utils"
	"Game/Window"
	"Game/Server"
	"Game/Utils/Log"
	"Game/Heroes/Users"
	"Game/Interface/Menu"
	"Game/Components/Map"
	"Game/Components/Start"
	"Game/Components/States"
	"Game/Interface/GameProcess"
	"Game/Interface/JoinLobbyMenu"
	"Game/Interface/LobbyWaitRoom"
	"Game/Interface/CreationLobbyMenu"

	"github.com/gookit/color"
	"github.com/faiface/pixel/pixelgl"
)

var (
	frames = 0
	second = time.Tick(time.Second)
)

func choseActionGate(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
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
		GameProcess.CreateGame(userConfig, winConf, currState, mapComponents)
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

	//Gets info from user to place his name and server's name
	username, server := Start.GetStartInfo()
	conn := Server.GetConnection(server)

	//Create window and place all the components
	winConf := Window.CreateWindow()

	//Loads all the components
	winConf.LoadAllImageComponents()
	winConf.LoadAllTextAreas()
	winConf.LoadAvailableHeroImages()

	//Draws background image to start game
	winConf.DrawBackgroundImage()

	//Get user's spawn place
	randomSpawn := Utils.GetRandomSpawn()

	//Configures user's info
	userConfig := Users.User{
		Username: username,
		Conn: conn,
		X: int(randomSpawn.X),
		Y: int(randomSpawn.Y),
		Game: &Users.Game{ReadWriteUpdate: make(chan string)},
		HeroPicture: Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
		CurrentFrameMatrix: []string{"0", "0", "0", "0"},
	}

	//Sets current state at 'StartMenu'
	currState := States.States{StartMenu: true, ComponentsStates: new(States.ComponentsStates)}

	//Configures map
	mapComponents := Map.MapConf(new(Map.MapC))
	mapComponents.Init()
	mapComponents.ConfAll(&winConf, userConfig)

	winConf.Components.AvPlacesForSpaws = mapComponents.GetAnailizer().AnalizeAvailablePlaces()

	//Initiates logger
	log := Log.Logger(new(Log.Log))
	log.Init(&userConfig)

	//Starts pinger
	go log.GetPing()

	for !winConf.Win.Closed(){

		//Shows statistics about user if argument is placed
		log.Show()
	
		frames++
		select{
		case <- second:
			//Sets title of the window with frame rate
			winConf.Win.SetTitle(fmt.Sprintf("Hide and seek| %d", frames))
			frames = 0
		default:
			//Upgrades background
			winConf.UpdateBackground()

			//Goes to the action gate to chose an important one
			choseActionGate(&winConf, &currState, &userConfig, mapComponents)
			
			winConf.Win.Update()
		}
	}
}

func main(){
	pixelgl.Run(run)
	color.Green.Println("Goodbye!")
}