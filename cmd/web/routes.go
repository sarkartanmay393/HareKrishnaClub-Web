package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	_ = app
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	return mux
}
