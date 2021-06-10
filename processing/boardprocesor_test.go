package processing

import (
	"golang-toy-robot/model"
	"reflect"
	"testing"
)

func TestPlaceCommand(t *testing.T) {
	initialTable := model.Table{MaxCoordinate: model.Coordinate{X: 5, Y: 5}, Robot: nil}
	command := model.PlaceCommand{Coordinate: model.Coordinate{X: 1, Y: 1}, Facing: model.North}

	resultTable := ProcessCommand(initialTable, command)

	expectedTable := model.Table{MaxCoordinate: model.Coordinate{X: 5, Y: 5}, Robot: &model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.North}}

	// I believe DeepEqual is required because == will only compare pointers for equality when something is a pointer to a value.
	// In this case the expectedTable has its own Robot created in place
	if !reflect.DeepEqual(resultTable, expectedTable) {
		t.Errorf("Test failure got: %s", resultTable)
	}
}

func TestReportCommand(t *testing.T) {
	initialTable := model.Table{MaxCoordinate: model.Coordinate{X: 5, Y: 5}, Robot: &model.Robot{Position: model.Coordinate{X: 1, Y: 1}, Facing: model.North}}
	command := model.ReportCommand{}

	resultTable := ProcessCommand(initialTable, command)

	if !reflect.DeepEqual(resultTable, initialTable) {
		t.Errorf("Test failure got: %s", resultTable)
	}
}
