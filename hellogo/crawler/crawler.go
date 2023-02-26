package main

import (
	"fmt"
	"log"
	"os"

	colly "github.com/gocolly/colly/v2"
)

// go run crawler.go 1000
func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Missing problem number parameter.")
		return
	}
	problem := os.Args[1]

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	var result string

	url := "https://www.acmicpc.net/problem/"
	url += problem

	c.OnHTML("div#problem-body", func(e *colly.HTMLElement) {
		result = e.ChildText("div#problem_description")
		result += e.ChildText("div#problem_input")
		result += e.ChildText("div#problem_output")

		for i := 1; i <= 5; i++ {
			txt_input := e.ChildText(fmt.Sprintf("pre#sample-input-%d", i))
			txt_output := e.ChildText(fmt.Sprintf("pre#sample-output-%d", i))
			if len(txt_input) > 0 {
				result += fmt.Sprintf("\n[INPUT %d]\n%s", i, txt_input)
				result += fmt.Sprintf("\n[OUTPUT %d]\n%s", i, txt_output)
			} else {
				continue
			}
		}

		result = "/**\n" + url + "\n\n" + result + "\n*/\n"

		result += "package main\n\n"
		result += "import (\n"
		result += "  \"fmt\"\n"
		result += ")\n\n"

		result += "func main() {\n"
		result += "  \n"
		result += "}"
	})

	// Start scraping
	c.Visit(url)

	// Wait until threads are finished
	c.Wait()

	// Write(Append) to file
	defer writeToFile(problem, result)
}

// problem: 문제번호
// desc: 문제설명
func writeToFile(problem, desc string) {
	// If the
	fname := "problems/" + problem + ".go"
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(desc)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
