package StartMenu

import (
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
	"github.com/faiface/pixel/pixelgl"
)

type StartMenu struct{
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf        *Window.WindowConfig

	currState      *States.States
	
	userConfig     *Users.User

	mapComponents  Map.MapConf
}

func (s *StartMenu)Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
	s.winConf       = winConf
	s.currState     = currState
	s.userConfig    = userConfig
	s.mapComponents = mapComponents
}

func (s *StartMenu)ProcessNetworking(){
	//WARNING: it is not implemented!
}

func (s *StartMenu)ProcessKeyboard(){
	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 320 && s.winConf.Win.MousePosition().Y <= 415) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		s.currState.MainStates.SetCreateLobbyMenu()
	}
	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 144 && s.winConf.Win.MousePosition().Y <= 242) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		s.currState.MainStates.SetJoinLobbyMenu()
	}
}

func (s *StartMenu)ProcessTextInput(){
	//WARNING: it is not implemented!
}

func (s *StartMenu)ProcessMusic(){
	//WARNING: it is not implemented!
}

func (s *StartMenu)DrawAnnouncements(){
	//WARNING: it is not implemented!
}

func (s *StartMenu)DrawElements(){
	s.winConf.DrawStartMenuBG()
}

func (s *StartMenu)Run(){

	s.DrawElements()
	
	s.ProcessKeyboard()
}