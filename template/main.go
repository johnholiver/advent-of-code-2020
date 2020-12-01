package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2020/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2020", "1")
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
		strings = append(strings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(strings)

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

	fmt.Print(strings)

	return fmt.Sprint(0)
}

func isValid(passwordRule string) bool {
	tokens := strings.Split(passwordRule, " ")
	between := strings.Split(tokens[0], "-")
	aim := []rune(strings.Split(tokens[1], ":")[0])[0]
	pass := tokens[2]

	aimCnt := 0
	for _, char := range pass {
		if char == aim {
			aimCnt++
		}
	}
	lowLimit, _ := strconv.Atoi(between[0])
	upLimit, _ := strconv.Atoi(between[1])

	return aimCnt >= lowLimit && aimCnt <= upLimit
}

func isValid2(passwordRule string) bool {
	tokens := strings.Split(passwordRule, " ")
	between := strings.Split(tokens[0], "-")
	lowPos, _ := strconv.Atoi(between[0])
	upPos, _ := strconv.Atoi(between[1])

	aim := []rune(strings.Split(tokens[1], ":")[0])[0]
	pass := tokens[2]

	aimCnt := 0
	for pos, char := range pass {
		if char == aim && (lowPos-1 == pos || upPos-1 == pos) {
			aimCnt++
		}
	}

	return aimCnt == 1
}
