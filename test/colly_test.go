package test

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"testing"
)

type RateS struct {
	Rate	string
}
//testing uupool for making sure the json response
func Test_json(t *testing.T) {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://uupool.cn/api/getPosCoin.php")
}

// testing uupool for making sure the json response
func Test_html(t *testing.T) {
	c := colly.NewCollector()
	c.OnHTML("div[class='col-lg-3 col-md-4'] div:nth-child(2) ul li:nth-child(3) span", func(e *colly.HTMLElement) {
		fmt.Println("~~~~~~~~~~~")
		fmt.Println( e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start
	// scraping on https://hackerspaces.org
	c.Visit("http://dwarfpool.com/eth/address?wallet=0x12bd26dadd552233a86be4a9696f999362953912")
}

func Test_json_map(t *testing.T) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting=", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) { //online.xrB11x36.rate
		num := gjson.GetBytes(r.Body, "workers")
		a := num.Map()
		fmt.Println("~~~~~~~num~~~~~~~~~~~")
		fmt.Println(num)
		for k,v := range a{
			fmt.Println(k)
			fmt.Println(v.Get("count_invalid_share"))
			fmt.Println(v.Get("count_share"))
		}
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://eth.pandaminer.com/api/api/workers/0xddf76a8dd8aae44e904a21b74b7552c83e2cea8d")
}

func Test_json2(t *testing.T) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting=", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) { //online.xrB11x36.rate
		num := gjson.Get(string(r.Body), "data.statistics.currentStatistics.reportedHashrate")
		fmt.Println("~~~~~~~num~~~~~~~~~~~")

		fmt.Println(num.String())

	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://api.ethermine.org/miner/0xfd69143d360c804b0cde31116d44bc04767a01ab/dashboard")
}
func Test_api(t *testing.T) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting=", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) { //online.xrB11x36.rate
		num := gjson.Get(string(r.Body), "total_hashrate")

		fmt.Println("~~~~~~~num~~~~~~~~~~~")
		fmt.Println(num.String())
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("http://dwarfpool.com/eth/api?wallet=0x12bd26dadd552233a86be4a9696f999362953912&email=eth@example.com")
}

