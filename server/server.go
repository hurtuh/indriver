package server

import (
	"github.com/gorilla/mux"
	"github.com/hurtuh/indriver/domain"
	"net/http"
)

type Server struct {
	router   *mux.Router
	handlers domain.Handlers
}

func InitServer(handlers domain.Handlers) (s *Server) {
	s = new(Server)
	s.handlers = handlers
	s.router = mux.NewRouter()

	s.router.StrictSlash(true)

	s.router.Methods("POST").Path("/add").HandlerFunc(s.handlers.AddRecord)
	s.router.Methods("GET").Path("/get").HandlerFunc(s.handlers.GetRecord)
	s.router.Methods("PUT").Path("/edit").HandlerFunc(s.handlers.EditRecord)
	s.router.Methods("DELETE").Path("/delete").HandlerFunc(s.handlers.DelRecord)

	return s
}

func (s *Server) StartServer(port string) error {
	return http.ListenAndServe(port, s.router)
}