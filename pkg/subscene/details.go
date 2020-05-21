package subscene

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"strconv"
	"strings"
)

type SubtitleDetails struct {
	Releases     []string `json:"releases"`
	Description  string   `json:"description"`
	User         string   `json:"user"`
	UserRating   int      `json:"userRating"`
	DateTime     string   `json:"datetime"`
	Files        int      `json:"files"`
	Rate         int      `json:"rate"`
	RateCount    int      `json:"rateCount"`
	Downloads    int      `json:"downloads"`
	DownloadLink string   `json:"downloadLink"`
	Comments     int      `json:"comments"`
	CommentLink  string   `json:"commentLink"`
}

func Details(name, id string) (SubtitleDetails, error) {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	sub := SubtitleDetails{
		CommentLink: "https://comments.jeded.com/comments/" + id,
	}

	c.OnHTML("a#downloadButton", func(e *colly.HTMLElement) {
		splittedPath := strings.Split(strings.TrimSpace(e.Attr("href")), "/")
		sub.DownloadLink = "/download/?token=" + splittedPath[len(splittedPath) - 1]
	})

	c.OnHTML("#content > div.subtitle li.release", func(e *colly.HTMLElement) {
		releases := []string{}

		e.ForEach("div", func(_ int, e *colly.HTMLElement) {
			releases = append(releases, strings.TrimSpace(e.Text))
		})

		sub.Releases = releases
	})

	c.OnHTML("#content > div.subtitle li.author a", func(e *colly.HTMLElement) {
		sub.User = strings.TrimSpace(e.Text)

		ratingTitle := strings.TrimSpace(e.ChildAttr("span.rating-bar", "title"))
		rateString := strings.TrimSpace(strings.TrimPrefix(ratingTitle, "Combined subtitle rating: "))

		rate, _ := strconv.Atoi(rateString)

		sub.UserRating = rate
	})

	c.OnHTML("#content > div.subtitle div.details > div.commentsContainer h3 a", func(e *colly.HTMLElement) {
		commentsStrings := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(e.Text), "View Comments ("), ")")
		sub.Comments, _ = strconv.Atoi(commentsStrings)
	})

	c.OnHTML("#content > div.subtitle li.comment-wrapper div.comment", func(e *colly.HTMLElement) {
		sub.Description = strings.TrimSpace(e.Text)
	})

	c.OnHTML("#details ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, e *colly.HTMLElement) {
			title := strings.TrimSpace(e.ChildText("strong:nth-child(1)"))

			if title == "Online:" {
				dateTimeString := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(e.Text), "Online:"))
				sub.DateTime = strings.Split(dateTimeString, "\n")[0]
			}

			if title == "Files:" {
				filesText := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(e.Text), "Files:"))
				FilesString := strings.Split(filesText, " ")
				sub.Files, _ = strconv.Atoi(FilesString[0])
			}

			if title == "Downloads:" {
				downloadsText := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(e.Text), "Downloads:"))
				downloadsString := strings.ReplaceAll(downloadsText, ",", "")
				sub.Downloads, _ = strconv.Atoi(downloadsString)
			}
		})

		sub.Rate, _ = strconv.Atoi(strings.TrimSpace(e.ChildText("span[itemprop=ratingValue]")))
		sub.RateCount, _ = strconv.Atoi(strings.TrimSpace(e.ChildText("span[itemprop=ratingCount]")))
	})

	err := c.Visit("https://subscene.com/subtitles/" + name + "/farsi_persian/" + id)

	return sub, err
}
