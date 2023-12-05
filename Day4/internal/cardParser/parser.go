package cardParser

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	ID             int
	WinningNumbers []int
	CardNumbers    []int
}

func ParseCardFile(filePath string) ([]Card, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cards := string(data)
	var parsedCards []Card
	for _, line := range strings.Split(cards, "\n") {
		if line == "" {
			continue
		}
		card, err := parseCard(line)
		if err != nil {
			fmt.Printf("Failed to parse input %s\n", err)
			return nil, err
		}
		card.ID = len(parsedCards) + 1
		parsedCards = append(parsedCards, card)
	}
	return parsedCards, nil
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
		card.WinningNumbers = append(card.WinningNumbers, number)
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
		card.CardNumbers = append(card.CardNumbers, number)
	}
	return card, joinedErr
}
