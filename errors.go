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
	102: {Message: "failed to parse student details for test"},
	103: {Message: "student details validation failed"},
	104: {Message: "adding student details failed"},
	105: {Message: "test complete"},
	106: {Message: "failed to get question"},
	107: {Message: "one of ref no, question id or ans is empty"},
	108: {Message: "failed to update answare"},
	109: {Message: "failed to get question"},
	110: {Message: "ref no is empty"},
	111: {Message: "get result failed"},
	112: {Message: "Answare is already updated for ref no"},
	113: {Message: "Failed to get all student results"},
	114: {Message: "no records"},
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

func apiErrGetTestDetailsFailed() *apiErr {
	return newApiErr(101, nil)
}

func apiErrStartTestFailed() *apiErr {
	return newApiErr(102, nil)
}

func apiErrStudDetailsValidationFail(err error) *apiErr {
	return newApiErr(103, err)
}

func apiErrStudDetailsInsertionFail(err error) *apiErr {
	return newApiErr(104, err)
}

func apiErrNoMoreQuestions() *apiErr {
	return newApiErr(105, nil)
}

func apiErrGetQuestionFailed() *apiErr {
	return newApiErr(106, nil)
}

func apiErrEmptyAnsParaValues() *apiErr {
	return newApiErr(107, nil)
}

func apiErrAnsUpdateFailed() *apiErr {
	return newApiErr(108, nil)
}

func apiErrGetSlNoFailed() *apiErr {
	return newApiErr(109, nil)
}

func apiErrRefNoEmpty() *apiErr {
	return newApiErr(110, nil)
}

func apiErrGetResultFailed() *apiErr {
	return newApiErr(111, nil)
}

func apiErrAnsAlreadyUpdated() *apiErr {
	return newApiErr(112, nil)
}

func apiErrGetAllResultsFailed() *apiErr {
	return newApiErr(113, nil)
}

func apiErrNoRecords() *apiErr {
	return newApiErr(114, nil)
}

// helpers

func (ae *apiErr) render(c echo.Context) error {
	return c.JSONPretty(ae.httpStatusCode, ae, jsonIndent)
}
