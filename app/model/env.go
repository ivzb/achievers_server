package model

type Env struct {
	Request *Request
	DB      DBSourcer
	Log     Logger
	Token   Tokener
}
