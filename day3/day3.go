package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Gilver() (int, int) {
	text, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	input := string(text)
	enabled := true
	silver, gold := 0, 0
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, finding := range r.FindAllStringSubmatch(input, -1) {
		if strings.Contains(finding[0], "mul") {
			first, _ := strconv.Atoi(finding[1])
			second, _ := strconv.Atoi(finding[2])

			if enabled {
				gold += first * second
			}

			silver += first * second
		} else if strings.Contains(finding[0], "don't()") {
			enabled = false
		} else if strings.Contains(finding[0], "do()") {
			enabled = true
		}
	}

	fmt.Println(silver, gold)
	return silver, gold
}
