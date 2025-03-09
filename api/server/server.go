package server

import (
	"net/http"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/config"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/controllers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/db"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/routes"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router          *chi.Mux
	port            string
	emailController controllers.IMailController
}

func NewServer(port string) *Server {
	r := chi.NewRouter()
	db := db.NewZinchSearchDB()
	emailService := services.NewMailService(db)
	emailController := controllers.NewMailController(emailService)

	return &Server{
		router:          r,
		port:            port,
		emailController: emailController,
	}
}

func (server *Server) Start() {
	server.setupMiddleware()
	server.setupRoutes()
	http.ListenAndServe(":"+server.port, server.router)
}

func (server *Server) setupMiddleware() {
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.CleanPath)
	server.router.Use(middleware.Compress(5, "application/json"))

	server.router.Use(middlewares.SetupCORS())
	server.router.Use(middlewares.JsonResponseMiddleware)
	server.router.Use(middlewares.CleanQueryMiddleware)

}

func (server *Server) setupRoutes() {
	mailsRoutes := routes.MailsRouter(server.emailController)
	api := chi.NewRouter()
	api.Mount("/mails", mailsRoutes)

	if config.IS_DEVELOPMENT {
		api.Route("/test", func(r chi.Router) {
			r.Use(middlewares.Paginator(constants.PAGINATOR_MAXSIZE))
			r.Get("/query", controllers.TestQueryBuilderfunc)
		})
	}

	server.router.Mount("/api", api)

	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Working"))
	})
}
