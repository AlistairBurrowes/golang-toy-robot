package processing

import (
	"fmt"
	"golang-toy-robot/model"
	"reflect"
	"testing"
)

func TestReportRobot(t *testing.T) {
	input := model.Robot{Position: model.Coordinate{1, 1}, Facing: model.North}
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
		{model.Robot{Position: model.Coordinate{1, 1}, Facing: model.North}, model.Robot{Position: model.Coordinate{1, 2}, Facing: model.North}},
		{model.Robot{Position: model.Coordinate{1, 1}, Facing: model.East}, model.Robot{Position: model.Coordinate{2, 1}, Facing: model.East}},
		{model.Robot{Position: model.Coordinate{2, 2}, Facing: model.South}, model.Robot{Position: model.Coordinate{2, 1}, Facing: model.South}},
		{model.Robot{Position: model.Coordinate{2, 2}, Facing: model.West}, model.Robot{Position: model.Coordinate{1, 2}, Facing: model.West}},

		// does not go off the edge of the table
		{model.Robot{Position: model.Coordinate{1, 5}, Facing: model.North}, model.Robot{Position: model.Coordinate{1, 5}, Facing: model.North}},
		{model.Robot{Position: model.Coordinate{5, 1}, Facing: model.East}, model.Robot{Position: model.Coordinate{5, 1}, Facing: model.East}},
		{model.Robot{Position: model.Coordinate{1, 1}, Facing: model.South}, model.Robot{Position: model.Coordinate{1, 1}, Facing: model.South}},
		{model.Robot{Position: model.Coordinate{1, 1}, Facing: model.West}, model.Robot{Position: model.Coordinate{1, 1}, Facing: model.West}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			result := move(tt.input, model.Coordinate{5, 5})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("with %s got %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
