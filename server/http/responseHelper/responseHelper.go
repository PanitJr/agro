package responseHelper

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ResponseHelper TODO
type Response struct {
	ResultCode        int         `json:"resultCode,omitempty"`
	ResultDescription string      `json:"resultDescription,omitempty"`
	ResultData        interface{} `json:"resultData,omitempty"`
	Error             error       `json:"error,omitempty"`
}

const (
	ResultCodeSuccess             = 20000
	ResultCodeCreated             = 20100
	ResultCodeBadRequset          = 40000
	ResultCodeUnauthorized        = 40100
	ResultCodeForbidden           = 40300
	ResultCodeNotFound            = 40401
	ResultCodeInternalServerError = 50000
	ResultCodeNotImplemented      = 50100
	ResultCodeServiceUnavailable  = 50300
	ResultCodeGatewayTimeout      = 50400
)

var (
	ErrNotFound            = errors.New("Not Found")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrBadRequset          = errors.New("Bad Requset")
	ErrForbidden           = errors.New("Forbidden")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotImplemented      = errors.New("Not Implemented")
	ErrServiceUnavailable  = errors.New("Service Unavailable")
	ErrGatewayTimeout      = errors.New("Gateway Timeout")
)

// ResultDescription string
var ResultDescription = map[int]string{
	ResultCodeSuccess:             "Success",
	ResultCodeCreated:             "Created Success",
	ResultCodeBadRequset:          "Bad Request",
	ResultCodeUnauthorized:        "Unauthorized",
	ResultCodeForbidden:           "Forbidden",
	ResultCodeNotFound:            "Not Found",
	ResultCodeInternalServerError: "Internal Server Error",
	ResultCodeNotImplemented:      "Not Implemented",
	ResultCodeServiceUnavailable:  "Service Unavailable",
	ResultCodeGatewayTimeout:      "Gateway Timeout",
}

func HTTPJSONResponse(result interface{}, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httpCode, resultCode := handlerError(err)
	w.WriteHeader(httpCode)
	resp := &Response{
		ResultCode:        resultCode,
		ResultDescription: ResultDescription[resultCode],
		ResultData:        result,
	}
	json.NewEncoder(w).Encode(resp)
	return
}

func handlerError(err error) (int, int) {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound, ResultCodeNotFound
	case ErrBadRequset:
		return http.StatusBadRequest, ResultCodeBadRequset
	case ErrForbidden:
		return http.StatusForbidden, ResultCodeForbidden
	case ErrUnauthorized:
		return http.StatusUnauthorized, ResultCodeUnauthorized
	case ErrNotImplemented:
		return http.StatusNotImplemented, ResultCodeNotImplemented
	case ErrServiceUnavailable:
		return http.StatusServiceUnavailable, ResultCodeServiceUnavailable
	case ErrGatewayTimeout:
		return http.StatusGatewayTimeout, ResultCodeGatewayTimeout
	case nil:
		return http.StatusOK, ResultCodeSuccess
	default:
		return http.StatusInternalServerError, ResultCodeInternalServerError
	}
}
