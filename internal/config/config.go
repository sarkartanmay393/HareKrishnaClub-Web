package config

import (
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/scraper"
	"html/template"
)

type AppConfig struct {
	CacheLoaded   bool
	TemplateCache map[string]*template.Template
	ScrapedBlogs  *[]scraper.Blog
}
