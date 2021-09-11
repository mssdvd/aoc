package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var adapters []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		a, err := strconv.Atoi(s.Text())
		check(err)
		adapters = append(adapters, a)
	}
	check(s.Err())
	sort.Ints(adapters)
	fmt.Printf("%+v\n", adapters) // output for debug
	var oneJolt int
	var threeJolt int
	// Count diff between lowest-rated adapter and the outlet
	oneJolt++
	for i := 0; i < len(adapters)-1; i++ {
		switch adapters[i+1] - adapters[i] {
		case 1:
			oneJolt++
		case 3:
			threeJolt++
		}
	}
	// Count diff between highest-rated adapter and the device
	threeJolt++
	fmt.Println("Number of 1-jolt * 3-jolt diffs:", oneJolt*threeJolt)
}
