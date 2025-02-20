package apierrors

import (
	"errors"
	"net/http"
)

var (
	ErrDataBaseRequest     = errors.New("the request could not be created")
	ErrCreateRequestFailed = errors.New("failed to create HTTP request")
	ErrDoRequestFailed     = errors.New("failed to do HTTP request")
	ErrRequestFailed       = errors.New("HTTP request failed")
	ErrInvalidDataEntered  = errors.New("the data entered or the query form is not valid")
	ErrSerializedFields    = errors.New("there was a error when processing the fields")
	ErrPaginatorNotIsSetup = errors.New("paginator middleware is not setup")
	ErrPaginatorSize       = errors.New("paginator middleware is not setup")
	ErrTimeoutError        = errors.New("the request took too long to process")
	ErrConnectionClosed    = errors.New("the connection was closed before the request could be completed")
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

	ErrResponseMailNotFound = NewResponseError(
		http.StatusNotFound,
		"No mail was found",
		"MailNotFoundError",
		nil,
	)

	ErrResponseIndexNotFound = NewResponseError(
		http.StatusInternalServerError,
		"There are issues connecting to the server. Please contact technical support",
		"IndexNotFound",
		nil,
	)

	ErrResponseFaildParsedError = NewResponseError(
		http.StatusBadRequest,
		"The submitted query is not valid",
		"FailedParsedField",
		nil,
	)
)
