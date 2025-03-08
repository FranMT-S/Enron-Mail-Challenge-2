package shared

import (
	"errors"
	"fmt"
)

func MailError(path string, err error) error {
	msg := fmt.Sprintf("Failed to index mail with path %v, error:%v", path, err.Error())
	newError := errors.New(msg)

	return newError
}
