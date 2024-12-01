package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var left, right []int

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		l, err := strconv.Atoi(s.Text())
		check(err)
		left = append(left, l)

		s.Scan()

		r, err := strconv.Atoi(s.Text())
		check(err)
		right = append(right, r)
	}
	check(s.Err())

	slices.Sort(left)
	slices.Sort(right)

	var sum1, sum2 int
	occ := make(map[int]int)

	for i, l := range left {
		if l > right[i] {
			sum1 += l - right[i]
		} else {
			sum1 += right[i] - l
		}

		o, ok := occ[l]
		if !ok {
			for _, r := range right {
				if r > l {
					break
				}
				if r == l {
					o++
				}
			}
			occ[l] = o
		}
		sum2 += o * l
	}

	fmt.Printf("Part One: %v\n", sum1)
	fmt.Printf("Part Two: %v\n", sum2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
