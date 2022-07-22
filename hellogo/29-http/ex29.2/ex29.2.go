package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query() // 쿼리인수 가져오기
	name := values.Get("name") // 특정 키 값이 있는지 확인
	if name == "" {
		name = "World"
	}

	id, _ := strconv.Atoi(values.Get("id")) // id값을 가져와서 int형 타입 변환
	fmt.Fprintf(w, "Hello %s! id: %d", name, id)
}


// http://localhost:3000/bar?name=Lalalala&id=123
func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}