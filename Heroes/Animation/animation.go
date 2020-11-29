package Animation

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"Game/Heroes/Users"
	"strconv"
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
	switch user.CurrentFrame {
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
	if user.CurrentFrame == 5{
		user.CurrentFrame = 0
	}else{
		user.CurrentFrame++
	}
	return frame
}

func ConvertToStringList(R pixel.Rect)[]string{
	MinX := fmt.Sprintf("%f", R.Min.X)
	MinY := fmt.Sprintf("%f", R.Min.Y)
	MaxX := fmt.Sprintf("%f", R.Max.X)
	MaxY := fmt.Sprintf("%f", R.Max.Y)
	return []string{
		MinX, MinY, MaxX, MaxY,
	}
}

func CompareEqualBeetwenSlices(a []string, b[]string)bool{
	for index, value := range a{
		if value != b[index]{
			return false
		}
	}
	return true
}

func ConvertStringSliceToRect(StringSlice []string)pixel.Rect{
	MinX, err := strconv.ParseFloat(StringSlice[0], 64)
	if err != nil{
		panic(err)
	}
	MinY, err := strconv.ParseFloat(StringSlice[1], 64)
	if err != nil{
		panic(err)
	}
	MaxX, err := strconv.ParseFloat(StringSlice[2], 64)
	if err != nil{
		panic(err)
	}
	MaxY, err := strconv.ParseFloat(StringSlice[3], 64)
	if err != nil{
		panic(err)
	}

	return pixel.R(
		MinX, MinY, MaxX, MaxY,
	)
}

// func ConvertCoordinatesToFloat(Coord string)float64{
// 	result, err := strconv.ParseFloat(Coord, 64)
// 	if err != nil{
// 		panic(err)
// 	}
// 	return result
// }

func ChangeAnimation(user *Users.User, image pixel.Picture, win *pixelgl.Window){
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	user.UpdationRun++
	if user.UpdationRun == 5{
		user.UpdationRun = 0
		user.CurrentFrameMatrix = ConvertToStringList(GetFrame(user))
		if !CompareEqualBeetwenSlices(user.CurrentFrameMatrix, ConvertToStringList(pixel.R(0, 0, 0, 0))){
			sprite.Set(image, ConvertStringSliceToRect(user.CurrentFrameMatrix))
			sprite.Draw(win, pixel.IM.Moved(pixel.V(
				float64(user.X),
				float64(user.Y),
			)))
		}	
	}
	if !CompareEqualBeetwenSlices(user.CurrentFrameMatrix, ConvertToStringList(pixel.R(0, 0, 0, 0))){
		sprite.Set(image, ConvertStringSliceToRect(user.CurrentFrameMatrix))
		sprite.Draw(win, pixel.IM.Moved(pixel.V(
			float64(user.X),
			float64(user.Y),
		)))
	}else{
		sprite.Set(image, getFrame(17, 50, .9, 0))
		sprite.Draw(win, pixel.IM.Moved(pixel.V(
			float64(user.X),
			float64(user.Y),
		)))
	}
}

func MoveAndChangeAnim(userConfig *Users.User, win *pixelgl.Window, availableHeroImages map[string]pixel.Picture){
	ChangeAnimation(userConfig, availableHeroImages[userConfig.HeroPicture], win)
}