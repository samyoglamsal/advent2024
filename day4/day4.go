package day4

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SIDE_LENGTH = 140
)

func findXMAS(board [][]rune, row int, col int) int {
	XMAS := []rune("XMAS")
	count := 0
	found := true

	// Left
	if col >= 3 {
		for i := 0; i < 4; i++ {
			if board[row][col-i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Up
	if row >= 3 {
		for i := 0; i < 4; i++ {
			if board[row-i][col] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Down
	if row <= 136 {
		for i := 0; i < 4; i++ {
			if board[row+i][col] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Right
	if col <= 136 {
		for i := 0; i < 4; i++ {
			if board[row][col+i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true

	// Upper left
	if col >= 3 && row >= 3 {
		for i := 0; i < 4; i++ {
			if board[row-i][col-i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Upper right
	if col <= 136 && row >= 3 {
		for i := 0; i < 4; i++ {
			if board[row-i][col+i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Bottom left
	if col >= 3 && row <= 136 {
		for i := 0; i < 4; i++ {
			if board[row+i][col-i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true
	// Bottom right
	if col <= 136 && row <= 136 {
		for i := 0; i < 4; i++ {
			if board[row+i][col+i] != XMAS[i] {
				found = false
			}
		}

		if found {
			count += 1
		}
	}
	found = true

	return count
}

func findCrossMAS(board [][]rune, row int, col int) int {
	count := 0

	if board[row][col] != 'A' {
		return 0
	}

	if row > 0 && col > 0 && row < 139 && col < 139 {
		if board[row-1][col-1] == 'M' && board[row+1][col+1] == 'S' && board[row-1][col+1] == 'M' && board[row+1][col-1] == 'S' {
			count += 1
		} else if board[row-1][col-1] == 'S' && board[row+1][col+1] == 'M' && board[row-1][col+1] == 'S' && board[row+1][col-1] == 'M' {
			count += 1
		} else if board[row-1][col-1] == 'M' && board[row+1][col+1] == 'S' && board[row-1][col+1] == 'S' && board[row+1][col-1] == 'M' {
			count += 1
		} else if board[row-1][col-1] == 'S' && board[row+1][col+1] == 'M' && board[row-1][col+1] == 'M' && board[row+1][col-1] == 'S' {
			count += 1
		}
	}

	return count
}

func Gilver() {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	board := make([][]rune, 140)
	for i := range board {
		board[i] = make([]rune, 140)
	}

	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		board[row] = []rune(scanner.Text())
		row += 1
	}

	silver, gold := 0, 0
	for i := 0; i < SIDE_LENGTH; i++ {
		for j := 0; j < SIDE_LENGTH; j++ {
			silver += findXMAS(board, i, j)
			gold += findCrossMAS(board, i, j)
		}
	}

	fmt.Println("Silver:", silver)
	fmt.Println("Gold:", gold)
}
