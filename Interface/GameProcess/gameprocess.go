package GameProcess



func KeyBoardButtonListener(userConfig *Users.User, win *pixelgl.Window){
	if (win.MousePosition().X > 18 && win.MousePosition().X < 141) && (win.MousePosition().Y > 74 && win.MousePosition().Y < 102) && win.Pressed(pixelgl.MouseButtonLeft){}
	if win.Pressed(pixelgl.KeyW){
		userConfig.Y += 7
	}else if win.Pressed(pixelgl.KeyA){
		userConfig.X -= 7
	}else if win.Pressed(pixelgl.KeyS){
		userConfig.Y -= 7
	}else if win.Pressed(pixelgl.KeyD){
		userConfig.X += 7
	} 
}


func ReDraw(otherUsers *[]*Users.User, win *pixelgl.Window, availableHeroImages map[string]pixel.Picture){
	for _, value := range *otherUsers{
		Animation.MoveAndChangeAnim(value, win, availableHeroImages)
	}
}

func ChangePos(userConfig *Users.User, win *pixelgl.Window, availableHeroImages map[string]pixel.Picture){
	KeyBoardButtonListener(userConfig, win)
	Animation.MoveAndChangeAnim(userConfig, win, availableHeroImages)
}


func RunMainPipeLine(readWriteChan chan string, userConfig *Users.User, winConf *WindowConfig, availableHeroImages map[string]pixel.Picture){
	otherUsers := []*Users.User{}
	
	select{
	case response := <- readWriteChan:
		UpdateBackground(winConf.win, winConf.bgsprite)
		UnparseCurrent(response, userConfig)
		ChangePos(userConfig, winConf.win, availableHeroImages)
		UnparseOthers(response, *userConfig, &otherUsers)
		ReDraw(&otherUsers, winConf.win, availableHeroImages)
		parsedMessage := ParseConfig(userConfig, otherUsers, response)
		userConfig.Conn.Write([]byte(parsedMessage))
	}
}