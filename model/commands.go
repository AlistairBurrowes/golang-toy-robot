package model

import (
	"fmt"

	rapid "pgregory.net/rapid"
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

func GenPlaceCommand(t *rapid.T) Command {
	return PlaceCommand{
		Coordinate: rapid.Custom(GenCoordinate).Draw(t, "PlaceCommandPosition").(Coordinate),
		Facing: rapid.Custom(GenDirection).Draw(t, "PlaceCommandFacing").(Direction),
	}
}

func GenCommand(t *rapid.T) Command {
	gen := rapid.OneOf(
		rapid.Just(MoveCommand{}),
		rapid.Just(LeftCommand{}),
		rapid.Just(RightCommand{}),
		rapid.Just(ReportCommand{}),
		rapid.Custom(GenPlaceCommand),
	)
	return gen.Draw(t, "Command").(Command)
}