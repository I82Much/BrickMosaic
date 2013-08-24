package BrickMosaic

// Where 0, 0 = upper left hand corner
type Piece interface {
	Extent() []Location
}

// TODO(ndunn): finish this, sort by price, size (number of grid lcoations it takes up)
// By is the type of a "less" function that defines the ordering of its Piece pieces
type By func(p1, p2 *Piece) bool

type RectPiece struct {
	numRows, numCols int
}

func (r RectPiece) Extent() []Location {
	locs := make([]Location, r.numRows*r.numCols)
	for row := 0; row < r.numRows; row++ {
		for col := 0; col < r.numCols; col++ {
			locs = append(locs, Location{row, col})
		}
	}
	return locs
}

// TODO(ndunn): support non-rectangular shapes
// TODO(ndunn): color?
type MosaicPiece struct {
	Brick BrickPiece
	// In whatever orientation the mosaic is facing. e.g. a 2x4 brick when viewed above has size 2x4.
	// When viewed from the side, it has size 3x4 (3 plates high, 4 bricks wide)
	locs []Location
}

func (r MosaicPiece) Extent() []Location {
	return r.locs
}

func StudsUpPiece(piece BrickPiece) MosaicPiece {
	// Studs up, so rows = width, cols = length
	r := RectPiece{piece.width, piece.length}
	return MosaicPiece{
		piece,
		r.Extent(),
	}
}

func StudsTopPiece(piece BrickPiece) MosaicPiece {
	// Studs to the top on side, so rows = height, cols = length
	r := RectPiece{piece.height, piece.length}
	return MosaicPiece{
		piece,
		r.Extent(),
	}
}

func StudsRightPiece(piece BrickPiece) MosaicPiece {
	// Studs to the right on its side, so rows = length, cols = height
	r := RectPiece{piece.length, piece.height}
	return MosaicPiece{
		piece,
		r.Extent(),
	}
}

func PiecesForOrientation(o ViewOrientation, pieces []BrickPiece) []Piece {
	result := make([]Piece, len(pieces))
	switch o {
	case StudsUp:
		for i, p := range pieces {
			result[i] = StudsUpPiece(p)
		}
	case StudsTop:
		for i, p := range pieces {
			result[i] = StudsTopPiece(p)
		}
	case StudsRight:
		for i, p := range pieces {
			result[i] = StudsRightPiece(p)
		}
	}
	return result
}
