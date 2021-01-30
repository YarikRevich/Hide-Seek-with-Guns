package ConfigParsers

import (
	"Game/Heroes/Users"
	"Game/Window"
	//"fmt"
	"Game/Server"
//	"strconv"
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
		if value.PersonalInfo.Username == newUser.PersonalInfo.Username{
			return true
		}
	}
	return false
}

func IsCurrUser(newUser *Users.User, userConfig *Users.User)bool{
	if newUser.PersonalInfo.Username == userConfig.PersonalInfo.Username{
		return true
	}
	return false
}

//Parses gotten request to user config
type ConfigParser interface{
	Init(*Window.WindowConfig, *Users.User)
	//Parse(*Server.GameRequest)[]byte
	ApplyConfig(*Server.GameRequest)
	Unparse(*Server.GameRequest)*Users.User
	Commit(*Users.User)
}

type CP struct{
	w *Window.WindowConfig
	u *Users.User
}

func(c *CP) Init(w *Window.WindowConfig, u *Users.User){
	w.GameProcess.OtherUsers = []*Users.User{}
	c.w = w
	c.u = u
}


func(c *CP) Unparse(m *Server.GameRequest)*Users.User{
	return &Users.User{
		Pos:          (*Users.Pos)(&m.Pos),
		GameInfo:     (*Users.GameInfo)(&m.GameInfo),
		PersonalInfo: (*Users.PersonalInfo)(&m.PersonalInfo),
		Animation:    (*Users.Animation)(&m.Animation),
		Networking:   (*Users.Networking)(&m.Networking),
	}
}

func(c *CP) ApplyConfig(con *Server.GameRequest){
	c.u.GameInfo = (*Users.GameInfo)(&con.GameInfo)
}

func(c *CP) Commit(u *Users.User){
	if !IsCurrUser(u, c.u){
		c.w.GameProcess.OtherUsers = append(c.w.GameProcess.OtherUsers, u)
	}
}
// func UnparseCurrent(response string, userConfig *Users.User){
// 	splitedResp := strings.Split(response, "/::/")
// 	var chosenPart string
// 	for _, value := range splitedResp{
// 		if strings.Contains(value, userConfig.PersonalInfo.Username){
// 			chosenPart = value
// 		}
// 	}
// 	if strings.Contains(chosenPart, "~"){
// 		chosenPart = strings.Split(chosenPart, "~")[1]
// 	}
// 	textToReturn := strings.Split(chosenPart, "/")
// 	X, err := strconv.Atoi(textToReturn[1])
// 	if err != nil{
// 		panic(err)
// 	}
// 	Y, err := strconv.Atoi(textToReturn[2])
// 	if err != nil{
// 		panic(err)
// 	}
// 	currentFrame, err := strconv.Atoi(textToReturn[4])
// 	if err != nil{
// 		panic(err)
// 	}
// 	updationRun, err := strconv.Atoi(textToReturn[3])
// 	if err != nil{
// 		panic(err)
// 	}
// 	currentFrameMatrixSplited := strings.Split(textToReturn[5], "|")

// 	cFM1, err := strconv.ParseFloat(currentFrameMatrixSplited[0], 64)
// 	if err != nil{
// 		panic(err)
// 	}
// 	cFM2, err := strconv.ParseFloat(currentFrameMatrixSplited[1], 64)
// 	if err != nil{
// 		panic(err)
// 	}
// 	cFM3, err := strconv.ParseFloat(currentFrameMatrixSplited[2], 64)
// 	if err != nil{
// 		panic(err)
// 	}
// 	cFM4, err := strconv.ParseFloat(currentFrameMatrixSplited[3], 64)
// 	if err != nil{
// 		panic(err)
// 	}

// 	userConfig.Pos.X = X
// 	userConfig.Pos.Y = Y
// 	userConfig.PersonalInfo.HeroPicture = textToReturn[6]
// 	userConfig.Animation.CurrentFrame = currentFrame
// 	userConfig.Animation.UpdationRun = updationRun
// 	userConfig.Animation.CurrentFrameMatrix = []float64{cFM1, cFM2, cFM3, cFM4}
// }


// func UnparseOthers(response string, currentUser Users.User, winConf *Window.WindowConfig){
// 	splitedUsers := strings.Split(response, "/::/")
// 	for _, value := range splitedUsers{
// 		if !strings.Contains(value, currentUser.PersonalInfo.Username){
// 			splitedUserConf := strings.Split(value, "/")
// 			newUser := Users.User{
// 				PersonalInfo: &Users.PersonalInfo{Username: splitedUserConf[0]},
// 				Animation:    new(Users.Animation),
// 				GameInfo:     new(Users.GameInfo),
// 				Pos:          new(Users.Pos),
// 			}
// 			X, err := strconv.Atoi(splitedUserConf[1])
// 			if err != nil{
// 				panic(err)
// 			}
// 			Y, err := strconv.Atoi(splitedUserConf[2])
// 			if err != nil{
// 				panic(err)
// 			}
// 			currentFrame, err := strconv.Atoi(splitedUserConf[4])
// 			if err != nil{
// 				panic(err)
// 			}
// 			updationRun, err := strconv.Atoi(splitedUserConf[3])
// 			if err != nil{
// 				panic(err)
// 			}
// 			currentFrameMatrixSplited := strings.Split(splitedUserConf[5], "|")

// 			cFM1, err := strconv.ParseFloat(currentFrameMatrixSplited[0], 64)
// 			if err != nil{
// 				panic(err)
// 			}
// 			cFM2, err := strconv.ParseFloat(currentFrameMatrixSplited[1], 64)
// 			if err != nil{
// 				panic(err)
// 			}
// 			cFM3, err := strconv.ParseFloat(currentFrameMatrixSplited[2], 64)
// 			if err != nil{
// 				panic(err)
// 			}
// 			cFM4, err := strconv.ParseFloat(currentFrameMatrixSplited[3], 64)
// 			if err != nil{
// 				panic(err)
// 			}

// 			newUser.Pos.X = X
// 			newUser.Pos.Y = Y
// 			newUser.Animation.CurrentFrame = currentFrame
// 			newUser.PersonalInfo.HeroPicture = splitedUserConf[6]
// 			newUser.Animation.UpdationRun = updationRun
// 			newUser.Animation.CurrentFrameMatrix = []float64{cFM1, cFM2, cFM3, cFM4}
			
// 			if !IsAppended(&newUser, winConf){
// 				winConf.GameProcess.OtherUsers = append(winConf.GameProcess.OtherUsers, &newUser)
// 			}
// 		}
// 	}
// }

// func UnparseUsers(response string)[]string{
// 	cleanedResp := strings.Split(response, "~/")[1]
// 	splitedUsers := strings.Split(cleanedResp, "/::/")
// 	var result []string
// 	for _, value := range splitedUsers{
// 		splitedUserConf := strings.Split(value, "/")
// 		result = append(result, splitedUserConf[0])
// 	}
// 	return result
// }

// func ParseConfig(currUser *Users.User)string{
// 	return fmt.Sprintf("UpdateUser///%s~/%s/%d/%d/%d/%d/%s|%s|%s|%s/%s", 
// 			currUser.PersonalInfo.LobbyID, 
// 			currUser.PersonalInfo.Username,
// 			currUser.Pos.X, 
// 			currUser.Pos.Y, 
// 			currUser.Animation.UpdationRun,
// 			currUser.Animation.CurrentFrame, 
// 			currUser.Animation.CurrentFrameMatrix[0],
// 			currUser.Animation.CurrentFrameMatrix[1],
// 			currUser.Animation.CurrentFrameMatrix[2],
// 			currUser.Animation.CurrentFrameMatrix[3],
// 			currUser.PersonalInfo.HeroPicture,
// 		)
// }
