package main

import (
	"bufio"
	"os"
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
	for s.Scan() {

	}
	check(s.Err())
}
