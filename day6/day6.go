package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sync"
	"time"

	. "github.com/samyoglamsal/advent2024/util"
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

func loopCheck(grid [][]rune, pos position, start position, counter *int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	if grid[pos.y][pos.x] == EMPTY {
		grid[pos.y][pos.x] = OBSTACLE

		g := guard{pos: position{x: start.x, y: start.y}, dir: direction{0, -1}}
		visited := make(map[guard]bool)
		visited[g] = true

		for {
			g.move(grid)
			if g.outOfBounds(len(grid)) {
				break
			}

			if visited[g] {
				m.Lock()
				*counter += 1
				m.Unlock()
				break
			} else {
				visited[g] = true
			}
		}

		grid[pos.y][pos.x] = EMPTY
	}
}

func Silver() []position {
	start := time.Now()
	defer MeasureExecutionTime(start, "silver")

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
	defer MeasureExecutionTime(start, "gold")

	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]rune, 0)

	// silver := 0
	scanner := bufio.NewScanner(file)
	var startPosition position
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)

		if slices.Contains(line, '^') {
			startPosition = position{slices.Index(line, '^'), len(grid)}
		}
	}

	positions := Silver()

	gold := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	for _, pos := range positions {
		if grid[pos.y][pos.x] == OBSTACLE || grid[pos.y][pos.x] == GUARD {
			continue
		} else {
			cpy := make([][]rune, len(grid))
			for i := range grid {
				cpy[i] = make([]rune, len(grid[i]))
				copy(cpy[i], grid[i])
			}

			wg.Add(1)
			go loopCheck(cpy, pos, startPosition, &gold, &wg, &m)
		}
	}

	wg.Wait()
	fmt.Println("Gold: ", gold)
}
