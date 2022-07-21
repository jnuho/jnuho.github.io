package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/"
	, func(w http.ResponseWriter
		, r *http.Request) {
			fmt.Fprint("Hello world")
	})

	http.ListenAndServe(":3000", nil)
}