package handler

import (
	"gitlab.pandaminer.com/scar/apper/logger"
	"github.com/gocolly/colly"
	"time"
)

var log = logger.Log

func MatchHTML(url, path string, seq int, duration time.Duration) (res []string, err error) {

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Info("pipe ", seq, " Visiting=", r.URL.String())
	})
	c.OnHTML(path, func(element *colly.HTMLElement) {
		res = append(res, element.Text)
	})
	c.SetRequestTimeout(duration)
	// Start scraping on https://hackerspaces.org
	c.Visit(url)
	return
}
