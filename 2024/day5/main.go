package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	before string
	after  string
}

type (
	Rules   map[Rule]struct{}
	Update  []string
	Updates []Update
)

func main() {
	safetyManualBytes, err := os.ReadFile("./input")
	check(err)

	safetyManual := strings.Split(string(safetyManualBytes), "\n")

	// Set
	rules := make(Rules)

	var i int
	var rule string
	for i, rule = range safetyManual {
		if rule == "" {
			i++
			break
		}

		pair := strings.Split(rule, "|")
		rules[Rule{pair[0], pair[1]}] = struct{}{}
	}

	var updates Updates

	for i := i; i < len(safetyManual); i++ {
		if safetyManual[i] == "" {
			break
		}

		update := strings.Split(safetyManual[i], ",")
		updates = append(updates, update)
	}

	var sum1 int
	var sum2 int

	for _, update := range updates {
		valid, invalidRule := isValid(rules, update)

		if valid {
			mid, err := strconv.Atoi(update[len(update)/2])
			check(err)
			sum1 += mid
		} else {
			var before, after int
			for !valid {
				before = slices.Index(update, invalidRule.before)
				after = slices.Index(update, invalidRule.after)

				update[before], update[after] = update[after], update[before]

				valid, invalidRule = isValid(rules, update)
			}
			mid, err := strconv.Atoi(update[len(update)/2])
			check(err)
			sum2 += mid
		}
	}

	fmt.Println("Part One:", sum1)
	fmt.Println("Part Two:", sum2)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isValid(rules Rules, update Update) (bool, Rule) {
	for n := range update {
		page := update[n]

		for i := 0; i < n; i++ {
			rule := Rule{update[i], page}
			if _, ok := rules[rule]; !ok {
				return false, rule
			}
		}

		for i := n + 1; i < len(update); i++ {
			rule := Rule{page, update[i]}
			if _, ok := rules[rule]; !ok {
				return false, rule
			}
		}
	}

	return true, Rule{}
}
