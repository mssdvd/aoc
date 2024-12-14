package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Topo []string

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var topo Topo
	var trailheads []Point

	s := bufio.NewScanner(f)
	for s.Scan() {
		topo = append(topo, s.Text())

		for x, height := range s.Text() {
			if height == '0' {
				trailheads = append(trailheads, Point{x, len(topo) - 1})
			}
		}
	}
	check(s.Err())

	var score, rating int
	for _, th := range trailheads {
		s, r := topo.FindTrail(th)
		score += s
		rating += r
	}

	fmt.Println("Part One:", score)
	fmt.Println("Part Two:", rating)
}

func (t Topo) FindTrail(h Point) (int, int) {
	trailMaxLengths := make(map[Point]int)
	return t.findTrail(h, 0, trailMaxLengths)
}

func (t Topo) findTrail(h Point, length int, trailMaxLengths map[Point]int) (int, int) {
	if t[h.y][h.x] == '9' {
		if l, ok := trailMaxLengths[h]; !ok || length > l {
			trailMaxLengths[h] = length
			return 1, 1
		}
		return 0, 1
	}

	nextHeight := t[h.y][h.x] + 1
	var score, rating int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0 && j == 0) || (i+j != -1 && i+j != 1) {
				continue
			}
			newY, newX := h.y+i, h.x+j
			if newY >= 0 && newY < len(t) && newX >= 0 && newX < len(t[0]) {
				height := t[newY][newX]
				if height >= '0' && height <= '9' && height == nextHeight {
					s, r := t.findTrail(Point{newX, newY}, length+1, trailMaxLengths)
					score += s
					rating += r
				}
			}
		}
	}

	return score, rating
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
