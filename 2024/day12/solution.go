package main

import (
	adventofcode "adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	garden := getGardenAsSliceOfStrings("./simpleInput.txt")

	fmt.Println("Garden: ", garden)

	// iterate over garden (slice of strings)
	// for each character, check if there are any adjacent identical characters, check the same for those characters until can be found
	// then resume loop

	for y, row := range garden {
		for x, char := range row {
			fmt.Println((string).char)
		}
	}

	fmt.Println("PART 1:")
	fmt.Println("PART 2:")
}

func getGardenAsSliceOfStrings(in string) []string {
	content := adventofcode.GetFileContent(in)

	// windows line endings
	return strings.Split(content, "\r\n")
}
