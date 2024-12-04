package main

import (
	"bufio"
	"fmt"
	"os"
)

const word = "XMAS"

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var input []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		input = append(input, s.Text())
	}
	check(s.Err())

	var sum1, sum2 int

	for i := range len(input) {
		for j := range len(input[0]) {

			// ns == NorthSouth; ew == EastWest
			for ns := -1; ns < 2; ns++ {
			EastWestLoop:
				for ew := -1; ew < 2; ew++ {
					if ns == 0 && ew == 0 {
						continue
					}

					y, x := i, j
					var valid1 int

					for _, l := range word {
						if y < 0 || y >= len(input) || x < 0 || x >= len(input[0]) {
							continue EastWestLoop
						}

						if input[y][x] == byte(l) {
							valid1++
						}
						y += ns
						x += ew
					}

					if valid1 == len(word) {
						sum1++
					}

				}
			}

			if i == 0 || i == len(input)-1 || j == 0 || j == len(input[0])-1 {
				continue
			}

			var valid2 int

			if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' || input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' {
				valid2++
			}
			if input[i][j] == 'A' {
				valid2++
			}
			if input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' || input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M' {
				valid2++
			}

			if valid2 == 3 {
				sum2++
			}
		}
	}

	fmt.Println("Part One:", sum1)
	fmt.Println("Part Two:", sum2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
