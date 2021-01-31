package Animation

import(
	"fmt"
	"github.com/faiface/pixel"
)

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

func compareEqualBeetwenSlices(a []string, b[]string)bool{
	for index, value := range a{
		if value != b[index]{
			return false
		}
	}
	return true
}

func convertStringSliceToRect(floatlist []float64)pixel.Rect{
	return pixel.R(
		floatlist[0], floatlist[1], floatlist[2], floatlist[3],
	)
}


func getFrame(frameWidth float64, frameHeight float64, xGrid float64, yGrid int) pixel.Rect{
	return pixel.R(
		float64(xGrid)*frameWidth,
		float64(yGrid)*frameHeight,
		float64(xGrid + 1)*frameWidth,
		float64(yGrid + 1)*frameHeight,
	)
}
