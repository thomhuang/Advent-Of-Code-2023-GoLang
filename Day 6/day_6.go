package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	var pt1Grid [][]rune
	var pt2Grid [][]rune
	row := 0
	pt1Start := make([]int, 2)
	pt2Start := make([]int, 2)
	for sc.Scan() {
		line := sc.Text()

		pt1Grid = append(pt1Grid, []rune(line))
		pt2Grid = append(pt2Grid, []rune(line))
		for col, _ := range line {
			if pt1Grid[row][col] == '^' {
				pt1Start[0] = row
				pt1Start[1] = col

				pt2Start[0] = row
				pt2Start[1] = col

				pt1Grid[row][col] = '.'
				pt2Grid[row][col] = '.'
			}
		}
		row += 1
	}

	// part1
	getDistinctPositions(pt1Grid, pt1Start)
	// part2
	getBadObstructions(pt2Grid, pt2Start)
}

func getDistinctPositions(grid [][]rune, pos []int) {
	dir := 0
	distinctPos := 1
	// set start to visited
	grid[pos[0]][pos[1]] = 'V'

	// assumed no cycles
	for true {
		nextRow := pos[0] + dirs[dir][0]
		nextCol := pos[1] + dirs[dir][1]
		// if we've reached outside of the allowed cells, we're done
		if !isValid(len(grid), len(grid[0]), nextRow, nextCol) {
			break
		} else {
			// change direction given obstruction
			if grid[nextRow][nextCol] == '#' {
				dir = (dir + 1) % 4
			} else {
				// update position
				pos[0] = nextRow
				pos[1] = nextCol
				// if we haven't visited, increment distinct pos
				if grid[nextRow][nextCol] != 'V' {
					grid[nextRow][nextCol] = 'V'
					distinctPos++
				}
			}
		}
	}

	fmt.Println(distinctPos)
}

func getBadObstructions(grid [][]rune, start []int) {
	badObstructions := 0
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == '.' && (start[0] != r && start[1] != c) {
				// turn current position, that isn't current guard position, to an obstruction
				// to see if it creates a cycle eventually
				grid[r][c] = '#'
				if containsCycle(grid, start) {
					badObstructions++
				}

				// revert and attempt again
				grid[r][c] = '.'
			}
		}
	}

	fmt.Println(badObstructions)
}

func containsCycle(grid [][]rune, start []int) bool {
	// keep track of visited cells, with hash of position + direction
	// we know that if we run into combination of position + same moving direction, we've entered a cylce
	visited := make(map[string]bool)

	// ensure we aren't updating start by ref ...
	curr := make([]int, 2)
	copy(curr, start)
	dir := 0
	for true {
		// hash
		currKey := fmt.Sprintf("%d-%d-%d", curr[0], curr[1], dir)
		if visited[currKey] {
			return true
		}
		visited[currKey] = true

		nextRow := curr[0] + dirs[dir][0]
		nextCol := curr[1] + dirs[dir][1]
		if !isValid(len(grid), len(grid[0]), nextRow, nextCol) {
			break
		}
		// if we run into an obstacle, change direction
		if grid[nextRow][nextCol] == '#' {
			dir = (dir + 1) % 4
		} else {
			// otherwise iterate to next row/col
			curr[0], curr[1] = nextRow, nextCol
		}
	}

	return false
}

func isValid(maxRow, maxCol, row, col int) bool {
	return row >= 0 && row < maxRow && col >= 0 && col < maxCol
}
