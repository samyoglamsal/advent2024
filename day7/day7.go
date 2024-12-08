package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bruteForce(target int, current int, inputs []int, gold bool) bool {
	if len(inputs) == 0 {
		return current == target
	} else {
		return bruteForce(target, current+inputs[0], inputs[1:], gold) || bruteForce(target, current*inputs[0], inputs[1:], gold) || (gold && bruteForce(target, concat(current, inputs[0]), inputs[1:], gold))
	}
}

func concat(a, b int) int {
	c := b
	digits := 0
	for c != 0 {
		c /= 10
		digits += 1
	}

	result := a
	for i := 0; i < digits; i++ {
		result *= 10
	}

	return result + b
}

func solve() {
	file, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	silver, gold := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		tokens := strings.Split(input, ":")
		result, _ := strconv.Atoi(tokens[0])

		values := make([]int, 0)
		tokens = strings.Split(strings.TrimSpace(tokens[1]), " ")
		for _, token := range tokens {
			number, _ := strconv.Atoi(token)
			values = append(values, number)
		}

		possibleSilver := bruteForce(result, values[0], values[1:], false)
		possibleGold := bruteForce(result, values[0], values[1:], true)

		if possibleSilver {
			silver += result
		}
		if possibleGold {
			gold += result
		}
	}

	fmt.Println("Silver: ", silver)
	fmt.Println("Gold: ", gold)
}
