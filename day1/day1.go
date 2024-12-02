package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samyoglamsal/advent2024/collections"
)

const DELIMITER = "   "

func getPair(input string) (int, int) {
	tokens := strings.Split(input, DELIMITER)
	left, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	return left, right
}

func Part1() {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	leftHeap, rightHeap := collections.IntHeap{}, collections.IntHeap{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left, right := getPair(scanner.Text())
		leftHeap.Push(left)
		rightHeap.Push(right)
	}

	heap.Init(&leftHeap)
	heap.Init(&rightHeap)

	sum := 0

	for leftHeap.Len() > 0 {
		left := heap.Pop(&leftHeap).(int)
		right := heap.Pop(&rightHeap).(int)

		if left > right {
			sum += left - right
		} else {
			sum += right - left
		}

	}

	if scanner.Err() != nil {
		panic("Error reading file")
	}

	fmt.Printf("Part1 solution: %v\n", sum)
}

func Part2() {
	leftCounts, rightCounts := make(map[int]int), make(map[int]int)

	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left, right := getPair(scanner.Text())
		leftCounts[left]++
		rightCounts[right]++

	}

	sum := 0

	for key, value := range leftCounts {
		if rightCounts[key] > 0 {
			sum += key * value * rightCounts[key]
		}
	}

	fmt.Printf("Part2 solution: %v\n", sum)
}
