package ConfigParsers

import (
	"Game/Heroes/Users"
	"Game/Window"
	"fmt"
	"strconv"
	"strings"
)

func IsUsersInfo(response string)bool{
	if strings.Contains(response, "GetUsersInfo"){
		return true
	}
	return false
}

func IsAppended(newUser *Users.User, winConf *Window.WindowConfig)bool{
	for _, value := range winConf.GameProcess.OtherUsers{
		if value.Username == newUser.Username{
			return true
		}
	}
	return false
}

func UnparseCurrent(response string, userConfig *Users.User){
	splitedResp := strings.Split(response, "/::/")
	var chosenPart string
	for _, value := range splitedResp{
		if strings.Contains(value, userConfig.Username){
			chosenPart = value
		}
	}
	if strings.Contains(chosenPart, "~"){
		chosenPart = strings.Split(chosenPart, "~")[1]
	}
	textToReturn := strings.Split(chosenPart, "/")
	X, err := strconv.Atoi(textToReturn[1])
	if err != nil{
		panic(err)
	}
	Y, err := strconv.Atoi(textToReturn[2])
	if err != nil{
		panic(err)
	}
	currentFrame, err := strconv.Atoi(textToReturn[4])
	if err != nil{
		panic(err)
	}
	updationRun, err := strconv.Atoi(textToReturn[3])
	if err != nil{
		panic(err)
	}
	currentFrameMatrixSplited := strings.Split(textToReturn[5], "|")
	currentFrameMatrix := []string{
		currentFrameMatrixSplited[0],
		currentFrameMatrixSplited[1],
		currentFrameMatrixSplited[2],
		currentFrameMatrixSplited[3],
	}

	userConfig.X = X
	userConfig.Y = Y
	userConfig.HeroPicture = textToReturn[6]
	userConfig.CurrentFrame = currentFrame
	userConfig.UpdationRun = updationRun
	userConfig.CurrentFrameMatrix = currentFrameMatrix
}


func UnparseOthers(response string, currentUser Users.User, winConf *Window.WindowConfig){
	splitedUsers := strings.Split(response, "/::/")
	for _, value := range splitedUsers{
		if !strings.Contains(value, currentUser.Username){
			splitedUserConf := strings.Split(value, "/")
			newUser := Users.User{Username: splitedUserConf[0]}
			X, err := strconv.Atoi(splitedUserConf[1])
			if err != nil{
				panic(err)
			}
			Y, err := strconv.Atoi(splitedUserConf[2])
			if err != nil{
				panic(err)
			}
			currentFrame, err := strconv.Atoi(splitedUserConf[4])
			if err != nil{
				panic(err)
			}
			updationRun, err := strconv.Atoi(splitedUserConf[3])
			if err != nil{
				panic(err)
			}
			currentFrameMatrixSplited := strings.Split(splitedUserConf[5], "|")
			currentFrameMatrix := []string{
				currentFrameMatrixSplited[0],
				currentFrameMatrixSplited[1],
				currentFrameMatrixSplited[2],
				currentFrameMatrixSplited[3],
			}

			newUser.X = X
			newUser.Y = Y
			newUser.CurrentFrame = currentFrame
			newUser.HeroPicture = splitedUserConf[6]
			newUser.UpdationRun = updationRun
			newUser.CurrentFrameMatrix = currentFrameMatrix
			
			if !IsAppended(&newUser, winConf){
				winConf.GameProcess.OtherUsers = append(winConf.GameProcess.OtherUsers, &newUser)
			}
		}
	}
}

func UnparseUsers(response string)[]string{
	cleanedResp := strings.Split(response, "~/")[1]
	splitedUsers := strings.Split(cleanedResp, "/::/")
	var result []string
	for _, value := range splitedUsers{
		splitedUserConf := strings.Split(value, "/")
		result = append(result, splitedUserConf[0])
	}
	return result
}

func ParseConfig(currUser *Users.User)string{
	return fmt.Sprintf("UpdateUser///%s~/%s/%d/%d/%d/%d/%s|%s|%s|%s/%s", 
			currUser.LobbyID, 
			currUser.Username,
			currUser.X, 
			currUser.Y, 
			currUser.UpdationRun,
			currUser.CurrentFrame, 
			currUser.CurrentFrameMatrix[0],
			currUser.CurrentFrameMatrix[1],
			currUser.CurrentFrameMatrix[2],
			currUser.CurrentFrameMatrix[3],
			currUser.HeroPicture,
		)
}
