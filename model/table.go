package model

import (
	"fmt"

	rapid "pgregory.net/rapid"
)

type Table struct {
	MaxCoordinate Coordinate
	Robot         *Robot
}

func (c Table) String() string { return fmt.Sprintf("Table %s,%s", c.MaxCoordinate, c.Robot) }

func GenTable(t *rapid.T) Table {
	robotGen := rapid.Custom(GenRobot)
	return Table{
		MaxCoordinate: Coordinate{X: 5, Y: 5},
		Robot: rapid.Custom(rapid.Ptr(robotGen, true)).Draw(t, "TableRobot").(*Robot),
	}
}