package States

type States struct{
	StartMenu bool
	CreateLobbyMenu bool
	JoinLobbyMenu bool
	WaitRoom bool
	Game bool
	ComponentsStates *ComponentsStates
}

type ComponentsStates struct{
	PlayGameSound bool
}

func (s *States)SetStartMenu(){
	// Sets state to 'StartMenu'

	s.StartMenu = true
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetCreateLobbyMenu(){
	//Sets state to 'CreateLobbyMenu'

	s.StartMenu = false
	s.CreateLobbyMenu = true
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetJoinLobbyMenu(){
	//Sets state to 'JoinLobbyMenu'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = true
	s.WaitRoom = false
	s.Game = false
}

func (s *States)SetWaitRoom(){
	//Sets state to 'WaitRoom'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = true
	s.Game = false
}

func (s *States)SetGame(){
	//Sets state to 'Game'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = true
}