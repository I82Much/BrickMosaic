package BrickMosaic

import (
	"reflect"
	"testing"
)

// Make sure all bricks are accounted for
func TestAllBricks(t *testing.T) {
	allPieces := make(map[BrickPiece]bool)
	for _, b := range Bricks {
		allPieces[b] = true
	}
	for _, p := range Plates {
		allPieces[p] = true
	}
	foundPieces := make(map[BrickPiece]bool)
	for _, b := range Pieces {
		foundPieces[b] = true
	}
	if len(foundPieces) != len(allPieces) {
		t.Errorf("got %d pieces wanted %d pieces", len(foundPieces), len(allPieces))
	}
	if !reflect.DeepEqual(foundPieces, allPieces) {
		t.Errorf("got %v wanted %v", foundPieces, allPieces)
	}
}

type pieceLocWant struct {
	name  string
	piece BrickPiece
	want  Piece
}

func TestStudsUpPiece(t *testing.T) {
	for _, test := range []pieceLocWant{
		{
			"1x1",
			OneByOne,
			RectPiece{1, 1},
		},
		{
			"2x4",
			TwoByFour,
			RectPiece{2, 4},
		},
		{
			"1x4 plate",
			OneByFourPlate,
			RectPiece{1, 4},
		},
	} {
		if got := StudsUpPiece(test.piece); !reflect.DeepEqual(got.Extent(), test.want.Extent()) {
			t.Errorf("%q want %v got %v", test.name, test.want, got)
		}
	}
}

func TestStudsTopPiece(t *testing.T) {
	for _, test := range []pieceLocWant{
		{
			"1x1",
			OneByOne,
			// 3 plates high, 1 across
			RectPiece{3, 1},
		},
		{
			"2x4",
			TwoByFour,
			// 3 plates high, 4 across
			RectPiece{3, 4},
		},
		{
			"1x4 plate",
			OneByFourPlate,
			// 1 plate high, 4 across
			RectPiece{1, 4},
		},
	} {
		if got := StudsTopPiece(test.piece); !reflect.DeepEqual(got.Extent(), test.want.Extent()) {
			t.Errorf("%q want %v got %v", test.name, test.want.Extent(), got.Extent())
		}
	}

}

func TestStudsRightPiece(t *testing.T) {
	for _, test := range []pieceLocWant{
		{
			"1x1",
			OneByOne,
			// 1 brick high, 3 plates long
			RectPiece{1, 3},
		},
		{
			"2x4",
			TwoByFour,
			// 4 bricks high, 3 plates long
			RectPiece{4, 3},
		},
		{
			"1x4 plate",
			OneByFourPlate,
			// 4 bricks high, 1 plate long
			RectPiece{4, 1},
		},
	} {
		if got := StudsRightPiece(test.piece); !reflect.DeepEqual(got.Extent(), test.want.Extent()) {
			t.Errorf("%q want %v got %v", test.name, test.want.Extent(), got.Extent())
		}
	}

}
