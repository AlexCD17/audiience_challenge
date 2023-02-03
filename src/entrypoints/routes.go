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
	log.Println("Setting up router...")
	s.router.Use(ipValidatorMiddleware)
	s.router.Use(verifyMiddleware)
	s.router.Methods("Get").Path("/estimate").HandlerFunc(s.GetEstimate)
	log.Println("Ready: ")
}

func NewServer(estimateServices services.IService, router *mux.Router) *Server {
	log.Println("Creating server...")
	return &Server{
		estimate: estimateServices,
		router:   router,
	}
}
