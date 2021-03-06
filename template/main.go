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

	return fmt.Sprint(0)
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

	return fmt.Sprint(0)
}
