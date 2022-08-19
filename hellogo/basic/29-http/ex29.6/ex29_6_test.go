package main

import (
	"io"
	"net/http/httptest"
	"net/http"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // '/' 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) // '/student' 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := &Student{}

	// res.Body 파싱 -> Student 타입
	err := json.NewDecoder(res.Body).Decode(student)
	assert.Nil(err) // 결과 확인
	assert.Equal("Abc", student.Name)
	assert.Equal(18, student.Age)
	assert.Equal(87, student.Score)
}