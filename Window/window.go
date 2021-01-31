package Window

import (
	"Game/Heroes/Users"
	"Game/Utils"
	"time"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type MenuImages struct {
	StartMenuBG                      *pixel.Sprite
	CreatioLobbyMenuBG               *pixel.Sprite
	JoinLobbyMenuBG                  *pixel.Sprite
	WaitRoomMenuBG                   *pixel.Sprite
	WaitRoomJoinBG                   *pixel.Sprite
	Game                             *pixel.Sprite
	HorDoor                          *pixel.Sprite
	VerDoor                          *pixel.Sprite
	Darkness                         *pixel.Sprite
	GoldChest                        *pixel.Sprite
	CreationLobbyMenuBGPressedButton *pixel.Sprite
	WaitRoomPressedButton            *pixel.Sprite
	StartMenuPressedCreateButton     *pixel.Sprite
	StartMenuPressedJoinButton       *pixel.Sprite
	HPHeart                          *pixel.Sprite
	ElementsPanel                    *pixel.Sprite
}

type TextAreas struct {
	GameLogo               *text.Text
	WriteIDTextArea        *text.Text
	CreateLobbyInput       *CreateLobbyInput
	NewMembersAnnouncement *text.Text
	NewMembersTextArea     *text.Text
	CurrentLobbyIDArea     *text.Text
	JoinLobbyAnnouncement  *text.Text
	JoinLobbyInput         *JoinLobbyInput
	LobbyDoesNotExistError *text.Text
}

type JoinLobbyInput struct {
	InputLobbyIDTextArea *text.Text
	WrittenText          []string
}

type CreateLobbyInput struct {
	InputLobbyIDTextArea *text.Text
	WrittenText          []string
}

type WindowUpdation struct {
	StartMenuFrame     int
	CreationMenuFrame  int
	JoinLobbyMenuFrame int
	WaitRoomFrame      int
}

type WaitRoom struct {
	RoomType       string
	GettingUpdates bool
	NewMembers     []string
}

type StartMenu struct {
	DrawedTemporally []float64
	DrawCounter      int
	Regime           int
}

type GameProcess struct {
	OtherUsers []*Users.User
}

type Components struct {
	AvailableHeroImages map[string]*pixel.Sprite
	AvailableWeaponImages map[string]*pixel.Sprite
	AvailableWeaponIconImages map[string]*pixel.Sprite
	AvPlacesForSpaws    []pixel.Vec
}

type Cam struct {
	CamPos  pixel.Vec
	CamZoom float64
}

type WindowConfig struct {
	Win            *pixelgl.Window
	BGImages       *MenuImages
	TextAreas      *TextAreas
	WindowUpdation *WindowUpdation
	Components     *Components
	WindowError    *WindowError
	WaitRoom       *WaitRoom
	StartMenu      *StartMenu
	GameProcess    *GameProcess
	Cam            *Cam
}

type WindowError struct {
	LobbyDoesNotExist bool
	LobbyErrorStop    time.Time
	LobbyErrorText    string
}

func CreateWindow() WindowConfig {

	cfg := pixelgl.WindowConfig{
		Title:  "Hide and seek",
		Bounds: pixel.R(0, 0, 950, 550),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	textArea := TextAreas{
		CreateLobbyInput: new(CreateLobbyInput),
		JoinLobbyInput:   new(JoinLobbyInput),
	}
	return WindowConfig{
		Win:            win,
		BGImages:       new(MenuImages),
		TextAreas:      &textArea,
		WindowUpdation: new(WindowUpdation),
		WindowError:    new(WindowError),
		Components:     new(Components),
		WaitRoom:       new(WaitRoom),
		GameProcess:    new(GameProcess),
		StartMenu:      new(StartMenu),
		Cam:            new(Cam),
	}
}

func (winConf *WindowConfig) UpdateBackground() {
	winConf.Win.Clear(colornames.Black)
}

func (winConf *WindowConfig) DrawStartMenuBG() {
	winConf.BGImages.StartMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadAvailableHeroImages() {
	winConf.Components.AvailableHeroImages = Utils.GetAvailableHeroImages()
}

func (winConf *WindowConfig) LoadAvailableWeaponImages(){
	winConf.Components.AvailableWeaponImages = Utils.GetAvailableWeaponImages()
}

func (winConf *WindowConfig) LoadAvailableWeaponIconImages(){
	winConf.Components.AvailableWeaponIconImages = Utils.GetAvailableWeaponIconImages()
}

func (winConf *WindowConfig) DrawErrorText() {
	if winConf.WindowError.LobbyDoesNotExist {
		end := time.Now()
		if end.Sub(winConf.WindowError.LobbyErrorStop).Seconds() <= 2.0 {
			winConf.TextAreas.LobbyDoesNotExistError.Clear()
			winConf.TextAreas.LobbyDoesNotExistError.Write([]byte(winConf.WindowError.LobbyErrorText))
			winConf.TextAreas.LobbyDoesNotExistError.Draw(winConf.Win, pixel.IM.Scaled(winConf.TextAreas.LobbyDoesNotExistError.Orig, 2.5))
		}
	}
}

func (winConf *WindowConfig) LoadAllTextAreas() {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	winConf.TextAreas.GameLogo = text.New(pixel.V(730, 400), atlas)
	winConf.TextAreas.WriteIDTextArea = text.New(pixel.V(220, 460), atlas)
	winConf.TextAreas.CreateLobbyInput.InputLobbyIDTextArea = text.New(pixel.V(285, 332), atlas)
	winConf.TextAreas.NewMembersAnnouncement = text.New(pixel.V(240, 495), atlas)
	winConf.TextAreas.NewMembersTextArea = text.New(pixel.V(330, 410), atlas)
	winConf.TextAreas.NewMembersTextArea.Color = colornames.Black
	winConf.TextAreas.NewMembersTextArea.LineHeight = 18
	winConf.TextAreas.CurrentLobbyIDArea = text.New(pixel.V(670, 20), atlas)
	winConf.TextAreas.JoinLobbyInput.InputLobbyIDTextArea = text.New(pixel.V(300, 373), atlas)
	winConf.TextAreas.JoinLobbyAnnouncement = text.New(pixel.V(270, 495), atlas)
	winConf.TextAreas.LobbyDoesNotExistError = text.New(pixel.V(257, 250), atlas)
}

func (winConf *WindowConfig) DrawBackgroundImage() {
	//Draws background image

	image, err := Utils.LoadImage("SysImages/Menues/StartMenu.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.StartMenuBG = sprite
}

func (winConf *WindowConfig) LoadCreationLobbyMenuBG() {
	image, err := Utils.LoadImage("SysImages/Menues/CreateLobbyImage.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.CreatioLobbyMenuBG = sprite
}

func (winConf *WindowConfig) DrawCreationLobbyMenuBG() {
	winConf.BGImages.CreatioLobbyMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadWaitRoomMenuBG() {
	image, err := Utils.LoadImage("SysImages/Menues/WaitRoom.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.WaitRoomMenuBG = sprite
}

func (winConf *WindowConfig) DrawWaitRoomMenuBG() {
	winConf.BGImages.WaitRoomMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadJoinLobbyMenu() {
	image, err := Utils.LoadImage("SysImages/Menues/JoinLobbyMenu.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.JoinLobbyMenuBG = sprite
}

func (winConf *WindowConfig) DrawJoinLobbyMenuBG() {
	winConf.BGImages.JoinLobbyMenuBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadWaitRoomJoinBG() {
	image, err := Utils.LoadImage("SysImages/Menues/WaitRoomJoin.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.WaitRoomJoinBG = sprite
}

func (winConf *WindowConfig) DrawWaitRoomJoinBG() {
	winConf.BGImages.WaitRoomJoinBG.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadGameBackground() {
	image, err := Utils.LoadImage("SysImages/GameProcess/Game.png")
	if err != nil {
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

func (winConf *WindowConfig) DrawGameBackground() {
	winConf.BGImages.Game.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadHorDoor() {
	//Loads horizontal door

	image, err := Utils.LoadImage("SysImages/GameProcess/HorDoor.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.HorDoor = sprite
}

func (winConf *WindowConfig) LoadVerDoor() {
	//Loads vertical door

	image, err := Utils.LoadImage("SysImages/GameProcess/VerDoor.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.VerDoor = sprite
}

func (winConf WindowConfig) DrawHorDoor(coords pixel.Vec) {
	//Draws horizontal doors at passed coords

	winConf.BGImages.HorDoor.Draw(winConf.Win, pixel.IM.Moved(coords))
}

func (winConf WindowConfig) DrawVerDoor(coords pixel.Vec) {
	//Draws vertical doors at passed coords

	winConf.BGImages.VerDoor.Draw(winConf.Win, pixel.IM.Moved(coords))
}

func (winConf *WindowConfig) LoadDarkness() {
	image, err := Utils.LoadImage("SysImages/GameProcess/darkness.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.Darkness = sprite
}

func (winConf WindowConfig) DrawDarkness(coords pixel.Vec) {
	//Draws darkness on the background.

	winConf.BGImages.Darkness.Draw(winConf.Win, pixel.IM.Scaled(coords, 1.6).Moved(coords))
}

func (winConf *WindowConfig) LoadGoldChest() {
	image, err := Utils.LoadImage("SysImages/GameProcess/goldchest.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.GoldChest = sprite
}

func (winConf WindowConfig) DrawGoldChest() {
	for _, value := range winConf.Components.AvPlacesForSpaws {
		winConf.BGImages.GoldChest.Draw(winConf.Win, pixel.IM.Moved(value))
	}
}

func (winConf *WindowConfig) LoadCreationLobbyMenuBGPressedButton() {
	image, err := Utils.LoadImage("SysImages/Menues/CreateLobbyImagePressedButton.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.CreationLobbyMenuBGPressedButton = sprite
}

func (winConf *WindowConfig) DrawCreationLobbyMenuBGPressedButton() {
	winConf.BGImages.CreationLobbyMenuBGPressedButton.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadWaitRoomPressedButton() {
	image, err := Utils.LoadImage("SysImages/Menues/WaitRoomPressedButton.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.WaitRoomPressedButton = sprite
}

func (winConf *WindowConfig) DrawWaitRoomPressedButton() {
	winConf.BGImages.WaitRoomPressedButton.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadStartMenuPressedCreateButton() {
	image, err := Utils.LoadImage("SysImages/Menues/StartMenuPressedCreateButton.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.StartMenuPressedCreateButton = sprite
}

func (winConf *WindowConfig) DrawStartMenuPressedCreateButton() {
	winConf.BGImages.StartMenuPressedCreateButton.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadStartMenuPressedJoinButton() {
	image, err := Utils.LoadImage("SysImages/Menues/StartMenuPressedJoinButton.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.StartMenuPressedJoinButton = sprite
}

func (winConf *WindowConfig) DrawStartMenuPressedJoinButton() {
	winConf.BGImages.StartMenuPressedJoinButton.Draw(winConf.Win, pixel.IM.Moved(winConf.Win.Bounds().Center()))
}

func (winConf *WindowConfig) LoadHPHeart() {
	picture, err := Utils.LoadImage("SysImages/GameProcess/ElementsPanel/hp-heart.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(picture, picture.Bounds())
	winConf.BGImages.HPHeart = sprite
}

func (winConf *WindowConfig) DrawHPHeart(vector pixel.Vec) {
	cam := pixel.IM.Moved(winConf.Win.Bounds().Center().Sub(winConf.BGImages.StartMenuBG.Picture().Bounds().Center()))
	winConf.Win.SetMatrix(cam)
	winConf.BGImages.HPHeart.Draw(winConf.Win, pixel.IM.Moved(vector).Scaled(winConf.BGImages.HPHeart.Picture().Bounds().Center(), 0.4))
	cam = pixel.IM.Scaled(winConf.Cam.CamPos, winConf.Cam.CamZoom).Moved(winConf.Win.Bounds().Center().Sub(winConf.Cam.CamPos))
	winConf.Win.SetMatrix(cam)
}

func (winConf *WindowConfig) LoadElemtsPanel(){
	image, err := Utils.LoadImage("SysImages/GameProcess/ElementsPanel/ElementsPanel.png")
	if err != nil{
		log.Fatalln(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	winConf.BGImages.ElementsPanel = sprite
}

func (winConf *WindowConfig) DrawElementsPanel(){
	cam := pixel.IM.Moved(winConf.Win.Bounds().Center().Sub(winConf.BGImages.StartMenuBG.Picture().Bounds().Center()))
	winConf.Win.SetMatrix(cam)
	winConf.BGImages.ElementsPanel.Draw(winConf.Win, pixel.IM.Moved(pixel.V(475, 532.5)))
	cam = pixel.IM.Scaled(winConf.Cam.CamPos, winConf.Cam.CamZoom).Moved(winConf.Win.Bounds().Center().Sub(winConf.Cam.CamPos))
	winConf.Win.SetMatrix(cam)
}

func (winConf *WindowConfig) DrawWeaponIcon(wn string){
	cam := pixel.IM.Moved(winConf.Win.Bounds().Center().Sub(winConf.BGImages.StartMenuBG.Picture().Bounds().Center()))
	winConf.Win.SetMatrix(cam)
	winConf.Components.AvailableWeaponIconImages[wn].Draw(winConf.Win, pixel.IM.Moved(pixel.V(870, 525)))
	cam = pixel.IM.Scaled(winConf.Cam.CamPos, winConf.Cam.CamZoom).Moved(winConf.Win.Bounds().Center().Sub(winConf.Cam.CamPos))
	winConf.Win.SetMatrix(cam)
}

func (winConf WindowConfig) LoadAllImageComponents() {
	//It loads all the images for working.

	winConf.LoadCreationLobbyMenuBG()
	winConf.LoadJoinLobbyMenu()
	winConf.LoadWaitRoomMenuBG()
	winConf.LoadWaitRoomJoinBG()
	winConf.LoadGameBackground()
	winConf.LoadCreationLobbyMenuBGPressedButton()
	winConf.LoadWaitRoomPressedButton()
	winConf.LoadStartMenuPressedCreateButton()
	winConf.LoadStartMenuPressedJoinButton()
	winConf.LoadHorDoor()
	winConf.LoadVerDoor()
	winConf.LoadDarkness()
	winConf.LoadGoldChest()
	winConf.LoadHPHeart()
	winConf.LoadElemtsPanel()
}
