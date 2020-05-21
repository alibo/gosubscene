package subscene

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"strings"
)

type Subtitle struct {
	Lang    string `json:"lang"`
	Release string `json:"release"`
	User    string `json:"user"`
	Comment string `json:"comment"`
	URL     string `json:"url"`
}

func ListSubtitles(name string) ([]Subtitle, error) {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	cookies := []*http.Cookie{
		{Name: "LanguageFilter", Value: "46"},
		{Name: "SortSubtitlesByDate", Value: "true"},
	}
	c.SetCookies("https://subscene.com", cookies)

	list := []Subtitle{}

	c.OnHTML("#content > div.subtitles.byFilm > div.content.clearfix > table > tbody > tr", func(e *colly.HTMLElement) {
		url := e.ChildAttr("td.a1 a[href]", "href")

		if url == "" {
			return
		}

		list = append(list, Subtitle{
			Lang:    strings.TrimSpace(e.ChildText("td.a1 span.l")),
			Release: strings.TrimSpace(e.ChildText("td.a1 span:nth-child(2)")),
			User:    strings.TrimSpace(e.ChildText("td.a5")),
			Comment: strings.TrimSpace(e.ChildText("td.a6")),
			URL:     url,
		})
	})

	err := c.Visit("https://subscene.com/subtitles/" + name)

	return list, err
}
