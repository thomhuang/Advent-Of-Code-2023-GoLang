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
	for sc.Scan() {
		text := sc.Text()
		for i := 0; i < len(text); i++ {
			if text[i] == 'm' && text[i+1] == 'u' && text[i+2] == 'l' && text[i+3] == '(' {
				i += 4
				first := parseNumber(text, &i)
				if i+2 < len(text) {
					if text[i] == ',' {
						i += 1
						second := parseNumber(text, &i)
						if i < len(text) && text[i] == ')' {
							if first != 0 && second != 0 {
								total += first * second
							}
						}
					}
				}

			}
		}
	}

	fmt.Println(total)
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
