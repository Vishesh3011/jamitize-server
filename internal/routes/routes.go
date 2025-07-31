package routes

import (
	"example/internal/controller"
	"example/internal/core/application"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	application.Application
}

func NewServer(app application.Application) *Server {
	return &Server{
		Application: app,
	}
}

func (s *Server) AddRoutes(router *http.ServeMux) {
	baseUrl := "/apis"
	Get(router, baseUrl+"/health", HandleRequest(s.Application, controller.HealthCheck))
}

func (s *Server) Start() *Server {
	router := http.NewServeMux()
	s.AddRoutes(router)
	handler := ChainMiddleware(router, CORSMiddleware)
	s.server = &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Unable to start server: %v", err)
		}
	}()
	return s
}
