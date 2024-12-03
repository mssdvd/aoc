package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	memoryBytes, err := os.ReadFile("./input")
	check(err)
	memory := string(memoryBytes)

	sum1 := sumMuls(memory)
	fmt.Println("Part One:", sum1)

	var sum2, start int
	var flushed bool
	for i := range len(memory) {
		if flushed && memory[i:min(i+4, len(memory))] == "do()" {
			start = i
			flushed = false
		}
		if !flushed && (memory[i:min(i+7, len(memory))] == "don't()" || i == len(memory)-1) {
			sum2 += sumMuls(memory[start:i])
			flushed = true
		}

	}
	fmt.Println("Part Two:", sum2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sumMuls(memory string) int {
	reMul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	muls := reMul.FindAllStringSubmatch(memory, -1)

	var sum int
	for _, mul := range muls {
		a, err := strconv.Atoi(mul[1])
		check(err)
		b, err := strconv.Atoi(mul[2])
		check(err)
		sum += a * b
	}

	return sum
}
