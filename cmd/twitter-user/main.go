package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Tweet struct {
	Username string
	UserId   int
	Tweets   string
}

func main() {
	c := colly.NewCollector()
	t := &Tweet{}

	c.Limit(&colly.LimitRule{
		DomainRegexp: "twitter.com",
		Delay:        15 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.Headers.Set("User-Agent", "1 Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Save("response.html"))
	})

	c.OnHTML(`#id__nf07z2bfra > span`, func(e *colly.HTMLElement) {
		t.Username = "username"
		t.Tweets = e.Text
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://twitter.com/elonmusk/status/1786826907297906699")

	fmt.Println("Hello, World!", *t)
}
