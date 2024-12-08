package day8

import (
	"fmt"
	"time"

	. "github.com/samyoglamsal/advent2024/util"
)

const (
	EMPTY = '.'
)

func Gilver() {
	start := time.Now()
	defer MeasureExecutionTime(start, "Gilver")

	lines := ReadInput("inputs/day8.txt")

	sideLength := len(lines)
	antennas := make(map[rune][]Position)
	silverAntinodes := make(map[Position]bool)
	goldAntinodes := make(map[Position]bool)

	for i, line := range lines {
		for j, c := range line {
			if c != EMPTY {
				if antennas[c] == nil {
					antennas[c] = make([]Position, 0)
				}
				antennas[c] = append(antennas[c], Position{X: j, Y: i})
			}
		}
	}

	for _, signal := range antennas {
		for i := 0; i < len(signal)-1; i++ {
			for j := i + 1; j < len(signal); j++ {
				distance := signal[i].Subtract(signal[j])

				for k := 0; ; k++ {
					closeAntinode := signal[i].NSubtract(distance, k)
					farAntinode := signal[i].NAdd(distance, k)
					if signal[i].Subtract(distance).Equals(signal[j]) {
						closeAntinode, farAntinode = farAntinode, closeAntinode
					}

					if !closeAntinode.OutOfBounds(sideLength) {
						goldAntinodes[closeAntinode] = true

						if k == 1 {
							silverAntinodes[closeAntinode] = true
						}
					}

					if !farAntinode.OutOfBounds(sideLength) {
						goldAntinodes[farAntinode] = true

						if k == 2 {
							silverAntinodes[farAntinode] = true
						}
					}

					if farAntinode.OutOfBounds(sideLength) && closeAntinode.OutOfBounds(sideLength) {
						break
					}
				}
			}
		}
	}

	fmt.Println("Silver: ", len(silverAntinodes))
	fmt.Println("Gold: ", len(goldAntinodes))
}
