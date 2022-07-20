package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should return 81, but returned %d\n", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should return 9, but returned %d\n", rst)
	}
}

func TestSquare3(t *testing.T) {
	assert := assert.New(t)
	// 테스트 함수 호출
	assert.Equal(81, square(9),"9^2 = 81 결과가 나와야 함")
	assert.Equal(9, square(3),"3^2 = 9 결과가 나와야 함")
}