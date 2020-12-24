package Utils

import (
	"fmt"
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
)

func MessageIsEmpty(message []byte)bool{
	//Checks whether returned resp is empty

	emptyMessage := make([]byte, 144)
	if bytes.Compare(message, emptyMessage) == 0{
		return true
	}
	return false
}

func CleanGottenResponse(resp string)string{
	//Cleanes passed just returned resp and returns cleaned version.

	var cleanedResponse string
	for _, value := range resp{
		if ascii.IsPrint(byte(value)){
			cleanedResponse += string(value)
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

func CheckErrorResp(resp string)bool{

	cleanedResp := CleanGottenResponse(resp)
	splitedOne := strings.Split(cleanedResp, "@")
	if len(splitedOne) > 1{
		if splitedOne[0] == "error"{
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

