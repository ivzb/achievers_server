package env

import (
	"net/http"

	"github.com/ivzb/achievers_server/app/db"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/logger"
	"github.com/ivzb/achievers_server/app/shared/token"
	"github.com/ivzb/achievers_server/app/shared/uuid"
)

type Env struct {
	Request *http.Request
	DB      db.DBSourcer
	Log     logger.Loggerer
	Token   token.Tokener
	Config  *config.Config
	UserID  string
	UUID    uuid.UUIDer
}
