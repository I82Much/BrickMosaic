package grid

// Where 0, 0 = upper left hand corner
type Piece interface {
	Extent() []Location
}

// TODO(ndunn): finish this, sort by price, size (number of grid lcoations it takes up)
// By is the type of a "less" function that defines the ordering of its Piece pieces
type By func(p1, p2 *Piece) bool

type RectPiece struct {
	NumRows int
	NumCols int
}

func (r RectPiece) Extent() []Location {
	locs := make([]Location, r.NumRows*r.NumCols)
	for row := 0; row < r.NumRows; row++ {
		for col := 0; col < r.NumCols; col++ {
			locs = append(locs, Location{row, col})
		}
	}
	return locs
}


