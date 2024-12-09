package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var sum1, sum2 int
	var p1, p2 [][]string

	s := bufio.NewScanner(f)
	for s.Scan() {
		eqStr := strings.FieldsFunc(s.Text(), func(r rune) bool { return r == ' ' || r == ':' })
		var eq []int
		for _, elStr := range eqStr {
			el, err := strconv.Atoi(elStr)
			check(err)
			eq = append(eq, el)
		}

		nOps := len(eq) - 2

		p1 = cartesianProduct(p1, []rune{'+', '*'}, nOps)
		p2 = cartesianProduct(p2, []rune{'+', '*', '|'}, nOps)

		sum1 += validEqs(p1, eq)
		sum2 += validEqs(p2, eq)
	}
	check(s.Err())

	fmt.Println("Part One:", sum1)
	fmt.Println("Part Two:", sum2)
}

func cartesianProduct(p [][]string, alph []rune, n int) [][]string {
	for k := len(p); k <= n; k++ {

		prod := []string{""}
		for i := 0; i < k; i++ {
			var newProd []string
			for _, prefix := range prod {
				for _, char := range alph {
					newProd = append(newProd, prefix+string(char))
				}
			}
			prod = newProd
		}
		p = append(p, prod)
	}

	return p
}

func validEqs(p [][]string, eq []int) int {
	ops := eq[1:]

	var sum int
	for _, o := range p[len(ops)-1] {
		res := ops[0]

		for i, c := range o {
			switch c {
			case '+':
				res += ops[i+1]
			case '*':
				res *= ops[i+1]
			case '|':
				var err error
				res, err = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(ops[i+1]))
				check(err)
			}
			if res > eq[0] {
				break
			}
		}

		if res == eq[0] {
			sum += eq[0]
			break
		}
	}
	return sum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
