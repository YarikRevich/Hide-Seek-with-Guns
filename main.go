package main

import (
	"Game/Heroes/Users"
	"Game/Interface/CreationLobbyMenu"
	"Game/Interface/GameProcess"
	"Game/Interface/GameProcess/Map"
	"Game/Interface/JoinLobbyMenu"
	"Game/Interface/LobbyWaitRoom"
	"Game/Interface/Menu"
	"Game/Server"
	"Game/Utils"
	"Game/Utils/Log"
	"Game/Window"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/gookit/color"
)

var (
	frames = 0
	second = time.Tick(time.Second)
)

func getMainServer()string{
	file, err := os.OpenFile("config.txt", os.O_CREATE|os.O_RDONLY, 0755)
	if err != nil{
		panic(err)
	}
	buff := make([]byte, 4096)
	file.Read(buff)
	var cleanedBuff []byte
	for _, value := range buff{
		if value == 0{
			continue
		}
		cleanedBuff = append(cleanedBuff, value)
	}
	if len(cleanedBuff) == 0{
		fmt.Println("Config file is empty! Please write your own main server adress or copy it from the repo!")
		os.Exit(0)
	}
	return string(cleanedBuff)
}

func connectToMainServer(adress string)*net.UDPConn{
	splittedAdress := strings.Split(strings.Replace(adress, "\n", "", 1), ":")
	ip, portStr := splittedAdress[0], splittedAdress[1]
	port, err := strconv.Atoi(portStr)
	if err != nil{
		panic(err)
	}
	udpaddr := net.UDPAddr{
		Port: port,
		IP: net.ParseIP(ip),
	}
	conn, err := net.DialUDP("udp", nil, &udpaddr)
	if err != nil{
		panic(err)
	}
	return conn
}

func getAvailableServers(conn *net.UDPConn)map[int]string{
	conn.Write([]byte("CheckServers"))
	buff := make([]byte, 4096)
	conn.Read(buff)
	var cleanedBuff []byte
	for _, value := range buff{
		if value == 0{
			continue
		}
		cleanedBuff = append(cleanedBuff, value)
	}
	result := map[int]string{}
	for index, value := range strings.SplitAfter(string(cleanedBuff), " "){
		result[index+1] = value
	}
	return result
}

func formatAvailableServersList(availableservers map[int]string){
	for index, value := range availableservers{
		ip := strings.Split(value, ":")[0]
		pink := color.Magenta.Darken().Render
		fmt.Printf("%s %s: %d)\n", pink("=>"), ip, index)
	}
}

func choseCoresspondingServer(listServers map[int]string)string{
	for{
		color.Yellow.Print("Write the number of server: ")
		var server int
		fmt.Scan(&server)
		value, ok := listServers[server]
		if ok{
			return value
		}
		color.Red.Println("Such one is not available!")
	}
}

func getStartInfo()(string, string){
	color.Green.Println("Chose the server to play on!")
	listServers := getAvailableServers(connectToMainServer(getMainServer()))
	formatAvailableServersList(listServers)
	server := choseCoresspondingServer(listServers)
	fmt.Printf("Chosen server is %s\n", strings.Split(server, ":")[0])
	color.Green.Println("Write your username!")
	var username string
	fmt.Scan(&username)
	return username, server
}

func choseActionGate(winConf *Window.WindowConfig, currState *Users.States, userConfig *Users.User, camBorder Map.CamBorder){
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
		GameProcess.CreateGame(userConfig, winConf, camBorder)
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

	username, server := getStartInfo()

	winConf := Window.CreateWindow()
	winConf.DrawBackgroundImage()
	winConf.LoadCreationLobbyMenuBG()
	winConf.LoadJoinLobbyMenu()
	winConf.LoadWaitRoomMenuBG()
	winConf.LoadWaitRoomJoinBG()
	winConf.LoadGameBackground()
	winConf.LoadHorDoor()
	winConf.LoadVerDoor()
	winConf.DrawAllTextAreas()
	winConf.LoadAvailableHeroImages()
	conn := Server.GetConnection(server)

	randomSpawn := Utils.GetRandomSpawn()

	userConfig := Users.User{
		Username: username,
		Conn: conn,
		X: int(randomSpawn.X),
		Y: int(randomSpawn.Y),
		Game: &Users.Game{ReadWriteUpdate: make(chan string)},
		HeroPicture: Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
		CurrentFrameMatrix: []string{"0", "0", "0", "0"},
	}

	CurrState := Users.States{StartMenu: true}

	camBorder := Map.CamBorder(&Map.CB{})
	camBorder.Init(winConf.BGImages.Game)

	winConf.SetCam(userConfig, camBorder)

	log := Log.Logger(&Log.Log{})
	log.Init(&userConfig)

	for !winConf.Win.Closed(){

		//Shows statistics about user if argument is placed
		log.Show()

		frames++
		select{
		case <- second:
			winConf.Win.SetTitle(fmt.Sprintf("Hide and seek| %d", frames))
			frames = 0
		default:
			winConf.UpdateBackground()
			choseActionGate(&winConf, &CurrState, &userConfig, camBorder)
			winConf.Win.Update()
		}
	}
}

func main(){
	pixelgl.Run(run)
	color.Green.Println("Goodbye!")
}