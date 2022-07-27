package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil) // '/students' 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(4, len(list))
	assert.Equal("zzz", list[0].Name)
	assert.Equal("aaa", list[1].Name)
}

func TestJsonHandler2(t *testing.T) { 
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student/1", nil) // id=1 학생 조회

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("zzz", student.Name)


	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/student/2", nil) // id=2 학생 조회
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)
}

func TestJsonHandler3(t *testing.T) { 
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		// 새로운 학생 데이터 등록
		strings.NewReader(`{"Id":0, "Name": "nnnn", "Age": 15, "Score": 78}`),
	)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/student/5", nil)
	// 추가된 학생 데이터
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("nnnn", student.Name)
}

func TestJsonHandler4(t *testing.T) { 
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/student/1", nil)

	// DELETE 요청
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(3, len(list))
	assert.Equal("aaa", list[0].Name)
}