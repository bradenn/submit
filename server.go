package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"submit/config"
	"submit/logger"
)

// The Server type contains the core http hosting components.
type Server struct {
	Router   *chi.Mux
	Database *gorm.DB
	JWTAuth  *jwtauth.JWTAuth
}

func NewServer() Server {
	secret := os.Getenv("secret")

	server := Server{
		Router:  chi.NewRouter(),
		JWTAuth: jwtauth.New("HS256", []byte(secret), nil),
	}

	return server
}

func (s *Server) Init() error {
	// Load configuration from the environment
	config.Init()
	// Load router

	// Load routes
	// Load database
	// Load session
	return nil
}

// Run starts the http api server.
func (s *Server) Run() error {
	// Create url string from config variables
	address := fmt.Sprintf("%s:%s", os.Getenv("host"), os.Getenv("port"))
	logger.Info("Server listening on %d routes at %s.", len(s.Router.Routes()), address)
	// Begin listening on the specified address
	err := http.ListenAndServe(address, s.Router)
	if err != nil {
		logger.Error(err)
	}

	return nil
}
