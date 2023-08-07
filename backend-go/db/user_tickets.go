package db

import (
	"database/sql"
	"fmt"

	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/jmoiron/sqlx"
)

type UserTickets interface {
	GetUser(chain int, id int) (models.UserTickets, error)
	GetLatestRafflePotID(chain int) (sql.NullInt16, error)
	GetLatestRafflePotIDSeconds(chain int) (sql.NullInt16, error)
	Insert(chain int, user models.UserTickets) error
	SetWinnerForPot(chain int, potID uint16, ticketId uint32) error
	SetPotRaffleTimestamp(chain int, potID uint16) error
	GetUserTicketsForPot(chain int, walletAddress string, potID uint16) (models.UserPotTickets, error)
	GetUserPotsWithRaffleTimestamp(chain int, walletAddress string) ([]models.PotsWithRaffleTimestamp, error)
}

type UserTicketsWrapper struct {
	handler *sqlx.DB
}

func (d *DB) UserTickets() UserTickets {
	return &UserTicketsWrapper{
		handler: d.db,
	}
}

func (w *UserTicketsWrapper) GetUser(chain int, id int) (models.UserTickets, error) {
	var user models.UserTickets

	err := w.handler.Get(&user, `SELECT * FROM user_ticket WHERE id = $1 and chain = $2`, id, chain)
	return user, err
}

func (w *UserTicketsWrapper) GetLatestRafflePotID(chain int) (sql.NullInt16, error) {
	var potID sql.NullInt16

	err := w.handler.QueryRow(`SELECT pot_id FROM user_ticket WHERE chain = $1 AND raffle_timestamp > current_date - interval '30' day ORDER BY raffle_timestamp DESC LIMIT 1;`, chain).Scan(&potID)
	return potID, err
}

// TODO: For testing REMOVE
func (w *UserTicketsWrapper) GetLatestRafflePotIDSeconds(chain int) (sql.NullInt16, error) {
	var potID sql.NullInt16

	err := w.handler.QueryRow(`SELECT pot_id FROM user_ticket WHERE chain = $1 AND raffle_timestamp > now() - interval '30' second ORDER BY raffle_timestamp DESC LIMIT 1;`, chain).Scan(&potID)
	return potID, err
}

func (w *UserTicketsWrapper) Insert(chain int, user models.UserTickets) error {
	user.Chain = chain
	_, err := w.handler.NamedExec(`INSERT INTO user_ticket (wallet_address, ticket_id, pot_id, raffle_timestamp)
	VALUES (:wallet_address, :ticket_id, :pot_id, :raffle_timestamp)`, user)

	return err
}

func (w *UserTicketsWrapper) SetWinnerForPot(chain int, potID uint16, ticketId uint32) error {
	r, err := w.handler.Exec(`UPDATE user_ticket SET is_winner = true WHERE pot_id = $1 AND ticket_id = $2 and chain = $3`, potID, ticketId, chain)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}

func (w *UserTicketsWrapper) SetPotRaffleTimestamp(chain int, potID uint16) error {
	r, err := w.handler.Exec(`UPDATE user_ticket SET raffle_timestamp = now() WHERE pot_id = $1 and chain = $2`, potID, chain)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}

func (w *UserTicketsWrapper) GetUserTicketsForPot(chain int, walletAddress string, potID uint16) (models.UserPotTickets, error) {
	sqlStatement := `SELECT hu.wallet_address, hu.pot_id, hu.ticket_id, hu.is_winner, COUNT(*) OVER() FROM user_ticket as hu WHERE chain = $1 AND LOWER(wallet_address) = LOWER($2) and pot_id = $3`

	rows, err := w.handler.Query(sqlStatement, chain, walletAddress, potID)
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
func (w *UserTicketsWrapper) GetUserPotsWithRaffleTimestamp(chain int, walletAddress string) ([]models.PotsWithRaffleTimestamp, error) {
	var pots []models.PotsWithRaffleTimestamp
	err := w.handler.Select(&pots, `SELECT pot_id, max(raffle_timestamp) as "raffle_timestamp" 
	FROM user_ticket WHERE chain = $1 AND lower(wallet_address) = lower($2) GROUP BY pot_id`, chain, walletAddress)

	return pots, err

}
