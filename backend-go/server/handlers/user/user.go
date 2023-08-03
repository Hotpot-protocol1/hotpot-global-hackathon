package user

import (
	"net/http"
	"strconv"

	"github.com/Hotpot-protocol1/hotpot-global/server/errs"
	"github.com/labstack/echo"
)

const paramID = "id"

func (h *Handler) GetUser(c echo.Context) error {
	userIDString := c.Param(paramID)
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.IncorrectBodyErr)
	}

	user, err := h.userDB.GetUser(userID)
	if err != nil {
		h.log.WithError(err).WithField("userID", userID).Error("Failed to get user")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}
	return c.JSON(http.StatusOK, user)
}
