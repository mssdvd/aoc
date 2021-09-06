package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(passwords []string) int {
	var valid int
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, password := range passwords {
		var matches int
		for _, field := range fields {
			if regexp.MustCompile(field).MatchString(password) {
				matches++
			}
		}
		if matches == 7 {
			valid++
		}
	}
	return valid
}

type Document struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid int
	cid int
}

func part2(passwords []string) int {
	var valid int
LoopPass:
	for _, password := range passwords {
		fields := strings.FieldsFunc(password, func(r rune) bool {
			if r == ':' || r == ' ' || r == '\n' {
				return true
			}
			return false
		})
		var matches int
		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]
			switch key {
			case "byr":
				year,err := strconv.Atoi(value)
				check(err)
				if year < 1920 || 2002 < year {
					continue LoopPass
				}
			case "iyr":
				year,err := strconv.Atoi(value)
				check(err)
				if year < 2010 || 2020 < year {
					continue LoopPass
				}
			case "eyr":
				year,err := strconv.Atoi(value)
				check(err)
				if year < 2020 || 2030 < year {
					continue LoopPass
				}
			case "hgt":
				height, err := strconv.Atoi(value[:len(value)-2])
				check(err)
				unit := value[len(value)-2:]

				var min int
				var max int
				if unit == "cm" {
					min = 150
					max = 193
				} else if unit == "in" {
					min = 59
					max = 76
				}
				if height < min || max < height {
					continue LoopPass
				}
			case "hcl":
				if !regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(value) {
					continue LoopPass
				}
			case "ecl":
				if !regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$").MatchString(value) {
					continue LoopPass
				}
			case "pid":
				if !regexp.MustCompile("^[0-9]{9}$").MatchString(value) {
					continue LoopPass
				}
			case "cid":
				matches--
			}
			matches++
		}
		if matches != 7 {
			continue
		}
		valid++
	}
	return valid
}

func main() {
	f, err := ioutil.ReadFile("./input")
	check(err)

	passwords := regexp.MustCompile("(?m)"+"^$").Split(string(f), -1)

	fmt.Println("Part 1:", part1(passwords))
	fmt.Println("Part 2:", part2(passwords))
}
