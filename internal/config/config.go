package config

import (
	"html/template"

	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/scraper"
)

type AppConfig struct {
	CacheLoaded     bool
	TemplateCache   map[string]*template.Template
	ScrapedBlogs    *[]scraper.Blog
	ScrapedPoetries *[]scraper.Poetry
}
