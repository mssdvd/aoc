package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func travel(grid []string, right int, down int) int {
	lineLen := len(grid[0])

	var j, trees int

	for i := down; i < len(grid); i += down {
		j = (j + right) % lineLen
		if grid[i][j] == '#' {
			trees++
		}
	}

	return trees
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var grid []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	check(s.Err())

	a := travel(grid, 1, 1)
	b := travel(grid, 3, 1)
	c := travel(grid, 5, 1)
	d := travel(grid, 7, 1)
	e := travel(grid, 1, 2)

	fmt.Println("a: 1 right 1 down:", a)
	fmt.Println("b: 3 right 1 down:", b)
	fmt.Println("c: 5 right 1 down:", c)
	fmt.Println("d: 7 right 1 down:", d)
	fmt.Println("e: 1 right 2 down:", e)

	fmt.Println("a*b*c*d*e =", a*b*c*d*e)
}
