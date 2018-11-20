package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type apiErr struct {
	Error apiErrDetails `json:"error"`

	rawErr         error
	httpStatusCode int
}

type apiErrDetails struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

const (
	jsonIndent = "    "
	errsPrefix = "TT-Core"
)

var apiErrors = map[int]apiErrDetails{
	101: {Message: "failed to start test"},
}

func newApiErr(code int, errIn error, errors ...string) *apiErr {
	e := apiErr{
		rawErr: errIn,
	}
	errDetails, ok := apiErrors[code]
	if !ok {
		errDetails = apiErrDetails{Code: "UNKNOWN"}
	}

	e.Error = errDetails
	if e.Error.Code == "" {
		e.Error.Code = fmt.Sprintf("%s-%04d", errsPrefix, code)
	}

	if e.httpStatusCode == 0 {
		e.httpStatusCode = http.StatusInternalServerError
	}
	e.Error.Errors = errors

	msgs := []string{}
	if errDetails.Message != "" {
		msgs = append(msgs, errDetails.Message)
	}
	if errIn != nil {
		msgs = append(msgs, errIn.Error())
	}
	e.Error.Message = strings.Join(msgs, "; ")

	return &e
}

func apiErrEmptyHashValue(err error) *apiErr {
	return newApiErr(107, err)
}

func apiErrNoRecordFound(err error) *apiErr {
	return newApiErr(107, err)
}

func (apierr *apiErr) respondOnErr(c echo.Context) error {
	return c.JSONPretty(apierr.httpStatusCode, apierr, jsonIndent)
}
