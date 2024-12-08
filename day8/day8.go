package day8

import (
	"fmt"

	. "github.com/samyoglamsal/advent2024/util"
)

const (
	EMPTY = '.'
)

func Gilver() {
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
				distance := Position{X: signal[i].X - signal[j].X, Y: signal[i].Y - signal[j].Y}

				for k := 0; ; k++ {
					if signal[i].X+distance.X == signal[j].X && signal[i].Y+distance.Y == signal[j].Y {
						closeAntinode := Position{X: signal[i].X - k*distance.X, Y: signal[i].Y - k*distance.Y}
						farAntinode := Position{X: signal[i].X + k*distance.X, Y: signal[i].Y + k*distance.Y}

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

					} else if signal[i].X-distance.X == signal[j].X && signal[i].Y-distance.Y == signal[j].Y {
						closeAntinode := Position{X: signal[i].X + k*distance.X, Y: signal[i].Y + k*distance.Y}
						farAntinode := Position{X: signal[i].X - k*distance.X, Y: signal[i].Y - k*distance.Y}

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
	}

	fmt.Println("Silver: ", len(silverAntinodes))
	fmt.Println("Gold: ", len(goldAntinodes))
}
