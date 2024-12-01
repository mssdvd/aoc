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

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()
	a, err := strconv.Atoi(s.Text())
	check(err)

	s.Scan()
	bus := strings.Split(s.Text(), ",")

	var earliestBusID int
	min := a

	for _, b := range bus {
		if b == "x" {
			continue
		}
		b, err := strconv.Atoi(b)
		check(err)

		w := (a/b+1)*b - a
		if w < min {
			min, earliestBusID = w, b
		}
	}

	// TODO: Part2
	// See the Chinese Remainder Theorem

	// From HN: had no idea WTF the Chinese Remainder Theorem was and
	// just worked out the issue iteratively. All you had to know was
	// that if you are trying to find some large number that is
	// divisible by other prime numbers, the deltas between the
	// candidates will be the product of the numbers.  as in, if you
	// are trying to find a number that is divisible by both 37 and
	// 41, you really just need to find numbers divisible by 1517.

	fmt.Println("ID * wait:", earliestBusID*min)
}
