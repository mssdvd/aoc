package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var mask string
	mem1 := make(map[int]int)
	mem2 := make(map[int]int)

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()

		if l[:7] == "mask = " {
			mask = l[7:]
			continue
		}

		r := regexp.MustCompile(`\d+`).FindAllString(l, 2)

		addrInt, err := strconv.Atoi(r[0])
		check(err)

		valInt, err := strconv.Atoi(r[1])
		check(err)

		valB := []byte(fmt.Sprintf("%036b", valInt))
		addrB := []byte(fmt.Sprintf("%036b", addrInt))
		var XInd []int

		for i := 0; i < 36; i++ {
			if mask[i] == '1' {
				addrB[i] = '1'
			}
			if mask[i] == 'X' {
				XInd = append(XInd, i)
				continue
			}
			valB[i] = mask[i]
		}

		for i := 0; i < 1<<len(XInd); i++ {
			a := addrB
			for j, x := range XInd {
				b := (i & (1 << j))
				var c byte
				if b == 0 {
					c = '0'
				} else {
					c = '1'
				}
				a[x] = c
			}
			addr, err := strconv.ParseInt(string(a), 2, 0)
			check(err)
			mem2[int(addr)] = valInt
		}

		valInt64, err := strconv.ParseInt(string(valB), 2, 0)
		check(err)

		mem1[addrInt] = int(valInt64)
	}
	check(s.Err())

	var c int
	for _, m := range mem1 {
		c += m
	}
	fmt.Println("Sum (value mask):", c)

	c = 0
	for _, m := range mem2 {
		c += m
	}
	fmt.Println("Sum (address mask):", c)
}
