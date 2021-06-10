package model

import (
	"fmt"
)

type Robot struct {
	Position Coordinate
	Facing   Direction
}

func (c Robot) String() string { return fmt.Sprintf("Robot %s,%s", c.Position, c.Facing) }
