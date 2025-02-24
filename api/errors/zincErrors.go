package apierrors

import (
	"errors"
	"fmt"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/config"
)

var (
	ErrZincIdNotFound = errors.New("id not found")
	ErrZincTypeError  = "x_content_parse_exception"
)

type ZincSearchError struct {
	MessageError string `json:"error"`
}

func (ze ZincSearchError) Error() string {
	return ze.MessageError
}

type ZincSearchErrorWithType struct {
	MessageError ZincBodyError `json:"error"`
}

type ZincBodyError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
	Cause  string `json:"cause"`
}

func IsZincError(zincErr ZincSearchError) *ResponseError {
	var ErrZincIndexNotExist = fmt.Errorf("index %v does not exists", config.CFG.DatabaseName)

	if zincErr.MessageError == ErrZincIdNotFound.Error() {
		return ErrResponseMailNotFound
	} else if zincErr.MessageError == ErrZincIndexNotExist.Error() {
		return ErrResponseIndexNotFound
	}

	return nil
}

func IsZincErrorWithType(zincErr ZincSearchErrorWithType) *ResponseError {
	if zincErr.MessageError.Type == ErrZincTypeError {
		return ErrResponseFaildParsedError.WithLogError(errors.New(zincErr.MessageError.Cause))
	}

	return nil
}
