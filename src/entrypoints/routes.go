package entrypoints

import (
	"audiience_challenge/repositories/rates"
	"audiience_challenge/services"
	"github.com/gorilla/mux"
	"log"
)

type Server struct {
	estimate services.IService
	router   *mux.Router
	repo     *rates.Repository
}

func (s *Server) SetupRouter() {
	s.router.Use(ipValidatorMiddleware)
	s.router.Use(verifyMiddleware)
	s.router.Methods("Get").Path("/estimate").HandlerFunc(s.GetEstimate)
}

func NewServer(estimateServices services.IService, router *mux.Router) *Server {
	log.Println("creating server")
	return &Server{
		estimate: estimateServices,
		router:   router,
	}
}
