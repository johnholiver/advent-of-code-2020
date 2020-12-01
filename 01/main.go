package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2020/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
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

	var elements []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		elements = append(elements, i)
	}
	e1, e2 := twoElementsThatSum2020(elements)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(e1 * e2)
}

func twoElementsThatSum2020(expenseReport []int) (int, int) {
	for i := 0; i < len(expenseReport); i++ {
		e1 := expenseReport[i]
		for j := i + 1; j < len(expenseReport); j++ {
			e2 := expenseReport[j]
			if e1+e2 == 2020 {
				return e1, e2
			}
		}
	}
	return 0, 0
}

func threeElementsThatSum2020(expenseReport []int) (int, int, int) {
	for i := 0; i < len(expenseReport); i++ {
		e1 := expenseReport[i]
		for j := i + 1; j < len(expenseReport); j++ {
			e2 := expenseReport[j]
			for k := j + 1; k < len(expenseReport); k++ {
				e3 := expenseReport[k]
				if e1+e2+e3 == 2020 {
					return e1, e2, e3
				}
			}
		}
	}
	return 0, 0, 0
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)

	var elements []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		elements = append(elements, i)
	}

	e1, e2, e3 := threeElementsThatSum2020(elements)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(e1 * e2 * e3)
}
