package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Allowed to cannoect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "DELETE", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-CSRF-Token", "Accpect"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           250,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	mux.Post("/handle", app.HandleSubmission)

	return mux
}
