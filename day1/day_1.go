package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N int = 1000

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	leftNums := make([]int, N)
	rightNums := make([]int, N)
	freq := make(map[int]int)

	counter := 0
	for sc.Scan() {
		line := sc.Text()

		num1, num2 := parseLine(line)

		leftNums[counter] = num1
		rightNums[counter] = num2

		freq[num2] += 1
		counter += 1
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	totalBalance := 0
	similarityScore := 0
	for i := 0; i < N; i++ {
		totalBalance += absDiff(leftNums[i], rightNums[i])
		similarityScore += leftNums[i] * freq[leftNums[i]]
	}

	fmt.Println(totalBalance)
	fmt.Println(similarityScore)
}

func absDiff(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -diff
	}

	return diff
}

func parseLine(text string) (num1, num2 int) {
	nums := strings.Fields(text)
	if len(nums) < 2 {
		panic("Could not parse the line!")
	}

	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		panic("Could not parse the first number as an int")
	}

	num2, err = strconv.Atoi(nums[1])
	if err != nil {
		panic("Could not parse the second number as an int")
	}

	return num1, num2
}
