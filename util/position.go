package util

type Position struct {
	X int
	Y int
}

// Assumes square input
func (p Position) OutOfBounds(sideLength int) bool {
	return p.X < 0 || p.X >= sideLength || p.Y < 0 || p.Y >= sideLength
}
