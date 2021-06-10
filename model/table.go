package model

import (
	"fmt"
)

type Table struct {
	MaxCoordinate Coordinate
	Robot         *Robot
}

func (c Table) String() string { return fmt.Sprintf("Table %s,%s", c.MaxCoordinate, c.Robot) }
