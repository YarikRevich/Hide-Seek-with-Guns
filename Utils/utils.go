package Utils

import (
	"fmt"
	"time"
	"math/rand"
	"image"
	"os"
	"os/exec"
	"strings"
	"github.com/faiface/pixel"
	_ "image/png"
)

func MessageIsEmpty(message []byte)bool{
	//Checks whether returned resp is empty

	for _, value := range message{
		if value != 0{
			return false
		}
	}
	return true
}

func CleanGottenResponse(resp []byte)string{
	//Cleanes passed just returned resp and returns cleaned version.

	var cleanedResponse string
	for _, value := range resp{
		if value != 0{
			cleanedResponse += string(value)
		}
	}
	return cleanedResponse
}

func IsOkResp(resp string)bool{
	if resp == "1"{
		return true
	}
	return false
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

func GetRandomHeroImage(availableHeroImages map[string]*pixel.Sprite)string{
	//Choses random hero image from the map of all the available hero images.

	var imageNames []string
	for key := range availableHeroImages{
		imageNames = append(imageNames, key)
	}
	return imageNames[GetRandNum(len(imageNames))]
}

func GetAvailableHeroImages()map[string]*pixel.Sprite{
	/* Saves to map all the available hero images in
	   current directory. Choses files with png extension
	*/

	listOfHeroes := make(map[string]*pixel.Sprite)
	CommInstanse := exec.Command("ls", "./SysImages")

	result, err := CommInstanse.Output()
	if err != nil{
		panic(err)
	}
	splitedResults := strings.Split(string(result), "\n")
	for _, value := range splitedResults{
		if len(value) > 0{
			if strings.HasSuffix(value, ".png") && strings.Contains(value, "hero"){
				fileName := strings.Split(value, ".")[0]
				image, err := LoadImage(fmt.Sprintf("./SysImages/%s", value))
				if err != nil{
					panic(err)
				}
				sprite := pixel.NewSprite(image, image.Bounds())
				listOfHeroes[fileName] = sprite
			}
		}
	}
	return listOfHeroes
}

func CheckErrorResp(resp []byte)bool{

	cleanedResp := CleanGottenResponse(resp)
	splitedOne := strings.Split(cleanedResp, "@")
	if len(splitedOne) > 1{
		if splitedOne[0] == "error"{
			return true
		}
	}
	return false
}

func CheckLobbyIsReady(resp []byte)bool{
	cleanedResp := CleanGottenResponse(resp)
	splitedOne := strings.Split(cleanedResp, "@")
	if len(splitedOne) > 1{
		if splitedOne[1] == "lobby is ready"{
			return true
		}
	}
	return false
}

func GetRandomSpawn()pixel.Vec{
	spawnPlaces := []pixel.Vec{
		pixel.V(-166, -182),
		pixel.V(1174, 748),
		pixel.V(1124, -182),
		pixel.V(-231, 768),
	}
	return spawnPlaces[GetRandNum(len(spawnPlaces))]
}

func RemoveIndex(s []string, index int)[]string{
	return append(s[:index], s[index+1:]...)
}

