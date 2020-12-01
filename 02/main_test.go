package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	validCnt := 0
	for _, v := range input {
		if isValid(v) {
			validCnt++
		}
	}

	assert.Equal(t, validCnt, 2)
}

func Test_isValid2(t *testing.T) {
	input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	validCnt := 0
	for _, v := range input {
		if isValid2(v) {
			validCnt++
		}
	}

	assert.Equal(t, validCnt, 1)
}
