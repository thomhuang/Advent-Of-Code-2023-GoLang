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
	row := 0
	for sc.Scan() {
		line := sc.Text()

		if len(grid) == 0 {
			grid = make([][]rune, len(line))
			for i := range len(line) {
				grid[i] = make([]rune, len(line))
			}
		}

		for col, curr := range line {
			grid[row][col] = curr
		}

		row += 1
	}

	xmasCount := countXMAS(grid)
	xMASCount := countX_MAS(grid)

	fmt.Println(xmasCount)
	fmt.Println(xMASCount)
}

func countXMAS(grid [][]rune) int {
	count := 0

	xmas := [4]rune{'X', 'M', 'A', 'S'}

	directions := [12][2]int{
		// forward
		{1, 0},
		// backward
		{-1, 0},
		// up
		{0, 1},
		// down
		{0, -1},
		// left up
		{-1, 1},
		// left down
		{-1, -1},
		// right up
		{1, 1},
		// right down
		{1, -1},
	}

	for r := range len(grid) {
		for c := range len(grid[r]) {
			for _, d := range directions {
				d_r := d[0]
				d_c := d[1]
				isMatching := true

				for offSet := range len(xmas) {
					// d_r, d_c accounts for one index movement
					currRow := r + offSet*d_r
					currCol := c + offSet*d_c

					if !validSearch(len(grid[0]), len(grid), currRow, currCol) || grid[currRow][currCol] != xmas[offSet] {
						// direction goes out of bounds or current position given direction is not matching 'XMAS'
						isMatching = false
						break
					}
				}

				// if direction movement is valid, count it
				if isMatching {
					count++
				}
			}
		}
	}

	return count
}

func countX_MAS(grid [][]rune) int {
	count := 0

	directions := [][2][2]int{
		{
			// each of these are (r, c) pairs
			// directions[][i] are are diagonal from each other
			{-1, 1},
			{1, -1},
		},
		{
			{-1, -1},
			{1, 1},
		},
	}

	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] != 'A' {
				continue
			}

			isMatching := true
			for _, d := range directions {
				row1_d := r + d[0][0]
				col1_d := c + d[0][1]

				row2_d := r + d[1][0]
				col2_d := c + d[1][1]

				if !validSearch(len(grid), len(grid[0]), row1_d, col1_d) || !validSearch(len(grid), len(grid[0]), row2_d, col2_d) {
					isMatching = false
					break
				}

				/*
					Any of these are possible:

					M M       M S      S M      S S
					 A    OR   A   OR   A   OR   A
					S S       M S      S M      M M

					and so on, thus we check if one of them are valid, then check the other direction
				*/
				if (grid[row1_d][col1_d] == 'M' && grid[row2_d][col2_d] == 'S') || (grid[row1_d][col1_d] == 'S' && grid[row2_d][col2_d] == 'M') {
					continue
				}

				isMatching = false
				break
			}

			if isMatching {
				count++
			}
		}
	}
	return count
}

func validSearch(maxRow, maxCol, r, c int) bool {
	return r >= 0 && r < maxRow && c >= 0 && c < maxCol
}
