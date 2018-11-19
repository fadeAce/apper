package handler

import (
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"encoding/json"
	"time"
)

func MatchJSON(url, path string, seq int, duration time.Duration) (res []byte, err error) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Info("pipe ", seq, " Visiting=", r.URL.String())
	})
	c.OnResponse(func(response *colly.Response) {
		result := gjson.GetBytes(response.Body, path)
		res, _ = json.Marshal(result.Raw)
	})
	c.SetRequestTimeout(duration)
	// Start scraping on https://hackerspaces.org
	c.Visit(url)
	return
}
