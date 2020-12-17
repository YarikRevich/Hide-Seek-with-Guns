package GameProcess

import (
	"github.com/faiface/pixel"
	"strings"
	"Game/Utils"
	"fmt"
	_ "time"
	"Game/Window"
	"Game/Interface/GameProcess/ConfigParsers"
	"Game/Heroes/Users"
	"Game/Heroes/Animation"
	"github.com/faiface/pixel/pixelgl"
	"Game/Interface/GameProcess/Map"
)

func KeyBoardButtonListener(userConfig *Users.User, winConf *Window.WindowConfig){

	heroBorder := Map.HeroBorder(&Map.HB{})
	heroBorder.Init(winConf.BGImages.Game)
	camBorder := Map.CamBorder(&Map.CB{})
	camBorder.Init(winConf.BGImages.Game)

	if winConf.Win.Pressed(pixelgl.KeyW){
		if userConfig.Y <= heroBorder.Top(){
			userConfig.Y += 5
		}
		if (winConf.Cam.CamPos.Y*2) < camBorder.Top(){
			if userConfig.Y >= int(winConf.Win.Bounds().Center().Y){
				winConf.Cam.CamPos.Y += 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyA){
		if userConfig.X >= heroBorder.Left(){
			userConfig.X -= 5
		}
		if winConf.Cam.CamPos.X >= camBorder.Left(){
			if userConfig.X <= int(winConf.Win.Bounds().Center().X){
				winConf.Cam.CamPos.X -= 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyS){
		if userConfig.Y >= heroBorder.Bottom(){
			userConfig.Y -= 5
		}
		if winConf.Cam.CamPos.Y >= camBorder.Bottom(){
			if userConfig.Y <= int(winConf.Win.Bounds().Center().Y){	
				winConf.Cam.CamPos.Y -= 5
			}
		}
	}else if winConf.Win.Pressed(pixelgl.KeyD){
		if userConfig.X <= heroBorder.Right(){
			userConfig.X += 5
		}
		if (winConf.Cam.CamPos.X*2) != camBorder.Right(){
			if userConfig.X >= int(winConf.Win.Bounds().Center().X){	
				winConf.Cam.CamPos.X += 5
			}
		}
	} 
}

func ListenToCollisions(userConfig *Users.User){
	//Firstly looks at the hero's coords
	//and due to them checks whether they
	//are the collisions and does some changes.

	collisions := Map.Collisions(&Map.C{})
	collisions.Init()
	if collisions.IsCollision(pixel.V(float64(userConfig.X), float64(userConfig.Y))){
		collisions.React(pixel.V(float64(userConfig.X), float64(userConfig.Y)))
	}
}

func ReDraw(otherUsers *[]*Users.User, winConf *Window.WindowConfig){
	for _, value := range *otherUsers{
		Animation.MoveAndChangeAnim(value, winConf)
	}
}

func ChangePos(userConfig *Users.User, winConf *Window.WindowConfig){
	KeyBoardButtonListener(userConfig, winConf)
	Animation.MoveAndChangeAnim(userConfig, winConf)
}

func ListenToUsersInfo(userConfig *Users.User)string{
	
	buff := make([]byte, 4096)
	userConfig.Conn.Read(buff)
	return string(buff)
}

func CreateGame(userConfig *Users.User, winConf *Window.WindowConfig){

	formattedReq := fmt.Sprintf("GetUsersInfo///%s~", userConfig.LobbyID)
	userConfig.Conn.Write([]byte(formattedReq))
	response := ListenToUsersInfo(userConfig)
	Window.DrawGameBackground(*winConf)


	//Draws main hero
	ChangePos(userConfig, winConf)
	parsedMessage := ConfigParsers.ParseConfig(userConfig)
	userConfig.Conn.Write([]byte(parsedMessage))

	ListenToCollisions(userConfig)

	if ConfigParsers.IsUsersInfo(response){
		if cleaned := Utils.CleanGottenResponse(strings.Split(response, "~/")[1]); len(cleaned) != 0{
			
			//Draws other heroes
			otherUsers := []*Users.User{}
			ConfigParsers.UnparseOthers(cleaned, *userConfig, &otherUsers)
			ReDraw(&otherUsers, winConf)
		}
	}
	Window.UpdateCam(winConf)
}