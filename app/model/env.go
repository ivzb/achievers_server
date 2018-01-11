package model

import (
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/config"
)

type Env struct {
	Request *http.Request
	DB      DBSourcer
	Log     Logger
	Token   Tokener
	Config  *config.Config
	UserID  string
}
