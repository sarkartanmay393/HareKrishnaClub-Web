package render

import (
	"bytes"
	"fmt"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions map[string]interface{}
var pathToTemplates = "web/templates"

var app *config.AppConfig

func AttachConfig(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.ScrapedBlogs = app.ScrapedBlogs
	td.ScrapedPoetries = app.ScrapedPoetries
	td.LoadMore = false

	_ = r
	return td
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if len(pages) == 0 {
		log.Fatalln("No pages found", pages)
	}

	if err != nil {
		log.Println("ERR\t", err)
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("ERR\t", err)
			return cache, err
		}

		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			log.Println("ERR\t", err)
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				log.Println("ERR\t", err)
				return cache, err
			}
		}

		cache[name] = ts
	}

	if len(cache) == 0 {
		log.Fatalln("No templates found", cache)
	}

	return cache, nil
}

func TemplateRender(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template
	var err error

	if app.CacheLoaded {
		cache = app.TemplateCache
		// log.Println(app.CacheLoaded, cache)
	} else {
		cache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(app.CacheLoaded, cache)
	}

	t, ok := cache[tmpl]
	td.ActivePage = tmpl
	if !ok {
		log.Fatalln("Could not get template from cache")
	}

	buffer := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err = t.Execute(buffer, td)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}
}
