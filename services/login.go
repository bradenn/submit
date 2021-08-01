package services

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"submit/templates"
)

func LoginRoute(router chi.Router) {
	router.Get("/", viewLogin)
}

func viewLogin(writer http.ResponseWriter, reader *http.Request) {
	templates.Write("login.gohtml", writer, nil)
}
