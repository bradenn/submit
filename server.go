package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"submit/jwt"
	"submit/logger"
	"submit/services"
	"submit/templates"
)

type Server struct {
	Router chi.Router
}

func NewServer() Server {
	server := Server{
		Router: chi.NewRouter(),
	}

	server.configureMiddleware()

	return server
}

func (s *Server) Run(address string) {
	err := services.NewDatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer services.TerminateDatabase()

	logger.Info("Server listening on %d routes at %s.", len(s.Router.Routes()), address)
	err = http.ListenAndServe(address, s.Router)
	if err != nil {
		logger.Error("Server Error: %s", err.Error())
	}
}


func (s *Server) configureMiddleware() {
	s.Router.Use(logger.Middleware)
	// middleware.Heartbeat()
	s.Router.Group(func(r chi.Router) {
		jwt.Init()
		r.Use(jwtauth.Verifier(jwt.Auth))
		r.Use(jwtauth.Authenticator)
		r.Get("/home", func(writer http.ResponseWriter, request *http.Request) {
			_, k, _ := jwtauth.FromContext(request.Context())
			_, err := writer.Write([]byte(fmt.Sprintf("Welcome my princess, %s.", k["userId"])))
			if err != nil {
				fmt.Println(err)
			}
		})
	})

	s.Router.Get("/status", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("stated."))
	})

	fs := http.FileServer(http.Dir("static"))
	s.Router.Handle("/static/*", http.StripPrefix("/static/", fs))

	s.Router.Route("/submission", services.SubmissionRoute)

	s.Router.Route("/login", services.LoginRoute)

	s.Router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		templates.Write("root.gohtml", writer, nil)
	})
}
