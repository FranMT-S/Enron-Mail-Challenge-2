package controllers

import (
	"context"
	"net/http"
	"strings"
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
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()

	query := r.URL.Query().Get("query")
	sortParam := r.URL.Query().Get("sort")
	sort := strings.Split(sortParam, ",")
	page, err := middlewares.GetPageFromContext(r)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	size, err := middlewares.GetSizeFromContext(r)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	sortFields := helpers.CreateSortFields(sort)

	if len(sortFields) == 0 || sortParam == "" {
		sortFields = []string{"-date"}
	}

	hits, err := mc.emailService.GetMailsHitsAndTotal(ctx, query, page, size, sortFields)

	response(hits, err, w, r)
}

func (mc MailController) GetMail(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	hits, err := mc.emailService.GetMailByID(ctx, id)

	response(hits, err, w, r)
}

func response[T any](response T, err *apierrors.ResponseError, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	render.JSON(w, r, response)
}
