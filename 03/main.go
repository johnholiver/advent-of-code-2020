package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2020/pkg/input"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2020", "3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var strings []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		strings = append(strings, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprint(traverseForest(strings, 3, 1, 0, 0))
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := traverseForest(strings, 1, 1, 0, 0)
	total *= traverseForest(strings, 3, 1, 0, 0)
	total *= traverseForest(strings, 5, 1, 0, 0)
	total *= traverseForest(strings, 7, 1, 0, 0)
	total *= traverseForest(strings, 1, 2, 0, 0)

	return fmt.Sprint(total)
}

func traverseForest(forest []string, diffX, diffY, currentX, currentY int) int {
	hasTree := 0
	currentX = (currentX + diffX) % len(forest[currentY])
	currentY = currentY + diffY

	if currentY > len(forest)-1 {
		return hasTree
	}

	if rune(forest[currentY][currentX]) == '#' {
		hasTree = 1
	}

	return traverseForest(forest, diffX, diffY, currentX, currentY) + hasTree
}
