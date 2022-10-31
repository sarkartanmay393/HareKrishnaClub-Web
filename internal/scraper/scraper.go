package scraper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

type Blog struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	//Author      string `json:"author"`
	//AuthorImage string `json:"author_image"`
	CoverImage string `json:"cover_image"`
	//ExtraRef    string `json:"extra_ref"`
	//PostDate    string `json:"post_date"`
	URL string `json:"url"`
}

// ScapeIskcondesiretree scrapes for https://iskcondesiretree.com
func ScapeIskcondesiretree() ([]Blog, error) {
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
	return blogs, nil
}
