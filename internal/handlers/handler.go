package handlers

import (
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/models"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/render"
	"net/http"
)

// HomeHandler redirects to blogs on main page.
func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blogs", 303)
}

// BlogsHandler writes blog.tmpl to browser.
func (repo *Repository) BlogsHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "blogs.page.tmpl", &models.TemplateData{})
}

// PoetryHandler writes poetry.tmpl to browser.
func (repo *Repository) PoetryHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "poetry.page.tmpl", &models.TemplateData{})
}

// StoriesHandler writes poetry.tmpl to browser.
func (repo *Repository) StoriesHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "stories.page.tmpl", &models.TemplateData{})
}

// LordHandler writes poetry.tmpl to browser.
func (repo *Repository) LordHandler(w http.ResponseWriter, r *http.Request) {
	render.TemplateRender(w, r, "lord.page.tmpl", &models.TemplateData{})
}
