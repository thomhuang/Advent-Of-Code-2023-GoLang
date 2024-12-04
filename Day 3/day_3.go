package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	total := 0
	instructedTotal := 0
	do := true
	for sc.Scan() {
		text := sc.Text()
		for i := 0; i < len(text); i++ {
			if i+4 < len(text) && text[i] == 'm' && text[i+1] == 'u' && text[i+2] == 'l' && text[i+3] == '(' {
				i += 4
				first := parseNumber(text, &i)
				if i+1 < len(text) {
					if text[i] == ',' {
						i++
						second := parseNumber(text, &i)
						if i < len(text) && text[i] == ')' {
							if first != 0 && second != 0 {
								total += first * second
							}
							if do && first != 0 && second != 0 {
								instructedTotal += first * second
							}
						}
					}
				}
			} else if i+1 < len(text) && text[i] == 'd' && text[i+1] == 'o' {
				if i+2 < len(text) {
					i += 2
					parseInstruction(text, &i, &do)
				}
			}
		}
	}

	fmt.Println(total)
	fmt.Println(instructedTotal)
}

func parseNumber(text string, i *int) int {
	val := 0
	for *i < len(text) && val < 1000 && unicode.IsDigit(rune(text[*i])) {
		val = val*10 + int(text[*i]-'0')
		*i++
	}

	if val > 0 && val < 1000 {
		return val
	}

	return 0
}

func parseInstruction(text string, i *int, prevInstruction *bool) {
	if *i+1 < len(text) && text[*i] == '(' && text[*i+1] == ')' {
		*i += 1
		*prevInstruction = true
		return
	} else if *i+4 < len(text) && text[*i] == 'n' && text[*i+1] == '\'' && text[*i+2] == 't' && text[*i+3] == '(' && text[*i+4] == ')' {
		*i += 4
		*prevInstruction = false
	}
}
