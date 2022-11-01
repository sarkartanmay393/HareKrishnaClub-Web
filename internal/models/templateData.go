package models

import (
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/scraper"
)

type TemplateData struct {
	CSRFToken    string
	ScrapedBlogs *[]scraper.Blog

	SuccessMessage string
	ErrorMessage   string
	WarningMessage string
	ActivePage     string
	LoadMore       bool
}
