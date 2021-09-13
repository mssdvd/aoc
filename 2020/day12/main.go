package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var istrs []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		istrs = append(istrs, s.Text())
	}
	check(s.Err())

	x, y := withoutWaypoint(istrs, 0)

	fmt.Println("Manhattan distance:", math.Abs(float64(x))+math.Abs(float64(y)))

	x, y = waypoint(istrs, 10, 1)
	fmt.Println("Manhattan distance (w/ waypoint):", math.Abs(float64(x))+math.Abs(float64(y)))
}

func withoutWaypoint(istrs []string, dir float64) (x, y int) {
	for _, i := range istrs {
		a := i[0]
		v, err := strconv.Atoi(i[1:])
		check(err)

		switch a {
		case 'N':
			y += v
		case 'S':
			y -= v
		case 'E':
			x += v
		case 'W':
			x -= v
		case 'R':
			dir -= float64(v)
		case 'L':
			dir += float64(v)
		case 'F':
			x += v * int(math.Cos(dir*(math.Pi/180.0)))
			y += v * int(math.Sin(dir*(math.Pi/180.0)))
		}
	}
	return
}

func waypoint(istrs []string, waypointX, waypointY int) (x, y int) {
	wX, wY := waypointX, waypointY

	for _, i := range istrs {
		a := i[0]
		v, err := strconv.Atoi(i[1:])
		check(err)

		switch a {
		case 'N':
			wY += v
		case 'S':
			wY -= v
		case 'E':
			wX += v
		case 'W':
			wX -= v
		case 'R', 'L':
			if a == 'R' {
				v = -v
			}
			v := float64(v) * (math.Pi / 180.0)
			oldwX := wX
			wX = wX*int(math.Cos(v)) - wY*int(math.Sin(v))
			wY = wY*int(math.Cos(v)) + oldwX*int(math.Sin(v))
		case 'F':
			x += wX * v
			y += wY * v
		}
	}
	return
}
