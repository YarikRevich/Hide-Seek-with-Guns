package Utils

import (
	"github.com/galsondor/go-ascii"
	"time"
	"math/rand"
	"image"
	"os"
	"os/exec"
	"strings"
	"bytes"
	_ "image/png"
	"github.com/faiface/pixel"
	_ "github.com/faiface/pixel/pixelgl"
)

func MessageIsEmpty(message []byte)bool{
	//Checks whether returned resp is empty

	emptyMessage := make([]byte, 144)
	if bytes.Compare(message, emptyMessage) == 0{
		return true
	}
	return false
}

func SymbolIsAvailable(symbol string)bool{
	//Checks whether passed symbol is an available symbol.

	availableSymbols := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", ".", ":", "/", "@",
	}
	for _, value := range availableSymbols{
		if value == symbol{
			return true
		}
	}
	return false
}

func CleanGottenResponse(resp string)string{
	//Cleanes passed just returned resp and returns cleaned version.

	cleanedResponse := ""
	for _, value := range resp{
		if strings.Split(resp, "@")[0] != "error"{
			if SymbolIsAvailable(string(value)){
				cleanedResponse += string(value)
			}else{
				return cleanedResponse
			}
		}else{
			if ascii.IsASCII(byte(value)) || string(value) == "/"{
				cleanedResponse += string(value)
			}else{
				return cleanedResponse
			}
		}
	}
	return cleanedResponse
}

func LoadImage(path string)(pixel.Picture, error){
	//Loads image by passed path in arg.

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
	//Returns random num chosen by randomiser.

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GetRandomHeroImage(availableHeroImages map[string]pixel.Picture)string{
	//Choses random hero image from the map of all the available hero images.

	var imageNames []string
	for key, _ := range availableHeroImages{
		imageNames = append(imageNames, key)
	}
	return imageNames[GetRandNum(len(imageNames))]
}

func GetAvailableHeroImages()map[string]pixel.Picture{
	/* Saves to map all the available hero images in
	   current directory. Choses files with png extension
	*/

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