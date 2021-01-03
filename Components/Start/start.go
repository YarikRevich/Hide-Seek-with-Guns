package Start

import (
	"os"
	"fmt"
	"net"
	"time"
	"errors"
	"strconv"
	"strings"
	"github.com/gookit/color"
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
		color.Red.Println("There is no connection for the internet")
		os.Exit(0)
	}
	conn.SetReadDeadline(time.Now().Add(1000 * time.Millisecond))
	conn.SetWriteDeadline(time.Now().Add(1000 * time.Millisecond))
	return conn
}

func getAvailableServers(conn *net.UDPConn)map[int]string{
	conn.Write([]byte("CheckServers"))
	buff := make([]byte, 4096)
	conn.Read(buff)
	conn.Close()
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

func formatAvailableServersList(availableservers map[int]string)error{
	value, _ := availableservers[1]
	if len(availableservers) == 1 && len(value) == 0{
		color.Red.Println("There are no available servers right now!")
		return errors.New("there are no available server")
	}
	for index, value := range availableservers{
		ip := strings.Split(value, ":")[0]
		pink := color.Magenta.Darken().Render
		fmt.Printf("%s %s: %d)\n", pink("=>"), ip, index)
	}
	return nil
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

func GetStartInfo()(string, string){
	color.Green.Println("Chose the server to play on!")
	listServers := getAvailableServers(connectToMainServer(getMainServer()))
	err := formatAvailableServersList(listServers)
	var server string
	if err == nil{
		server = choseCoresspondingServer(listServers)
		fmt.Printf("Chosen server is %s\n", strings.Split(server, ":")[0])
	}

	color.Green.Println("Write your username!")
	var username string
	fmt.Scan(&username)
	return username, server
}