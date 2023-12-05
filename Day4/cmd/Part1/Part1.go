package main

import (
	"fmt"
	"os"

	parser "Day4/v2/internal/cardParser"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("Day 4 Part 1")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pass the input as a command line parameter.")
		return
	}
	total := 0
	cards, err := parser.ParseCardFile(args[0])
	if err != nil {
		fmt.Printf("Failed to parse card file %s\n", err)
		return
	}
	for _, card := range cards {
		winningCount := 0
		for _, guess := range card.CardNumbers {
			if !slices.Contains(card.WinningNumbers, guess) {
				continue
			}
			if winningCount == 0 {
				winningCount = 1
				continue
			}
			winningCount *= 2
		}
		total += winningCount
	}
	fmt.Printf("Total %d\n", total)
}
