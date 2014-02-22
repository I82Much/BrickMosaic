// package grid represents a 2d array of coordinates and their state as we build the mosaic. Its coordinates
// are divorced from the physical coordinate system. For instance, we treat all mosaics as the same grid
// coordinate system, regardless of whether we are building them top down, or from the top down, or from
// left to right. This grid abstracts away the physical dimensions of the bricks and allows us to concentrate
// on the core algorithm. This is handled via the Piece interface defined in piece.go.
//
// In each orientation, a column of the grid has width equal to the smallest width piece, and a
// row is as high as the smallest height piece.
// For instance, if we are looking top down at a mosaic, each row and column is of equal size.
// If we are building a studs up mosaic, the row is equal in height to a single plate, and the column
// would be equal to the width of a single 1x1. If we are building a studs right mosaic, the row is
// equal to the width of a single 1x1 and the column would be equal to the height of a plate.
//
// The grid is indexed such that [0][0] is the upper left corner and [numRows-1][numColumns-1] is the
// lower right corner.
package BrickMosaic

import (
	"fmt"
	"strings"
)

// GridSolver is the interface for fitting pieces into the given grid.
type GridSolver func (g *Grid, pieces []Piece) (Solution, error)

// Location represents one cell in the grid.
type Location struct {
	Row, Col int
}

// Add adds one location to another (vector addition).
func (loc Location) Add(loc2 Location) Location {
	return Location{loc.Row + loc2.Row, loc.Col + loc2.Col}
}

// State is an enum representing the state of a location in the grid.
type State int

const (
	// Empty indicates that nothing is in the grid location, nor should there be.
	Empty State = iota
	// ToBeFilled indicates that there is nothing currently in the grid location, but there should be.
	ToBeFilled State = Empty + 1
	// Filled indicates that there is already something in the grid location.
	Filled State = ToBeFilled + 1
)

// Grid is an abstract representation of the mosaic to assemble.
type Grid struct {
	Rows, Cols int
	State            [][]State
}

// Solution encapsulates the original requested grid to solve, as well as the solution to that grid,
// mapping location to the brick that goes there.
type Solution struct {
	Original Grid
	Pieces   map[Location]Piece
}

// NewGrid returns an empty grid of size numRows by numCols.
func NewGrid(numRows, numCols int) Grid {
	grid := make([][]State, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]State, numCols)
	}
	return Grid{
		Rows: numRows, 
		Cols: numCols, 
		State: grid,
	}
}

// WithState is a convenience function for creating
// a grid whose contents are entirely filled with state s.
func WithState(numRows, numCols int, s State) Grid {
	grid := NewGrid(numRows, numCols)
	grid.Fill(s)
	return grid
}

// Fill mutates the given grid to be completely filled with state s.
func (g *Grid) Fill(s State) {
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			g.State[row][col] = s
		}
	}
}

// Find returns all of the locations that have the given state,
// in row major order.
func (g *Grid) Find(s State) []Location {
	locs := make([]Location, 0)
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			if g.State[row][col] == s {
				locs = append(locs, Location{row, col})
			}
		}
	}
	return locs
}

// Any determines whether any entry in the grid has state s.
func (g *Grid) Any(s State) bool {
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			if g.State[row][col] == s {
				return true
			}
		}
	}
	return false
}

// outOfBounds determines if the given row/col is out of bounds (not a valid index into the data structure).
func (g *Grid) outOfBounds(row, col int) bool {
	return row < 0 || row >= g.Rows || col < 0 || col >= g.Cols
}

// Get returns the state at row, col in the given grid. If the given
// row, col arguments are out of bouns, the method returns Empty.
func (g *Grid) Get(row, col int) State {
	if g.outOfBounds(row, col) {
		return Empty
	}
	return g.State[row][col]
}

// Set sets the state at the given (row, column) pair in the grid. If it's out of
// bounds, this does nothing.
func (g *Grid) Set(row, col int, state State) {
	if g.outOfBounds(row, col) {
		return
	}
	g.State[row][col] = state
}

// PieceFits determines if the given piece can fit at the desired location
// in the grid, where loc is the upper left hand corner of the piece. Note that
// orientation is already baked into the Extent() of the piece, which is why it is
// not an argument to this method.
func (g *Grid) PieceFits(piece Piece, loc Location) bool {
	relativeLocations := piece.Extent()
	for _, relLoc := range relativeLocations {
		absLoc := relLoc.Add(loc)
		if gridEntry := g.Get(absLoc.Row, absLoc.Col); gridEntry != ToBeFilled {
			return false
		}
	}
	return true
}

// Clone returns a copy of the grid.
func (g *Grid) Clone() Grid {
	grid := NewGrid(g.Rows, g.Cols)
	for row := 0; row < grid.Rows; row++ {
		for col := 0; col < grid.Cols; col++ {
			grid.State[row][col] = g.State[row][col]
		}
	}
	return grid
}

// String returns a string representation of the grid, suitable for display in a terminal.
func (g Grid) String() string {
	result := "[\n"
	for _, row := range g.State {
		// Remove spaces
		result += strings.Replace(fmt.Sprintf("%v\n", row), " ", "", -1)
	}
	result += "\n]"
	return result
}

// TODO(ndunn): Remove, not very necessary given the mosaic rendering in svg
func (solution Solution) String() string {
	// TODO(ndunn): support more than 52 pieces
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	pieces := make([][]string, solution.Original.Rows)
	for row := 0; row < solution.Original.Rows; row++ {
		pieces[row] = make([]string, solution.Original.Cols)
		for col := 0; col < solution.Original.Cols; col++ {
			pieces[row][col] = "_"
		}
	}

	pieceIndex := -1
	// Figure out which pieces fill up which spaces in the grid
	for loc, piece := range solution.Pieces {
		pieceIndex++
		pieceIndex = pieceIndex % len(pieces)
		pieceLetter := alphabet[pieceIndex]
		// The locations of the extent are relative to upper left hand corner of the piece
		for _, relLoc := range piece.Extent() {
			absLoc := relLoc.Add(loc)
			pieces[absLoc.Row][absLoc.Col] = pieceLetter
		}
	}

	result := "  ["
	// column headers
	for i := 0; i < solution.Original.Cols; i++ {
		if i%3 == 0 {
			result += "|"
		} else {
			result += " "
		}
	}
	result += "\n"

	for i, row := range pieces {
		// Remove spaces
		if i%4 == 0 {
			result += fmt.Sprintf("%2d", i+1) + strings.Replace(fmt.Sprintf("%v\n", row), " ", "", -1)
		} else {
			result += "  " + strings.Replace(fmt.Sprintf("%v\n", row), " ", "", -1)
		}
	}
	result += "\n]"
	return result
}
