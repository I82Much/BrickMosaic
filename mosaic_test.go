package BrickMosaic

import (
	"reflect"
	"testing"

	"github.com/I82Much/BrickMosaic/palette"
)

type gridTest struct {
	name             string
	input            map[Location]palette.BrickColor
	numRows, numCols uint
	want             map[palette.BrickColor]Grid
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
			make(map[Location]palette.BrickColor),
			0, 0,
			make(map[palette.BrickColor]Grid),
		},
		{
			"One Color - 1x1",
			map[Location]palette.BrickColor{
				Location{0, 0}: palette.White,
			},
			1, 1,
			map[palette.BrickColor]Grid{
				palette.White: MakeFilledGrid(1, 1, ToBeFilled),
			},
		},
		{
			"Two colors",
			map[Location]palette.BrickColor{
				Location{0, 0}: palette.White,
				Location{0, 1}: palette.Black,
			},
			1, 2,
			map[palette.BrickColor]Grid{
				palette.White: whiteGrid,
				palette.Black: blackGrid,
			},
		},
	} {
		if got := makeGrids(test.numRows, test.numCols, test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("for test %q got %v wanted %v", test.name, got, test.want)
		}
	}

}
