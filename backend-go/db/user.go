package db

import (
	"fmt"

	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUser(id int) (models.User, error)
	Insert(user models.User) error
	SetWinnerForPot(potID uint16, ticketId uint32) error
}

type UserWrapper struct {
	handler *sqlx.DB
}

func (d *DB) User() User {
	return &UserWrapper{
		handler: d.db,
	}
}

func (w UserWrapper) GetUser(id int) (models.User, error) {
	var user models.User

	err := w.handler.Get(&user, `SELECT * FROM hotpot_user WHERE id = $1`, id)
	return user, err
}

func (w UserWrapper) Insert(user models.User) error {
	_, err := w.handler.NamedExec(`INSERT INTO hotpot_user (wallet_address, ticket_id, pot_id)
	VALUES (:wallet_address, :ticket_id, :pot_id)`, user)

	return err
}

func (w UserWrapper) SetWinnerForPot(potID uint16, ticketId uint32) error {
	r, err := w.handler.Exec(`UPDATE hotpot_user SET is_winner = true WHERE pot_id = $1 AND ticket_id = $2`, potID, ticketId)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}
