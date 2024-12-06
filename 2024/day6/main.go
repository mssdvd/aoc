package main

import (
	"bytes"
	"fmt"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Guard struct {
	x, y int
}

func main() {
	labSlice, err := os.ReadFile("./input_example")
	check(err)

	lab := bytes.Split(labSlice, []byte("\n"))
	lab = lab[:len(lab)-1]

	var guard Guard
	dir := North
	nrows := len(lab)
	ncols := len(lab[0])

	var row []byte
	for guard.y, row = range lab {
		if guard.x = bytes.IndexByte(row, '^'); guard.x != -1 {
			break
		}
	}

	visited := make(map[Guard]Direction)

	for {
		visited[guard] = dir

		nextGuard := next(guard, dir)

		if nextGuard.x < 0 || nextGuard.x == ncols || nextGuard.y < 0 || nextGuard.y == nrows {
			break
		}

		if lab[nextGuard.y][nextGuard.x] == '#' {
			dir = (dir + 1) % 4
			continue
		}
		guard = nextGuard

	}

	fmt.Println("Part One:", len(visited))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func next(g Guard, dir Direction) Guard {
	switch dir {
	case North:
		return Guard{g.x, g.y - 1}
	case East:
		return Guard{g.x + 1, g.y}
	case South:
		return Guard{g.x, g.y + 1}
	case West:
		return Guard{g.x - 1, g.y}
	default:
		return Guard{g.x, g.y}
	}
}
