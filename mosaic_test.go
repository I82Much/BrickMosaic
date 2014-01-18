package BrickMosaic

import (
	"reflect"
	"testing"
)

type gridTest struct {
	name             string
	input            map[Location]BrickColor
	numRows, numCols uint
	want             map[BrickColor]Grid
}

func TestMakeGrids(t *testing.T) {
	// Two color test - one row, two columns. First color is white, second is black
	whiteGrid := MakeGrid(1, 2)
	blackGrid := MakeGrid(1, 2)
	whiteGrid.Set(0, 0, ToBeFilled)
	blackGrid.Set(0, 1, ToBeFilled)

	for _, test := range []gridTest{
		{
			"Empty",
			make(map[Location]BrickColor),
			0, 0,
			make(map[BrickColor]Grid),
		},
		{
			"One Color - 1x1",
			map[Location]BrickColor{
				Location{0, 0}: White,
			},
			1, 1,
			map[BrickColor]Grid{
				White: MakeFilledGrid(1, 1, ToBeFilled),
			},
		},
		{
			"Two colors",
			map[Location]BrickColor{
				Location{0, 0}: White,
				Location{0, 1}: Black,
			},
			1, 2,
			map[BrickColor]Grid{
				White: whiteGrid,
				Black: blackGrid,
			},
		},
	} {
		if got := makeGrids(test.numRows, test.numCols, test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("for test %q got %v wanted %v", test.name, got, test.want)
		}
	}

}
