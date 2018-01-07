package model

type Env struct {
	Request Requester
	DB      DBSource
	UserId  string
	Token   Tokener
	Log     Logger
	Form    Former
}
