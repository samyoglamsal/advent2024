package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Gilver() {
	file, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	prerequisites := make(map[int][]int)
	prerequisitesDone := false
	silver, gold := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			prerequisitesDone = true
			continue
		}

		if !prerequisitesDone {
			pair := strings.Split(input, "|")
			before, _ := strconv.Atoi(pair[0])
			after, _ := strconv.Atoi(pair[1])

			if prerequisites[after] == nil {
				prerequisites[after] = make([]int, 0)
			}

			prerequisites[after] = append(prerequisites[after], before)
		} else {
			tokens := strings.Split(input, ",")
			valid := true
			update := make([]int, 0)
			for i, token := range tokens {
				number, _ := strconv.Atoi(token)
				var correctPosition int = len(update)

				if i != 0 {
					for j := range update {
						index := len(update) - j - 1
						prior := update[index]
						if slices.Contains(prerequisites[prior], number) {
							valid = false
							correctPosition = index
						}
					}
				}

				if valid {
					update = append(update, number)
				} else {
					update = slices.Insert(update, correctPosition, number)
				}
			}

			if valid {
				silver += update[len(update)/2]
			} else {
				gold += update[len(update)/2]
			}

		}
	}
	fmt.Println("Silver: ", silver)
	fmt.Println("Gold: ", gold)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
