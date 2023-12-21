package main

import (
	"errors"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mapSet struct {
	seedInputs []string
	maps       [7][]mappedRange
}

type mappedRange struct {
	// The start of the mapped range
	min int
	// The end of the mapped range
	max int
	// Where the min would map to
	destStart int
}

func (mRange mappedRange) contains(val int) bool {
	slog.Debug("Checking map range", "min", mRange.min, "max", mRange.max, "val", val)
	return !(val < mRange.min || val > mRange.max)
}

func main() {
	// Just wanted to try out slog I guess
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
	slog.SetDefault(logger)
	slog.Info("Day 5 Part 1")
	args := os.Args[1:]
	if len(args) < 1 {
		slog.Error("Please provide a map input as a command line arg")
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		slog.Error("Failed to open map", "file", args[0], "err", err.Error())
		return
	}
	maps, err := parseMapFile(string(data))
	if err != nil {
		// Error messages should be logged in children calls
		return
	}
	minLocation := -1
	for _, seed := range maps.seedInputs {
		slog.Debug("Parsing seed", "seed", seed)
		current, err := strconv.Atoi(seed)
		if err != nil {
			slog.Error("Failed to convert string to int", "string", seed, "error", err)
		}
		for i := 0; i < 7; i++ {
			okay := false
			for checkRange := 0; checkRange < len(maps.maps[i]); checkRange++ {
				checkMap := maps.maps[i][checkRange]
				okay = okay || checkMap.contains(current)
				if okay {
					current = checkMap.destStart + (current - checkMap.min)
					slog.Debug("Mapped Successfully", "val", current, "dest", current)
					break
				}
			}
		}
		slog.Debug("Seeds minimum", "minimum", current)
		if minLocation > current || minLocation == -1 {
			slog.Debug("new minimum", "minimum", current)
			minLocation = current
		}
	}
	slog.Info("Found lowest location", "location", minLocation)
}

func parseMapFile(data string) (mapSet, error) {
	slog.Debug("Parsing map file")
	result := mapSet{}
	seedListMatcher := regexp.MustCompile(`seeds: ([\d+ ]*)`)
	matches := seedListMatcher.FindStringSubmatch(data)
	if matches == nil {
		slog.Error("Failed to find seed input list in map input")
		return result, errors.New("Failed to find seed input list in map input")
	}
	slog.Debug("Found seed list", "list", matches[1])
	result.seedInputs = strings.Split(matches[1], " ")
	mapsMatcher := regexp.MustCompile(`map:\n([\d+ \n]+)`)
	mapMatches := mapsMatcher.FindAllStringSubmatch(data, -1)
	if mapMatches == nil || len(mapMatches) != 7 {
		slog.Error("Failed to parse maps in map input", "length", len(mapMatches))
		return result, errors.New("Failed to parse maps in map input")
	}
	for i, match := range mapMatches {
		slog.Debug("Parsing map", "map", match[1])
		err := parseMap(match[1], &result.maps[i])
		if err != nil {
			slog.Error("Failed to parse map", "map", len(match[1]))
			return result, err
		}
	}
	return result, nil
}

func parseMap(mapInput string, mapping *[]mappedRange) error {
	lines := strings.Split(mapInput, "\n")
	lineCount := 0
	for _, line := range lines {
		if line == "" {
			// My regex will always pick up an extra \n too many, Too bad!
			continue
		}
		lineCount++
	}
	lines = lines[:lineCount]
	slog.Debug("Making new mapping", "size", lineCount)
	newMapping := make([]mappedRange, lineCount)
	*mapping = newMapping
	for i, line := range lines {
		mappingDetails := strings.Split(line, " ")
		destStr, srcStr, lengthStr := mappingDetails[0], mappingDetails[1], mappingDetails[2]
		length, errLen := strconv.Atoi(lengthStr)
		dest, errDest := strconv.Atoi(destStr)
		src, errSrc := strconv.Atoi(srcStr)
		if errLen != nil || errSrc != nil || errDest != nil {
			return errors.Join(errLen, errSrc, errDest)
		}
		newMapping[i] = mappedRange{
			min:       src,
			max:       src + length,
			destStart: dest,
		}
	}
	return nil
}
