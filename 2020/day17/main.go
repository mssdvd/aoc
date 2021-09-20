package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid3D struct {
	m       [][][]bool
	w, h, l int
}

type Grid4D struct {
	m          [][][][]bool
	w, h, l, t int
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var input [][]bool
	s := bufio.NewScanner(f)
	for s.Scan() {
		var line []bool
		for _, el := range s.Text() {
			if el == '#' {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
		}
		input = append(input, line)
	}
	check(s.Err())

	// 20x20x20 is enough "infinite" for me
	g := newGrid3D(20, 20, 20)
	g.firstGen3D(input)

	var n int
	for i := 0; i < 6; i++ {
		n = g.newGen3D()
	}

	fmt.Println("3D:", n)

	h := newGrid4D(20, 20, 20, 20)
	h.firstGen4D(input)

	n = 0
	for i := 0; i < 6; i++ {
		n = h.newGen4D()
	}

	fmt.Println("4D:", n)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func newGrid3D(w, h, l int) *Grid3D {
	m := make([][][]bool, h)
	for i := range m {
		m[i] = make([][]bool, w)
		for j := range m[i] {
			m[i][j] = make([]bool, l)
		}
	}
	return &Grid3D{m, w, h, l}
}

func (g *Grid3D) firstGen3D(input [][]bool) {
	pivotGrid := struct{ x, y int }{g.w / 2, g.h / 2}
	pivotInput := struct{ x, y int }{len(input[0]) / 2, len(input) / 2}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] {
				g.m[pivotGrid.y-pivotInput.y+y][pivotGrid.x-pivotInput.x+x][0] = true
			}
		}
	}
}

func (g *Grid3D) newGen3D() int {
	new := newGrid3D(g.w, g.h, g.l)
	var n int

	for y, plane := range g.m {
		for x, line := range plane {
			for z, cube := range line {
				numNeigh := g.numOfNeighbours3D(x, y, z)
				if (cube && (numNeigh == 2 || numNeigh == 3)) || (!cube && numNeigh == 3) {
					new.m[y][x][z] = true
					n++
				}
			}
		}
	}
	g.m = new.m
	return n
}

func (g *Grid3D) numOfNeighbours3D(x, y, z int) int {
	var neighY, neighX, neighZ, n int
	for a := -1; a < 2; a++ {
		for b := -1; b < 2; b++ {
			for c := -1; c < 2; c++ {
				if a == 0 && b == 0 && c == 0 {
					continue
				}

				neighY = (y + a + g.h) % g.h
				neighX = (x + b + g.w) % g.w
				neighZ = (z + c + g.l) % g.l

				if g.m[neighY][neighX][neighZ] {
					n++
				}
			}
		}
	}
	return n
}

func newGrid4D(w, h, l, t int) *Grid4D {
	m := make([][][][]bool, h)
	for i := range m {
		m[i] = make([][][]bool, w)
		for j := range m[i] {
			m[i][j] = make([][]bool, l)
			for k := range m[i][j] {
				m[i][j][k] = make([]bool, t)
			}
		}
	}
	return &Grid4D{m, w, h, l, t}
}

func (g *Grid4D) firstGen4D(input [][]bool) {
	pivotGrid := struct{ x, y int }{g.w / 2, g.h / 2}
	pivotInput := struct{ x, y int }{len(input[0]) / 2, len(input) / 2}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] {
				g.m[pivotGrid.y-pivotInput.y+y][pivotGrid.x-pivotInput.x+x][0][0] = true
			}
		}
	}
}

func (g *Grid4D) newGen4D() int {
	new := newGrid4D(g.w, g.h, g.l, g.t)
	var n int

	for y, space := range g.m {
		for x, plane := range space {
			for z, line := range plane {
				for w, cube := range line {
					numNeigh := g.numOfNeighbours4D(x, y, z, w)
					if (cube && (numNeigh == 2 || numNeigh == 3)) || (!cube && numNeigh == 3) {
						new.m[y][x][z][w] = true
						n++
					}

				}
			}
		}
	}
	g.m = new.m
	return n
}

func (g *Grid4D) numOfNeighbours4D(x, y, z, w int) int {
	var n int
	for a := -1; a < 2; a++ {
		for b := -1; b < 2; b++ {
			for c := -1; c < 2; c++ {
				for d := -1; d < 2; d++ {
					if a == 0 && b == 0 && c == 0 && d == 0 {
						continue
					}

					neighY := (y + a + g.h) % g.h
					neighX := (x + b + g.w) % g.w
					neighZ := (z + c + g.l) % g.l
					neighW := (w + d + g.t) % g.t

					if g.m[neighY][neighX][neighZ][neighW] {
						n++
					}
				}
			}
		}
	}
	return n
}
