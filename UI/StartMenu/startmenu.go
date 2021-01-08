package StartMenu

import (
	"fmt"
	"math"
	"Game/Window"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
	
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type StartMenu struct {
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf *Window.WindowConfig

	currState *States.States

	userConfig *Users.User

	mapComponents Map.MapConf
}

func (s *StartMenu) Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf) {
	s.winConf = winConf
	s.currState = currState
	s.userConfig = userConfig
	s.mapComponents = mapComponents
}

func (s *StartMenu) ProcessNetworking() {
	//WARNING: it is not implemented!
}

func (s *StartMenu) ProcessKeyboard() {

	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 320 && s.winConf.Win.MousePosition().Y <= 415){
		s.winConf.DrawStartMenuPressedCreateButton()
	}

	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 144 && s.winConf.Win.MousePosition().Y <= 242){
		s.winConf.DrawStartMenuPressedJoinButton()
	}

	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 320 && s.winConf.Win.MousePosition().Y <= 415) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft) {
		s.currState.MainStates.SetCreateLobbyMenu()
	}
	if (s.winConf.Win.MousePosition().X >= 379 && s.winConf.Win.MousePosition().X <= 590) && (s.winConf.Win.MousePosition().Y >= 144 && s.winConf.Win.MousePosition().Y <= 242) && s.winConf.Win.Pressed(pixelgl.MouseButtonLeft) {
		s.currState.MainStates.SetJoinLobbyMenu()
	}
}

func (s *StartMenu) ProcessTextInput() {
	//WARNING: it is not implemented!
}

func (s *StartMenu) ProcessMusic() {
	//WARNING: it is not implemented!
}

func (s *StartMenu) DrawAnnouncements() {
	//Here it draws logo with additional animation

	s.winConf.TextAreas.GameLogo.Clear()
	IsOld := func(value float64, list []float64) bool {
		//Checks whether new pos is not already used

		for _, i := range list {
			if value == i {
				return true
			}
		}
		return false
	}

	max := float64(4)

	ChangeLogo := func(i float64){
		//It draws logo applying a new color
		//and a new rotation pos appending new
		//params to the 'already used' list

		fmt.Fprint(s.winConf.TextAreas.GameLogo, "Hide&Seek")
		s.winConf.TextAreas.GameLogo.Color = colornames.Orange
		s.winConf.TextAreas.GameLogo.Draw(
			s.winConf.Win, 
			pixel.IM.Scaled(s.winConf.TextAreas.GameLogo.Orig, max-i).Rotated(s.winConf.TextAreas.GameLogo.Orig, -.6),
		)
		s.winConf.StartMenu.DrawedTemporally = append(s.winConf.StartMenu.DrawedTemporally, math.Round(i*10)/10)
	}

	switch s.winConf.StartMenu.Regime {
	case 0:
		//This regime makes logo smaller

		for i := max; i != 0; i = i - 0.1 {
			if !IsOld(math.Round(i*10)/10, s.winConf.StartMenu.DrawedTemporally){
				ChangeLogo(i)
				s.winConf.TextAreas.GameLogo.Orig = pixel.V(s.winConf.TextAreas.GameLogo.Orig.X-math.Round(i*10)/10, s.winConf.TextAreas.GameLogo.Orig.Y+math.Round(i*10)/10)
				break
			}
		}
		if len(s.winConf.StartMenu.DrawedTemporally) == int(max*10+1) {
			s.winConf.StartMenu.DrawedTemporally = []float64{}
			s.winConf.StartMenu.Regime = 1
		}
	case 1:
		//This regime makes logo bigger

		for i := 0.0; i != max; i = i + 0.1 {
			if !IsOld(math.Round(i*10)/10, s.winConf.StartMenu.DrawedTemporally) {
				ChangeLogo(i)
				s.winConf.TextAreas.GameLogo.Orig = pixel.V(s.winConf.TextAreas.GameLogo.Orig.X+math.Round(i*10)/10, s.winConf.TextAreas.GameLogo.Orig.Y-math.Round(i*10)/10)
				break
			}
		}
		if len(s.winConf.StartMenu.DrawedTemporally) == int(max*10+1) {
			s.winConf.StartMenu.DrawedTemporally = []float64{}
			s.winConf.StartMenu.Regime = 0
		}
	}
}

func (s *StartMenu) DrawElements() {
	s.winConf.DrawStartMenuBG()
}

func (s *StartMenu) Run() {

	s.DrawElements()

	s.ProcessKeyboard()

	s.DrawAnnouncements()
}
