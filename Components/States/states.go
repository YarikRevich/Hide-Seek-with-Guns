package States

type States struct{
	MainStates       *MainStates
	MusicStates      *MusicStates
	SendStates       *SendStates
	NetworkingStates *NetworkingStates
}

type MainStates struct{
	StartMenu       bool
	CreateLobbyMenu bool
	JoinLobbyMenu   bool
	WaitRoom        bool
	Game            bool
}

type MusicStates struct{
	PlayGameSound bool
}

type SendStates struct{
	//This struct contains all the states
	//for sending request to the network for
	//'create' and 'join' rooms. E.g you have
	//pressed 'create lobby' button, now, there 'createroom'
	//state is active and it says there will be sent a request
	//to get the confirmation about the lobby creation

	//It is a state for 'join lobby' button to send a request
	//to get the confirmation about joining to lobby
	JoinRoom   bool

	//It is a state for 'create lobby' button to send a request
	//to get the confirmation about lobby creation
	CreateRoom bool
}

type NetworkingStates struct{
	LobbyWaitRoom bool
	GameProcess   bool
}


func (s *MainStates)SetStartMenu(){
	// Sets state to 'StartMenu'

	s.StartMenu = true
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *MainStates)SetCreateLobbyMenu(){
	//Sets state to 'CreateLobbyMenu'

	s.StartMenu = false
	s.CreateLobbyMenu = true
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = false
}

func (s *MainStates)SetJoinLobbyMenu(){
	//Sets state to 'JoinLobbyMenu'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = true
	s.WaitRoom = false
	s.Game = false
}

func (s *MainStates)SetWaitRoom(){
	//Sets state to 'WaitRoom'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = true
	s.Game = false
}

func (s *MainStates)SetGame(){
	//Sets state to 'Game'

	s.StartMenu = false
	s.CreateLobbyMenu = false
	s.JoinLobbyMenu = false
	s.WaitRoom = false
	s.Game = true
}