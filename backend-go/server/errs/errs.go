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
	IncorrectBodyErr  = ErrResp{"INCORRECT_BODY", http.StatusBadRequest}
	IncorrectChainErr = ErrResp{"INCORRECT_CHAIN", http.StatusBadRequest}
)
