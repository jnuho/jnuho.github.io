package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Student struct {
	Name string
	Age int
	Score int
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{"Abc", 18, 87}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprint(w, "Hello World")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}