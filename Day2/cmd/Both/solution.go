package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	GameID   int
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}

func main() {
	fmt.Println("Part 1 and 2")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No filepath provided")
		return
	}
	filepath := args[0]
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read file %s err %s\n", filepath, err)
		return
	}
	lines := strings.Split(string(data), "\n")
	part1Count := 0
	part2Count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		game, err := parseGame(line)
		if err != nil {
			fmt.Printf("Error when parsing game %s\n", err)
		}
		part2Count += game.MaxRed * game.MaxBlue * game.MaxGreen
		if game.MaxRed > 12 || game.MaxGreen > 13 || game.MaxBlue > 14 {
			continue
		}
		part1Count += game.GameID
	}
	fmt.Printf("Part 1 count %d\n", part1Count)
	fmt.Printf("Part 2 count %d\n", part2Count)
}

func parseGame(line string) (Game, error) {
	game := Game{}
	gameRegex := regexp.MustCompile(`Game (\d+)`)
	revealRegex := regexp.MustCompile(`(?:(\d+) (green|red|blue))`)

	gameNumber := gameRegex.FindStringSubmatch(line)[1]
	gameNumberInt, err := strconv.Atoi(gameNumber)
	if err != nil {
		return game, err
	}
	game.GameID = gameNumberInt
	reveals := revealRegex.FindAllStringSubmatch(line, -1)
	for _, match := range reveals {
		colourInt := &game.MaxRed
		colourCount, err := strconv.Atoi(match[1])
		if err != nil {
			return game, err
		}
		switch match[2] {
		case "blue":
			colourInt = &game.MaxBlue
		case "green":
			colourInt = &game.MaxGreen
		}
		if colourCount > *colourInt {
			*colourInt = colourCount
		}
	}
	return game, nil
}
