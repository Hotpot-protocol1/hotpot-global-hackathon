package orders

import "github.com/ethereum/go-ethereum/signer/core/apitypes"

var marketplaceOrderTypes = apitypes.Types{
	"EIP712Domain": {
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	},
	"Order": {
		{Name: "offerer", Type: "address"},
		{Name: "offerItem", Type: "OfferItem"},
		{Name: "royalty", Type: "RoyaltyData"},
		{Name: "salt", Type: "uint256"},
	},
	"OfferItem": {
		{Name: "offerToken", Type: "address"},
		{Name: "offerTokenId", Type: "uint256"},
		{Name: "offerAmount", Type: "uint256"},
		{Name: "endTime", Type: "uint256"},
	},
	"RoyaltyData": {
		{Name: "royaltyPercent", Type: "uint256"},
		{Name: "royaltyRecipient", Type: "address"},
	},
}

var marketplacePendingAmountsTypes = apitypes.Types{
	"EIP712Domain": {
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	},
	"PendingAmountData": {
		{Name: "offererPendingAmount", Type: "uint256"},
		{Name: "buyerPendingAmount", Type: "uint256"},
		{Name: "orderHash", Type: "bytes32"},
	},
}
