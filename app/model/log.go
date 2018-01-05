package model

import (
	l "log"
)

type Logger interface {
	Message(string)
	Error(error)
}

type Log struct {
}

func NewLogger() *Log {
	return &Log{}
}

func (log *Log) Message(message string) {
	l.Println(message)
}

func (log *Log) Error(err error) {
	l.Println(err.Error())
}
