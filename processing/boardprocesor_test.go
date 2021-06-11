package processing

import (
	"golang-toy-robot/model"
	"reflect"
	"testing"

	rapid "pgregory.net/rapid"
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

func TestRobotCannotGoOffBoard(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		table := rapid.Custom(model.GenTable).Draw(t, "table").(model.Table)
		command := rapid.Custom(model.GenCommand).Draw(t, "table").(model.Command)
		resultTable := ProcessCommand(table, command)

		if resultTable.Robot != nil &&
		  (resultTable.Robot.Position.X < 1 || resultTable.Robot.Position.X > 5 || resultTable.Robot.Position.Y < 1 || resultTable.Robot.Position.Y > 5) {
			// the rapid logging will not log into points https://github.com/flyingmutant/rapid/issues/7
			// so we print out what table we eneded up with
			t.Logf("Initial table: %s", table)
			t.Errorf("Test failure got: %s", resultTable)
		}
	})
}