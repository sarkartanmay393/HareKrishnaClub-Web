package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	_ = app
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)

	mux.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	return mux
}
