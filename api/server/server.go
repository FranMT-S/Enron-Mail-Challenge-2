package server

import (
	"net/http"

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

type Must struct {
	MatchAll struct{} `json:"match_all"`
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

	api := chi.NewRouter()
	api.Mount("/mails", routes.MailsRouter(server.emailController))

	server.router.Mount("/api", api)
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Working"))
	})
}
