package apierrors

import (
	"fmt"
	"net/http"
)

var (
	ErrDataBaseRequest     = fmt.Errorf("the request could not be created")
	ErrCreateRequestFailed = fmt.Errorf("failed to create HTTP request")
	ErrDoRequestFailed     = fmt.Errorf("failed to do HTTP request")
	ErrRequestFailed       = fmt.Errorf("HTTP request failed")
	ErrInvalidDataEntered  = fmt.Errorf("the data entered or the query form is not valid")
	ErrSerializedFields    = fmt.Errorf("there was a error when processing the fields")
	ErrPaginatorNotIsSetup = fmt.Errorf("paginator middleware is not setup")
	ErrPaginatorSize       = fmt.Errorf("paginator middleware is not setup")
	ErrTimeoutError        = fmt.Errorf("the request took too long to process")
	ErrConnectionClosed    = fmt.Errorf("the connection was closed before the request could be completed")
)

var (
	ErrResponsePaginatorNotIsSetup = NewResponseError(
		http.StatusInternalServerError,
		"There was a problem when processing the pagination",
		"PaginationMiddlewareNotSetupError",
		nil,
	)

	ErrResponseFailedProcessingQuery = NewResponseError(
		http.StatusInternalServerError,
		"The was a problem processing the query",
		"FailedProcessingQueryError",
		nil,
	)

	ErrResponseRequestNotProcessed = NewResponseError(
		http.StatusBadRequest,
		"The request could not be processed",
		"RequestNotProcessedError",
		nil,
	)

	ErrResponseRequestTimeOut = NewResponseError(
		http.StatusRequestTimeout,
		"The server took too much to respond. Try again later",
		"RequestTimeOutError",
		ErrTimeoutError,
	)

	ErrResponseDateFormatNotValid = NewResponseError(
		http.StatusBadRequest,
		"The date sent is not valid",
		"DateFormatNotValidError",
		nil,
	)

	ErrResponseDataBaseNotConnection = NewResponseError(
		http.StatusServiceUnavailable,
		"The server is not temporarily available. Try again later",
		"DataBaseNotConnectionError",
		ErrConnectionClosed,
	)

	ErrResponseRequestCancelled = NewResponseError(
		http.StatusRequestTimeout,
		"Request canceled",
		"RequestCancelledError",
		nil,
	)
)
