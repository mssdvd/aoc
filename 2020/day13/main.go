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

	fmt.Println("ID * wait:", earliestBusID*min)
}
