package helpers

import (
	"net/http"
	"time"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
)

// Get string Timezone from 'Time-Zone' headers or 'tz' url params with header priority. if no exists return a empty string
func GetTimeZone(r *http.Request) string {
	timeZone := ""

	timeZone = r.Header.Get("Time-Zone")
	if timeZone != "" {
		return timeZone
	}

	timeZone = r.URL.Query().Get("tz")

	return timeZone
}

func DateParsed(value, timeZone string, layout []string) (time.Time, *apierrors.ResponseError) {
	var loc *time.Location
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}

	for _, la := range layout {
		date, err := time.ParseInLocation(la, value, loc)
		if err != nil {
			continue
		}
		return date, nil
	}

	resErr := apierrors.ErrResponseDateFormatNotValid.WithLogError(err)
	resErr.Message = resErr.Message + ".Value:" + value
	return time.Time{}, resErr
}

func GetDateWithTimeOneMinuteNextDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 0, 0, date.Location())
}
