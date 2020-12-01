package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_recursiveFuelPerModule(t *testing.T) {
	i, j, k := threeElementsThatSum2020([]int{1, 2, 2017})
	assert.Equal(t, i, 1)
	assert.Equal(t, j, 2)
	assert.Equal(t, k, 2017)
	i, j = twoElementsThatSum2020([]int{2019, 1})
	assert.Equal(t, i, 2019)
	assert.Equal(t, j, 1)
}
