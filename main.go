package main

import (
	"Game/Components/Map"
	"Game/Components/Start"
	"Game/Components/States"
	"Game/Heroes/Users"
	"Game/Server"
	"Game/UI"
	"Game/UI/CreationLobbyMenu"
	"Game/UI/GameProcess"
	"Game/UI/JoinLobbyMenu"
	"Game/UI/LobbyWaitRoom"
	"Game/UI/StartMenu"
	"Game/Utils"
	"Game/Utils/Log"
	"Game/Window"
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/gookit/color"
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
	case currState.MainStates.StartMenu:
		stage := UI.Stage(new(StartMenu.StartMenu))
		stage.Init(winConf, currState, userConfig, mapComponents)
		stage.Run()
	case currState.MainStates.CreateLobbyMenu:
		stage := UI.Stage(new(CreationLobbyMenu.CreationLobbyMenu))
		stage.Init(winConf, currState, userConfig, mapComponents)
		stage.Run()
	case currState.MainStates.JoinLobbyMenu:
		stage := UI.Stage(new(JoinLobbyMenu.JoinLobbyMenu))
		stage.Init(winConf, currState, userConfig, mapComponents)
		stage.Run()
	case currState.MainStates.WaitRoom:
		stage := UI.Stage(new(LobbyWaitRoom.LobbyWaitRoom))
		stage.Init(winConf, currState, userConfig, mapComponents)
		stage.Run()
	case currState.MainStates.Game:
		stage := UI.Stage(new(GameProcess.GameProcess))
		stage.Init(winConf, currState, userConfig, mapComponents)
		stage.Run()
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
		Username:           username,
		Conn:               conn,
		X:                  int(randomSpawn.X),
		Y:                  int(randomSpawn.Y),
		Game:               &Users.Game{ReadWriteUpdate: make(chan string)},
		HeroPicture:        Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
		CurrentFrameMatrix: []string{"0", "0", "0", "0"},
	}

	//Sets current state at 'StartMenu'
	currState := States.States{
		MainStates:       new(States.MainStates),
		MusicStates:      new(States.MusicStates),
		SendStates:       new(States.SendStates), 
		NetworkingStates: new(States.NetworkingStates),
	}
	currState.MainStates.SetStartMenu()

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