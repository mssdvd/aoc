package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(fields []string) bool {
	min, e := strconv.Atoi(fields[0])
	check(e)
	max, e := strconv.Atoi(fields[1])
	check(e)
	char := fields[2]
	pass := fields[3]

	matches := strings.Count(pass, char)

	if min <= matches && matches <= max {
		return true
	}
	return false
}

func part2(fields []string) bool {
	first, e := strconv.Atoi(fields[0])
	check(e)
	second, e := strconv.Atoi(fields[1])
	check(e)
	char := fields[2]
	pass := fields[3]

	var matches int
	if string(pass[first-1]) == char {
		matches++
	}
	if string(pass[second-1]) == char {
		matches++
	}

	if matches == 1 {
		return true
	}
	return false
}

func main() {
	f, e := os.Open("./input")
	check(e)
	defer f.Close()

	var valid1, valid2 int
	s := bufio.NewScanner(f)
	for s.Scan() {
		entry := s.Text()

		fields := strings.FieldsFunc(entry, func(r rune) bool {
			if r == '-' || r == ' ' || r == ':' {
				return true
			}
			return false
		})

		if part1(fields) {
			valid1++
		}
		if part2(fields) {
			valid2++
		}
	}
	check(s.Err())

	fmt.Println("Valid passwords (part 1):", valid1)
	fmt.Println("Valid passwords (part 2):", valid2)
}
