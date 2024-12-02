package main

import (
	adventofcode "adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getReportsFromInput("input")

	numberOfSafeReports := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			numberOfSafeReports += 1
		}
	}

	fmt.Println("Number of safe reports (part 1): ", numberOfSafeReports)

	numberOfCorrectedSafeReports := 0

	for _, report := range reports {
		if !reportIsSafe(report) {
			for i := 0; i < len(report); i++ {

				// create new slice and copy all elements except the one at index i
				newReport := make([]int, len(report)-1)
				copy(newReport, report[:i])
				copy(newReport[i:], report[i+1:])

				if reportIsSafe(newReport) {
					numberOfCorrectedSafeReports += 1
					break
				}
			}
		} else {
			numberOfCorrectedSafeReports += 1
		}
	}

	fmt.Println("Number of corrected safe reports (part 2): ", numberOfCorrectedSafeReports)
}

func allAscending(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] >= report[i] {
			return false
		}
	}
	return true
}

func allDescending(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] <= report[i] {
			return false
		}
	}
	return true
}

func reportIsSafe(report []int) bool {
	if !allAscending(report) && !allDescending(report) {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		if math.Abs(float64(report[i])-float64(report[i+1])) > 3 {
			return false
		}
	}

	return true
}

func getReportsFromInput(in string) [][]int {
	content := adventofcode.GetFileContent(in)

	reports := [][]int{}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		report := []int{}
		for _, id := range strings.Split(line, " ") {
			convertedID, err := strconv.Atoi(id)
			if err == nil {
				report = append(report, convertedID)
			}
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return reports
}
