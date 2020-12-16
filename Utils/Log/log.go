package Log

import (
	"os"
	"log"
)

type Logger interface{
	Setup()
	Update()
	Show()
}

type Log struct{
	userX int
	userY int
}


func (l *Log)Setup(){}

func (l *Log)Update(){}


func (l Log)Show(){}