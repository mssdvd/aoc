package main

import (
	"bufio"
	"fmt"
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

	var sum int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var stack []int
		var opStack []rune
		// Shunting-yard algorithm
		for _, el := range s.Text() {
			val, err := strconv.Atoi(string(el))
			if err == nil {
				stack = append(stack, val)
				continue
			}
			switch el {
			case '*', '+':
				for checkOpenParen(opStack) && checkPrecedence(el, opStack[len(opStack)-1]) {
					calc(&stack, &opStack)
				}
				opStack = append(opStack, el)
			case '(':
				opStack = append(opStack, el)
			case ')':
				for checkOpenParen(opStack) {
					calc(&stack, &opStack)
				}
				opStack = opStack[:len(opStack)-1]
			case ' ':
				continue
			}
		}
		for checkOpenParen(opStack) {
			calc(&stack, &opStack)
		}
		sum += stack[0]
	}
	check(s.Err())

	fmt.Printf("sum: %v\n", sum)
}

// checkPrecedence returns true if b has greater precedence than a or if they are equal
func checkPrecedence(a, b rune) bool {
	return a == '*' && b == '+' || a == b
}

func checkOpenParen(opStack []rune) bool {
	return len(opStack) > 0 && opStack[len(opStack)-1] != '('
}

func calc(stack *[]int, opStack *[]rune) {
	n := len(*stack)
	var op rune
	op, *opStack = (*opStack)[len(*opStack)-1], (*opStack)[:len(*opStack)-1]
	switch op {
	case '*':
		(*stack)[n-2] = (*stack)[n-2] * (*stack)[n-1]
	case '+':
		(*stack)[n-2] = (*stack)[n-2] + (*stack)[n-1]
	}
	*stack = (*stack)[:n-1]
}
