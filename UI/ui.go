package UI

import (
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
)

type Stage interface{
	//It is the main interface for state managing.
	//To run it, it should be started via 'Init' with placed 
	//params(they can be not used in each state, but anyway
	//it is such called convention:))

	//Inits stage with corrisponding params
	Init(*Window.WindowConfig, *States.States, *Users.User, Map.MapConf)

	//Method which works for networking getting the newest comming info
	ProcessNetworking()

	//Method which works to process pressed buttons
	ProcessKeyboard()

	//Method which works to process input from keyboard
	ProcessTextInput()

	//Method which works to process music in game
	ProcessMusic()

	//Method which works to draw all the announcements in scene
	DrawAnnouncements()

	//Method which works to draw all the other elements
	DrawElements()

	//Method which runs chosen scene
	Run()
}