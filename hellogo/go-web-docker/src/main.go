package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)

	log.Println("Serving---go http server http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		log.Printf("ERROR Happend! %v\n", err)
	}
}

/**
curl http://localhost:8080
Hello World!
*/
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
