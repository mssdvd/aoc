package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var grid [][]rune
	s := bufio.NewScanner(f)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}
	check(s.Err())

	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	var occupied, prev int
	for {
		occupied += simulateSeats(&grid, 4, true)
		if occupied == prev {
			break
		}
		prev = occupied
	}
	fmt.Println("Seats that end up occupied:", occupied)

	occupied = 0
	prev = occupied
	for {
		occupied += simulateSeats(&gridCopy, 5, false)
		if occupied == prev {
			break
		}
		prev = occupied
	}
	fmt.Println("Seats that end up occupied:", occupied)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func simulateSeats(grid *[][]rune, maxAdjSeats int, onlyAdj bool) (changes int) {
	tempGrid := make([][]rune, len(*grid))
	for i := range *grid {
		tempGrid[i] = make([]rune, len((*grid)[i]))
		copy(tempGrid[i], (*grid)[i])
	}
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[y]); x++ {
			switch (*grid)[y][x] {
			case 'L':
				if countAdjacentSeats(grid, y, x, onlyAdj) == 0 {
					tempGrid[y][x] = '#'
					changes++
				}
			case '#':
				if countAdjacentSeats(grid, y, x, onlyAdj) >= maxAdjSeats {
					tempGrid[y][x] = 'L'
					changes--
				}
			}
		}
	}
	*grid = tempGrid
	return
}

func countAdjacentSeats(grid *[][]rune, y int, x int, onlyAdj bool) (occupied int) {
	var seatY, seatX int
	for a := -1; a < 2; a++ {
	xLoop:
		for b := -1; b < 2; b++ {
			seatY = y
			seatX = x
			for {
				seatY += a
				seatX += b
				if (a == 0 && b == 0) || seatX < 0 || seatX >= len((*grid)[0]) || seatY < 0 || seatY >= len(*grid) {
					continue xLoop
				}
				seat := (*grid)[seatY][seatX]
				if seat == '#' {
					occupied++
				}
				if onlyAdj || seat != '.' {
					break
				}
			}
		}
	}
	return
}
