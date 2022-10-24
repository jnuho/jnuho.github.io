package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
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

func wrapFunc(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(w, req)
		logger.Printf("wrapFunc : %s %s | Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		//fn(w, req)
		h.ServeHTTP(w, req)
		logger.Printf("wrapHandlerFunc : %s %s | Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func handlerJSON(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
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

func handlerGOB(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
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
	mainRouter.HandleFunc("/", wrapFunc(handler))
	mainRouter.HandleFunc("/json", wrapHandlerFunc(handlerJSON))
	mainRouter.HandleFunc("/gob", wrapHandlerFunc(handlerGOB))

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}
