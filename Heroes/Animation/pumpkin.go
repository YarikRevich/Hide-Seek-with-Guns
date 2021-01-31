package Animation

import (
	"Game/Heroes/Users"
	"Game/Window"

	"github.com/faiface/pixel"
)

type P struct{
	WinConf *Window.WindowConfig
	User    *Users.User
}

func (h *P) Frame()pixel.Rect{
	var frame pixel.Rect
	switch h.User.Animation.CurrentFrame {
	case 1:
		frame = getFrame(17, 50, .9, 0)
	case 2:
		frame = getFrame(17, 50, 1.9, 0)
	case 3:
		frame = getFrame(17, 50, 2.85, 0)
	case 4:
		frame = getFrame(17, 50, 4.7, 0)
	case 5:
		frame = getFrame(17, 50, 6.6, 0)
	}
	if h.User.Animation.CurrentFrame == 5{
		h.User.Animation.CurrentFrame = 0
	}else{
		h.User.Animation.CurrentFrame++
	}
	return frame
}

func (h *P) Move(){
	image := h.WinConf.Components.AvailableHeroImages[h.User.PersonalInfo.HeroPicture]
	h.User.Animation.HeroIconUpdation++

	if h.User.Animation.HeroIconUpdation == 5{
		h.User.Animation.HeroIconUpdation = 0
		h.User.Animation.CurrentFrameMatrix = convertRectToIntList(h.Frame())
		if !compareEqualBeetwenSlices(convertToStrList(h.User.Animation.CurrentFrameMatrix), convertRectToStrList(pixel.R(0, 0, 0, 0))){
			image.Set(image.Picture(), convertStringSliceToRect(h.User.Animation.CurrentFrameMatrix))
			image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(
				float64(h.User.Pos.X),
				float64(h.User.Pos.Y),
			)))
		}	
	}
	if !compareEqualBeetwenSlices(convertToStrList(h.User.Animation.CurrentFrameMatrix), convertRectToStrList(pixel.R(0, 0, 0, 0))){
		image.Set(image.Picture(), convertStringSliceToRect(h.User.Animation.CurrentFrameMatrix))
		image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(
			float64(h.User.Pos.X),
			float64(h.User.Pos.Y),
		)))
	}else{
		image.Set(image.Picture(), getFrame(17, 50, .9, 0))
		image.Draw(h.WinConf.Win, pixel.IM.Moved(pixel.V(
			float64(h.User.Pos.X),
			float64(h.User.Pos.Y),
		)))
	}
}

func NewPumpkinAnimator(winConf *Window.WindowConfig, userConfig *Users.User)IconfAnimator{
	return &P{
		WinConf:  winConf,
		User:     userConfig,
	}
}