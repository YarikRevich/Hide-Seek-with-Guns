package ConfigParsers

import (
	"fmt"
	"strconv"
	"strings"
	"Game/Heroes/Users"
)

func UnparseCurrent(response string, userConfig *Users.User){
	splitedResp := strings.Split(response, "::/")
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
	X, err := strconv.Atoi(textToReturn[0])
	if err != nil{
		panic(err)
	}
	Y, err := strconv.Atoi(textToReturn[1])
	if err != nil{
		panic(err)
	}
	currentFrame, err := strconv.Atoi(textToReturn[2])
	if err != nil{
		panic(err)
	}
	updationRun, err := strconv.Atoi(textToReturn[3])
	if err != nil{
		panic(err)
	}
	currentFrameMatrixSplited := strings.Split(textToReturn[4], "|")
	currentFrameMatrix := []string{
		currentFrameMatrixSplited[0],
		currentFrameMatrixSplited[1],
		currentFrameMatrixSplited[2],
		currentFrameMatrixSplited[3],
	}

	userConfig.X = X
	userConfig.Y = Y
	userConfig.HeroPicture = textToReturn[5]
	userConfig.CurrentFrame = currentFrame
	userConfig.UpdationRun = updationRun
	userConfig.CurrentFrameMatrix = currentFrameMatrix
}


func UnparseOthers(response string, currentUser Users.User, otherUsers *[]*Users.User){
	splitedLobbyID := strings.Split(response, "~")
	splitedUsers := strings.Split(splitedLobbyID[1], "::/")
	for _, value := range splitedUsers{
		if !strings.Contains(value, currentUser.Username){
			splitedUserConf := strings.Split(value, "/")
			newUser := Users.User{Username: splitedUserConf[len(splitedUserConf)-1]}
			X, err := strconv.Atoi(splitedUserConf[0])
			if err != nil{
				panic(err)
			}
			Y, err := strconv.Atoi(splitedUserConf[1])
			if err != nil{
				panic(err)
			}
			currentFrame, err := strconv.Atoi(splitedUserConf[2])
			if err != nil{
				panic(err)
			}
			updationRun, err := strconv.Atoi(splitedUserConf[3])
			if err != nil{
				panic(err)
			}
			currentFrameMatrixSplited := strings.Split(splitedUserConf[4], "|")
			currentFrameMatrix := []string{
				currentFrameMatrixSplited[0],
				currentFrameMatrixSplited[1],
				currentFrameMatrixSplited[2],
				currentFrameMatrixSplited[3],
			}

			newUser.X = X
			newUser.Y = Y
			newUser.CurrentFrame = currentFrame
			newUser.HeroPicture = splitedUserConf[5]
			newUser.UpdationRun = updationRun
			newUser.CurrentFrameMatrix = currentFrameMatrix
			*otherUsers = append(*otherUsers, &newUser)
		}
	}
}

func ParseConfig(currUser *Users.User, otherUsers []*Users.User, response string)string{
	var result []string
	lobbyNum := strings.Split(response, "~")[0]
	result = append(result, fmt.Sprintf("UpdateUser///%s~%d/%d/%d/%d/%s|%s|%s|%s/%s/%s", 
			lobbyNum, 
			currUser.X, 
			currUser.Y, 
			currUser.CurrentFrame, 
			currUser.UpdationRun,
			currUser.CurrentFrameMatrix[0],
			currUser.CurrentFrameMatrix[1],
			currUser.CurrentFrameMatrix[2],
			currUser.CurrentFrameMatrix[3],
			currUser.HeroPicture,
			currUser.Username,
		))
	for _, value := range otherUsers{
		result = append(result, fmt.Sprintf("%d/%d/%d/%d/%s|%s|%s|%s/%s/%s",  
				value.X, 
				value.Y, 
				value.CurrentFrame, 
				value.UpdationRun,
				value.CurrentFrameMatrix[0],
				value.CurrentFrameMatrix[1],
				value.CurrentFrameMatrix[2],
				value.CurrentFrameMatrix[3],
				value.HeroPicture,
				value.Username,
		))
	}
	return strings.Join(result, "::/")
}
