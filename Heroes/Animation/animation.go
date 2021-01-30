package Animation

import (
	"fmt"
	"Game/Window"
	"Game/Heroes/Users"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func getFrame(frameWidth float64, frameHeight float64, xGrid float64, yGrid int) pixel.Rect{
	return pixel.R(
		float64(xGrid)*frameWidth,
		float64(yGrid)*frameHeight,
		float64(xGrid + 1)*frameWidth,
		float64(yGrid + 1)*frameHeight,
	)
}

func GetFrame(user *Users.User)pixel.Rect{
	var frame pixel.Rect
	switch user.Animation.CurrentFrame {
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
	if user.Animation.CurrentFrame == 5{
		user.Animation.CurrentFrame = 0
	}else{
		user.Animation.CurrentFrame++
	}
	return frame
}

func convertRectToStrList(R pixel.Rect)[]string{
	MinX := fmt.Sprintf("%f", R.Min.X)
	MinY := fmt.Sprintf("%f", R.Min.Y)
	MaxX := fmt.Sprintf("%f", R.Max.X)
	MaxY := fmt.Sprintf("%f", R.Max.Y)
	return []string{
		MinX, MinY, MaxX, MaxY,
	}
}

func convertRectToIntList(R pixel.Rect)[]float64{
	return []float64{
		R.Min.X, R.Min.Y, R.Max.X, R.Max.Y,
	}
}

func convertToStrList(intlist []float64)[]string{
	var strlist []string
	for _, i := range intlist{
		num := fmt.Sprintf("%f", i)
		strlist = append(strlist, num)
	}
	return strlist
}

func CompareEqualBeetwenSlices(a []string, b[]string)bool{
	for index, value := range a{
		if value != b[index]{
			return false
		}
	}
	return true
}

func convertStringSliceToRect(floatlist []float64)pixel.Rect{
	// MinX, err := strconv.ParseFloat(StringSlice[0], 64)
	// if err != nil{
	// 	panic(err)
	// }
	// MinY, err := strconv.ParseFloat(StringSlice[1], 64)
	// if err != nil{
	// 	panic(err)
	// }
	// MaxX, err := strconv.ParseFloat(StringSlice[2], 64)
	// if err != nil{
	// 	panic(err)
	// }
	// MaxY, err := strconv.ParseFloat(StringSlice[3], 64)
	// if err != nil{
	// 	panic(err)
	// }
	return pixel.R(
		floatlist[0], floatlist[1], floatlist[2], floatlist[3],
	)
}

func ChangeAnimation(user *Users.User, image *pixel.Sprite, win *pixelgl.Window){
	user.Animation.UpdationRun++
	if user.Animation.UpdationRun == 5{
		user.Animation.UpdationRun = 0
		user.Animation.CurrentFrameMatrix = convertRectToIntList(GetFrame(user))
		if !CompareEqualBeetwenSlices(convertToStrList(user.Animation.CurrentFrameMatrix), convertRectToStrList(pixel.R(0, 0, 0, 0))){
			image.Set(image.Picture(), convertStringSliceToRect(user.Animation.CurrentFrameMatrix))
			image.Draw(win, pixel.IM.Moved(pixel.V(
				float64(user.Pos.X),
				float64(user.Pos.Y),
			)))
		}	
	}
	if !CompareEqualBeetwenSlices(convertToStrList(user.Animation.CurrentFrameMatrix), convertRectToStrList(pixel.R(0, 0, 0, 0))){
		image.Set(image.Picture(), convertStringSliceToRect(user.Animation.CurrentFrameMatrix))
		image.Draw(win, pixel.IM.Moved(pixel.V(
			float64(user.Pos.X),
			float64(user.Pos.Y),
		)))
	}else{
		image.Set(image.Picture(), getFrame(17, 50, .9, 0))
		image.Draw(win, pixel.IM.Moved(pixel.V(
			float64(user.Pos.X),
			float64(user.Pos.Y),
		)))
	}
}

func MoveAndChangeAnim(userConfig *Users.User, winConf *Window.WindowConfig){
	ChangeAnimation(userConfig, winConf.Components.AvailableHeroImages[userConfig.PersonalInfo.HeroPicture], winConf.Win)
}