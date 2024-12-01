package adventofcode

import (
	"log"
	"os"
)

func GetFileContent(path string) string {
	content, error := os.ReadFile(path)

	if error != nil {
		log.Fatal(error)
	}

	return string(content)
}

func CountOccurrences(slice []int, item int) int {
	count := 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}
