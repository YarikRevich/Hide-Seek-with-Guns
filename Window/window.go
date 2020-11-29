package Window

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"Game/Utils"
)

type WindowConfig struct{
	win *pixelgl.Window
	bgsprite *pixel.Sprite
}

func CreateWindow()WindowConfig{

	cfg := pixelgl.WindowConfig{
		Title: "Hide and seek",
		Bounds: pixel.R(0, 0, 640, 480),
		Resizable: true,
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil{
		panic(err)
	}
	return WindowConfig{win: win}
}

func UpdateBackground(win *pixelgl.Window, sprite *pixel.Sprite){
	win.Clear(colornames.Black)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}


func DrawBackgroundImage(win *pixelgl.Window, winConf *WindowConfig)*pixel.Sprite{
	//Draws background image 

	image, err := Utils.LoadImage("test.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, image.Bounds())
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	winConf.bgsprite = sprite
}

