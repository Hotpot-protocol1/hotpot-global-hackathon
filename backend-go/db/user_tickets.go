package db

import (
	"fmt"

	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/jmoiron/sqlx"
)

type UserTickets interface {
	GetUser(chain int, id int) (models.UserTickets, error)
	GetLatestRafflePotInfo(chain int) (models.PotWithRaffleTimestamp, error)
	GetLatestRafflePotInfoSeconds(chain int) (models.PotWithRaffleTimestamp, error)
	GetPotTicketLeaderboard(chain, potID int) ([]models.UserLeaderboard, error)
	Insert(chain int, user models.UserTickets) error
	SetWinnerForPot(chain int, potID uint16, ticketId uint32) error
	GetWinnersForPot(chain int, potID uint16) ([]models.Winner, error)
	SetPotRaffleTimestamp(chain int, potID uint16) error
	GetUserTicketsForPot(chain int, walletAddress string, potID uint16) (models.UserPotTickets, error)
	GetUserPotsWithRaffleTimestamp(chain int, walletAddress string) ([]models.PotWithRaffleTimestamp, error)
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

func (w *UserTicketsWrapper) GetLatestRafflePotInfo(chain int) (models.PotWithRaffleTimestamp, error) {
	var potWithRaffleTs models.PotWithRaffleTimestamp

	err := w.handler.QueryRow(`SELECT pot_id, raffle_timestamp FROM user_ticket WHERE chain = $1 AND raffle_timestamp > current_date - interval '30' day ORDER BY raffle_timestamp DESC LIMIT 1;`, chain).Scan(&potWithRaffleTs.PotId, potWithRaffleTs.RaffleTimestamp)
	return potWithRaffleTs, err
}

// TODO: For testing REMOVE
func (w *UserTicketsWrapper) GetLatestRafflePotInfoSeconds(chain int) (models.PotWithRaffleTimestamp, error) {
	var potInfo models.PotWithRaffleTimestamp

	err := w.handler.QueryRow(`SELECT pot_id, raffle_timestamp FROM user_ticket WHERE chain = $1 AND raffle_timestamp > now() - interval '30' second ORDER BY raffle_timestamp DESC LIMIT 1;`, chain).Scan(&potInfo.PotId, &potInfo.RaffleTimestamp)
	return potInfo, err
}

func (w *UserTicketsWrapper) GetPotTicketLeaderboard(chain, potID int) ([]models.UserLeaderboard, error) {
	sqlStatement := `SELECT wallet_address, pot_id, COUNT(*) as num_of_tickets FROM user_ticket WHERE chain = $1 AND pot_id = $2 GROUP BY pot_id, wallet_address`

	rows, err := w.handler.Query(sqlStatement, chain, potID)
	if err != nil {
		return []models.UserLeaderboard{}, err
	}

	defer rows.Close()

	leaderboard := []models.UserLeaderboard{}
	for rows.Next() {
		var entry models.UserLeaderboard

		err := rows.Scan(&entry.WalletAddress, &entry.PotID, &entry.NumOfTickets)
		if err != nil {
			return []models.UserLeaderboard{}, err
		}

		leaderboard = append(leaderboard, entry)
	}

	return leaderboard, rows.Err()
}

func (w *UserTicketsWrapper) Insert(chain int, user models.UserTickets) error {
	user.Chain = chain
	_, err := w.handler.NamedExec(`INSERT INTO user_ticket (wallet_address, ticket_id, pot_id, raffle_timestamp, chain)
	VALUES (:wallet_address, :ticket_id, :pot_id, :raffle_timestamp, :chain) ON CONFLICT (ticket_id, pot_id) DO NOTHING`, user)

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

func (w *UserTicketsWrapper) GetWinnersForPot(chain int, potID uint16) ([]models.Winner, error) {
	var pots []models.Winner
	err := w.handler.Select(&pots, `SELECT wallet_address, ticket_id FROM user_ticket WHERE chain = $1 AND pot_id = $2 and is_winner`, chain, potID)

	return pots, err
}

func (w *UserTicketsWrapper) SetPotRaffleTimestamp(chain int, potID uint16) error {
	r, err := w.handler.Exec(`UPDATE user_ticket SET raffle_timestamp = now() WHERE pot_id = $1 and chain = $2 and raffle_timestamp is null`, potID, chain)
	if err != nil {
		return err
	}

	if ra, err := r.RowsAffected(); ra == 0 || err != nil {
		return fmt.Errorf("nothing affected: Err %v", err)
	}

	return nil
}

func (w *UserTicketsWrapper) GetUserTicketsForPot(chain int, walletAddress string, potID uint16) (models.UserPotTickets, error) {
	sqlStatement := `SELECT ut.wallet_address, ut.pot_id, ut.ticket_id, ut.is_winner, COUNT(*) OVER() FROM user_ticket as ut WHERE chain = $1 AND LOWER(wallet_address) = LOWER($2) and pot_id = $3`

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

func (w *UserTicketsWrapper) GetUserPotsWithRaffleTimestamp(chain int, walletAddress string) ([]models.PotWithRaffleTimestamp, error) {
	var pots []models.PotWithRaffleTimestamp
	err := w.handler.Select(&pots, `SELECT pot_id, max(raffle_timestamp) as "raffle_timestamp" 
	FROM user_ticket WHERE chain = $1 AND lower(wallet_address) = lower($2) GROUP BY pot_id`, chain, walletAddress)

	return pots, err
}
