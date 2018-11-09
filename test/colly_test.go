package test

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"reflect"
	"testing"
)

type RateS struct {
	Rate	string
}
//testing uupool for making sure the json response
//func Test_json(t *testing.T) {
//	c := colly.NewCollector()
//	c.OnHTML("workers", func(e *colly.HTMLElement) {
//		fmt.Println(e.Text)
//
//	})
//
//	//c.OnResponse(func(r *colly.Response) {
//	//	fmt.Println(string(r.Body))
//	//})
//	// Start scraping on https://hackerspaces.org
//	c.Visit("http://dwarfpool.com/eth/api?wallet=0x980ca6dc75f7e2eec8a5211e10277361564b1d91&email=eth@example.com")
//}
//
//// testing uupool for making sure the json response
//func Test_html(t *testing.T) {
//	c := colly.NewCollector()
//	c.OnHTML("div[class='col-lg-3 col-md-4'] div:nth-child(2) ul li:nth-child(3) span", func(e *colly.HTMLElement) {
//		fmt.Println("~~~~~~~~~~~")
//		fmt.Println( e.Text)
//	})
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL.String())
//	})
//
//	// Start
//	// scraping on https://hackerspaces.org
//	c.Visit("http://dwarfpool.com/eth/address?wallet=0x12bd26dadd552233a86be4a9696f999362953912")
//}
//
//
//func Test_json2(t *testing.T) {
//	c := colly.NewCollector()
//
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting=", r.URL.String())
//	})
//	c.OnResponse(func(r *colly.Response) { //online.xrB11x36.rate
//		num := gjson.Get(string(r.Body), "data.statistics.currentStatistics.reportedHashrate")
//		fmt.Println("~~~~~~~num~~~~~~~~~~~")
//
//		fmt.Println(num.String())
//
//	})
//	// Start scraping on https://hackerspaces.org
//	c.Visit("https://api.ethermine.org/miner/0xfd69143d360c804b0cde31116d44bc04767a01ab/dashboard")
//}
func Test_api(t *testing.T) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting=", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) { //online.xrB11x36.rate
		num := gjson.Get(string(r.Body), "workers")
		fmt.Println(num.String())
		num2 :=gjson.Parse(num.String())
		mpp := num2.Map()
		for _,v := range  mpp {
			fmt.Println(v.Get("alive").Bool())
			fmt.Println("type:", reflect.TypeOf(v.Get("alive").Bool()))
		}
		fmt.Println(len(mpp))

	})
	// Start scraping on https://hackerspaces.org
	c.Visit("http://dwarfpool.com/eth/api?wallet=0x5f51b076339bdbb4b56dbc22e8890125d13b2fdd&email=eth@example.com")
}
