package scraper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"time"
)

type Story struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	CoverImage  string    `json:"cover_image"`
	Body        string    `json:"body"`
	LastUpdated time.Time `json:"last_updated"`
}

// ScrapeStoriesFromIshaFoundation scrapes for https://isha.sadhguru.org
func ScrapeStoriesFromIshaFoundation() ([]Story, error) {
	// links
	link1 := "https://isha.sadhguru.org/us/en/wisdom?contentType=article&combine=Krishna"
	var stories []Story

	mainCollector := colly.NewCollector(colly.CacheDir("./.cache"), colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0"))

	mainCollector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	mainCollector.OnHTML("div.css-1xiu77m > div.css-174o3wj", func(e *colly.HTMLElement) {
		fmt.Print(e.DOM.Nodes)

		var story Story
		story.CoverImage = e.ChildAttr("a.chakra-link > div.card-image > div.css-2b4ntw > img", "src")
		story.URL = fmt.Sprintf("https://isha.sadhguru.org%s", e.ChildAttr("a", "href"))
		story.Title = e.ChildText("div.isha-card-podcast-title > a > div.css-45whdv")
		story.Body = e.ChildText("div.isha-rp-desc > a > div.css-1gcyif5")
		story.LastUpdated = time.Now()
		stories = append(stories, story)
	})

	mainCollector.Visit(link1)

	if len(stories) == 0 {
		return nil, fmt.Errorf("no story found\n")
	}

	log.Println("Successfully scraped Isha Foundation Few Krishna Stories.")

	// Saving scraped data for future references
	err := SaveScrapedData("stories-IshaFoundation.json", stories)
	if err != nil {
		log.Println("Failed to save scraped data")
	}

	return stories, nil
}

type Blog struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	//Author      string `json:"author"`
	//AuthorImage string `json:"author_image"`
	CoverImage string `json:"cover_image"`
	//ExtraRef    string `json:"extra_ref"`
	//PostDate    string `json:"post_date"`
	URL         string    `json:"url"`
	LastUpdated time.Time `json:"last_updated"`
}

// ScapeIskconDesireTree scrapes for https://iskcondesiretree.com
func ScapeIskconDesireTree() ([]Blog, error) {
	// links
	link1 := "https://iskcondesiretree.com/profiles/blogs/list/tag/lord+krishna"
	var blogs []Blog

	mainCollector := colly.NewCollector(colly.CacheDir("./.cache"))

	mainCollector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	// Find and visit all links and titles
	mainCollector.OnHTML("article.blogListPage-entry", func(e *colly.HTMLElement) {
		var blog Blog
		blog.URL = e.ChildAttr("header.entry-headline > div.media-body > h3.entry-title > a", "href")
		blog.Title = e.ChildText("header.entry-headline > div.media-body > h3.entry-title > a")
		blog.CoverImage = e.ChildAttr("section.entry-content > p > img", "src")

		blog.Body = "\"One who offers prayers to the Lord to fulfill his different desires must know that the highest perfectional fulfillment of desire is to go back home, back to Godhead.\""
		blog.LastUpdated = time.Now()
		blogs = append(blogs, blog)
	})

	// If next page is available
	pageCounter := 1
	mainCollector.OnHTML("div.pagination > a >  i.icon-next", func(e *colly.HTMLElement) {
		url :=
			fmt.Sprintf("https://iskcondesiretree.com/profiles/blogs/list/tag/lord+krishna?page=%d", pageCounter)
		pageCounter++
		mainCollector.Visit(url)
	})

	mainCollector.OnScraped(func(r *colly.Response) {
	})

	mainCollector.Visit(link1)

	if len(blogs) == 0 {
		return nil, fmt.Errorf("no blogs found")
	}

	log.Println("Successfully scraped Iskcon Desire Tree blogs.")

	// Saving scraped data for future references
	err := SaveScrapedData("blogs-IskconDesireTree.json", blogs)
	if err != nil {
		log.Println("Failed to save scraped data")
	}

	return blogs, nil
}

type Poetry struct {
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	CoverImage  string    `json:"cover_image"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"last_updated"`
}

// ScapeAllPoetry scrapes for https://allpoetry.com
func ScapeAllPoetry() ([]Poetry, error) {
	searchResult, err := Scrape("hare krishna poems allpoetry", "com", "en", 10, 10, 1)
	if err != nil {
		return []Poetry{}, err
	}

	poetries := []Poetry{}
	for _, result := range searchResult {
		if strings.Contains(result.SearchURL, "allpoetry.com") {
			poetry := Poetry{
				Title:       result.SearchTitle,
				URL:         result.SearchURL,
				Description: result.SearchDesc,
				LastUpdated: time.Now(),
			}

			poetries = append(poetries, poetry)
		}
	}

	log.Println("Successfully scraped allpoetry.com using Google search")
	// Saving scraped data for future references
	err = SaveScrapedData("poetry-AllPoetry.json", poetries)
	if err != nil {
		log.Println("Failed to save scraped data")
	}
	return poetries, nil
}
