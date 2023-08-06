package tickets

import (
	"net/http"
	"strconv"

	"github.com/Hotpot-protocol1/hotpot-global/server/errs"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const paramPotID = "pot_id"
const paramWalletAddress = "wallet_address"

func (h *Handler) GetUserTicketsForPot(c echo.Context) error {
	walletAddr := c.Param(paramWalletAddress)
	potIDString := c.Param(paramPotID)
	potID, err := strconv.Atoi(potIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	userTickets, err := h.userDB.GetUserTicketsForPot(walletAddr, uint16(potID))
	if err != nil {
		h.log.WithError(err).WithFields(logrus.Fields{"potID": potID, "wallet": walletAddr}).Error("Failed to get user tickets for pot")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, userTickets)
}

func (h *Handler) GetPotsWithRaffleTimestamp(c echo.Context) error {
	walletAddr := c.Param(paramWalletAddress)
	pots, err := h.userDB.GetUserPotsWithRaffleTimestamp(walletAddr)
	if err != nil {
		h.log.WithError(err).Error("Failed to get pots with raffle timestamp")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, pots)
}
