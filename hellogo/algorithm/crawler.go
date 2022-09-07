package main

import (
  "fmt"
	"log"
  "os"
  "strings"
  "strconv"

  "github.com/gocolly/colly/v2"
)

func main() {
  //if len(os.Args) < 2 {
    //fmt.Println("[ERROR] Missing problem number parameter.")
    //return
  //}
  //problem := os.Args[1]

  c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	//var result string

	url := "https://leetcode.com/problems/"
  //https://leetcode.com/problems
  //div role = "row"
  //  div role="cell"[1]
  // 5번쨰 child a.href
  //  /problems/two-sum
  //url += problem

  problems := make(map[int]string)

  c.OnHTML("div.truncate a", func(e *colly.HTMLElement) {
    
    os := strings.Index(e.Text, ".")
    idx,_ := strconv.Atoi(e.Text[:pos])

    problems[idx] = e.Attr("href")
		//result = "/**\n" + url + "\n\n" + result + "\n*/\n"
		
		//result += "package main\n\n"
		//result += "import (\n"
		//result += "  \"fmt\"\n"
		//result += ")\n\n"
		//result += "func main() {\n"
		//result += "  \n"
		//result += "}"
  })


	// Start scraping
  c.Visit(url)

	// Wait until threads are finished
	c.Wait()

  for key,val := range problems {
    fmt.Println(key, val)
  }
	// Write(Append) to file
	//defer writeToFile(problem, result)
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
