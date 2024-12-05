package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	numMap := make(map[int]map[int]bool)
	pageUpdates := [][]int{}
	invalidPageUpdates := [][]int{}
	for sc.Scan() {
		line := sc.Text()

		if len(line) == 0 {
			continue
		} else if len(line) == 5 {
			first, second := parseOrder(line)
			if len(numMap[first]) == 0 {
				numMap[first] = make(map[int]bool)
			}

			numMap[first][second] = true
		} else {
			pageUpdates = append(pageUpdates, parseUpdates(line))
		}
	}

	middlePageSum := 0
	for _, pageUpdate := range pageUpdates {
		isValid := true
		for i, page := range pageUpdate {
			for j := i + 1; j < len(pageUpdate); j++ {
				if !numMap[page][pageUpdate[j]] && numMap[pageUpdate[j]][page] {
					fmt.Println(pageUpdate, page, pageUpdate[j])
					isValid = false
					break
				}
			}
			if !isValid {
				invalidPageUpdates = append(invalidPageUpdates, pageUpdate)
				break
			}
		}

		if isValid {
			middlePageSum += pageUpdate[len(pageUpdate)/2]
		}
	}

	fmt.Println(middlePageSum)
}

func parseOrder(line string) (num1, num2 int) {
	pages := strings.Split(line, "|")
	if len(pages) != 2 {
		panic("Attempted parsing page order and failed!")
	}

	num1, err := strconv.Atoi(pages[0])
	if err != nil {
		panic(err)
	}

	num2, err = strconv.Atoi(pages[1])
	if err != nil {
		panic(err)
	}

	return num1, num2
}

func parseUpdates(line string) []int {
	pageOrder := strings.Split(line, ",")
	update := make([]int, len(pageOrder))
	for i, page := range pageOrder {
		curr, err := strconv.Atoi(page)
		if err != nil {
			panic(err)
		}

		update[i] = curr
	}

	return update
}
