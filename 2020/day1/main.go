package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(entries []int) string {
	// O(n²) incoming
	var result string
	for i := 0; i < len(entries); i++ {
		a := entries[i]
		for j := 0; j < len(entries); j++ {
			b := entries[j]
			if a+b == 2020 {
				result = fmt.Sprintf("%d + %d = %d", a, b, a*b)
				return result
			}
		}
	}
	return ""
}

func part2(entries []int) string {
	// O(n³) incoming
	var result string
	for i := 0; i < len(entries); i++ {
		a := entries[i]
		for j := 0; j < len(entries); j++ {
			b := entries[j]
			for k := 0; k < len(entries); k++ {
				c := entries[k]
				if a+b+c == 2020 {
					result = fmt.Sprintf("%d + %d + %d = %d", a, b, c, a*b*c)
					return result
				}
			}
		}
	}
	return ""
}

func main() {
	f, e := os.Open("./input")
	check(e)
	defer f.Close()

	var entries []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		entry, e := strconv.Atoi(s.Text())
		check(e)
		entries = append(entries, entry)
	}
	check(s.Err())

	fmt.Println("Part 1:", part1(entries))
	fmt.Println("Part 2:", part2(entries))
}
