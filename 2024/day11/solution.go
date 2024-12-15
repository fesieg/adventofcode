package main

import (
	adventofcode "adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type StoneMap map[int]int

func main() {
	input := adventofcode.GetFileContent("./input.txt")

	stones := parseInput(string(input))

	fmt.Println("PART 1:", sumValues(blinkMapNTimes(stones, 25)))
	fmt.Println("PART 2:", sumValues(blinkMapNTimes(stones, 75)))
}

func parseInput(input string) []int {
	parts := strings.Fields(input)
	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		result[i] = num
	}
	return result
}

func getChangedStones(stone int) []int {
	stoneStr := strconv.Itoa(stone)

	if stoneStr == "0" {
		return []int{1}
	}

	length := len(stoneStr)
	if length%2 == 0 {
		mid := length / 2
		firstHalf := removeLeadingZeroes(stoneStr[:mid])
		secondHalf := removeLeadingZeroes(stoneStr[mid:])

		firstNum, _ := strconv.Atoi(firstHalf)
		secondNum, _ := strconv.Atoi(secondHalf)
		return []int{firstNum, secondNum}
	}

	stoneInt := stone * 2024
	return []int{stoneInt}
}

func blinkMapNTimes(inputStones []int, blinks int) StoneMap {
	stoneMap := make(StoneMap)

	for _, stone := range inputStones {
		stoneMap[stone]++
	}

	for blinks > 0 {
		newMap := make(StoneMap)
		for stone, count := range stoneMap {
			changedStones := getChangedStones(stone)
			for _, newStone := range changedStones {
				newMap[newStone] += count
			}
		}
		stoneMap = newMap
		blinks--
	}

	return stoneMap
}

func removeLeadingZeroes(input string) string {
	i := 0
	for i < len(input) && input[i] == '0' {
		i++
	}
	return input[i:]
}

func sumValues(stoneMap StoneMap) int {
	sum := 0
	for _, count := range stoneMap {
		sum += count
	}
	return sum
}
