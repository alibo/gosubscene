package subscene

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"strings"
)

type SearchItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type SearchResult struct {
	Found      bool                    `json:"found"`
	Categories map[string][]SearchItem `json:"categories"`
}

func Search(name string) (SearchResult, error) {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	query := map[string]string{
		"query": name,
		"l":     "",
	}

	result := SearchResult{
		Categories: map[string][]SearchItem{},
		Found:      false,
	}

	c.OnHTML("div.search-result h2", func(e *colly.HTMLElement) {
		category := strings.TrimSpace(e.Text)
		links := e.DOM.Parent().Find("a[href]")

		var items []SearchItem

		links.Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")

			items = append(items, SearchItem{
				Title: strings.TrimSpace(s.Text()),
				URL:   url,
			})

			result.Found = true
		})

		result.Categories[category] = items
	})

	err := c.Post("https://subscene.com/subtitles/searchbytitle", query)

	return result, err
}
