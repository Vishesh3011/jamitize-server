package routes

import (
	"example/internal/controller"
	"example/internal/core/application"
	"net/http"
)

type Server struct {
	server *http.Server
	application.Application
}

func NewServer(app application.Application) *Server {
	s := &http.Server{
		Addr: ":8080",
	}
	return &Server{
		server:      s,
		Application: app,
	}
}

func (s *Server) AddRoutes(mux *http.ServeMux) {
	baseUrl := "/apis"
	mux.HandleFunc(http.MethodGet+" "+baseUrl+"/health", HandleRequest(s.Application, controller.HealthCheck))
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	s.AddRoutes(mux)
	s.server.Handler = mux

	return s.server.ListenAndServe()
}
