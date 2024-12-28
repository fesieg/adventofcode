package main

import (
	adventofcode "adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Bot struct {
	pos  Position
	xVel int
	yVel int
}

func main() {
	lines := getContentAsListOfStrings("./input.txt")

	var bots []Bot

	// for the simple input
	// xLimit, yLimit := 11, 7

	xLimit, yLimit := 101, 103

	for _, line := range lines {
		botX, botY, botXVel, botYVel := getBotFromLine(line)
		bots = append(bots, Bot{Position{botX, botY}, botXVel, botYVel})
	}

	for move := 0; move < 100; move++ {
		for i := range bots {
			moveBot(&bots[i], xLimit, yLimit)
		}
	}

	q1, q2, q3, q4 := countBotsInQuadrants(bots, xLimit, yLimit)
	fmt.Println("PART 1: ", q1*q2*q3*q4)
}

// DEBUG
// func printGrid(bots []Bot, xLimit, yLimit int) {
// 	for y := 0; y < yLimit; y++ {
// 		for x := 0; x < xLimit; x++ {
// 			fmt.Print(countBotPositions(bots, xLimit, yLimit)[x][y], " ")
// 		}
// 	}
// }

func moveBot(bot *Bot, xlim, ylim int) {
	newX := bot.pos.x + bot.xVel

	// if new position is out of bounds, teleport to other side
	if newX < 0 {
		newX = newX + xlim
	} else if newX >= xlim {
		newX = newX - xlim
	}

	newY := bot.pos.y + bot.yVel

	if newY < 0 {
		newY = newY + ylim
	} else if newY >= ylim {
		newY = newY - ylim
	}

	bot.pos = Position{newX, newY}
}

func countBotsInQuadrants(bots []Bot, xLimit, yLimit int) (int, int, int, int) {
	midX, midY := xLimit/2, yLimit/2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, bot := range bots {
		x, y := bot.pos.x, bot.pos.y

		// skip robots on the middle lines
		if x == midX || y == midY {
			continue
		}

		// determine quadrant
		if x > midX && y < midY {
			q1++
		} else if x < midX && y < midY {
			q2++
		} else if x < midX && y > midY {
			q3++
		} else if x > midX && y > midY {
			q4++
		}
	}

	return q1, q2, q3, q4
}

func getBotFromLine(line string) (int, int, int, int) {
	// match numbers (optional negative sign, followed by digits)
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers[0], numbers[1], numbers[2], numbers[3]
}

func getContentAsListOfStrings(in string) []string {
	content := adventofcode.GetFileContent(in)

	// windows line endings
	return strings.Split(content, "\r\n")
}
