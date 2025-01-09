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

	var equations [][]int
	for sc.Scan() {
		line := sc.Text()

		equationSplit := strings.Split(line, ": ")
		result, resErr := strconv.Atoi(equationSplit[0])
		if resErr != nil {
			panic(resErr)
		}

		valuesString := strings.Split(equationSplit[1], " ")
		values := make([]int, len(valuesString))
		for i, val := range valuesString {
			intVal, valErr := strconv.Atoi(val)
			if valErr != nil {
				panic(valErr)
			}

			values[i] = intVal
		}

		equations = append(equations, append(values, result))
	}

	partOneResult := 0
	partTwoResult := 0
	for _, equation := range equations {
		fmt.Println(equation)
		if len(equation) >= 2 {
			if evaluate(equation[1:len(equation)-1], equation[0], equation[len(equation)-1]) {
				partOneResult += equation[len(equation)-1]
			}

			if evaluateWithConcat(equation[1:len(equation)-1], equation[0], equation[len(equation)-1]) {
				partTwoResult += equation[len(equation)-1]
			}
		}
	}

	fmt.Println(partOneResult)
	fmt.Println(partTwoResult)
}

func evaluate(values []int, path, target int) bool {
	if len(values) == 0 {
		return path == target
	}

	if path > target {
		return false
	}

	add := path + values[0]
	multiply := path * values[0]

	return evaluate(values[1:], add, target) || evaluate(values[1:], multiply, target)
}

func evaluateWithConcat(values []int, path, target int) bool {
	if len(values) == 0 {
		return path == target
	}

	if path > target {
		return false
	}

	add := path + values[0]
	multiply := path * values[0]
	concat, concatErr := strconv.Atoi(fmt.Sprintf("%d%d", path, values[0]))
	if concatErr != nil {
		panic(concatErr)
	}

	return evaluateWithConcat(values[1:], add, target) || evaluateWithConcat(values[1:], multiply, target) || evaluateWithConcat(values[1:], concat, target)
}
