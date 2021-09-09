package main

import (
	"fmt"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.ReadFile("./input")
	check(err)

	groups := regexp.MustCompile("(?m)"+"^$\n").Split(string(f), -1)

	var countAnyone, countEveryone int

	for _, group := range groups {
		answers := make(map[rune]int)

		for _, r := range group {
			answers[r]++
		}

		countAnyone += len(answers) - 1 // remove \n

		groupSize := answers['\n']

		for _, v := range answers {
			if v == groupSize {
				countEveryone++
			}
		}
		countEveryone-- // remove \n
	}

	fmt.Println("Part 1:", countAnyone)
	fmt.Println("Part 2:", countEveryone)
}
