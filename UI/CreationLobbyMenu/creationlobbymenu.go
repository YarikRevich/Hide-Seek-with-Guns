package CreationLobbyMenu

import (
	"fmt"
	"strings"
	"Game/Utils"
	"Game/Window"
	"Game/Server"
	"Game/Heroes/Users"
	"Game/Components/Map"
	"Game/Components/States"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type CreationLobbyMenu struct{
	//It is such called stage struct
	//it uses all the important methods
	//for the corrisponding 'Stage' interface

	winConf        *Window.WindowConfig

	currState      *States.States
	
	userConfig     *Users.User

	mapComponents  Map.MapConf

}

func (c *CreationLobbyMenu)Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf){
	c.winConf       = winConf
	c.currState     = currState
	c.userConfig    = userConfig
	c.mapComponents = mapComponents
}

func (c *CreationLobbyMenu)ProcessNetworking(){

	if c.currState.SendStates.CreateRoom{
		server := Server.Network(new(Server.N))

		writtenID := strings.Join(c.winConf.TextAreas.CreateLobbyInput.WrittenText, "")
		c.userConfig.LobbyID = writtenID
		server.Init(fmt.Sprintf("CreateLobby///%s", writtenID), c.userConfig.Conn, 1)
		server.Write()
		server.Read()

		server.Init(
			fmt.Sprintf(
				"AddToLobby///%s~/%s/%d/%d/%d/%d/0|0|0|0/%s", 
				writtenID,
				c.userConfig.Username,
				c.userConfig.X,
				c.userConfig.Y,
				c.userConfig.UpdationRun,
				c.userConfig.CurrentFrame,
				c.userConfig.HeroPicture,
			), 
			c.userConfig.Conn,
			1,
		)
		server.Write()
		response := server.Read()

		if !Utils.MessageIsEmpty(response){
			c.currState.MainStates.SetWaitRoom()
			c.currState.SendStates.CreateRoom = false
			c.winConf.WaitRoom.RoomType = "create"
			c.winConf.TextAreas.CreateLobbyInput.WrittenText = []string{}
		}
	}
}

func (c *CreationLobbyMenu)ProcessKeyboard(){

	if c.winConf.WindowUpdation.CreationMenuFrame % 5 == 0 && c.winConf.WindowUpdation.CreationMenuFrame != 0{
		if (c.winConf.Win.MousePosition().X >= 21 && c.winConf.Win.MousePosition().X <= 68) && (c.winConf.Win.MousePosition().Y >= 468 && c.winConf.Win.MousePosition().Y <= 511) && c.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
			c.winConf.TextAreas.CreateLobbyInput.WrittenText = []string{}
			c.currState.MainStates.SetStartMenu()
		}
	}

	if (c.winConf.Win.MousePosition().X >= 342 && c.winConf.Win.MousePosition().X <= 612) && (c.winConf.Win.MousePosition().Y >= 75 && c.winConf.Win.MousePosition().Y <= 172){
		c.winConf.DrawCreationLobbyMenuBGPressedButton()
	}

	if (c.winConf.Win.MousePosition().X >= 342 && c.winConf.Win.MousePosition().X <= 612) && (c.winConf.Win.MousePosition().Y >= 75 && c.winConf.Win.MousePosition().Y <= 172) && c.winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		c.currState.SendStates.CreateRoom = true
	}
}

func (c *CreationLobbyMenu)ProcessTextInput(){

	c.winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Clear()
	if c.winConf.Win.Pressed(pixelgl.KeyBackspace){
		if c.winConf.WindowUpdation.CreationMenuFrame % 8 == 0{
			if len(c.winConf.TextAreas.CreateLobbyInput.WrittenText) > 0{
				c.winConf.TextAreas.CreateLobbyInput.WrittenText = Utils.RemoveIndex(c.winConf.TextAreas.CreateLobbyInput.WrittenText, len(c.winConf.TextAreas.CreateLobbyInput.WrittenText)-1)
			}
		}
	}
	if len(c.winConf.Win.Typed()) != 0 && len(c.winConf.TextAreas.CreateLobbyInput.WrittenText) < 10{
		c.winConf.TextAreas.CreateLobbyInput.WrittenText = append(c.winConf.TextAreas.CreateLobbyInput.WrittenText, c.winConf.Win.Typed())
	}
	for _, value := range c.winConf.TextAreas.CreateLobbyInput.WrittenText{
		fmt.Fprintf(c.winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea, value)
	}
	c.winConf.WindowUpdation.CreationMenuFrame++
	c.winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Draw(c.winConf.Win, pixel.IM.Scaled(c.winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea.Orig, 3))
}

func (c *CreationLobbyMenu)ProcessMusic(){
	//WARNING: it is not implemented!
}

func (c *CreationLobbyMenu)DrawAnnouncements(){

	c.winConf.TextAreas.WriteIDTextArea.Clear()
	fmt.Fprintf(c.winConf.TextAreas.WriteIDTextArea, "Write your lobby ID!")
	c.winConf.TextAreas.WriteIDTextArea.Draw(c.winConf.Win, pixel.IM.Scaled(c.winConf.TextAreas.WriteIDTextArea.Orig, 4))

}

func (c *CreationLobbyMenu)DrawElements(){

	c.winConf.DrawCreationLobbyMenuBG()
	if c.currState.SendStates.CreateRoom{
		c.winConf.DrawCreationLobbyMenuBGPressedButton()
	}
}

func (c *CreationLobbyMenu)Run(){

	c.DrawElements()

	c.ProcessKeyboard()

	c.DrawAnnouncements()

	c.ProcessTextInput()

	c.ProcessNetworking()
}