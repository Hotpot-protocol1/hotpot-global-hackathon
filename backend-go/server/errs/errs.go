package errs

import (
	"errors"
	"net/http"
)

type ErrResp struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func (e ErrResp) ToError() error {
	return errors.New(e.Message)
}

var (
	InternalServerErr = ErrResp{"INTERNAL_SERVER_ERROR", http.StatusInternalServerError}
	BadSignature      = ErrResp{"BAD_SIGNATURE", http.StatusForbidden}
	IncorrectBodyErr  = ErrResp{"INCORRECT_BODY", http.StatusBadRequest}
	IncorrectChainErr = ErrResp{"INCORRECT_CHAIN", http.StatusBadRequest}
	NoRaffle          = ErrResp{"NO_RAFFLE", http.StatusNotFound}
	NoOrder           = ErrResp{"NO_ORDER", http.StatusNotFound}
)
