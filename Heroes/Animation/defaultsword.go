package Animation

import (
	"Game/Heroes/Users"
	"Game/Window"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type DS struct {
	WinConf *Window.WindowConfig
	User    *Users.User
}

func (h *DS) Frame() pixel.Rect {
	var frame pixel.Rect
	switch h.User.Animation.WeaponIconUpdation{
	case 0:
		frame = getFrame(18, 52, 1.7, 0)
	case 1:
		frame = getFrame(17, 50, 3, 0)
	case 2:
		frame = getFrame(22, 52, 3.7, 0)
	}
	return frame
}

func (h *DS) Move() {

	image := h.WinConf.Components.AvailableWeaponImages[h.User.GameInfo.WeaponName]

	if !(((h.WinConf.Win.Pressed(pixelgl.KeyW) ||
		h.WinConf.Win.Pressed(pixelgl.KeyA) ||
		h.WinConf.Win.Pressed(pixelgl.KeyS) ||
		h.WinConf.Win.Pressed(pixelgl.KeyD)) && h.WinConf.Win.JustPressed(pixelgl.KeySpace)) ||
		h.WinConf.Win.JustPressed(pixelgl.KeySpace)){
			image.Set(image.Picture(), h.Frame())
			image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(float64(h.User.Pos.X+8), float64(h.User.Pos.Y))))
			return
	}
	go func(){
		for{
			h.User.Animation.WeaponIconUpdation++

			if h.User.Animation.WeaponIconUpdation == 3 {
				h.User.Animation.WeaponIconUpdation = 0
				if !compareEqualBeetwenSlices(convertToStrList(convertRectToIntList(h.Frame())), convertRectToStrList(pixel.R(0, 0, 0, 0))){
					image.Set(image.Picture(), h.Frame())
					image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(float64(h.User.Pos.X+8), float64(h.User.Pos.Y))))
				}
				return
			}
			if !compareEqualBeetwenSlices(convertToStrList(convertRectToIntList(h.Frame())), convertRectToStrList(pixel.R(0, 0, 0, 0))){
				image.Set(image.Picture(), h.Frame())
				image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(float64(h.User.Pos.X+8), float64(h.User.Pos.Y))))
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func NewDefaultSwordAnimator(winConf *Window.WindowConfig, userConfig *Users.User) IconfAnimator {
	return &DS{
		WinConf: winConf,
		User:    userConfig,
	}
}
