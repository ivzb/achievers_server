package model

import (
	l "log"
)

type Logger interface {
	Log(string) error
}

type Log struct {
}

func NewLogger() *Log {
	return &Log{}
}

func (log *Log) Log(value string) error {
	l.Println(value)
	return nil
}
