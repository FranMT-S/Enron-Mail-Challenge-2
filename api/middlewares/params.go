package middlewares

import (
	"net/http"
	"net/url"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
)

// Sanitize the query by removing special characters
func CleanQueryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cleanedQuery := url.Values{}

		for key, values := range r.URL.Query() {
			cleanKey := helpers.CleanSpace(key)
			for _, value := range values {
				cleanValue := helpers.CleanSpace(value)
				cleanValue = helpers.SanitizeInput(cleanValue)
				cleanedQuery.Add(cleanKey, cleanValue)
			}
		}

		r.URL.RawQuery = cleanedQuery.Encode()
		next.ServeHTTP(w, r)
	})
}
