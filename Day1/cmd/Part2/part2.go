package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberMap map[string]string = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func main() {
	fmt.Println("Day 1 part 2")
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pass a file to parse as the day 1 input.")
		return
	}
	filepath := args[0]
	fmt.Printf("Reading %s\n", filepath)
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read file %s, err: %s\n", filepath, err)
		return
	}
	lines := strings.Split(string(data), "\n")
	runningCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		firstMatch := findFirstDigit(line)
		lastMatch := findLastDigit(line)
		fmt.Printf(
			"%s %s + %s\n",
			line,
			firstMatch,
			lastMatch,
		)
		stringValue := firstMatch + lastMatch
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			fmt.Printf(
				"Failed to turn %s into a number err: %s\n",
				stringValue,
				err,
			)
			continue
		}
		runningCount += value
	}
	fmt.Printf("Final count is %d\n", runningCount)
}

func findFirstDigit(line string) string {
	first := len(line)
	res := ""
	for symbol := range numberMap {
		if index := strings.Index(line, symbol); index < first && index != -1 {
			first = index
			res = numberMap[symbol]
		}
	}
	return res
}

func findLastDigit(line string) string {
	last := -1
	res := ""
	for symbol := range numberMap {
		if index := strings.LastIndex(line, symbol); index > last && index != -1 {
			last = index
			res = numberMap[symbol]
		}
	}
	return res
}
