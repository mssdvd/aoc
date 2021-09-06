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

func decodeSeat(code string, lb int, ub int) int {
	for i := 0; i < len(code); i++ {
		if code[i] == 'F' || code[i] == 'L' {
			ub = lb + (ub-lb)/2
		} else {
			lb += (ub-lb)/2 + 1
		}
	}
	return lb
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var max int
	var takenSeats [127*8+7]bool
	s := bufio.NewScanner(f)
	for s.Scan() {
		code := s.Text()
		seatID := decodeSeat(code[:7], 0, 127) * 8
		seatID += decodeSeat(code[7:], 0, 7)
		takenSeats[seatID] = true
		if seatID > max {
			max = seatID
		}
	}
	check(s.Err())

	fmt.Println("Max seatID:", max)

	for i := 1; i < len(takenSeats)-1; i++ {
		if takenSeats[i-1] && !takenSeats[i] && takenSeats[i+1] {
			fmt.Println("My seat is: ", i)
			break
		}
	}
}
