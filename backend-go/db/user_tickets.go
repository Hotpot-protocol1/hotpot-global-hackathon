package db

import (
	"fmt"

	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/jmoiron/sqlx"
)

type UserTickets interface {
	GetUser(id int) (models.UserTickets, error)
	Insert(user models.UserTickets) error
	SetWinnerForPot(potID uint16, ticketId uint32) error
	SetPotRaffleTimestamp(potID uint16) error
	GetUserTicketsForPot(walletAddress string, potID uint16) (models.UserPotTickets, error)
	GetUserPotsWithRaffleTimestamp(walletAddress string) ([]models.PotsWithRaffleTimestamp, error)
}

type UserTicketsWrapper struct {
	handler *sqlx.DB
}

func (d *DB) UserTickets() UserTickets {
	return &UserTicketsWrapper{
		handler: d.db,
	}
}

func (w UserTicketsWrapper) GetUser(id int) (models.UserTickets, error) {
	var user models.UserTickets

	err := w.handler.Get(&user, `SELECT * FROM user_ticket WHERE id = $1`, id)
	return user, err
}

func (w UserTicketsWrapper) Insert(user models.UserTickets) error {
	_, err := w.handler.NamedExec(`INSERT INTO user_ticket (wallet_address, ticket_id, pot_id, raffle_timestamp)
	VALUES (:wallet_address, :ticket_id, :pot_id, :raffle_timestamp)`, user)

	return err
}

func (w UserTicketsWrapper) SetWinnerForPot(potID uint16, ticketId uint32) error {
	r, err := w.handler.Exec(`UPDATE user_ticket SET is_winner = true WHERE pot_id = $1 AND ticket_id = $2`, potID, ticketId)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}

func (w UserTicketsWrapper) SetPotRaffleTimestamp(potID uint16) error {
	r, err := w.handler.Exec(`UPDATE user_ticket SET raffle_timestamp = now() WHERE pot_id = $1`, potID)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}

func (w UserTicketsWrapper) GetUserTicketsForPot(walletAddress string, potID uint16) (models.UserPotTickets, error) {
	sqlStatement := `SELECT hu.wallet_address, hu.pot_id, hu.ticket_id, hu.is_winner, COUNT(*) OVER() FROM user_ticket as hu WHERE LOWER(wallet_address) = LOWER($1) and pot_id = $2`

	rows, err := w.handler.Query(sqlStatement, walletAddress, potID)
	if err != nil {
		return models.UserPotTickets{}, err
	}

	defer rows.Close()

	potTickets := models.UserPotTickets{}
	potTickets.Tickets = make([]models.UserPotTicket, 0)

	for rows.Next() {
		var potTicket models.UserPotTicket

		err := rows.Scan(&potTickets.WalletAddress, &potTickets.PotID, &potTicket.TicketID, &potTicket.IsWinner, &potTickets.NumOfTickets)
		if err != nil {
			return models.UserPotTickets{}, err
		}

		potTickets.Tickets = append(potTickets.Tickets, potTicket)
	}

	return potTickets, rows.Err()
}
func (w UserTicketsWrapper) GetUserPotsWithRaffleTimestamp(walletAddress string) ([]models.PotsWithRaffleTimestamp, error) {
	var pots []models.PotsWithRaffleTimestamp
	err := w.handler.Select(&pots, `SELECT pot_id, max(raffle_timestamp) as "raffle_timestamp" 
	FROM user_ticket WHERE lower(wallet_address) = lower($1) GROUP BY pot_id`, walletAddress)

	return pots, err

}
