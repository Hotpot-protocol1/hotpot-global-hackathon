package tickets

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/Hotpot-protocol1/hotpot-global/server/errs"
	eventservice "github.com/Hotpot-protocol1/hotpot-global/services/event"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const paramPotID = "pot_id"
const paramWalletAddress = "wallet_address"
const queryParamChain = "chain"

func validateChain(chain string) (int, error) {
	switch chain {
	case "mainnet":
		return eventservice.ChainMainnet, nil
	case "sepolia":
		return eventservice.ChainSepolia, nil
	case "xdc":
		return eventservice.ChainXDC, nil
	}

	return 0, errors.New("chain doesn't exist")
}

func (h *Handler) GetUserTicketsForPot(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	walletAddr := c.Param(paramWalletAddress)
	potIDString := c.Param(paramPotID)
	potID, err := strconv.Atoi(potIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	userTickets, err := h.userDB.GetUserTicketsForPot(chain, walletAddr, uint16(potID))
	if err != nil {
		h.log.WithError(err).WithFields(logrus.Fields{"potID": potID, "wallet": walletAddr}).Error("Failed to get user tickets for pot")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, userTickets)
}

func (h *Handler) GetUserTicketsForCurrentPot(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	walletAddr := c.Param(paramWalletAddress)
	potID, err := h.infura.GetCurrentPot()
	if err != nil {
		h.log.WithError(err).WithFields(logrus.Fields{"potID": potID, "wallet": walletAddr}).Error("Failed to get current pot id")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	userTickets, err := h.userDB.GetUserTicketsForPot(chain, walletAddr, uint16(potID))
	if err != nil {
		h.log.WithError(err).WithFields(logrus.Fields{"potID": potID, "wallet": walletAddr}).Error("Failed to get user tickets for pot")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, userTickets)
}

func (h *Handler) GetPotsWithRaffleTimestamp(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}
	walletAddr := c.Param(paramWalletAddress)

	pots, err := h.userDB.GetUserPotsWithRaffleTimestamp(chain, walletAddr)
	if err != nil {
		h.log.WithError(err).Error("Failed to get pots with raffle timestamp")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, pots)
}

func (h *Handler) GetLatestRafflePotID(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	potID, err := h.userDB.GetLatestRafflePotID(chain)
	if err != nil {
		h.log.WithError(err).Error("Failed to get pots with raffle timestamp")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if !potID.Valid {
		return c.JSON(http.StatusOK, map[string]interface{}{"pot_id": nil})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"pot_id": potID.Int16})
}

// TODO: REMOVE test only
func (h *Handler) GetLatestRafflePotIDSeconds(c echo.Context) error {
	chain, err := validateChain(c.QueryParam(queryParamChain))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectChainErr)
	}

	potID, err := h.userDB.GetLatestRafflePotIDSeconds(chain)
	if err != nil && err != sql.ErrNoRows {
		h.log.WithError(err).Error("Failed to get pots with raffle timestamp")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if !potID.Valid {
		return c.JSON(http.StatusOK, map[string]interface{}{"pot_id": nil})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"pot_id": potID.Int16})
}
