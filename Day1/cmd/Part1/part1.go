package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
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
	numberFinder, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Printf("Invalid regex, err: %s\n", err)
		return
	}
	runningCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := numberFinder.FindAllString(line, -1)
		if matches == nil {
			fmt.Printf("No number found in %q\n", line)
			continue
		}
		value, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			fmt.Printf(
				"Can't convert string to a number. Err %s\n",
				err,
			)
			continue
		}
		runningCount += value
	}
	fmt.Printf("Final count is %d\n", runningCount)
}
