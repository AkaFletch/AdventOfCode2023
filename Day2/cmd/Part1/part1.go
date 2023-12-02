package main

import (
	"fmt"
	"os"
	"strings"
)

type Reveal struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	GameID      int
	GameReveals []Reveal
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
	for _, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
	}
}
