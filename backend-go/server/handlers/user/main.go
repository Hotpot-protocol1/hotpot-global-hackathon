package user

import (
	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log    *logrus.Entry
	userDB db.User
}

func New(dbHandler db.DBHandler, log *logrus.Entry) Handler {
	return Handler{
		log:    log,
		userDB: dbHandler.User(),
	}
}
