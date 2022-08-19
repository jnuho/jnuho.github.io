package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id int
	Name string
	Age int
	Score int
}

/**
	Student 리스트를 Id 기준정렬하기 위한, Id 정렬 인터페이스 정의
*/
type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var students map[int]Student
var lastId int

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // 학생 목록을 Id로 정렬
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list) // JSON 포맷으로 변경
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student) // JSON 데이터 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lastId++	// id를 증가시킨 후 앱에 등록
	student.Id = lastId

	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // id 해당 학생 없으면 에러
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK) // 삭제 성공 시, StatusOK 반환
}

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() // gorilla/mux 생성

	/** 새로운 핸들러 등록 */
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET") // 학생 리스트 조회
	mux.HandleFunc("/student/{id:[0-9]+}", GetStudentHandler).Methods("GET") // 학생 한명 조회
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST") // 학생 등록
	mux.HandleFunc("/student/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE") // 학생 삭제

	students = make(map[int]Student)
	students[1] = Student{1, "zzz", 10, 77}
	students[2] = Student{2, "aaa", 20, 88}
	students[3] = Student{3, "ccc", 30, 99}
	students[4] = Student{4, "bbb", 40, 50}
	lastId = 4

	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}