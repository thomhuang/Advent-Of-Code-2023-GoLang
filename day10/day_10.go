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

	var topography [][]int
	for sc.Scan() {
		line := sc.Text()

		var nums []int
		for _, num := range line {
			nums = append(nums, int(num-'0'))
		}

		topography = append(topography, nums)
	}

	fmt.Println(part1(topography))
	fmt.Println(part2(topography))
}

func part1(topography [][]int) int {
	top := make([][]int, len(topography))
	copy(top, topography)

	score := 0
	for r := 0; r < len(top); r++ {
		for c := 0; c < len(top[r]); c++ {
			// don't traverse if not trail head
			if top[r][c] != 0 {
				continue
			}

			// keep track of visited trails with maximal slope
			visited := make(map[[2]int]bool)
			score += dfs_part1(top, visited, -1, r, c)
		}
	}

	return score
}

func part2(topography [][]int) int {
	top := make([][]int, len(topography))
	copy(top, topography)

	score := 0
	for r := 0; r < len(top); r++ {
		for c := 0; c < len(top[r]); c++ {
			// don't traverse if not trail head
			if top[r][c] != 0 {
				continue
			}

			// part 2, can revisit a '9' from the same trailhead
			score += dfs_part2(top, -1, r, c)
		}
	}

	return score
}

func dfs_part1(topography [][]int, visited map[[2]int]bool, prev, r, c int) int {
	if r < 0 || r >= len(topography) || c < 0 || c >= len(topography[r]) {
		return 0
	}

	point := [2]int{r, c}
	curr := topography[r][c]
	if curr-1 != prev {
		return 0
	}

	if curr == 9 && prev == 8 {
		if visited[point] {
			return 0
		}
		visited[point] = true
		return 1
	}

	dir := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	ans := 0
	for _, d := range dir {
		d_x, d_y := d[0], d[1]
		next_r, next_c := r+d_x, c+d_y
		ans += dfs_part1(topography, visited, curr, next_r, next_c)
	}

	return ans
}

func dfs_part2(topography [][]int, prev, r, c int) int {
	if r < 0 || r >= len(topography) || c < 0 || c >= len(topography[r]) {
		return 0
	}

	curr := topography[r][c]
	if curr-1 != prev {
		return 0
	}

	if curr == 9 && prev == 8 {
		return 1
	}

	dir := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	ans := 0
	for _, d := range dir {
		d_x, d_y := d[0], d[1]
		next_r, next_c := r+d_x, c+d_y
		ans += dfs_part2(topography, curr, next_r, next_c)
	}

	return ans
}
