package processing

import (
	"reflect"
	"testing"

	rapid "pgregory.net/rapid"
)

func TestSliceReverseRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		slice := rapid.SliceOf(rapid.String()).Draw(t, "slice").([]string)
		doubleReversed := reverse(reverse(slice))

		if !reflect.DeepEqual(slice, doubleReversed) {
			t.FailNow()
		}
	})
}
