package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	amounts := regexp.MustCompile("([0-9]+)")

	var tokens1, tokens2 int

	for s.Scan() {
		var a, b, prize [2]float64

		for _, v := range [...]*[2]float64{&a, &b, &prize} {
			values := amounts.FindAllString(s.Text(), -1)

			for i := range 2 {
				num, err := strconv.ParseFloat(values[i], 64)
				check(err)
				v[i] = num
			}
			s.Scan()
		}

		for iter := range 2 {
			total := &tokens1
			if iter == 1 {
				prize[0] += 10000000000000
				prize[1] += 10000000000000
				total = &tokens2
			}

			j := (a[0]*prize[1] - a[1]*prize[0]) / (b[1]*a[0] - b[0]*a[1])
			i := (prize[0] - j*b[0]) / a[0]

			if i-math.Floor(i) > 0.01 || j-math.Floor(j) > 0.01 {
				continue
			}

			*total += int(3*i + j)
		}
	}
	check(s.Err())

	fmt.Println("Part One:", tokens1)
	fmt.Println("Part Two:", tokens2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
