package processing

import (
	"fmt"
	"golang-toy-robot/model"
)

func ProcessCommand(table model.Table, command model.Command) model.Table {
	switch c := command.(type) {
	case model.PlaceCommand:
		if coordinateInBounds(c.Coordinate, table.MaxCoordinate) {
			table.Robot = &model.Robot{Position: c.Coordinate, Facing: c.Facing}
		}
	case model.RightCommand:
		if table.Robot != nil {
			table.Robot.Facing = turnRight(table.Robot.Facing)
		}
	case model.LeftCommand:
		if table.Robot != nil {
			table.Robot.Facing = turnLeft(table.Robot.Facing)
		}
	case model.MoveCommand:
		if table.Robot != nil {
			updatedRobot := move(*table.Robot, table.MaxCoordinate)
			table.Robot = &updatedRobot
		}
	case model.ReportCommand:
		if table.Robot != nil {
			fmt.Println(reportRobot(*table.Robot))
		}
	}

	return table
}
