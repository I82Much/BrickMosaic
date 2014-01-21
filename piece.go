package BrickMosaic

// Piece is the virtual representation of a physical brick. By abstracting out the orientation of the
// brick in real world space, it allows us to use the same algorithm for creating mosaics (solving grids)
// in any orientation.
//
// For instance, say that we have a physical 2x4 brick. Depending on which way we orient the brick, it has
// different dimensions in the grid. If it is placed studs up, it is 3 plates (3 rows) high and 4 studs (4 columns)
// wide. This would be represented by a RectPiece with 3 rows and 4 columns, and its extent would be given by
// (0, 0), (0, 1), (0, 2), (0, 3)
// (1, 0), (1, 1), (1, 2), (1, 3)
// (2, 0), (2, 1), (2, 2), (2, 3).
//
// In picture form:
//    +--+   +--+   +--+   +--+
// +--+--+---+--+---+--+-- +--+--+  +
// |                             |  |
// |                             |  |
// |                             |  |  3 plates high
// |                             |  |
// |                             |  |
// |                             |  |
// +-----------------------------+  +
//
// +-----------------------------+
//
//               4 studs wide
//
// Say that we instead are looking down on the brick and building our mosaics that way, with the studs facing out
// towards the viewer. We can orient this in two directions - vertically or horizontally. In the vertical case,
// we would say that the Piece has 4 rows (4 studs) and 2 columns (2 studs).
//
//
//
//  +--------------+
//  |              |
//  | +--+    +--+ |
//  | |  |    |  | |
//  | +--+    +--+ |
//  |              |
//  | +--+    +--+ |
//  | |  |    |  | |
//  | +--+    +--+ |
//  |              |
//  | +--+    +--+ |
//  | |  |    |  | |
//  | +--+    +--+ |
//  |              |
//  | +--+    +--+ |
//  | |  |    |  | |
//  | +--+    +--+ |
//  |              |
//  +--------------+
//
// In the horizontal case it's 2 rows (2 studs) and 4 columns (4 studs).
//
// Note that the physical meaning of the dimension of each cell in the grid is dependent upon the
// orientation of the pieces, but completely unnecessary in the algorithm of filling in the pieces.
//
// http://www.asciiflow.com/#Draw6066488077853425315/674394801
type Piece interface {
	Extent() []Location
}

// TODO(ndunn): finish this, sort by price, size (number of grid lcoations it takes up)
// By is the type of a "less" function that defines the ordering of its Piece pieces
type By func(p1, p2 *Piece) bool

// RectPiece represents a rectangular piece.
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
