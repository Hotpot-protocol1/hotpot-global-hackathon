package db

import (
	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/jmoiron/sqlx"
)

type Orders interface {
	Insert(chain int, user models.ChainOrder) error
	GetOrder(chain int, orderHash string) (models.OrderFlattened, error)
}

type OrdersWrapper struct {
	handler *sqlx.DB
}

func (d *DB) Orders() Orders {
	return &OrdersWrapper{
		handler: d.db,
	}
}

func (w *OrdersWrapper) Insert(chain int, order models.ChainOrder) error {
	order.Chain = chain
	_, err := w.handler.NamedExec(`INSERT INTO hotpot_order (order_hash, sig, chain, offerer, offer_token, offer_token_id, offer_amount, end_time, royalty_percent, royalty_recipient, salt)
	VALUES (lower(:order_hash), :sig, :chain, :offerer, :offer_item.offer_token, :offer_item.offer_token_id, :offer_item.offer_amount, :offer_item.end_time, :royalty.royalty_percent, :royalty.royalty_recipient, :salt)`, order)

	return err
}

func (w *OrdersWrapper) GetOrder(chain int, orderHash string) (models.OrderFlattened, error) {
	var order models.OrderFlattened

	err := w.handler.Get(&order, `SELECT offerer, offer_token, offer_token_id, offer_amount, end_time, royalty_percent, royalty_recipient, salt FROM hotpot_order WHERE lower(order_hash) = lower($1) and chain = $2`, orderHash, chain)
	return order, err
}
