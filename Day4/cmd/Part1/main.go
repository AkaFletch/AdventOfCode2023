package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Card struct {
	winningNumbers []int
	cardNumbers    []int
}

func main() {
	fmt.Println("Day 4")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pass the input as a command line parameter.")
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("Failed to read file, err: %s\n", err)
	}
	cards := string(data)
	total := 0
	for _, line := range strings.Split(cards, "\n") {
		if line == "" {
			continue
		}
		card, err := parseCard(line)
		if err != nil {
			fmt.Printf("Failed to parse input %s\n", err)
			return
		}
		winningCount := 0
		for _, guess := range card.cardNumbers {
			if !slices.Contains(card.winningNumbers, guess) {
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

func parseCard(line string) (Card, error) {
	splitCard := strings.Split(line, "|")
	lhs := splitCard[0]
	rhs := splitCard[1]
	winningNumbers := strings.Split(strings.Split(lhs, ":")[1], " ")
	var joinedErr error = nil
	card := Card{}
	for _, stringNumber := range winningNumbers {
		if stringNumber == "" {
			continue
		}
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			joinedErr = errors.Join(joinedErr, err)
		}
		card.winningNumbers = append(card.winningNumbers, number)
	}

	cardNumbers := strings.Split(rhs, " ")
	for _, stringNumber := range cardNumbers {
		if stringNumber == "" {
			continue
		}
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			joinedErr = errors.Join(joinedErr, err)
		}
		card.cardNumbers = append(card.cardNumbers, number)
	}
	return card, joinedErr
}
