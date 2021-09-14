package main

import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	for _, nth := range [...]int{2020, 30000000} {
		fmt.Printf("%dth number spoken: %v\n", nth, game(nth))
	}
}

func game(nth int) int {
	input := [...]int{0, 14, 1, 3, 7, 9}
	turns := make(map[int]int)
	for i, t := range input {
		turns[t] = i + 1
	}

	last := input[len(input)-1]

	var temp int
	for t := len(turns) + 1; t <= nth; t++ {
		prev, ok := turns[last]
		if ok {
			temp = t - 1 - prev
		} else {
			temp = 0
		}
		turns[last] = t - 1
		last = temp
	}
	return last
}
