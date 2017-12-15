package model

type Env struct {
	DB     DBSource
	UserId string
	Token  TokenSource
}
