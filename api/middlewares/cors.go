package middlewares

import (
	"net/http"
	"strings"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/config"
	"github.com/go-chi/cors"
)

func SetupCORS() func(next http.Handler) http.Handler {
	allowerOrigins := strings.Split(config.CFG.AllowedOrigins, ",")
	return cors.Handler(cors.Options{
		AllowedOrigins:   allowerOrigins,
		AllowedMethods:   []string{http.MethodGet},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Time-Zone"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
}
