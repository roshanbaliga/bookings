package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/roshanbaliga/bookings/pkgs/config"
	"github.com/roshanbaliga/bookings/pkgs/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	m := chi.NewRouter()
	m.Use(middleware.Recoverer)
	m.Use(NoSurf)
	m.Use(Session)

	m.Get("/", handlers.Repo.Home)
	m.Get("/about", handlers.Repo.About)
	return m
}
