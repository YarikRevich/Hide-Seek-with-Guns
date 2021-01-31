package Server

import (
	"Game/Heroes/Users"
	"encoding/json"
	"errors"
	"log"
)

//Parser for start connection to the main server
type StartParser interface{

	Parse(message []*StartRequest)[]byte
	Unparse(message []byte)[]*StartRequest
}

//Signature for request to the main server
type StartRequest struct{
	Index           int
	Error           string 
	Type            string
	Body            []string
}

func (m *StartRequest) Parse(message []*StartRequest)[]byte{
	b, err := json.Marshal(message)
	if err != nil{
		log.Fatalln(err)
	}
	return b
}

func (m *StartRequest) Unparse(message []byte)[]*StartRequest{
	var unparsed []*StartRequest
	err := json.Unmarshal(message, &unparsed)
	if err != nil{
		log.Fatalln(err)
	}
	return unparsed
}

//Configurator of startrequest
func NewStartRequest(type_ string)[]*StartRequest{
	return []*StartRequest{
		&StartRequest{
			Index:           0,
			Error:           "0",
			Type:            type_,
			Body:            []string{},
		},
	}
}


//Interface for game requests
type GameParser interface {
	Parse([]*GameRequest)[]byte
	Unparse([]byte)([]*GameRequest, error)
}

//Signature for game request
type GameRequest struct{
	Error           string
	Type            string
	Pos             struct{
		X int
		Y int
	}
	GameInfo        struct{
		Health int
		WeaponName string
		WeaponRadius int
	}
	PersonalInfo    struct{
		LobbyID string
		Username string
		HeroPicture string
	}
	Animation       struct{
		HeroIconUpdation int
		HeroIconUpdationDelay int
		WeaponIconUpdation int
		WeaponIconUpdationDelay int
		CurrentFrame int
		CurrentFrameMatrix []float64
	}
	Networking      struct{
		Index int
	}
	Context         struct{
		Additional      []string
	}
}

func(g *GameRequest) Parse(message []*GameRequest)[]byte{
	b, err := json.Marshal(message)
	if err != nil{
		log.Fatalln(err)
	}
	return b
}

func(g *GameRequest) Unparse(message []byte)([]*GameRequest, error){
	var unparsed []*GameRequest
	err := json.Unmarshal(message, &unparsed)
	if err != nil{
		var unparsedsingle *GameRequest
		err := json.Unmarshal(message, &unparsedsingle)
		if err != nil{
			return nil, errors.New("unmarshable array")
		}
		return []*GameRequest{unparsedsingle}, nil
	}
	return unparsed, nil
}

//Configurator for game requests
func NewGameRequest(t string, userconfig *Users.User)[]*GameRequest{
	return []*GameRequest{
		{
			Type: t,
			Pos: *userconfig.Pos,
			GameInfo: *userconfig.GameInfo,
			PersonalInfo: *userconfig.PersonalInfo,
			Animation: *userconfig.Animation,
			Networking: *userconfig.Networking,
			Context: *userconfig.Context,
		},
	}
}