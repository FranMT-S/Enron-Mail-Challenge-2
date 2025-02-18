package controllers

import (
	"context"
	"net/http"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
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
	errCh := make(chan *apierrors.ResponseError)
	resCh := make(chan *models.EmailSummaryResponse)
	ctx := r.Context()

	go func() {
		query := r.URL.Query().Get("query")
		timeZone := helpers.GetTimeZone(r)
		page, err := middlewares.Paginator.GetPageFromContext(r)
		if err != nil {
			errCh <- err
			return
		}

		size, err := middlewares.Paginator.GetSizeFromContext(r)
		if err != nil {
			errCh <- err
			return
		}

		// time.Sleep(time.Duration(5) * time.Second)
		hits, err := mc.emailService.GetMailsHitsAndTotal(query, page, size, timeZone)
		if err != nil {
			errCh <- err
			return
		}

		resCh <- hits
	}()

	response(resCh, errCh, ctx, w, r)
}

func (mc MailController) GetMail(w http.ResponseWriter, r *http.Request) {
	errCh := make(chan *apierrors.ResponseError)
	resCh := make(chan *models.Email)
	ctx := r.Context()
	go func() {
		id := chi.URLParam(r, "id")
		hits, err := mc.emailService.GetMailByID(id)
		if err != nil {
			errCh <- err
			return
		}
		resCh <- hits
	}()

	response(resCh, errCh, ctx, w, r)
}

func response[T any](responseCh chan T, errCh chan *apierrors.ResponseError, ctx context.Context, w http.ResponseWriter, r *http.Request) {
	select {
	case hits := <-responseCh:
		render.JSON(w, r, hits)
	case err := <-errCh:
		apierrors.RenderJSON(w, err)
	case <-ctx.Done():
		apierrors.RenderJSON(w, apierrors.ErrResponseRequestCancelled)
	}
}
