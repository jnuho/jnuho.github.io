package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"log"
	"net/http"

	"github.com/unrolled/render"
)

type Success struct {
	Success bool
}

var (
	// Rendering JSON response
	rd     *render.Render
	logger = log.New(os.Stdout, "[TEST] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

type Data struct {
	Name  string
	Value float64
	TS    string
}

func wrapFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		logger.Printf("wrapFunc : %s %s | Took %s", r.Method, r.URL.Path, time.Since(start))
	}
}

func GetTgHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm() // Parses the request body
		log.Println(r.Form)
		x := r.Form.Get("email") // x will be "" if parameter is not set
		log.Println(x)
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
		// rd.JSON(w, http.StatusOK, Success{Success: true})
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func main() {
	// flag port default 3001
	ptr := flag.String("port", "3001", "Port to listen to")
	flag.Parse()
	port := ":" + *ptr

	http.Handle("/", http.FileServer(http.Dir("../../web")))
	http.HandleFunc("/targetgroup", wrapFunc(GetTgHandler))

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
