package main

import (
	"bufio"
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

type bagRequirement struct {
	quantity int
	bag      string
}

func countBags(rules map[string][]bagRequirement, bag string) int {
	if len(rules[bag]) == 0 {
		return 1
	}

	n := 1
	for _, v := range rules[bag] {
		n += v.quantity * countBags(rules, v.bag)
	}

	return n
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	rules := make(map[string][]bagRequirement)
	containShinyGold := make(map[string]struct{})

	s := bufio.NewScanner(f)
	var sb strings.Builder
	for s.Scan() {
		ruleFields := strings.Fields(s.Text())

		sb.WriteString(ruleFields[0])
		sb.WriteString(" ")
		sb.WriteString(ruleFields[1])
		bag := sb.String()
		sb.Reset()

		if ruleFields[4] == "no" {
			rules[bag] = []bagRequirement{}
			continue
		}

		for i := 4; i < len(ruleFields); i += 4 {
			quantity, err := strconv.Atoi(ruleFields[i])
			check(err)

			sb.WriteString(ruleFields[i+1])
			sb.WriteString(" ")
			sb.WriteString(ruleFields[i+2])
			requiredBag := sb.String()
			sb.Reset()

			if requiredBag == "shiny gold" {
				containShinyGold[bag] = struct{}{}
			}

			rules[bag] = append(rules[bag], bagRequirement{quantity, requiredBag})
		}
	}
	check(s.Err())

	var prev int
	// naïve
	for {
		// exit if nothing is changed
		if prev == len(containShinyGold) {
			break
		}
		prev = len(containShinyGold)

		for bag, req := range rules {
			for _, v := range req {
				for c := range containShinyGold {
					if v.bag == c {
						containShinyGold[bag] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println("Number of bag colors that can contain at least one shiny gold bag:",
		len(containShinyGold)) // output for debug

	// remove shiny gold from the count
	fmt.Println("Number of bags required inside shiny gold:", countBags(rules, "shiny gold")-1)
}
