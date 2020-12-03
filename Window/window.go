package Window

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"Game/Utils"
)


type MenuImages struct{
	StartMenuBG pixel.Picture
	CreatioLobbyMenuBG pixel.Picture
	JoinLobbyMenuBG pixel.Picture
	WaitRoomMenuBG pixel.Picture
	Game pixel.Picture
}

type TextAreas struct{
	WriteIDTextArea *text.Text
	InputLobbyIDTextArea *text.Text
	NewMembersAnnouncement *text.Text
	NewMembersTextArea *text.Text
	WrittenText []string
} 

type WindowUpdation struct{
	Frame int
}

type WaitRoom struct{
	GettingUpdates bool
	NewMembers []string
}

type WindowConfig struct{
	Win *pixelgl.Window
	BGImages *MenuImages
	TextAreas *TextAreas
	WindowUpdation *WindowUpdation
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
	return WindowConfig{Win: win, BGImages: new(MenuImages), TextAreas: new(TextAreas), WindowUpdation: new(WindowUpdation)}
}

func UpdateBackground(winConf *WindowConfig){
	winConf.Win.Clear(colornames.Black)
	sprite := pixel.NewSprite(winConf.BGImages.StartMenuBG, winConf.BGImages.StartMenuBG.Bounds())
	sprite.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func DrawAllTextAreas(winConf *WindowConfig){
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	winConf.TextAreas.WriteIDTextArea = text.New(pixel.V(220, 460), atlas)
	winConf.TextAreas.InputLobbyIDTextArea = text.New(pixel.V(285, 332), atlas)
	winConf.TextAreas.NewMembersAnnouncement = text.New(pixel.V(240, 495), atlas)
	winConf.TextAreas.NewMembersTextArea = text.New(pixel.V(330, 410), atlas)
	winConf.TextAreas.NewMembersTextArea.Color = colornames.Black
	winConf.TextAreas.NewMembersTextArea.LineHeight = 18
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
