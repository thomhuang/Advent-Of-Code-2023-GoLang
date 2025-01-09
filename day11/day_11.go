package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type key struct {
	stone  string
	blinks int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	line, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	stones := strings.Fields(string(line))

	memo := make(map[key]int)
	count := 0
	totalBlinks := 75
	for _, stone := range stones {
		count += dfs(stone, totalBlinks, memo)
	}

	fmt.Println(count)
}

// At first did brute force basically of calculating and storing stones on the fly ...
// Realized order doesn't matter and we'll run into repeats, thus memoization for given (stone, blink)
func dfs(stone string, blinks int, memo map[key]int) int {
	if blinks == 0 {
		return 1
	}

	if val, exists := memo[key{stone, blinks}]; exists {
		return val
	}

	total := 0
	switch {
	case stone == "0":
		total = dfs("1", blinks-1, memo)
		break
	case len(stone)%2 == 0:
		// Probably not the best way to do this .. convert to string --> int --> string to deal with leading 0s, etc.
		left, right := stone[:len(stone)/2], stone[len(stone)/2:]

		leftVal, _ := strconv.Atoi(left)
		rightVal, _ := strconv.Atoi(right)

		left, right = strconv.Itoa(leftVal), strconv.Itoa(rightVal)

		total = dfs(left, blinks-1, memo) + dfs(right, blinks-1, memo)
		break
	default:
		curr, _ := strconv.Atoi(stone)
		total = dfs(strconv.Itoa(curr*2024), blinks-1, memo)
	}

	memo[key{stone, blinks}] = total

	return total
}
