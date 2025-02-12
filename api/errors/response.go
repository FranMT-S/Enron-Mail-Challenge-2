package apierrors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseError struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	LogError error  `json:"-"`
}

func NewResponseError(status int, message string, code string, logError error) *ResponseError {
	return &ResponseError{
		Code:     code,
		Status:   status,
		Message:  message,
		LogError: logError,
	}
}

func (reserr ResponseError) Error() string {
	msg := fmt.Sprintf(`
	{
		"status":%v,
		"code":%v,
		"message":"%v"
	}`, reserr.Status, reserr.Code, reserr.Message)

	jsonData, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		return "failed converting error to json"
	}

	return string(jsonData)
}

func (resErr ResponseError) WithLogError(err error) *ResponseError {
	resErr.LogError = err
	return &resErr
}

func (resErr ResponseError) SetStatus(code int) *ResponseError {
	resErr.Status = code
	return &resErr
}

func RenderJSON(w http.ResponseWriter, re *ResponseError) {
	if re.LogError != nil {
		log.Println(re.LogError.Error())
	}

	json, err := json.Marshal(re)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errMsg := fmt.Sprintf(`
		{
			"status":%v,
			"code":%v,
			"message":"%v"
		}`, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		log.Println(re.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(re.Status)

	w.Write(json) //nolint:errcheck
}
