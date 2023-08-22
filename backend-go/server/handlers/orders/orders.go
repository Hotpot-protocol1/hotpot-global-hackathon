package orders

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/Hotpot-protocol1/hotpot-global/server/errs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	eventservice "github.com/Hotpot-protocol1/hotpot-global/services/contract"
)

const queryParamChain = "chain"

func validateChain(chain string) (int, error) {
	switch chain {
	case "mainnet":
		return eventservice.ChainMainnet, nil
	case "sepolia":
		return eventservice.ChainSepolia, nil
	case "xdc":
		return eventservice.ChainXDC, nil
	case "goerli":
		return eventservice.ChainGoerli, nil
	}

	return 0, errors.New("chain doesn't exist")
}

func (h *Handler) CreateOrder(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	type createOrderReq struct {
		Order     models.Order `json:"order"`
		Signature string       `json:"signature"`
	}

	var req createOrderReq
	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("Failed to decode create order request")
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	if len(req.Signature) <= 2 {
		h.log.WithError(err).Error("Failed to decode create order request")
		return c.JSON(http.StatusBadRequest, errs.BadSignature)
	}

	orderHash, err := getOrderHash(h.infura, req.Order)
	if err != nil {
		h.log.WithError(err).Error("Failed to get order hash")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	signatureBytes, err := hex.DecodeString(req.Signature[2:])
	if err != nil {
		h.log.WithError(err).Error("Failed to decode signature")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if !verifySignature(h.log, signatureBytes, orderHash, req.Order.Offerer) {
		h.log.Error("Failed to verify signature")
		return c.JSON(http.StatusForbidden, errs.BadSignature)
	}

	dbModel := models.ChainOrder{OrderHash: "0x" + hex.EncodeToString(orderHash), Signature: req.Signature, Order: req.Order}
	err = h.ordersDB.Insert(chain, dbModel)
	if err != nil {
		h.log.WithError(err).WithField("chain order", dbModel).Error("Failed to insert order request")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}

func getOrderHash(infura *eventservice.Infura, orderr models.Order) ([]byte, error) {
	type OfferItem struct {
		OfferToken   common.Address `json:"offerToken"` // collection address
		OfferTokenId *big.Int       `json:"offerTokenId"`
		OfferAmount  *big.Int       `json:"offerAmount"` // the amount of ether for the offerer
		EndTime      *big.Int       `json:"endTime"`     // offer expiration timestamp
	}

	type RoyaltyData struct {
		RoyaltyPercent   *big.Int       `json:"royaltyPercent"`
		RoyaltyRecipient common.Address `json:"royaltyRecipient"`
	}

	type orders struct {
		Offerer   common.Address `json:"offerer"` // signer address - he will receive ether after the token sell
		OfferItem OfferItem      `json:"offerItem"`
		Royalty   RoyaltyData    `json:"royalty"`
		Salt      *big.Int       `json:"salt"` // random number for extra entropy
	}

	order := orders{}
	order.OfferItem.OfferToken = common.HexToAddress(orderr.OfferItem.OfferToken)
	order.OfferItem.OfferTokenId = orderr.OfferItem.OfferTokenId
	order.OfferItem.OfferAmount = orderr.OfferItem.OfferAmount
	order.OfferItem.EndTime = orderr.OfferItem.EndTime
	order.Royalty.RoyaltyPercent = orderr.Royalty.RoyaltyPercent
	order.Royalty.RoyaltyRecipient = common.HexToAddress(orderr.Royalty.RoyaltyRecipient)
	order.Offerer = common.HexToAddress(orderr.Offerer)
	order.Salt = orderr.Salt

	fmt.Println("Order: ", order)
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(inrec, &inInterface)

	fmt.Println("Interface: ", inInterface)
	message := apitypes.TypedDataMessage(inInterface)
	fmt.Println("Message: ", message)
	domainSep, err := infura.GetDomainSeparator()
	if err != nil {
		return nil, fmt.Errorf("problem getting domain separator: %v", err)
	}

	name, version, chainId, verifyingContract, err := infura.GetDomain()
	if err != nil {
		return nil, fmt.Errorf("problem getting domain: %v", err)
	}

	fmt.Println("CHAIN ID IS: ", chainId)

	domain := apitypes.TypedDataDomain{
		Name:              name,
		Version:           version,
		ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}

	fmt.Println("Domain: ", domain)

	var typedData = apitypes.TypedData{
		Types:       marketplaceOrderTypes,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, fmt.Errorf("problem hashing message struct: %v", err)
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("problem hashing domain struct: %v", err)
	}
	fmt.Println("Is domain separator equal: ", domainSeparator.String(), " 2: ", hex.EncodeToString(domainSep), " is true: ", domainSeparator.String() == "0x"+hex.EncodeToString(domainSep))

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	fmt.Println("RAWWWW ", fmt.Sprintf("\x19\x01%s%s", "0x"+hex.EncodeToString(domainSep), typedDataHash.String()))
	hashBytes := crypto.Keccak256Hash(rawData)
	fmt.Println("HASHHHHHKECCCAAAAKK ", hashBytes.String())

	return hashBytes.Bytes(), nil
}

func verifySignature(log *logrus.Entry, signature, orderHash []byte, offerer string) bool {
	if len(signature) != 65 {
		log.Errorf("invalid signature length: %d", len(signature))
		return false
	}

	if signature[64] != 27 && signature[64] != 28 {
		log.Errorf("invalid recovery id: %d", signature[64])
		return false
	}
	signature[64] -= 27

	recovered, err := crypto.SigToPub(orderHash, signature)
	if err != nil {
		log.WithError(err).Error("invalid pub")
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	offeredAddr := common.HexToAddress(offerer)
	if !bytes.Equal(offeredAddr.Bytes(), recoveredAddr.Bytes()) {
		log.Error("addresses do not match offerer: ", offerer, " recoverer: ", recoveredAddr)
		return false
	}

	return true
}

func getPendingAmountsHash(infura *eventservice.Infura, buyerPendingAmount, sellerPendingAmount *big.Int, orderHash string) ([]byte, error) {
	orderHashValue := [32]byte{}
	copy(orderHashValue[:], []byte(orderHash[2:]))
	inInterface := map[string]interface{}{"offererPendingAmount": sellerPendingAmount, "buyerPendingAmount": buyerPendingAmount, "orderHash": orderHashValue}
	message := apitypes.TypedDataMessage(inInterface)
	fmt.Println("Message: ", message)

	name, version, chainId, verifyingContract, err := infura.GetDomain()
	if err != nil {
		return nil, fmt.Errorf("problem getting domain: %v", err)
	}

	domain := apitypes.TypedDataDomain{
		Name:              name,
		Version:           version,
		ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}

	fmt.Println("Domain: ", domain)

	var typedData = apitypes.TypedData{
		Types:       marketplacePendingAmountsTypes,
		PrimaryType: "PendingAmountData",
		Domain:      domain,
		Message:     message,
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, fmt.Errorf("problem hashing message struct: %v", err)
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("problem hashing domain struct: %v", err)
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hashBytes := crypto.Keccak256Hash(rawData)

	return hashBytes.Bytes(), nil
}

func (h *Handler) FulFillOrder(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	type fulfillOrderReq struct {
		OrderHash string `json:"order_hash"`
		Fulfiller string `json:"fulfiller"`
	}

	type pendingAmountResp struct {
		PotId         uint16   `db:"pot_id" json:"pot_id"`
		WalletAddress string   `db:"wallet_address" json:"wallet_address"`
		PendingAmount *big.Int `db:"pending_amount" json:"pending_amount"`
	}

	type fulfillOrderResp struct {
		OrderHash           string            `json:"order_hash"`
		PendingAmountHash   string            `json:"pending_amount_hash"`
		BuyerPendingAmount  pendingAmountResp `json:"buyer_pending_amount"`
		SellerPendingAmount pendingAmountResp `json:"seller_pending_amount"`
		models.Order
	}

	var req fulfillOrderReq
	var resp fulfillOrderResp
	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("Failed to decode fulfill order request")
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	order, err := h.ordersDB.GetOrder(chain, req.OrderHash)
	if err != nil {
		if err == sql.ErrNoRows {
			h.log.WithError(err).WithField("order hash", req.OrderHash).Error("No order with given hash")
			return c.JSON(http.StatusNotFound, errs.NoOrder)
		}

		h.log.WithError(err).WithField("order hash", req.OrderHash).Error("Failed to get order")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	potID, err := h.infura.GetCurrentPot()
	if err != nil {
		h.log.WithError(err).WithFields(logrus.Fields{"potID": potID, "wallet": req.Fulfiller}).Error("Failed to get current pot id")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	buyerPendingAmount, err := h.userTicketsDB.GetUserPendingAmount(chain, req.Fulfiller, potID)
	if err != nil {
		if err != sql.ErrNoRows {
			h.log.WithError(err).WithField("order hash", req.OrderHash).Error("Failed to get order")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		buyerPendingAmount.PendingAmount = "0"
		buyerPendingAmount.PotId = potID
		buyerPendingAmount.WalletAddress = req.Fulfiller
	}

	sellerPendingAmount, err := h.userTicketsDB.GetUserPendingAmount(chain, order.Offerer, potID)
	if err != nil {
		if err != sql.ErrNoRows {
			h.log.WithError(err).WithField("order hash", req.OrderHash).Error("Failed to get order")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		sellerPendingAmount.PendingAmount = "0"
		sellerPendingAmount.PotId = potID
		sellerPendingAmount.WalletAddress = order.Offerer
	}

	resp.Order = models.Order{
		Offerer: order.Offerer,
		Salt:    big.NewInt(order.Salt),
		OfferItem: models.OfferItem{
			OfferToken:   order.OfferToken,
			OfferTokenId: big.NewInt(order.OfferTokenId),
			OfferAmount:  big.NewInt(order.OfferAmount),
			EndTime:      big.NewInt(order.EndTime),
		},
		Royalty: models.RoyaltyData{
			RoyaltyPercent:   big.NewInt(order.RoyaltyPercent),
			RoyaltyRecipient: order.RoyaltyRecipient,
		},
	}
	resp.OrderHash = req.OrderHash
	buyerPendAmount, ok := new(big.Int).SetString(buyerPendingAmount.PendingAmount, 10)
	if !ok {
		h.log.WithField("buyer string", buyerPendingAmount.PendingAmount).Error("Failed to get pending amount for buyer")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}
	sellerPendAmount, ok := new(big.Int).SetString(sellerPendingAmount.PendingAmount, 10)
	if !ok {
		h.log.WithField("seller string", sellerPendingAmount.PendingAmount).Error("Failed to get pending amount for seller")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp.BuyerPendingAmount = pendingAmountResp{
		PotId:         buyerPendingAmount.PotId,
		WalletAddress: buyerPendingAmount.WalletAddress,
		PendingAmount: buyerPendAmount,
	}
	resp.SellerPendingAmount = pendingAmountResp{
		PotId:         sellerPendingAmount.PotId,
		WalletAddress: sellerPendingAmount.WalletAddress,
		PendingAmount: sellerPendAmount,
	}
	pendingAmountsHash, err := getPendingAmountsHash(h.infura, buyerPendAmount, sellerPendAmount, req.OrderHash)
	if err != nil {
		h.log.WithError(err).WithField("order hash", req.OrderHash).Error("Failed to get pending amount hash")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp.PendingAmountHash = hex.EncodeToString(pendingAmountsHash)

	return c.JSON(http.StatusOK, resp)
}
