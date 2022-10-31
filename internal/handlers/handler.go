package handlers

import (
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/models"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/render"
	"net/http"
)

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "home.page.tmpl", &models.TemplateData{})
}
