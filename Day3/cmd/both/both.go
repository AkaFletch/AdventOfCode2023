package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("Day 3 Part 1 and 2")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Provide input file path as command line arg")
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("Failed to read file %s err: %s", args[0], err)
	}
	engineSchematic := string(data)
	symbolRegex := regexp.MustCompile(`[^\d.\n]`)
	symbolIndexes := symbolRegex.FindAllStringIndex(engineSchematic, -1)
	part1Sum := 0
	part2Sum := 0
	for _, symbol := range symbolIndexes {
		indexes := findNumbersAround(engineSchematic, symbol[0])
		numbers, err := expandNumbers(engineSchematic, indexes)
		if err != nil {
			fmt.Printf("Failed to parse, err: %s", err)
		}
		for _, partNumber := range numbers {
			part1Sum += partNumber
		}
		if engineSchematic[symbol[0]] == '*' && len(numbers) == 2 {
			part2Sum += numbers[0] * numbers[1]
		}
	}
	fmt.Printf("Part 1 count is %d\n", part1Sum)
	fmt.Printf("Part 2 count is %d\n", part2Sum)
}

func findNumbersAround(engineSchematic string, index int) []int {
	// +1 since we have a \n in our string
	lineLength := len(strings.Split(engineSchematic, "\n")[0]) + 1
	var foundNumbers []int
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			lengthCheck := index+x+y*lineLength > len(engineSchematic)
			skipSymbol := x == 0 && y == 0
			numberCheck := !unicode.IsNumber(rune(engineSchematic[index+x+y*lineLength]))
			if lengthCheck || skipSymbol || numberCheck {
				continue
			}
			foundNumbers = append(foundNumbers, index+x+y*lineLength)
		}
	}
	return foundNumbers
}

func expandNumbers(engineSchematic string, indexes []int) ([]int, error) {
	var alreadyFilled []int
	var foundNumbers []int
	for _, index := range indexes {
		if slices.Contains(alreadyFilled, index) {
			continue
		}
		builtString := string(engineSchematic[index])
		// We need a greater than 0 check here to stop us going negative
		for i := -1; index+i >= 0 && unicode.IsDigit(rune(engineSchematic[index+i])); i-- {
			if slices.Contains(indexes, index+i) {
				alreadyFilled = append(alreadyFilled, index+i)
			}
			builtString = string(engineSchematic[index+i]) + builtString
		}
		// This can't go over the string length as \n will catch it
		for i := 1; unicode.IsDigit(rune(engineSchematic[index+i])); i++ {
			if slices.Contains(indexes, index+i) {
				alreadyFilled = append(alreadyFilled, index+i)
			}
			builtString = builtString + string(engineSchematic[index+i])
		}
		foundNumber, err := strconv.Atoi(builtString)
		if err != nil {
			return nil, err
		}
		foundNumbers = append(foundNumbers, foundNumber)
	}
	return foundNumbers, nil
}
