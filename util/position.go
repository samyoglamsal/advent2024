package util

type Position struct {
	X int
	Y int
}

// Assumes square input
func (p Position) OutOfBounds(sideLength int) bool {
	return p.X < 0 || p.X >= sideLength || p.Y < 0 || p.Y >= sideLength
}

func (p Position) Add(other Position) Position {
	return Position{p.X + other.X, p.Y + other.Y}
}

func (p Position) NAdd(other Position, n int) Position {
	return Position{X: p.X + n*other.X, Y: p.Y + n*other.Y}
}

func (p Position) Subtract(other Position) Position {
	return Position{p.X - other.X, p.Y - other.Y}
}

func (p Position) NSubtract(other Position, n int) Position {
	return Position{X: p.X - n*other.X, Y: p.Y - n*other.Y}
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}
