package Animation

import (
	// "fmt"
	// "Game/Window"
	// "Game/Heroes/Users"
	"Game/Heroes/Users"
	"Game/Window"

	"github.com/faiface/pixel"
	// "github.com/faiface/pixel/pixelgl"
)

type IconfAnimator interface{
	Frame()pixel.Rect
	Move()
}

type Animation interface{
	Move()
}

type A struct{
	Animator IconfAnimator
}

func (a *A) Move(){
	a.Animator.Move()
}

func NewIconAnimator(winConf *Window.WindowConfig, userConfig *Users.User)Animation{
	switch userConfig.PersonalInfo.HeroPicture{
	case "pumpkinhero":
		return &A{
			Animator: NewPumpkinAnimator(winConf, userConfig),
		}
	}
	return nil
}

func NewWeaponAnimator(winConf *Window.WindowConfig, userConfig *Users.User)Animation{
	switch userConfig.GameInfo.WeaponName{
	case "defaultweapon":
		return &A{
			Animator: NewDefaultSwordAnimator(winConf, userConfig),
		}
	}
	return nil
}