package model

import (
	"fmt"
)

type CommandType int

type Command interface {
	isCommand()
}

type PlaceCommand struct {
	Coordinate Coordinate
	Facing     Direction
}

type MoveCommand struct{}
type LeftCommand struct{}
type RightCommand struct{}
type ReportCommand struct{}

func (pc PlaceCommand) isCommand()  {}
func (pc MoveCommand) isCommand()   {}
func (pc LeftCommand) isCommand()   {}
func (pc RightCommand) isCommand()  {}
func (pc ReportCommand) isCommand() {}

func (c PlaceCommand) String() string  { return fmt.Sprintf("PLACE %s,%s", c.Coordinate, c.Facing) }
func (c MoveCommand) String() string   { return "MOVE" }
func (c LeftCommand) String() string   { return "LEFT" }
func (c RightCommand) String() string  { return "RIGHT" }
func (c ReportCommand) String() string { return "REPORT" }
