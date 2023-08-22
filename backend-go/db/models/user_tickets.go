package models

import (
	"time"
)

type UserTickets struct {
	ID              string     `db:"id" json:"id"`
	WalletAddress   string     `db:"wallet_address" json:"wallet_address"`
	TicketID        uint32     `db:"ticket_id" json:"ticket_id"`
	PotID           uint16     `db:"pot_id" json:"pot_id"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
	IsWinner        bool       `db:"is_winner" json:"is_winner"`
	RaffleTimestamp *time.Time `db:"raffle_timestamp" json:"raffle_timestamp"`
	PendingAmount   string     `db:"pending_amount" json:"pending_amount"`
	Chain           int        `db:"chain" json:"chain"`
}

type UserLeaderboard struct {
	NumOfTickets  int    `db:"num_of_tickets" json:"num_of_tickets"`
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	PotID         uint16 `db:"pot_id" json:"pot_id"`
}

type UserPotTickets struct {
	NumOfTickets  int
	WalletAddress string          `db:"wallet_address" json:"wallet_address"`
	PotID         uint16          `db:"pot_id" json:"pot_id"`
	Tickets       []UserPotTicket `db:"tickets" json:"tickets"`
}

type UserPotTicket struct {
	TicketID      uint32 `db:"ticket_id" json:"ticket_id"`
	IsWinner      bool   `db:"is_winner" json:"is_winner"`
	PendingAmount string `db:"pending_amount" json:"pending_amount"`
}

type PotWithRaffleTimestamp struct {
	PotId           uint16     `db:"pot_id" json:"pot_id"`
	RaffleTimestamp *time.Time `db:"raffle_timestamp" json:"raffle_timestamp"`
}

type Winner struct {
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	TicketID      uint32 `db:"ticket_id" json:"ticket_id"`
}

type PendingAmount struct {
	PotId         uint16 `db:"pot_id" json:"pot_id"`
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	PendingAmount string `db:"pending_amount" json:"pending_amount"`
}
