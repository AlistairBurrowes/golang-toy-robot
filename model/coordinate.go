package model

import (
	"fmt"

	rapid "pgregory.net/rapid"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string { return fmt.Sprintf("%d,%d", c.X, c.Y) }

func GenCoordinate(t *rapid.T) Coordinate {
	return Coordinate{
			X: rapid.IntRange(1, 5).Draw(t, "X").(int),
			Y: rapid.IntRange(1, 5).Draw(t, "Y").(int),
	}
}
