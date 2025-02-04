package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func LongRequestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	min := 2
	max := 7
	randomInt := rand.Intn(max+1-min) + min
	select {
	case <-time.After(time.Duration(randomInt) * time.Second): // Simula una tarea larga
		fmt.Println("Test Proceso completado ", randomInt)
		fmt.Fprintln(w, fmt.Sprintf("Test Proceso completado %v", randomInt))
	case <-ctx.Done(): // Maneja la cancelación de la petición
		http.Error(w, "Petición cancelada", http.StatusRequestTimeout)
	}
}
