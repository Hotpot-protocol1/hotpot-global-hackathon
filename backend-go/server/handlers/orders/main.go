package orders

import (
	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/sirupsen/logrus"

	eventservice "github.com/Hotpot-protocol1/hotpot-global/services/contract"
)

type Handler struct {
	log           *logrus.Entry
	userTicketsDB db.UserTickets
	ordersDB      db.Orders
	infura        *eventservice.Infura
}

func New(dbHandler db.DBHandler, log *logrus.Entry, infura *eventservice.Infura) Handler {
	return Handler{
		log:           log,
		userTicketsDB: dbHandler.UserTickets(),
		ordersDB:      dbHandler.Orders(),
		infura:        infura,
	}
}
