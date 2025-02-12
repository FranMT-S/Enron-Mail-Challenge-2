package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IMailController interface {
	GetMails(w http.ResponseWriter, r *http.Request)
	GetMail(w http.ResponseWriter, r *http.Request)
}

type MailController struct {
	emailService services.IMailService
}

func NewMailController(emailService services.IMailService) *MailController {
	return &MailController{emailService: emailService}
}

func (mc MailController) GetMails(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, err := middlewares.Paginator.GetPageFromContext(r)
	timeZone := helpers.GetTimeZone(r)

	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	size, err := middlewares.Paginator.GetSizeFromContext(r)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	hits, err := mc.emailService.GetMailsHitsAndTotal(query, page, size, timeZone)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	render.JSON(w, r, hits)
}

func (mc MailController) GetMail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	hits, err := mc.emailService.GetMailByID(id)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	render.JSON(w, r, hits)
}

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
