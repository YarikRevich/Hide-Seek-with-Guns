package Window

import (
	"Game/Heroes/Users"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"Game/Utils"
	"Game/Interface/GameProcess/Map"
	"time"
)


type MenuImages struct{
	StartMenuBG *pixel.Sprite
	CreatioLobbyMenuBG *pixel.Sprite
	JoinLobbyMenuBG *pixel.Sprite
	WaitRoomMenuBG *pixel.Sprite
	WaitRoomJoinBG *pixel.Sprite
	Game *pixel.Sprite
}

type TextAreas struct{
	WriteIDTextArea *text.Text
	CreateLobbyInput *CreateLobbyInput
	NewMembersAnnouncement *text.Text
	NewMembersTextArea *text.Text
	CurrentLobbyIDArea *text.Text
	JoinLobbyAnnouncement *text.Text
	JoinLobbyInput *JoinLobbyInput
	LobbyDoesNotExistError *text.Text
}

type JoinLobbyInput struct{
	InputLobbyIDTextArea *text.Text
	WrittenText []string
}

type CreateLobbyInput struct{
	InputLobbyIDTextArea *text.Text
	WrittenText []string
}

type WindowUpdation struct{
	StartMenuFrame int
	CreationMenuFrame int
	JoinLobbyMenuFrame int
	WaitRoomFrame int
}

type WaitRoom struct{
	RoomType string
	GettingUpdates bool
	NewMembers []string
}

type Components struct{
	AvailableHeroImages map[string]pixel.Picture
}

type Cam struct{
	CamPos pixel.Vec
	CamZoom float64
}

type WindowConfig struct{
	Win *pixelgl.Window
	BGImages *MenuImages
	TextAreas *TextAreas
	WindowUpdation *WindowUpdation
	Components *Components
	WindowError *WindowError
	WaitRoom *WaitRoom
	Cam *Cam
}

type WindowError struct{
	LobbyDoesNotExist bool
	LobbyErrorStop time.Time
	LobbyErrorText string
}

func CreateWindow()WindowConfig{

	cfg := pixelgl.WindowConfig{
		Title: "Hide and seek",
		Bounds: pixel.R(0, 0, 950, 550),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil{
		panic(err)
	}
	textArea := TextAreas{CreateLobbyInput: new(CreateLobbyInput), JoinLobbyInput: new(JoinLobbyInput)}
	return WindowConfig{Win: win, BGImages: new(MenuImages), TextAreas: &textArea, WindowUpdation: new(WindowUpdation), WindowError: new(WindowError), Components: new(Components), WaitRoom: new(WaitRoom), Cam: new(Cam)}
}

func collibrateBottom(borders Map.CamBorder, userConfig Users.User)float64{
	bottom := borders.Bottom()
	Y := userConfig.Y
	for{
		if float64(Y) >= bottom{
			return float64(Y)
		}
		Y++
	}
}

func collibrateTop(borders Map.CamBorder, userConfig Users.User)float64{
	top := borders.Top()
	Y := userConfig.Y
	for{
		if float64(Y) <= top{
			return float64(Y)
		}
		Y--
	}
}

func collibrateLeft(borders Map.CamBorder, userConfig Users.User)float64{
	left := borders.Left()
	X := userConfig.X
	for{
		if float64(X) >= left{
			return float64(X)
		}  
		X++
	}
} 

func collibrateRight(borders Map.CamBorder, userConfig Users.User)float64{
	right := borders.Right()
	X := userConfig.X
	for{
		if float64(X) <= right{
			return float64(X)
		}
		X--
	}
}

func collibrate(borders Map.CamBorder, userConfig Users.User)pixel.Vec{
	var X float64
	var Y float64
	Y = collibrateBottom(borders, userConfig)
	NewY := collibrateTop(borders, userConfig)
	if NewY != float64(userConfig.Y){
		Y = NewY
	}
	X = collibrateLeft(borders, userConfig)
	NewX := collibrateRight(borders, userConfig)
	if NewX != float64(userConfig.X){
		X = NewX
	}
	return pixel.V(X, Y)

}

func (winConf *WindowConfig)SetCam(userConfig Users.User, borders Map.CamBorder){
	coords := collibrate(borders, userConfig)
	winConf.Cam.CamPos = pixel.V(coords.X, coords.Y)
	winConf.Cam.CamZoom = 1.0
}

func (winConf *WindowConfig)UpdateCam(){
	Cam := pixel.IM.Scaled(winConf.Cam.CamPos, winConf.Cam.CamZoom).Moved(winConf.Win.Bounds().Center().Sub(winConf.Cam.CamPos))
	winConf.Win.SetMatrix(Cam)
}

func (winConf *WindowConfig)UpdateBackground(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.StartMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig)LoadAvailableHeroImages(){
	winConf.Components.AvailableHeroImages = Utils.GetAvailableHeroImages()
}

func (winConf *WindowConfig)DrawErrorText(){
	if winConf.WindowError.LobbyDoesNotExist{
		end := time.Now()
		if end.Sub(winConf.WindowError.LobbyErrorStop).Seconds() <= 2.0{
			winConf.TextAreas.LobbyDoesNotExistError.Clear()
	 		winConf.TextAreas.LobbyDoesNotExistError.Write([]byte(winConf.WindowError.LobbyErrorText))
	  		winConf.TextAreas.LobbyDoesNotExistError.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.LobbyDoesNotExistError.Orig, 2.5))
		}
	}
}

func (winConf *WindowConfig)DrawAllTextAreas(){
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	winConf.TextAreas.WriteIDTextArea = text.New(pixel.V(220, 460), atlas)
	winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea = text.New(pixel.V(285, 332), atlas)
	winConf.TextAreas.NewMembersAnnouncement = text.New(pixel.V(240, 495), atlas)
	winConf.TextAreas.NewMembersTextArea = text.New(pixel.V(330, 410), atlas)
	winConf.TextAreas.NewMembersTextArea.Color = colornames.Black
	winConf.TextAreas.NewMembersTextArea.LineHeight = 18
	winConf.TextAreas.CurrentLobbyIDArea = text.New(pixel.V(670, 20), atlas)
	winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea = text.New(pixel.V(300, 373), atlas)
	winConf.TextAreas.JoinLobbyAnnouncement = text.New(pixel.V(270, 495), atlas)
	winConf.TextAreas.LobbyDoesNotExistError =  text.New(pixel.V(257, 250), atlas)
}


func (winConf *WindowConfig)DrawBackgroundImage(){
	//Draws background image 

	image, err := Utils.LoadImage("SysImages/BackgroundImage.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.StartMenuBG = sprite
}

func (winConf *WindowConfig)LoadCreationLobbyMenuBG(){
	image, err := Utils.LoadImage("SysImages/CreatLobbImage.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.CreatioLobbyMenuBG = sprite
}

func (winConf *WindowConfig)DrawCreationLobbyMenuBG(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.CreatioLobbyMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig)LoadWaitRoomMenuBG(){
	image, err := Utils.LoadImage("SysImages/WaitRoom.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.WaitRoomMenuBG = sprite
}

func (winConf *WindowConfig)DrawWaitRoomMenuBG(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.WaitRoomMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig)LoadJoinLobbyMenu(){
	image, err := Utils.LoadImage("SysImages/JoinLobbyMenu.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.JoinLobbyMenuBG = sprite
}

func (winConf *WindowConfig)DrawJoinLobbyMenuBG(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.JoinLobbyMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig)LoadWaitRoomJoinBG(){
	image, err := Utils.LoadImage("SysImages/WaitRoomJoin.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.WaitRoomJoinBG = sprite
}

func (winConf *WindowConfig)DrawWaitRoomJoinBG(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.WaitRoomJoinBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig)LoadGameBackground(){
	image, err := Utils.LoadImage("SysImages/Game.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(
		image, 
		pixel.R(
			image.Bounds().Min.X,
			image.Bounds().Min.Y,
			image.Bounds().Max.X,
			image.Bounds().Max.Y, 
		),
	)
	winConf.BGImages.Game = sprite
}

func (winConf *WindowConfig)DrawGameBackground(){
	winConf.Win.Clear(colornames.Black)
	winConf.BGImages.Game.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}
