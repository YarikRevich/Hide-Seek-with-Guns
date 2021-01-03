package GameProcess

import (
	"fmt"
	_ "time"
	"strings"
	"Game/Utils"
	"Game/Window"
	"Game/Server"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"Game/Components/Sound"
	"Game/Components/States"
	Map "Game/Components/Map"
	"Game/Interface/GameProcess/ConfigParsers"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func KeyBoardButtonListener(userConfig *Users.User, winConf *Window.WindowConfig, mapComponents Map.MapConf){

	currPosition := pixel.V(float64(userConfig.X), float64(userConfig.Y))
	vector, _, ok := mapComponents.GetCollisions().IsDoor(currPosition)

	if ok{
		mapComponents.GetCollisions().DeleteDoor(vector)
	}
	if !mapComponents.GetCollisions().IsNearDeletedDoor(currPosition){
	 	mapComponents.GetCollisions().RecreateDeletedDoors()
	}

	mapComponents.GetCollisions().DrawDoors(winConf.DrawHorDoor, winConf.DrawVerDoor)
	
	switch {
	case winConf.Win.Pressed(pixelgl.KeyW):
		if mapComponents.GetCollisions().IsCollision(pixel.V(float64(userConfig.X), float64(userConfig.Y+2))){
			return
		}
		if userConfig.Y <= mapComponents.GetHeroBorder().Top(){
			userConfig.Y += 3
		}
		if winConf.Cam.CamPos.Y < mapComponents.GetCamBorder().Top(){
			if userConfig.Y >= int(winConf.Win.Bounds().Center().Y){
				winConf.Cam.CamPos.Y += 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyA):
		if mapComponents.GetCollisions().IsCollision(pixel.V(float64(userConfig.X-2), float64(userConfig.Y))){
			return
		}
		if userConfig.X >= mapComponents.GetHeroBorder().Left(){
			userConfig.X -= 3
		}
		if winConf.Cam.CamPos.X >= mapComponents.GetCamBorder().Left(){
			if userConfig.X <= int(winConf.Win.Bounds().Center().X){
				winConf.Cam.CamPos.X -= 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyS):
		if mapComponents.GetCollisions().IsCollision(pixel.V(float64(userConfig.X), float64(userConfig.Y-2))){
			return
		}
		if userConfig.Y >= mapComponents.GetHeroBorder().Bottom(){
			userConfig.Y -= 3
		}
		if winConf.Cam.CamPos.Y >= mapComponents.GetCamBorder().Bottom(){
			if userConfig.Y <= int(winConf.Win.Bounds().Center().Y){	
				winConf.Cam.CamPos.Y -= 5
			}
		}
	case winConf.Win.Pressed(pixelgl.KeyD):
		if mapComponents.GetCollisions().IsCollision(pixel.V(float64(userConfig.X+2), float64(userConfig.Y))){
			return
		}
		if userConfig.X <= mapComponents.GetHeroBorder().Right(){
			userConfig.X += 3
		}
		if winConf.Cam.CamPos.X <= mapComponents.GetCamBorder().Right(){
			if userConfig.X >= int(winConf.Win.Bounds().Center().X){	
				winConf.Cam.CamPos.X += 5
			}
		}
	} 
}

func ReDraw(otherUsers *[]*Users.User, winConf *Window.WindowConfig){
	for _, value := range *otherUsers{
		Animation.MoveAndChangeAnim(value, winConf)
	}
}

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig, mapComponents Map.MapConf){
	KeyBoardButtonListener(userConfig, winConf, mapComponents)
	Animation.MoveAndChangeAnim(userConfig, winConf)
}

func CreateGame(userConfig *Users.User, winConf *Window.WindowConfig, states *States.States, mapComponents Map.MapConf){

	server := Server.Network(new(Server.N))
	server.Init(fmt.Sprintf("GetUsersInfo///%s~", userConfig.LobbyID), userConfig.Conn, 0)
	server.Write()
	response := server.Read()

	sound := Sound.Sound(new(Sound.S))
	sound.Init(states)
	sound.Play()

	//Draws main hero
	
	winConf.DrawGameBackground()
	ChangePos(userConfig, winConf, mapComponents)
	//winConf.DrawGoldChest() : It is testing now :)

	server.Init(ConfigParsers.ParseConfig(userConfig), userConfig.Conn, 0)
	server.Write()

	if ConfigParsers.IsUsersInfo(string(response)){
		if cleaned := Utils.CleanGottenResponse(strings.Split(string(response), "~/")[1]); len(cleaned) != 0{
			
			//Draws other heroes
			otherUsers := []*Users.User{}
			ConfigParsers.UnparseOthers(cleaned, *userConfig, &otherUsers)
			ReDraw(&otherUsers, winConf)
		}
	}
	winConf.DrawDarkness(pixel.V((float64(userConfig.X)*2.5)-31, (float64(userConfig.Y)*2.5)-30))
	mapComponents.GetCam().UpdateCam()
}