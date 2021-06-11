package model

import (
	"fmt"

	rapid "pgregory.net/rapid"
)

type Robot struct {
	Position Coordinate
	Facing   Direction
}

func (c Robot) String() string { return fmt.Sprintf("Robot %s,%s", c.Position, c.Facing) }

func GenRobot(t *rapid.T) Robot {
	return Robot{
		Position: rapid.Custom(GenCoordinate).Draw(t, "RobotPosition").(Coordinate),
		Facing: rapid.Custom(GenDirection).Draw(t, "RobotDirection").(Direction),
	}
}