package handlers

import "github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func CreateNewRepository(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

func AttachRepository(repo *Repository) {
	Repo = repo
}
