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

func twoSum(nums []int, t int, cache map[int]int) bool {
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[t-nums[i]]; ok {
			return true
		}
		cache[nums[i]] = i
	}
	return false
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var input []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		l, err := strconv.Atoi(s.Text())
		check(err)
		input = append(input, l)
	}
	check(s.Err())

	const dim int = 25
	var window []int
	var i, invalid int
	cache := make(map[int]int)

	for _, n := range input {
		if len(window) == dim {
			if !twoSum(window, n, cache) {
				invalid = n
				break
			}
			window[i] = n
			delete(cache, i)
			i = (i + 1) % dim
		} else {
			window = append(window, n)
		}
	}

	fmt.Println("Invalid number:", invalid)

	var sum, head, min, max int
	var cont []int
	offset := 2

	for {
		cont = input[head : head+offset]
		sum = 0
		for _, v := range cont {
			sum += v
		}

		if sum == invalid {
			min = cont[0]
			for _, v := range cont {
				if v < min {
					min = v
				} else if v > max {
					max = v
				}
			}
			break
		}

		if head+offset == len(input) {
			head = 0
			offset++
		} else {
			head++
		}
	}

	fmt.Println("Encryption weaknes:", min+max)
}
