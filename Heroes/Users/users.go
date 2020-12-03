package Users

import (
	"net"
)

type User struct{
	Username string
	LobbyID string
	Conn net.Conn
	X int
	Y int
	HeroPicture string
	CurrentFrame int
	CurrentFrameMatrix []string
	UpdationRun int
}

type States struct{
	StartMenu bool
	CreateLobbyMenu bool
	JoinLobbyMenu bool
	WaitRoom bool
	Game bool
}

func (s *States)SetStartMenu(){
	s.StartMenu = true
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetCreateLobbyMenu(){
	s.StartMenu = false
	s.CreateLobbyMenu = true
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetJoinLobbyMenu(){
	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = true
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetWaitRoom(){
	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = true
	s.Game = false
}

func (s *States)SetGame(){
	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = true
}

