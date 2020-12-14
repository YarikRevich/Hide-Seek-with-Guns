package Window

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"Game/Utils"
	"time"
)


type MenuImages struct{
	StartMenuBG pixel.Picture
	CreatioLobbyMenuBG pixel.Picture
	JoinLobbyMenuBG pixel.Picture
	WaitRoomMenuBG pixel.Picture
	WaitRoomJoinBG pixel.Picture
	Game pixel.Picture
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

type WindowConfig struct{
	Win *pixelgl.Window
	BGImages *MenuImages
	TextAreas *TextAreas
	WindowUpdation *WindowUpdation
	Components *Components
	WindowError *WindowError
	WaitRoom *WaitRoom
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
	return WindowConfig{Win: win, BGImages: new(MenuImages), TextAreas: &textArea, WindowUpdation: new(WindowUpdation), WindowError: new(WindowError), Components: new(Components), WaitRoom: new(WaitRoom)}
}

func UpdateBackground(winConf *WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.StartMenuBG, winConf.BGImages.StartMenuBG.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func LoadAvailableHeroImages(winConf *WindowConfig){
	winConf.Components.AvailableHeroImages = Utils.GetAvailableHeroImages()
}

func DrawErrorText(winConf *WindowConfig){
	if winConf.WindowError.LobbyDoesNotExist{
		end := time.Now()
		if end.Sub(winConf.WindowError.LobbyErrorStop).Seconds() <= 2.0{
			winConf.TextAreas.LobbyDoesNotExistError.Clear()
	 		winConf.TextAreas.LobbyDoesNotExistError.Write([]byte(winConf.WindowError.LobbyErrorText))
	  		winConf.TextAreas.LobbyDoesNotExistError.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.LobbyDoesNotExistError.Orig, 2.5))
		}
	}
}

func DrawAllTextAreas(winConf *WindowConfig){
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


func DrawBackgroundImage(winConf *WindowConfig){
	//Draws background image 

	image, err := Utils.LoadImage("SysImages/BackgroundImage.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.StartMenuBG = image
}

func LoadCreationLobbyMenuBG(winConf *WindowConfig){
	image, err := Utils.LoadImage("SysImages/CreatLobbImage.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.CreatioLobbyMenuBG = image
}

func DrawCreationLobbyMenuBG(winConf WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.CreatioLobbyMenuBG, winConf.Win.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func LoadWaitRoomMenuBG(winConf *WindowConfig){
	image, err := Utils.LoadImage("SysImages/WaitRoom.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.WaitRoomMenuBG = image
}

func DrawWaitRoomMenuBG(winConf WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.WaitRoomMenuBG, winConf.Win.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func LoadJoinLobbyMenu(winConf *WindowConfig){
	image, err := Utils.LoadImage("SysImages/JoinLobbyMenu.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.JoinLobbyMenuBG = image
}

func DrawJoinLobbyMenuBG(winConf WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.JoinLobbyMenuBG, winConf.Win.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func LoadWaitRoomJoinBG(winConf *WindowConfig){
	image, err := Utils.LoadImage("SysImages/WaitRoomJoin.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.WaitRoomJoinBG = image

}

func DrawWaitRoomJoinBG(winConf WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.WaitRoomJoinBG, winConf.Win.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func LoadGameBackground(winConf *WindowConfig){
	image, err := Utils.LoadImage("SysImages/Game.png")
	if err != nil{
		panic(err)
	}
	winConf.BGImages.Game = image
}

func DrawGameBackground(winConf WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.Game, winConf.Win.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}
