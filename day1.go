package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// func main() {
// 	lines := readFileLines("input/day1.txt")
// 	part1Result, err := day1part1(lines)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Part 1: %d\n", part1Result)

// 	part2Result, err := day1part2(lines)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Part 2: %d\n", part2Result)
// }

func day1part1(lines []string) (int, error) {
	list1, list2, err := parseLists(lines)
	if err != nil {
		return -1, err
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var sum int
	for i := range list1 {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return sum, nil
}

func day1part2(lines []string) (int, error) {
	list1, list2, err := parseLists(lines)
	if err != nil {
		return -1, err
	}

	list2CountMap := make(map[int]int, len(list2))
	for _, v := range list2 {
		list2CountMap[v]++
	}

	similarityScore := 0
	for _, v := range list1 {
		similarityScore += v * list2CountMap[v]
	}

	return similarityScore, nil
}

func parseLists(lines []string) ([]int, []int, error) {
	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))
	for i, line := range lines {
		tokens := strings.Fields(line)

		if len(tokens) != 2 {
			return list1, list2, fmt.Errorf("incorrect number of arguments (expected 2, got %d)", len(tokens))
		}

		value1, err := strconv.Atoi(tokens[0])
		if err != nil {
			return list1, list2, fmt.Errorf("invalid integer on line %d: %s", i, tokens[0])
		}

		value2, err := strconv.Atoi(tokens[1])
		if err != nil {
			return list1, list2, fmt.Errorf("invalid integer on line %d: %s", i, tokens[1])
		}

		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	return list1, list2, nil
}
