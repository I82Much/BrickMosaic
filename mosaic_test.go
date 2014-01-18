package BrickMosaic

/*
import (
	"reflect"
	"testing"

	"github.com/I82Much/BrickMosaic/grid"
	"github.com/I82Much/BrickMosaic/palette"

)

type gridTest struct {
	name             string
	input            map[grid.Location]palette.BrickColor
	numRows, numCols uint
	want             map[palette.BrickColor]grid.Grid
}

func TestMakeGrids(t *testing.T) {
	// Two color test - one row, two columns. First color is white, second is black
	whiteGrid := grid.New(1, 2)
	blackGrid := grid.New(1, 2)
	whiteGrid.Set(0, 0, grid.ToBeFilled)
	blackGrid.Set(0, 1, grid.ToBeFilled)

	for _, test := range []gridTest{
		{
			"Empty",
			make(map[grid.Location]palette.BrickColor),
			0, 0,
			make(map[palette.BrickColor]grid.Grid),
		},
		{
			"One Color - 1x1",
			map[grid.Location]palette.BrickColor{
				grid.Location{0, 0}: palette.White,
			},
			1, 1,
			map[palette.BrickColor]grid.Grid{
				palette.White: grid.WithState(1, 1, grid.ToBeFilled),
			},
		},
		{
			"Two colors",
			map[grid.Location]palette.BrickColor{
				grid.Location{0, 0}: palette.White,
				grid.Location{0, 1}: palette.Black,
			},
			1, 2,
			map[palette.BrickColor]grid.Grid{
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
*/