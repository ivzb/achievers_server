package model

type Env struct {
	DB     DBSource
	UserId string
	Token  Tokener
	Log    Logger
	Form   Former
}
