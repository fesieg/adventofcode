package main

import (
	adventofcode "adventofcode/utils"
	"fmt"
	"strings"
)

type Region struct {
	plants     []Coordinate
	char       byte
	perimeters []Coordinate
}

type Coordinate struct {
	x int
	y int
}

func main() {
	garden := getGardenAsSliceOfStrings("./simpleInput.txt")
	regions := make([]Region, 0)
	currentRegion := Region{plants: []Coordinate{{0, 0}}, char: garden[0][0]}

	// determine regions
	for y := 0; y < len(garden); y++ {
		row := garden[y]
		for x := 1; x < len(row); x++ {
			char := row[x]
			// check in all four directions for same char

			if (char == currentRegion.char && !containsPlant(currentRegion.plants, Coordinate{x, y})) {
				currentRegion.plants = append(currentRegion.plants, Coordinate{x, y})
			} else {
				regions = append(regions, currentRegion)

				currentRegion = Region{plants: []Coordinate{{x, y}}, char: char}
			}
		}
	}

	for i := 0; i < len(regions); i++ {
		addPerimeter(&regions[i], garden)
	}

	fmt.Println("Regions: ", regions)

	// fmt.Println("PART 1:")
	// fmt.Println("PART 2:")
}

func addPerimeter(region *Region, garden []string) {
	for j := 0; j < len(region.plants); j++ {
		plant := region.plants[j]
		if plant.x > 0 {
			if garden[plant.y][plant.x-1] != region.char {
				region.perimeters = append(region.perimeters, Coordinate{plant.x - 1, plant.y})
			}
		}
		if plant.x < len(garden[0])-1 {
			if garden[plant.y][plant.x+1] != region.char {
				region.perimeters = append(region.perimeters, Coordinate{plant.x + 1, plant.y})
			}
		}
		if garden[plant.y-1][plant.x] != region.char {
			region.perimeters = append(region.perimeters, Coordinate{plant.x, plant.y - 1})
		}

		if plant.y < len(garden)-1 {
			if garden[plant.y+1][plant.x] != region.char {
				region.perimeters = append(region.perimeters, Coordinate{plant.x, plant.y + 1})
			}
		}
	}
}

func containsPlant(plants []Coordinate, plant Coordinate) bool {
	for i := 0; i < len(plants); i++ {
		if plants[i] == plant {
			return true
		}
	}
	return false
}

func getGardenAsSliceOfStrings(path string) []string {
	content := adventofcode.GetFileContent(path)

	// windows line endings
	return strings.Split(content, "\r\n")
}
