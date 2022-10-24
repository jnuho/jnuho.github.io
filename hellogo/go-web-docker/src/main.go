package main

import (
	"encoding/gob"
	"encoding/json"
	"os"
	"time"

	"log"
	"net/http"
)

var (
	port   = ":8080"
	logger = log.New(os.Stdout, "[TEST] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

type Data struct {
	Name  string
	Value float64
	TS    string
}

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		logger.Printf("wrapHandlerFunc : %s %s | Took %s", r.Method, r.URL.Path, time.Since(start))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Printf("GET: Hello World, PATH=%v", r.URL.Path)
	case "POST":
		log.Printf("POST: Hello World, PATH=%v", r.URL.Path)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func handlerJSON(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func handlerGOB(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := gob.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func main() {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", wrapHandlerFunc(handler))

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}
