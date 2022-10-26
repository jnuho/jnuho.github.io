package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/unrolled/render"
	"github.com/xuri/excelize/v2"
)

var (
	// Rendering JSON response
	rd     *render.Render
	logger = log.New(os.Stdout, "[TEST] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	urls   map[string]string
)

func init() {
	rd = render.New()
	urls = make(map[string]string, 0)
}

// Response
type Success struct {
	Success bool              `json:"Success"`
	Urls    map[string]string `json:"Urls"`
}

// Response
type SuccessHealth struct {
	Success bool `json:"Success"`
	Status  int  `json:"Status"`
}

//func (s ImageUris) Len() int {
//return len(s)
//}
//
//func (s ImageUris) Less(i, j int) bool {
//return s[j].PushedAt.Before(s[i].PushedAt)
//}
//
//func (s ImageUris) Swap(i, j int) {
//s[i], s[j] = s[j], s[i]
//}

// 스트링 포함여부 확인 (대소문자 무시 Case-Insesitive)
func ContainsStr(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func checkHealth(url string) string {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 399 {
		return "red"
	} else {
		return "blue"
	}
}

// GET 핸들러 - TG 이름 조회
func GetTgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}
	r.ParseForm()
	tgname := r.Form.Get("tgname")

	urls = make(map[string]string, 0)

	f, err := excelize.OpenFile("../../assets/WebURL.xlsx")
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := f.GetRows("Sheet0")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		log.Println(row[1])
		if i > 0 {
			if tgname != "" {
				if ContainsStr(row[1], tgname) {
					urls[row[1]] = checkHealth(row[1])
				}
			} else {
				urls[row[1]] = checkHealth(row[1])
			}
		}
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//lists = []string{
	//"http://10.148.143.198:7081/etrade/ping.jsp",
	//"http://edm.koreanair.com/PNS-PDF-webapp/tivoli.jsp",
	//"http://10.148.41.211:7080/PNS-PDF-webapp/tivoli.jsp",
	//}
	//for _, e := range lists {
	//if ContainsStr(e, tgname) {
	//urls[e] = checkHealth(e)
	//}
	//}

	rd.JSON(w, http.StatusOK, Success{Success: true, Urls: urls})
}

func GetTgHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}
	r.ParseForm()
	url := r.Form.Get("url")

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	status := 0
	resp, err := client.Get(url)
	if err != nil {
		status = 400
	} else {
		status = resp.StatusCode
	}

	rd.JSON(w, http.StatusOK, SuccessHealth{Success: true, Status: status})
}

func main() {
	// flag port default 3001
	ptr := flag.String("port", "3030", "Port to listen to")
	flag.Parse()
	port := ":" + *ptr

	http.Handle("/", http.FileServer(http.Dir("../../web")))
	http.HandleFunc("/targetgroup", GetTgHandler)
	http.HandleFunc("/targethealth", GetTgHealthHandler)
	log.Println("Go http server --- http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
