package models

import (
	"math/big"
)

type ChainOrder struct {
	Chain     int    `json:"chain" db:"chain"`
	OrderHash string `json:"order_hash" db:"order_hash"`
	Signature string `json:"sig" db:"sig"`
	Order
}

type OrderFlattened struct {
	Offerer          string `json:"offerer" db:"offerer"`         // signer address - he will receive ether after the token sell
	OfferToken       string `json:"offer_token" db:"offer_token"` // collection address
	OfferTokenId     int64  `json:"offer_token_id" db:"offer_token_id"`
	OfferAmount      int64  `json:"offer_amount" db:"offer_amount"` // the amount of ether for the offerer
	EndTime          int64  `json:"end_time" db:"end_time"`         // offer expiration timestamp
	RoyaltyPercent   int64  `json:"royalty_percent" db:"royalty_percent"`
	RoyaltyRecipient string `json:"royalty_recipient" db:"royalty_recipient"`
	Salt             int64  `json:"salt" db:"salt"` // random number for extra entropy
}

type Order struct {
	Offerer   string      `json:"offerer" db:"offerer"` // signer address - he will receive ether after the token sell
	OfferItem OfferItem   `json:"offer_item" db:"offer_item"`
	Royalty   RoyaltyData `json:"royalty" db:"royalty"`
	Salt      *big.Int    `json:"salt" db:"salt"` // random number for extra entropy
}

type OfferItem struct {
	OfferToken   string   `json:"offer_token" db:"offer_token"` // collection address
	OfferTokenId *big.Int `json:"offer_token_id" db:"offer_token_id"`
	OfferAmount  *big.Int `json:"offer_amount" db:"offer_amount"` // the amount of ether for the offerer
	EndTime      *big.Int `json:"end_time" db:"end_time"`         // offer expiration timestamp
}

type RoyaltyData struct {
	RoyaltyPercent   *big.Int `json:"royalty_percent" db:"royalty_percent"`
	RoyaltyRecipient string   `json:"royalty_recipient" db:"royalty_recipient"`
}
