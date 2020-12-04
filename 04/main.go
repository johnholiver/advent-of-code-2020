package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2020/pkg/input"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2020", "4")
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
		strings = append(strings, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := buildPassportSlice(strings)

	return fmt.Sprint(countPresentPassports(m))
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

	m := buildPassportSlice(strings)

	return fmt.Sprint(countValidPassports(m))
}

type Passport struct {
	Ecl string
	Pid string
	Eyr string
	Hcl string
	Byr string
	Iyr string
	Cid string
	Hgt string
}

func newPassport() Passport {
	return Passport{
		Ecl: "_",
		Pid: "_",
		Eyr: "_",
		Hcl: "_",
		Byr: "_",
		Iyr: "_",
		Cid: "_",
		Hgt: "_",
	}
}

func (p Passport) isPresent() bool {
	return p.Pid != "_" && p.Ecl != "_" && p.Byr != "_" && p.Eyr != "_" && p.Hcl != "_" && p.Hgt != "_" && p.Iyr != "_"
}
func (p Passport) isValid() bool {
	contains := func(s []string, searchterm string) bool {
		i := sort.SearchStrings(s, searchterm)
		return i < len(s) && s[i] == searchterm
	}

	nByr, _ := strconv.Atoi(p.Byr)
	byrValid := nByr >= 1920 && nByr <= 2002

	nIyr, _ := strconv.Atoi(p.Iyr)
	iyrValid := nIyr >= 2010 && nIyr <= 2020

	nEyr, _ := strconv.Atoi(p.Eyr)
	eyrValid := nEyr >= 2020 && nEyr <= 2030

	hgtValid := false
	lenHgt := len(p.Hgt)
	sHgt := p.Hgt[lenHgt-2 : lenHgt]
	if sHgt == "cm" {
		nHgt, _ := strconv.Atoi(p.Hgt[0:3])
		hgtValid = nHgt >= 150 && nHgt <= 193
	} else if sHgt == "in" {
		nHgt, _ := strconv.Atoi(p.Hgt[0:2])
		hgtValid = nHgt >= 59 && nHgt <= 76
	}

	hclValid := false
	hclValid = len(p.Hcl) == 7 &&
		p.Hcl[0] == '#' &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[1])) &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[2])) &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[3])) &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[4])) &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[5])) &&
		contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}, string(p.Hcl[6]))

	eclValid := contains([]string{"amb", "blu", "brn", "grn", "gry", "hzl", "oth"}, p.Ecl)

	pidValid := false
	if _, err := strconv.Atoi(p.Pid); err == nil && len(p.Pid) == 9 {
		pidValid = true
	}

	pValid := pidValid && eclValid && byrValid && eyrValid && hclValid && hgtValid && iyrValid
	fmt.Printf("Passport (%v): %v\n", p.Pid, pValid)
	if !pidValid {
		fmt.Printf("p.Pid: %v\n", p.Pid)
	}
	if !eclValid {
		fmt.Printf("p.Ecl: %v\n", p.Ecl)
	}
	if !byrValid {
		fmt.Printf("p.Byr: %v\n", p.Byr)
	}
	if !eyrValid {
		fmt.Printf("p.Eyr: %v\n", p.Eyr)
	}
	if !hclValid {
		fmt.Printf("p.Hcl: %v\n", p.Hcl)
	}
	if !hgtValid {
		fmt.Printf("p.Hgt: %v\n", p.Hgt)
	}
	if !iyrValid {
		fmt.Printf("p.Iyr: %v\n", p.Iyr)
	}

	return pValid
}
func buildPassportSlice(input []string) []Passport {
	pp := []Passport{}

	p := newPassport()
	for _, line := range input {
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			kv := strings.Split(token, ":")
			switch kv[0] {
			case "ecl":
				p.Ecl = kv[1]
				break
			case "pid":
				p.Pid = kv[1]
				break
			case "eyr":
				p.Eyr = kv[1]
				break
			case "hcl":
				p.Hcl = kv[1]
				break
			case "byr":
				p.Byr = kv[1]
				break
			case "iyr":
				p.Iyr = kv[1]
				break
			case "cid":
				p.Cid = kv[1]
				break
			case "hgt":
				p.Hgt = kv[1]
				break
			}
		}

		if line == "" {
			if p.Pid != "_" {
				pp = append(pp, p)
			}
			p = newPassport()
		}
	}
	if p.Pid != "_" {
		pp = append(pp, p)
	}

	return pp
}

func countPresentPassports(m []Passport) int {
	validCnt := 0
	for _, p := range m {
		if p.isPresent() {
			validCnt++
		}
	}
	return validCnt
}
func countValidPassports(m []Passport) int {
	validCnt := 0
	for _, p := range m {
		if p.isPresent() && p.isValid() {
			validCnt++
		}
	}
	return validCnt
}
