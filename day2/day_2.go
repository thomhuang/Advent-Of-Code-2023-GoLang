package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const MAX_DIFFERENCE int = 3

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	totalSafe := 0
	toleratedSafe := 0
	for sc.Scan() {
		line := sc.Text()

		levels := parseLine(line)

		if isSafe(levels) {
			totalSafe += 1
		} else if isToleratedSafe(levels) {
			toleratedSafe += 1
		}
	}

	totalToleratedSafe := totalSafe + toleratedSafe

	fmt.Println(totalSafe)
	fmt.Println(totalToleratedSafe)
}

func isSafe(levels []int) bool {
	return isIncreasingValid(levels) || isDecreasingValid(levels)
}

func isToleratedSafe(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		curr := slices.Concat(levels[:i], levels[i+1:])

		if isSafe(curr) {
			return true
		}
	}

	return false
}

func isIncreasingValid(levels []int) bool {
	isIncreasing := sort.SliceIsSorted(levels, func(i, j int) bool {
		return levels[i] < levels[j]
	})

	if !isIncreasing {
		return false
	}

	prev := levels[0]
	for i := 1; i < len(levels); i++ {
		if levels[i] == prev {
			return false
		}

		diff := levels[i] - prev
		if diff > MAX_DIFFERENCE {
			return false
		}

		prev = levels[i]
	}

	return true
}

func isDecreasingValid(levels []int) bool {
	isDecreasing := sort.SliceIsSorted(levels, func(i, j int) bool {
		return levels[i] > levels[j]
	})

	if !isDecreasing {
		return false
	}

	prev := levels[0]
	for i := 1; i < len(levels); i++ {
		if levels[i] == prev {
			return false
		}

		diff := prev - levels[i]
		if diff > MAX_DIFFERENCE {
			return false
		}

		prev = levels[i]
	}

	return true
}

func parseLine(line string) []int {
	levelVals := make([]int, 0)

	levels := strings.Split(line, " ")
	for _, level := range levels {
		curr, err := strconv.Atoi(level)
		if err != nil {
			panic("Could not parse level from line!")
		}

		levelVals = append(levelVals, curr)
	}

	return levelVals
}
