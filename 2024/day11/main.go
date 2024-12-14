package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l, err := os.ReadFile("./input")
	check(err)

	originalStones := strings.Split(string(l[:len(l)-1]), " ")

	fmt.Println("Part One:", countStones(originalStones, 25))
	fmt.Println("Part One:", countStones(originalStones, 75))
}

func countStones(originalStones []string, blinks int) int {
	counts := make(map[string]int)

	for _, s := range originalStones {
		counts[s] = 1
	}

	for range blinks {
		nextCounts := make(map[string]int)
		for stone, count := range counts {
			children := subStones(stone)
			for _, child := range children {
				nextCounts[child] += count
			}

		}
		counts = nextCounts
	}

	var total int

	for _, count := range counts {
		total += count
	}
	return total
}

func subStones(s string) []string {
	var children []string

	if s == "0" {
		children = append(children, "1")
	} else if s == "1" {
		children = append(children, "2024")
	} else if len(s)%2 == 0 {
		l := len(s) / 2

		children = append(children, s[:l])

		s = strings.TrimLeft(s[l:], "0")

		if s == "" {
			children = append(children, "0")
		} else {
			children = append(children, s)
		}
	} else {
		ns, err := strconv.Atoi(s)
		check(err)
		ns *= 2024
		children = append(children, strconv.Itoa(ns))
	}
	return children
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
