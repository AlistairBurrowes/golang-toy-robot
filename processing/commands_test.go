package processing

import (
	"fmt"
	"golang-toy-robot/model"
	"testing"
)

func TestParseCommand(t *testing.T) {
	var tests = []struct {
		input    string
		expected model.Command
	}{
		{"PLACE 1,1,NORTH", model.PlaceCommand{model.Coordinate{X: 1, Y: 1}, model.North}},
		{"MOVE", model.MoveCommand{}},
		{"LEFT", model.LeftCommand{}},
		{"RIGHT", model.RightCommand{}},
		{"REPORT", model.ReportCommand{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			result := ParseCommand(tt.input)
			if result != tt.expected {
				t.Errorf("with %s got %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
