package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DELIMITER = " "

func getReport(input string) []int {
	levels := strings.Split(input, DELIMITER)

	report := []int{}
	for _, level := range levels {
		levelInt, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}
		report = append(report, levelInt)
	}

	return report
}

func validateReport(report []int) bool {
	ascending, descending := true, true

	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		if difference < 0 {
			difference = -difference
		}

		if (difference < 1) || (difference > 3) {
			return false
		}

		if report[i] > report[i-1] {
			descending = false
		} else if report[i] < report[i-1] {
			ascending = false
		}
	}

	return ascending || descending
}

func validateReportWithDampener(report []int) bool {
	if validateReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		dampened := make([]int, 0)
		dampened = append(dampened, report[:i]...)
		dampened = append(dampened, report[i+1:]...)

		fmt.Println(i, dampened)
		if validateReport(dampened) {
			return true
		}
	}

	return false
}

func Part1() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := getReport(scanner.Text())
		if validateReport(report) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part1 solution: %v\n", count)
}

func Part2() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := getReport(scanner.Text())
		if validateReportWithDampener(report) {
			count++
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part2 solution: %v\n", count)
}
