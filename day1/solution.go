package main

import (
	adventofcode "adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sliceOne, sliceTwo := getSortedListsFromInput("input")

	partOne := partOne(sliceOne, sliceTwo)
	fmt.Println("Total distance (part 1): ", partOne)

	partTwo := partTwo(sliceOne, sliceTwo)
	fmt.Println("Similarity score (part 2): ", partTwo)
}

func partTwo(sliceOne, sliceTwo []int) int {
	similarityScore := 0

	for i := 0; i < len(sliceOne); i++ {
		countInSliceTwo := adventofcode.CountOccurrences(sliceTwo, sliceOne[i])
		if countInSliceTwo > 0 {
			similarityScore += countInSliceTwo * sliceOne[i]
		}
	}

	return similarityScore
}

func partOne(sliceOne, sliceTwo []int) int {
	totalDistance := 0

	for i := 0; i < len(sliceOne); i++ {
		totalDistance += int(math.Abs(float64(sliceOne[i] - sliceTwo[i])))
	}

	return totalDistance
}

func getSortedListsFromInput(in string) ([]int, []int) {
	content := adventofcode.GetFileContent(in)

	sliceOne := []int{}
	sliceTwo := []int{}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		// split along three spaces specifically
		ids := strings.Split(scanner.Text(), "   ")

		firstIdInt, err := strconv.Atoi(strings.Trim(ids[0], ""))
		if err == nil {
			sliceOne = append(sliceOne, firstIdInt)
		}

		secondIdInt, err := strconv.Atoi(strings.Trim(ids[1], " "))
		if err == nil {
			sliceTwo = append(sliceTwo, secondIdInt)
		}
	}

	slices.Sort(sliceOne)
	slices.Sort(sliceTwo)

	return sliceOne, sliceTwo
}
