package processing

import (
	"golang-toy-robot/model"
	"reflect"
	"testing"

	rapid "pgregory.net/rapid"
)

func TestReportRobot(t *testing.T) {
	input := model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.North}
	result := reportRobot(input)
	expected := "1,1,NORTH"

	if result != expected {
		t.Errorf("Expected %s Got %s", expected, result)
	}
}

func TestMove(t *testing.T) {
	var tests = []struct {
		input    model.Robot
		expected model.Robot
	}{
		// basic move is working
		{model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.North}, model.Robot{Position: model.Coordinate{X: 1, Y: 2}, Facing: model.North}},
		{model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.East}, model.Robot{Position: model.Coordinate{X: 2, Y: 1}, Facing: model.East}},
		{model.Robot{Position: model.Coordinate{X: 2, Y: 2}, Facing: model.South}, model.Robot{Position: model.Coordinate{X: 2, Y: 1}, Facing: model.South}},
		{model.Robot{Position: model.Coordinate{X: 2, Y: 2}, Facing: model.West}, model.Robot{Position: model.Coordinate{X: 1, Y: 2}, Facing: model.West}},

		// does not go off the edge of the table
		{model.Robot{Position: model.Coordinate{X: 1, Y: 5}, Facing: model.North}, model.Robot{Position: model.Coordinate{X: 1, Y: 5}, Facing: model.North}},
		{model.Robot{Position: model.Coordinate{X: 5, Y: 1}, Facing: model.East}, model.Robot{Position: model.Coordinate{X: 5, Y: 1}, Facing: model.East}},
		{model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.South}, model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.South}},
		{model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.West}, model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.West}},
	}

	for _, tt := range tests {
		testname := tt.input.String()
		t.Run(testname, func(t *testing.T) {
			result := move(tt.input, model.Coordinate{X: 5, Y: 5})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("with %s got %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMoveRobotCannotGoOffBoard(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		maxCoordinate := model.Coordinate{X: 5, Y: 5}
		robot := rapid.Custom(model.GenRobot).Draw(t, "robot").(model.Robot)
		resultRobot := move(robot, maxCoordinate)

		if resultRobot.Position.X < 1 || resultRobot.Position.X > 5 || resultRobot.Position.Y < 1 || resultRobot.Position.Y > 5 {
			t.Errorf("Test failure got: %s", resultRobot)
		}
	})
}
