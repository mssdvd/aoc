package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var safe0, safe1 int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var levels []int

		for _, v := range strings.Split(s.Text(), " ") {
			l, err := strconv.Atoi(v)
			check(err)
			levels = append(levels, l)
		}

		if isSafe(levels) {
			safe0++
			continue
		}

		for pair := range len(levels) - 1 {
			l := slices.Clone(levels)
			l = slices.Delete(l, pair, pair+1)
			if isSafe(l) {
				safe1++
				break
			}

			r := slices.Clone(levels)
			r = slices.Delete(r, pair+1, pair+2)
			if isSafe(r) {
				safe1++
				break
			}
		}
	}
	check(s.Err())

	fmt.Println("Part One:", safe0)
	fmt.Println("Part Two:", safe0+safe1)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isSafe(levels []int) bool {
	var incr bool
	if levels[0] < levels[1] {
		incr = true
	}

	for i := 1; i < len(levels); i++ {
		diff, pos := abs(levels[i] - levels[i-1])
		if diff <= 0 || diff >= 4 || pos != incr {
			return false
		}
	}

	return true
}

func abs(n int) (int, bool) {
	if n < 0 {
		return -n, false
	}
	return n, true
}
