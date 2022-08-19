package main
import "net/http"

// http://localhost:3000/gopher.jpg
func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}