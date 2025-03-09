package apierrors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
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

// Used to log an error inside the ResponseError, which can be accessed via the LogError variable.
// The logged error does not appear in the JSON response.
func (resErr ResponseError) WithLogError(err error) *ResponseError {
	resErr.LogError = err
	return &resErr
}

func (resErr ResponseError) SetStatus(code int) *ResponseError {
	resErr.Status = code
	return &resErr
}

/*
RenderJSON sets the status code from the ResponseError and returns the error in JSON format.

- The status code is set based on the ResponseError's StatusCode field.
- The error message is returned in JSON format with the structure defined by ResponseError.
- Prints the original error contained in the `LogError` variable to the console for logging purposes.
*/
func RenderJSON(w http.ResponseWriter, re *ResponseError) {
	if re.LogError != nil {
		log.Println(constants.TextRed, re.LogError.Error(), constants.TextReset)
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
		log.Println(constants.TextYellow, re.Error(), constants.TextReset)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(re.Status)
	w.Write(json) //nolint:errcheck
}
