package routes

import (
	"errors"
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

	v1 := baseUrl + "/v1"

	user := v1 + "/users"
	Post(router, user, HandleRequest(s.Application, controller.CreateUserController))
	Post(router, user+"/login", HandleRequest(s.Application, controller.LoginUserController))
}

func (s *Server) Start() *http.Server {
	router := http.NewServeMux()
	s.AddRoutes(router)
	handler := ChainMiddleware(router, CORSMiddleware)
	s.server = &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Unable to start server: %v", err)
		}
	}()
	log.Println("Server started on http://localhost:8080")
	return s.server
}

func (s *Server) Server() *http.Server {
	return s.server
}
