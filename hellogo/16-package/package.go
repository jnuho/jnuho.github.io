package main

import (
	"fmt"
	"math/rand"

	// 겹치는 패키지 이름 별칭으로 묶기
	"text/template"
	htemplate "html/template"

	// 패키지를 사용하지는 않지만 부과효과 떄문에 import 하는 경우

	"database/sql"
	_ "github.com/mattn/go-sqlite3" // 밑줄 _ 이용하여 unused 에러 방지
	// "log"
	// "net/http"
)

type Service struct {
	db *sql.DB
}

func main() {

	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)

	fmt.Println(rand.Int())


	// Opening a driver typically will not attempt to connect to the database.
	/*
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.setmaxidleconns(50)
	db.setmaxopenconns(50)

	s := &service{db: db}

	http.listenandserve(":8080", s)
	*/
}