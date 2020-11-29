package Icons

import (
	_ "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"Game/Utils"
)

func GetHeroImage() pixel.Picture{
	heroimage, err := Utils.LoadImage("testhero.png")
	if err != nil{
		panic(err)
	}
	return heroimage
}