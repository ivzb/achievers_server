package model

import "net/http"

type Env struct {
	Request *http.Request
	DB      DBSourcer
	Log     Logger
	Token   Tokener
	UserID  string
}
