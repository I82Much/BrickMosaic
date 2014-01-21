package BrickMosaic

import (
	"fmt"
	"testing"
)

type fakeOrig struct{}

func (f fakeOrig) Orientation() ViewOrientation {
	return StudsOut
}
func (f fakeOrig) NumRows() int {
	return 5
}
func (f fakeOrig) NumCols() int {
	return 5
}
func (f fakeOrig) Color(row, col int) BrickColor {
	return BrightRed
}

type fakePlan struct {
	counter int
}

func (f *fakePlan) Orig() Ideal {
	return fakeOrig{}
}

func (f *fakePlan) Pieces() []PlacedBrick {
	var pieces []PlacedBrick
	for row := 0; row < f.Orig().NumRows(); row++ {
		for col := 0; col < f.Orig().NumCols(); col++ {
			pieces = append(pieces, f.Piece(row, col))
		}
	}
	return pieces
}

func (f *fakePlan) Piece(row, col int) PlacedBrick {
	f.counter++
	return PlacedBrick{
		Id: f.counter,
	}
}

func (f *fakePlan) Inventory() Inventory {
	return Inventory{}
}

func TestTerminalRender(t *testing.T) {
	w := WriterRenderer{}
	fmt.Println(w.Render(&fakePlan{}))
}
