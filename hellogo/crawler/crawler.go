package main

import (
  "fmt"
	"log"
  "os"

  "github.com/gocolly/colly/v2"
)

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
		result += "\n[INPUT 1]\n" + e.ChildText("pre#sample-input-1")
		result += "\n[OUTPUT 1]\n" + e.ChildText("pre#sample-output-1")
		result += "\n[INPUT 2]\n" + e.ChildText("pre#sample-input-2")
		result += "\n[OUTPUT 2]\n" + e.ChildText("pre#sample-output-2")
		result += "\n[INPUT 3]\n" + e.ChildText("pre#sample-input-3")
		result += "\n[OUTPUT 3]\n" + e.ChildText("pre#sample-output-3")
		result += "\n[INPUT 4]\n" + e.ChildText("pre#sample-input-4")
		result += "\n[OUTPUT 4]\n" + e.ChildText("pre#sample-output-4")
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
	f, err := os.OpenFile(fname, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
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