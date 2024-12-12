package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	var grid [][]rune
	for sc.Scan() {
		text := sc.Text()

		grid = append(grid, []rune(text))
	}

	for _, line := range grid {
		fmt.Println(string(line))
	}

	antennaMap := make(map[rune][][2]int)
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] != '.' {
				antennaMap[grid[i][j]] = append(antennaMap[grid[i][j]], [2]int{i, j})
			}
		}
	}

	antiMap := make(map[[2]int]bool)
	singleAntiNodeCount := 0
	for _, coords := range antennaMap {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				first := coords[i]
				second := coords[j]

				rise := second[0] - first[0]
				run := second[1] - first[1]

				antiNode1 := [2]int{first[0] - rise, first[1] - run}
				antiNode2 := [2]int{second[0] + rise, second[1] + run}

				if isValid(grid, antiNode1) && !antiMap[antiNode1] {
					antiMap[antiNode1] = true
					singleAntiNodeCount++
				}

				if isValid(grid, antiNode2) && !antiMap[antiNode2] {
					antiMap[antiNode2] = true
					singleAntiNodeCount++
				}
			}
		}
	}

	multipleAntiMap := make(map[[2]int]bool)
	multipleAntiNodeCount := 0
	for _, coords := range antennaMap {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				first := coords[i]
				second := coords[j]

				rise := second[0] - first[0]
				run := second[1] - first[1]

				antiNode1 := [2]int{first[0], first[1]}
				for isValid(grid, antiNode1) {
					if !multipleAntiMap[antiNode1] {
						multipleAntiNodeCount++
					}
					multipleAntiMap[antiNode1] = true

					antiNode1[0] -= rise
					antiNode1[1] -= run
				}

				antiNode2 := [2]int{second[0], second[1]}
				for isValid(grid, antiNode2) {
					if !multipleAntiMap[antiNode2] {
						multipleAntiNodeCount++
					}
					multipleAntiMap[antiNode2] = true

					antiNode2[0] += rise
					antiNode2[1] += run
				}
			}
		}
	}

	fmt.Println(singleAntiNodeCount)
	fmt.Println(multipleAntiNodeCount)
}

func isValid(grid [][]rune, point [2]int) bool {
	return point[0] >= 0 && point[0] < len(grid) && point[1] >= 0 && point[1] < len(grid[0])
}
