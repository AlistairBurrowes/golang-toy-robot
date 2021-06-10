package processing

import (
	"golang-toy-robot/model"
	"regexp"
	"strconv"
)

func ParseCommand(input string) model.Command {
	var result model.Command

	switch input {
	case "MOVE":
		result = model.MoveCommand{}
	case "LEFT":
		result = model.LeftCommand{}
	case "RIGHT":
		result = model.RightCommand{}
	case "REPORT":
		result = model.ReportCommand{}
	default:
		result = parsePlace(input)
	}

	return result
}

var matchPlace = regexp.MustCompile(`^PLACE\s(?P<X>\d+),(?P<Y>\d+),(?P<Facing>[A-Z]+)$`)

func parsePlace(input string) model.Command {
	parseResult := matchPlace.FindStringSubmatch(input)

	if len(parseResult) == 4 {
		parsedX := parseCoordinateAxis(parseResult[1])
		parsedY := parseCoordinateAxis(parseResult[2])
		parsedFacing := parseFacing(parseResult[3])

		if parsedX != nil && parsedY != nil && parsedFacing != nil {
			return model.PlaceCommand{Coordinate: model.Coordinate{X: *parsedX, Y: *parsedY}, Facing: *parsedFacing}
		}
	}

	return nil
}

func parseCoordinateAxis(input string) *int {
	i, err := strconv.Atoi(input)

	if err == nil {
		return &i
	} else {
		return nil
	}
}

func parseFacing(input string) *model.Direction {
	var result *model.Direction

	switch input {
	case "NORTH":
		d := model.North
		result = &d
	case "EAST":
		d := model.East
		result = &d
	case "SOUTH":
		d := model.South
		result = &d
	case "WEST":
		d := model.West
		result = &d
	}

	return result
}
