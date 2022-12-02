package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/marcos-nsantos/portfolio-api/internal/project"
	"gorm.io/gorm"
)

type Server struct {
	Router   *chi.Mux
	Database *gorm.DB
	Project  project.Service
}

func CreateNewServer(database *gorm.DB) *Server {
	s := &Server{
		Router:   chi.NewRouter(),
		Database: database,
	}
	s.Project = project.NewServices(project.NewRepo(s.Database))
	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.AllowContentType("application/json"))
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	s.Router.Use(middleware.Heartbeat("/"))

	// Mount all handlers here
	s.Router.Post("/projects", s.createProject)
	s.Router.Get("/projects/{id}", s.getProject)
}
