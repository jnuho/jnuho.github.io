package main

import (
	"log"

  "github.com/gocolly/colly/v2"
)

func main() {
  c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	var result string

	url := "https://raw.githubusercontent.com/jnuho/jnuho.github.io/master/hellogo/leetcode/1-two-sum/README.md"

  c.OnHTML("pre", func(e *colly.HTMLElement) {
		log.Println(e)
		result = e.Text
		
  })
	log.Println(result)


	// Start scraping
  c.Visit(url)

	// Wait until threads are finished
	c.Wait()
}

