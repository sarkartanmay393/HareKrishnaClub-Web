package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/config"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/handlers"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/render"
	"github.com/sarkartanmay393/HareKrishnaClub-Web/internal/scraper"
)

var portNumber = 8080

var app config.AppConfig

func main() {

	var err error

	// Data from the scraped websites
	scrapedBlogs, scrapedPoetries, scrapedStories := ScrapeNecessaries()

	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("In main.go", err)
	}

	app.ScrapedBlogs = &scrapedBlogs
	app.ScrapedPoetries = &scrapedPoetries
	app.ScrapedStories = &scrapedStories

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

// ScrapeNecessaries scrapes the necessary data from the websites.
func ScrapeNecessaries() ([]scraper.Blog, []scraper.Poetry, []scraper.Story) {
	scrapedBlogs, err := scraper.ScapeIskconDesireTree()
	if err != nil {
		log.Fatal("scrapedBlogs Failed", err)
	}
	scrapedPoetries, err := scraper.ScapeAllPoetry()
	if err != nil {
		log.Fatal("scrapedPoetries Failed", err)
	}
	scrapedStories, err := scraper.ScrapeStoriesFromIshaFoundation()
	fmt.Println(len(scrapedStories))
	if err != nil {
		log.Fatal("scrapedStories Failed", err)
	}

	return scrapedBlogs, scrapedPoetries, scrapedStories
}
