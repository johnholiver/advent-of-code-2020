package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var input03 = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func Test_traverseForest(t *testing.T) {
	forest := strings.Split(input03, "\n")

	assert.Equal(t, 7, traverseForest(forest, 3, 1, 0, 0))
}

func Test_traverseForest2(t *testing.T) {
	forest := strings.Split(input03, "\n")

	total := traverseForest(forest, 1, 1, 0, 0)
	total *= traverseForest(forest, 3, 1, 0, 0)
	total *= traverseForest(forest, 5, 1, 0, 0)
	total *= traverseForest(forest, 7, 1, 0, 0)
	total *= traverseForest(forest, 1, 2, 0, 0)

	assert.Equal(t, 2, traverseForest(forest, 1, 1, 0, 0))
	assert.Equal(t, 7, traverseForest(forest, 3, 1, 0, 0))
	assert.Equal(t, 3, traverseForest(forest, 5, 1, 0, 0))
	assert.Equal(t, 4, traverseForest(forest, 7, 1, 0, 0))
	assert.Equal(t, 2, traverseForest(forest, 1, 2, 0, 0))

	assert.Equal(t, 336, total)
}
