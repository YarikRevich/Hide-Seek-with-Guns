package ConfigParsers

import (
	"Game/Heroes/Users"
	"Game/Window"
	"Game/Server"
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