package logger

import (
	"log"
)

type Loggerer interface {
	Message(string)
	Error(error)
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Message(message string) {
	log.Println(message)
}

func (logger *Logger) Error(err error) {
	log.Println(err.Error())
}
