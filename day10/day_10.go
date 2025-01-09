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
			nums = append(nums, int(num)-'0')
		}

		topography = append(topography, nums)
	}

	printTopography(topography)

	fmt.Println(part1(topography))
	fmt.Println(part2(topography))
}

func part1(topography [][]int) int {
	top := make([][]int, len(topography))
	copy(top, topography)

	score := 0
	for r := 0; r < len(top); r++ {
		for c := 0; c < len(top[r]); c++ {
			if top[r][c] != 0 {
				continue
			}

			visited := make(map[[2]int]bool)
			score += dfs_visited(top, visited, -1, r, c)
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
			// don't act if not a trail head
			if top[r][c] != 0 {
				continue
			}

			// don't double count if we 've reached
			score += dfs(top, -1, r, c)
		}
	}

	return score
}

func dfs_visited(topography [][]int, visited map[[2]int]bool, prev, r, c int) int {
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
		ans += dfs_visited(topography, visited, curr, next_r, next_c)
	}

	return ans
}

func dfs(topography [][]int, prev, r, c int) int {
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
		ans += dfs(topography, curr, next_r, next_c)
	}

	return ans
}

func printTopography(top [][]int) {
	for _, line := range top {
		fmt.Println(line)
	}
}
