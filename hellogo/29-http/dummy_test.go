package main

import (
	"io"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := http.NewServeMux()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := http.NewServeMux()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := Student{}

	// res.Body 파싱 -> Student 타입
	err := json.NewDecoder(res.Body).Decode(student)
	assert.Nil(err)
	assert.Equal("Abc" student.Name)
	assert.Equal(18, student.Age)
	assert.Equal(87, student.Score)
}