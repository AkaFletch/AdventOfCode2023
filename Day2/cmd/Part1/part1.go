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
	fmt.Println("Part 1")
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
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		game, err := parseGame(line)
		if err != nil {
			fmt.Printf("Error when parsing game %s\n", err)
		}
		if game.MaxRed > 12 || game.MaxGreen > 13 || game.MaxBlue > 14 {
			continue
		}
		count += game.GameID
	}
	fmt.Printf("Final count %d\n", count)
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
