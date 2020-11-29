package Utils

import (
	"time"
	"math/rand"
	"image"
	"os"
	"os/exec"
	"strings"
	_ "image/png"
	"github.com/faiface/pixel"
	_ "github.com/faiface/pixel/pixelgl"
)

func LoadImage(path string)(pixel.Picture, error){
	file, err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil{
		panic(err)
	}
	return pixel.PictureDataFromImage(img), nil
}

func GetRandNum(max int)int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GetRandomName(availableHeroImages map[string]pixel.Picture)string{
	var imageNames []string
	for key, _ := range availableHeroImages{
		imageNames = append(imageNames, key)
	}
	return imageNames[GetRandNum(len(imageNames))]

}

func GetAvailableHeroImages()map[string]pixel.Picture{
	listOfHeroes := make(map[string]pixel.Picture)
	CommInstanse := exec.Command("ls")
	result, err := CommInstanse.Output()
	if err != nil{
		panic(err)
	}
	splitedResults := strings.Split(string(result), "\n")
	for _, value := range splitedResults{
		if len(value) > 0{
			if strings.HasSuffix(value, ".png"){
				fileName := strings.Split(value, ".")[0]
				image, err := LoadImage(value)
				if err != nil{
					panic(err)
				}
				listOfHeroes[fileName] = image
			}
		}
	}
	return listOfHeroes
}