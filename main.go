package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/render"
	"net/http"
	_ "net/http"
	"submit/config"
	"submit/database"
	"submit/logger"
	"submit/model"
)

func main() {
	logger.Info("Submit v%s initializing...", "1.22.2")
	// Initialize environment configuration
	config.Init()
	// Initialize and connect to the PostgreSQL database
	store, err := database.NewDatabase()
	if err != nil {
		logger.Warn(err)
	}
	// Sync the code and database declarations, repair relations, etc.
	store.AutoMigrate(model.User{}, model.Submission{}, model.File{})
	router := chi.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "DB", store)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	router.Use(logger.Middleware)

	mapList := map[string]interface{}{}

	mapList["users"] = new(model.User)
	mapList["submissions"] = new(model.Submission)

	router.Route("/api", func(r chi.Router) {
		for k, v := range mapList {
			route := fmt.Sprintf("/%s", k)
			r.Route(route, func(r chi.Router) {
				model.CreateRoutes(v, r)
			})
		}
	})

	err = http.ListenAndServe(":3001", router)
	if err != nil {
		return
	}

}
