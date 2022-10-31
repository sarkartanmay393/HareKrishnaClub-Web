package main

import (
	"fmt"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/handlers"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/render"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/scraper"
	"log"
	"net/http"
	"os"
)

var portNumber = 8080

var app config.AppConfig

func main() {

	var err error
	scrapedBlogs, err := scraper.ScapeIskcondesiretree()
	if err != nil {
		log.Fatal("Scrape Failed", err)
	}
	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("In main.go", err)
	}

	app.ScrapedBlogs = &scrapedBlogs
	app.TemplateCache = cache
	app.CacheLoaded = true
	render.AttachConfig(&app)
	tmp := handlers.CreateNewRepository(&app)
	handlers.AttachRepository(tmp)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", portNumber),
		Handler: routes(&app),
	}

	log.Println("Starting server on port", portNumber)
	err = server.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}