package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	lines := readFileLines("input/day2.txt")
// 	part1Result, err := day2part1(lines)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Part 1: %d\n", part1Result)

// 	part2Result, err := day2part2(lines)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Part 2: %d\n", part2Result)
// }

func day2part1(lines []string) (int, error) {
	reports, err := day2parseLines(lines)
	if err != nil {
		return 0, err
	}

	safeCount := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
		}
	}
	return safeCount, nil
}

func day2part2(lines []string) (int, error) {
	reports, err := day2parseLines(lines)
	if err != nil {
		return 0, err
	}

	// For each report r, generate a set of new reports (r') each with one element removed from r
	// If any r' is safe, then r is safe
	safeCount := 0
	for _, report := range reports {
		// Capacity is arbitary, calculating the capacity properly would require iterating through the reports
		updatedReports := make([][]int, 0, len(reports)*10)
		for i := range report {
			updatedReport := make([]int, 0, len(report)-1)
			for j, level := range report {
				if i != j {
					updatedReport = append(updatedReport, level)
				}
			}
			updatedReports = append(updatedReports, updatedReport)
		}

		// If any of the modified reports is safe, then the original report is safe
		isSafe := false
		for _, updatedReport := range updatedReports {
			if isReportSafe(updatedReport) {
				isSafe = true
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount, nil
}

func day2parseLines(lines []string) ([][]int, error) {
	reports := make([][]int, 0, len(lines))
	for i, line := range lines {
		tokens := strings.Fields(line)

		if len(tokens) < 2 {
			return reports, fmt.Errorf("report %d contains less than 2 arguments", i)
		}

		report := make([]int, 0, len(tokens))
		for j, token := range tokens {
			level, err := strconv.Atoi(token)
			if err != nil {
				return reports, fmt.Errorf("invalid integer on line %d, token %d: %s", i, j, token)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func isReportSafe(report []int) bool {
	var prevDiff int
	for i := 1; i < len(report); i++ {
		currentLevel := report[i]
		prevLevel := report[i-1]

		currentDiff := currentLevel - prevLevel

		if currentDiff == 0 || currentDiff > 3 || currentDiff < -3 || (i > 1 && (currentDiff > 0) != (prevDiff > 0)) {
			return false
		}

		prevDiff = currentDiff
	}
	return true
}
