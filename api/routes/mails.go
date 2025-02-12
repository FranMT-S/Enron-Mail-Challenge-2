package routes

import (
	"net/http"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/controllers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/db"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func MailsRouter(mailController controllers.IMailController) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Paginator)
		r.Get("/", mailController.GetMails)
		r.Get("/{id}", mailController.GetMail)
		r.Get("/query", func(w http.ResponseWriter, r *http.Request) {
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
			fields := []string{models.FromField, models.ToField, models.DateField, models.SubjectField}
			Query, errRes := db.QueryBuilder(query, page, size, fields, timeZone)
			if errRes != nil {
				apierrors.RenderJSON(w, errRes)
				return
			}
			render.JSON(w, r, Query)
			return
		})
	})

	return r
}
