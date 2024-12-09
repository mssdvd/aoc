package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Field [][]byte

var (
	sum1 = make(map[Point]struct{})
	sum2 = make(map[Point]struct{})
)

func main() {
	fieldSlice, err := os.ReadFile("./input")
	check(err)

	var field1 Field = bytes.Split(fieldSlice, []byte("\n"))
	field1 = field1[:len(field1)-1]

	field2 := make(Field, len(field1))

	freqs := make(map[byte][]Point)

	for y, row := range field1 {
		field2[y] = append(field2[y], row...)

		for x, b := range row {
			if b != '.' {
				freqs[b] = append(freqs[b], Point{x, y})
			}
		}
	}

	for _, locs := range freqs {
		for i := 0; i < len(locs); i++ {
			for j := 0; j < i; j++ {
				antinodes(field1, field2, locs[i], locs[j])
			}
			for j := i + 1; j < len(locs); j++ {
				antinodes(field1, field2, locs[i], locs[j])
			}
		}
	}

	fmt.Println("Part One:", len(sum1))
	fmt.Println("Part Two", len(sum2))
}

func antinodes(field1, field2 Field, locA, locB Point) {
	diff := Point{locA.x - locB.x, locA.y - locB.y}

	antinodes := []Point{{locA.x + diff.x, locA.y + diff.y}, {locB.x + diff.x, locB.y + diff.y}}

	for _, an := range antinodes {
		if an.y >= 0 && an.x >= 0 && an.y < len(field1) && an.x < len(field1[0]) && an != locA && an != locB {
			f := &field1[an.y][an.x]
			if *f == '.' {
				*f = '#'
			}
			sum1[an] = struct{}{}
		}
	}

	addNewAntinodes := func(diffY, diffX int) {
		y := antinodes[1].y
		x := antinodes[1].x
		for y >= 0 && x >= 0 && y < len(field2) && x < len(field2[0]) {
			f := &field2[y][x]
			if *f == '.' {
				*f = '#'
			}
			sum2[Point{x, y}] = struct{}{}

			y += diffY
			x += diffX
		}
	}

	addNewAntinodes(diff.y, diff.x)
	addNewAntinodes(-diff.y, -diff.x)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f Field) String() string {
	var b strings.Builder

	for _, row := range f {
		b.WriteString(string(row))
		b.WriteString("\n")
	}
	return b.String()
}
