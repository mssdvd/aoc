package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type instr struct {
	op  string
	arg int
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	var text []instr

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Fields(s.Text())

		currentInstr := instr{op: line[0]}

		arg, err := strconv.Atoi(line[1])
		check(err)
		currentInstr.arg = arg

		text = append(text, currentInstr)
	}
	check(s.Err())

	acc, _, _ := executeCode(text)
	fmt.Println("Value of the accumulator before loop:", acc)

	acc, i := findBadInstr(text)
	fmt.Println("Value of the accumulator after termination:", acc)
	fmt.Println("Bad instruction:", text[i])
}

func executeCode(text []instr) (acc int, i int, err error) {
	executed := make(map[int]bool)

	for {
		if i == len(text) {
			break
		}
		if executed[i] {
			return acc, i, errors.New("LOOP")
		}
		executed[i] = true
		switch text[i].op {
		case "acc":
			acc += text[i].arg
			i++
		case "jmp":
			i += text[i].arg
		case "nop":
			i++
		}
	}
	return acc, i, err
}

func findBadInstr(text []instr) (int, int) {
	// Array of nop and jmp
	var lineToTest []int
	for i, v := range text {
		if v.op == "nop" || v.op == "jmp" {
			lineToTest = append(lineToTest, i)
		}
	}
	var i int
	for _, v := range lineToTest {
		if text[v].op == "nop" {
			text[v].op = "jmp"
		} else {
			text[v].op = "nop"
		}
		acc, _, err := executeCode(text)
		if err == nil {
			fmt.Printf("Found after %d/%d iteractions\n", i, len(lineToTest))
			return acc, v
		}

		// Restore instruction
		if text[v].op == "nop" {
			text[v].op = "jmp"
		} else {
			text[v].op = "nop"
		}
		i++
	}
	return 0, -1
}
