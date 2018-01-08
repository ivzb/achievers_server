package model

type Env struct {
	Request *Request
	DB      DBSource
	Log     Logger
	Token   Tokener
}
