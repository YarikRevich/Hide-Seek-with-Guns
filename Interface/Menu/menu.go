package Menu

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

func CreateStartMenu(winConf WindowConfig, userConfig *Users.User){
	image, err := Utils.LoadImage("button.png")
	if err != nil{
		panic(err)
	}
	sprite := pixel.NewSprite(image, winConf.win.Bounds())
	sprite.Draw(winConf.win, pixel.IM.Moved(winConf.win.Bounds().Center()))
	checkUserInterfaceActivity
}