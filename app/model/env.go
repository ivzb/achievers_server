package model

type Env struct {
	DB      DBSource
	UserId  string
	Tokener Tokener
	Logger  Logger
	Former  Former
}
