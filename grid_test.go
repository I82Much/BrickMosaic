package BrickMosaic

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type getTest struct {
		name     string
		g        Grid
		row, col int
		want     State
	}
	various := MakeGrid(5, 6)
	various.state[4][3] = ToBeFilled
	various.state[3][2] = Filled
	for _, test := range []getTest{
		{"empty", MakeGrid(0, 0), 0, 0, Empty},
		{"negative row", MakeGrid(0, 0), -1, 0, Empty},
		{"negative col", MakeGrid(0, 0), 0, -1, Empty},
		{"out of bounds", MakeGrid(5, 5), -1, 45, Empty},
		{"to be filled", various, 4, 3, ToBeFilled},
		{"filled", various, 3, 2, Filled},
	} {
		if got := test.g.Get(test.row, test.col); got != test.want {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got)
		}
	}
}

func TestPieceFits(t *testing.T) {
	type fitTest struct {
		name     string
		g        Grid
		p        Piece
		row, col int
		want     bool
	}
	oneByOne := RectPiece{1, 1}
	twoByTwo := RectPiece{2, 2}
	twoByFour := RectPiece{2, 4}

	various := MakeGrid(5, 6)
	various.state[4][3] = ToBeFilled
	various.state[3][2] = Filled
	for _, test := range []fitTest{
		{"empty - no space", MakeFilledGrid(5, 5, Empty), oneByOne, 0, 0, false},
		{"filled case", MakeFilledGrid(5, 5, Filled), oneByOne, 0, 0, false},
		{"to be filled - with space (0, 0)", MakeFilledGrid(5, 5, ToBeFilled), oneByOne, 0, 0, true},
		{"to be filled - with space (4, 4)", MakeFilledGrid(5, 5, ToBeFilled), oneByOne, 4, 4, true},
		{"to be filled - with space, out of bounds (5, 4)", MakeFilledGrid(5, 5, ToBeFilled), oneByOne, 5, 4, false},
		{"2x2 doesn't fit in 1x1 spot", MakeFilledGrid(1, 1, ToBeFilled), twoByTwo, 0, 0, false},
		{"2x2 does fit in 2x2 spot", MakeFilledGrid(2, 2, ToBeFilled), twoByTwo, 0, 0, true},
		{"2x2 does fit in 2x2 spot, but not when offset", MakeFilledGrid(2, 2, ToBeFilled), twoByTwo, -1, 0, false},
		{"2x4 fits in horizontal grid", MakeFilledGrid(2, 4, ToBeFilled), twoByFour, 0, 0, true},
		{"2x4 does not fit in vertical grid", MakeFilledGrid(4, 2, ToBeFilled), twoByFour, 0, 0, false},
	} {
		if got := test.g.PieceFits(test.p, Location{test.row, test.col}); got != test.want {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got)
		}
	}
}

func TestSolve(t *testing.T) {
	type solveTest struct {
		name   string
		g      Grid
		p      []Piece
		want   map[Location]Piece
		hasErr bool
	}
	oneByOne := RectPiece{1, 1}
	oneByFour := RectPiece{1, 4}
	twoByTwo := RectPiece{2, 2}
	twoByFour := RectPiece{2, 4}
	fourByOne := RectPiece{4, 1}
	for _, test := range []solveTest{
		{
			"cannot be solved - no pieces",
			MakeFilledGrid(1, 1, ToBeFilled),
			[]Piece{},
			make(map[Location]Piece),
			true,
		},
		{
			"trivially solved - one piece",
			MakeFilledGrid(1, 1, ToBeFilled),
			[]Piece{oneByOne},
			map[Location]Piece{
				Location{0, 0}: oneByOne,
			},
			false,
		},
		{
			"trivially solved - one piece, 2x2",
			MakeFilledGrid(2, 2, ToBeFilled),
			[]Piece{twoByFour, twoByTwo, oneByOne},
			map[Location]Piece{
				Location{0, 0}: twoByTwo,
			},
			false,
		},
		{
			"5 x 5 grid - 2 2x4",
			MakeFilledGrid(5, 5, ToBeFilled),
			[]Piece{twoByFour, twoByTwo, oneByFour, fourByOne, oneByOne},
			map[Location]Piece{
				Location{0, 0}: twoByFour,
				Location{2, 0}: twoByFour,
				Location{0, 4}: fourByOne,
				Location{4, 0}: oneByFour,
				Location{4, 4}: oneByOne,
			},
			false,
		},
	} {
		got, err := test.g.Solve(test.p)
		if err != nil && !test.hasErr {
			t.Errorf("for %q wanted no error got %v", test.name, err)
		} else if err == nil && test.hasErr {
			t.Errorf("for %q should have had an error", test.name)
		}
		if !reflect.DeepEqual(got.Pieces, test.want) {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got.Pieces)
		}
	}

}
