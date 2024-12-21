package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readFileLines("input/day3.txt")
	part1Result, err := day3part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", part1Result)

	part2Result, err := day3part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", part2Result)
}

func day3part1(lines []string) (int, error) {
	input := strings.Join(lines, "")

	r := regexp.MustCompile(`mul\((?P<operand1>\d{1,3}),(?P<operand2>\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	sum := 0
	for i, submatches := range matches {
		operand1String := submatches[r.SubexpIndex("operand1")]
		operand2String := submatches[r.SubexpIndex("operand2")]

		operand1, err := strconv.Atoi(operand1String)
		if err != nil {
			return sum, fmt.Errorf("invalid integer on line %d, 1st argument: %s", i, operand1String)
		}

		operand2, err := strconv.Atoi(operand2String)
		if err != nil {
			return sum, fmt.Errorf("invalid integer on line %d, 2nd argument: %s", i, operand2String)
		}

		sum += operand1 * operand2
	}
	return sum, nil
}

func day3part2(lines []string) (int, error) {
	input := strings.Join(lines, "")

	r := regexp.MustCompile(`(?P<do_operator>do)\(\)|(?P<dont_operator>don't)\(\)|(?P<mul_operator>mul)\((?P<operand1>\d{1,3}),(?P<operand2>\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	sum := 0
	enabled := true
	for i, submatches := range matches {
		if submatches[r.SubexpIndex("do_operator")] != "" {
			enabled = true
		} else if submatches[r.SubexpIndex("dont_operator")] != "" {
			enabled = false
		} else if submatches[r.SubexpIndex("mul_operator")] != "" && enabled {
			operand1String := submatches[r.SubexpIndex("operand1")]
			operand2String := submatches[r.SubexpIndex("operand2")]

			operand1, err := strconv.Atoi(operand1String)
			if err != nil {
				return sum, fmt.Errorf("invalid integer on line %d, 1st argument: %s", i, operand1String)
			}

			operand2, err := strconv.Atoi(operand2String)
			if err != nil {
				return sum, fmt.Errorf("invalid integer on line %d, 2nd argument: %s", i, operand2String)
			}

			sum += operand1 * operand2
		}
	}
	return sum, nil
}
