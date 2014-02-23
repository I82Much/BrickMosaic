package BrickMosaic

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type solveTest struct {
		name   string
		g      Grid
		p      []MosaicPiece
		want   map[Location]MosaicPiece
		hasErr bool
	}
	oneByOne := StudsOutPiece(OneByOne)
	oneByFour := StudsOutPiece(OneByFour)
	twoByTwo := StudsOutPiece(TwoByTwo)
	twoByFour := StudsOutPiece(TwoByFour)
	r := RectPiece{4, 1}
	fourByOne := mosaicPiece{
		Brick: OneByFour,
		Rect:  r,
	}

	for _, test := range []solveTest{
		{
			"cannot be solved - no pieces",
			WithState(1, 1, ToBeFilled),
			[]MosaicPiece{},
			make(map[Location]MosaicPiece),
			true,
		},
		{
			"trivially solved - one piece",
			WithState(1, 1, ToBeFilled),
			[]MosaicPiece{oneByOne},
			map[Location]MosaicPiece{
				Location{0, 0}: oneByOne,
			},
			false,
		},
		{
			"trivially solved - one piece, 2x2",
			WithState(2, 2, ToBeFilled),
			[]MosaicPiece{twoByFour, twoByTwo, oneByOne},
			map[Location]MosaicPiece{
				Location{0, 0}: twoByTwo,
			},
			false,
		},
		{
			"5 x 5 grid - 2 2x4",
			WithState(5, 5, ToBeFilled),
			[]MosaicPiece{twoByFour, twoByTwo, oneByFour, fourByOne, oneByOne},
			map[Location]MosaicPiece{
				Location{0, 0}: twoByFour,
				Location{2, 0}: twoByFour,
				Location{0, 4}: fourByOne,
				Location{4, 0}: oneByFour,
				Location{4, 4}: oneByOne,
			},
			false,
		},
	} {
		got, err := GreedySolve(&test.g, test.p)
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
