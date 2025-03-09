package routes

import (
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/controllers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/go-chi/chi/v5"
)

// Returns a router with the routes configured for mail setup.
func MailsRouter(mailController controllers.IMailController) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Paginator(constants.PAGINATOR_MAXSIZE))
		r.Get("/", mailController.GetMails)
		r.Get("/{id}", mailController.GetMail)
	})

	return r
}
