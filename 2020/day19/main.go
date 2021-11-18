// https://adventofcode.com/2020/day/19
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()

	rules := make(map[string]string)

	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Text()) == 0 {
			break
		}

		split := strings.Split(s.Text(), ":")
		var sb strings.Builder
		for _, c := range strings.Fields(split[1]) {
			if c == "|" {
				sb.WriteString(c)
				continue
			}
			sb.WriteString("(")
			sb.WriteString(c)
			sb.WriteString(")")
		}

		rules[split[0]] = sb.String()
	}

	rule := strings.Join([]string{"^", rules["0"], "$"}, "")
	delete(rules, "0")

	rx, err := regexp.Compile("[0-9]+")
	check(err)
	for rx.MatchString(rule) {
		for key, value := range rules {
			rule = strings.ReplaceAll(rule, strings.Join([]string{"(", key, ")"}, ""), strings.Join([]string{"(", value, ")"}, ""))
		}
	}

	rule = strings.ReplaceAll(rule, "\"", "")

	r, err := regexp.Compile(rule)
	check(err)

	var count int
	for s.Scan() {
		if r.MatchString(s.Text()) {
			count++
		}
	}

	fmt.Println("count: ", count)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
