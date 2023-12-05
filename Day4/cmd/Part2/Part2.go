package main

import (
	"fmt"
	"os"
	"slices"

	parser "Day4/v2/internal/cardParser"
)

func main() {
	fmt.Println("Day 4 Part 2")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pass the input as a command line parameter.")
		return
	}
	parsedCards, err := parser.ParseCardFile(args[0])
	if err != nil {
		fmt.Printf("Failed to read file, err: %s\n", err)
		return
	}
	total, copies := 0, make([]int, len(parsedCards))
	for index := 0; index < len(parsedCards); index++ {
		card, count := parsedCards[index], 0
		for _, guess := range card.CardNumbers {
			if !slices.Contains(card.WinningNumbers, guess) {
				continue
			}
			count++
		}
		for i := 1; i < count+1; i++ {
			if index+i >= len(parsedCards) {
				continue
			}
			copies[index+i] += copies[index] + 1
		}
		total += 1 + copies[index]
	}
	fmt.Printf("Total %d\n", total)
}
