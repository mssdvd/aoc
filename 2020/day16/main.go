package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

	type validRanges struct {
		inf int
		sup int
	}

	ranges := make(map[string][2]validRanges)

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		fields := strings.Split(l, ":")

		if fields[0] == "" {
			break
		}

		v := strings.FieldsFunc(fields[1], func(r rune) bool {
			if r == '-' || r == ' ' {
				return true
			}
			return false
		})

		var vInt []int
		for _, j := range v {
			if j == "or" {
				continue
			}
			val, err := strconv.Atoi(j)
			check(err)
			vInt = append(vInt, val)
		}

		ranges[fields[0]] = [2]validRanges{
			{vInt[0], vInt[1]},
			{vInt[2], vInt[3]},
		}

	}

	// first ticket is mine
	var tickets [][]int
	for s.Scan() {
		l := s.Text()
		if l == "your ticket:" || l == "nearby tickets:" || l == "" {
			continue
		}

		fields := strings.Split(l, ",")
		var ticket []int
		for _, f := range fields {
			i, err := strconv.Atoi(f)
			check(err)
			ticket = append(ticket, i)
		}
		tickets = append(tickets, ticket)
	}

	check(s.Err())

	var wg sync.WaitGroup
	var m sync.Mutex
	var errors int64
	var validTickets [][]int

	for _, ticket := range tickets[1:] {
		wg.Add(1)

		go func(ticket []int, m *sync.Mutex) {
			defer wg.Done()
			for _, field := range ticket {
				var valid bool
				for _, r := range ranges {
					for _, el := range r {
						if field >= el.inf && field <= el.sup {
							valid = true
						}
					}
				}
				if !valid {
					atomic.AddInt64(&errors, int64(field))
					return
				}
			}
			m.Lock()
			validTickets = append(validTickets, ticket)
			m.Unlock()
		}(ticket, &m)
	}
	wg.Wait()

	fmt.Printf("errors: %v\n", errors)

	correctFields := make(map[int]string)

	// ranges = map[string][2]validRanges{
	// 	"class": {
	// 		{0, 1},
	// 		{4, 19},
	// 	},
	// 	"row": {
	// 		{0, 5},
	// 		{8, 19},
	// 	},
	// 	"seat": {
	// 		{0, 13},
	// 		{16, 19},
	// 	},
	// }

	// validTickets = [][]int{{3, 9, 18}, {15, 1, 5}, {5, 14, 9}}

	// TODO: Some fields fall in many ranges
	for i := 0; i < len(validTickets[0]); i++ {
	loopRanges:
		for k, r := range ranges {
			for j := 0; j < len(validTickets); j++ {
				var valid bool
				for _, el := range r {
					if validTickets[j][i] >= el.inf && validTickets[j][i] <= el.sup {
						valid = true
					}
				}
				if !valid {
					continue loopRanges
				}
			}
			correctFields[i] = k
			delete(ranges, k)
			break
		}
	}

	fmt.Printf("correctFields: %v\n", len(correctFields))

	departureMul := 1

	for col, field := range correctFields {
		if strings.Contains(field, "departure") {
			departureMul *= tickets[0][col]
		}
	}

	fmt.Printf("departureMul: %v\n", departureMul)
}
