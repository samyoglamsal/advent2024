package day8

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samyoglamsal/advent2024/util"
)

const (
	EMPTY = '.'
)

func Gilver() {
	file, err := os.Open("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sideLength := 0
	antennas := make(map[rune][]util.Position)
	silverAntinodes := make(map[util.Position]bool)
	goldAntinodes := make(map[util.Position]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := []rune(scanner.Text())

		for i, c := range input {
			if c != EMPTY {
				if antennas[c] == nil {
					antennas[c] = make([]util.Position, 0)
				}
				antennas[c] = append(antennas[c], util.Position{X: i, Y: sideLength})
			}
		}
		sideLength += 1
	}

	for _, signal := range antennas {
		for i := 0; i < len(signal)-1; i++ {
			for j := i + 1; j < len(signal); j++ {
				var closeAntinodes, farAntinodes []util.Position
				closeInBounds, farInBounds := true, true

				distance := util.Position{X: signal[i].X - signal[j].X, Y: signal[i].Y - signal[j].Y}
				if signal[i].X+distance.X == signal[j].X && signal[i].Y+distance.Y == signal[j].Y {
					for k := 0; closeInBounds; k++ {
						closeAntinode := util.Position{X: signal[i].X - k*distance.X, Y: signal[i].Y - k*distance.Y}
						if closeAntinode.OutOfBounds(sideLength) {
							closeInBounds = false
						} else {
							closeAntinodes = append(closeAntinodes, closeAntinode)
						}
					}

					for k := 0; farInBounds; k++ {
						farAntinode := util.Position{X: signal[i].X + k*distance.X, Y: signal[i].Y + k*distance.Y}
						if farAntinode.OutOfBounds(sideLength) {
							farInBounds = false
						} else {
							farAntinodes = append(farAntinodes, farAntinode)
						}
					}
				} else if signal[i].X-distance.X == signal[j].X && signal[i].Y-distance.Y == signal[j].Y {
					for k := 0; closeInBounds; k++ {
						closeAntinode := util.Position{X: signal[i].X + k*distance.X, Y: signal[i].Y + k*distance.Y}
						if closeAntinode.OutOfBounds(sideLength) {
							closeInBounds = false
						} else {
							closeAntinodes = append(closeAntinodes, closeAntinode)
						}
					}

					for k := 0; farInBounds; k++ {
						farAntinode := util.Position{X: signal[i].X - k*distance.X, Y: signal[i].Y - k*distance.Y}
						if farAntinode.OutOfBounds(sideLength) {
							farInBounds = false
						} else {
							farAntinodes = append(farAntinodes, farAntinode)
						}
					}
				}

				for i, closeAntinode := range closeAntinodes {
					goldAntinodes[closeAntinode] = true

					if i == 0 {
						silverAntinodes[closeAntinode] = true
					}
				}

				for i, farAntinode := range farAntinodes {
					goldAntinodes[farAntinode] = true

					if i == 0 {
						silverAntinodes[farAntinode] = true
					}
				}
			}
		}
	}

	fmt.Println("Silver: ", len(silverAntinodes))
	fmt.Println("Gold: ", len(goldAntinodes))
}
