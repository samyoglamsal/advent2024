package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

const (
	GUARD    = '^'
	OBSTACLE = '#'
	EMPTY    = '.'
)

type position struct {
	x int
	y int
}

type direction struct {
	dx int
	dy int
}

type guard struct {
	pos position
	dir direction
}

func measureExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func (g guard) outOfBounds(sideLength int) bool {
	return g.pos.x < 0 || g.pos.x > sideLength-1 || g.pos.y < 0 || g.pos.y > sideLength-1
}

func (g guard) canMove(sideLength int) bool {
	return g.pos.x > 0 && g.pos.x < sideLength-1 && g.pos.y > 0 && g.pos.y < sideLength-1
}

func (g *guard) move(grid [][]rune) {
	if g.canMove(len(grid)) {
		for grid[g.pos.y+g.dir.dy][g.pos.x+g.dir.dx] == OBSTACLE {
			g.dir.dx, g.dir.dy = -g.dir.dy, g.dir.dx
		}
	}

	g.pos.x += g.dir.dx
	g.pos.y += g.dir.dy
}

func Silver() []position {
	start := time.Now()
	defer measureExecutionTime(start, "silver")

	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)
	g := guard{pos: position{0, 0}, dir: direction{0, -1}}
	seen := make(map[position]bool)
	visited := make([]position, 0)

	visited = append(visited, g.pos)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)

		if slices.Contains(line, GUARD) {
			g.pos.x = slices.Index(line, GUARD)
			g.pos.y = len(grid)

			seen[g.pos] = true
			visited = append(visited, g.pos)
		}
	}

	for {
		g.move(grid)
		if g.outOfBounds(len(grid)) {
			break
		}

		if !seen[g.pos] {
			seen[g.pos] = true
			visited = append(visited, g.pos)
		}
	}

	fmt.Println("Silver: ", len(seen))
	return visited
}

func Gold() {
	start := time.Now()
	defer measureExecutionTime(start, "gold")

	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)

	// silver := 0
	scanner := bufio.NewScanner(file)
	var startX, startY int
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)

		if slices.Contains(line, '^') {
			startX = slices.Index(line, '^')
			startY = len(grid)
		}
	}

	positions := Silver()

	gold := 0
	for _, pos := range positions {
		if grid[pos.y][pos.x] == OBSTACLE || grid[pos.y][pos.x] == GUARD {
			continue
		} else {
			grid[pos.y][pos.x] = OBSTACLE

			g := guard{pos: position{x: startX, y: startY}, dir: direction{0, -1}}
			visited := make(map[guard]bool)
			visited[g] = true

			for {
				g.move(grid)
				if g.outOfBounds(len(grid)) {
					break
				}

				if visited[g] {
					gold += 1
					break
				} else {
					visited[g] = true
				}
			}

			grid[pos.y][pos.x] = EMPTY
		}
	}

	fmt.Println("Gold: ", gold)
}
