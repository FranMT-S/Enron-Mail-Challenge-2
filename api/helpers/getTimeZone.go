package helpers

import (
	"net/http"
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
