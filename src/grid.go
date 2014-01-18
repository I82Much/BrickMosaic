// grid.go encapsulates a 2d array of coordinates and their state as we build the mosaic. Its coordinates
// are divorced from actual physical coordinates. For instance, we treat all mosaics as the same grid
// coordinate system, regardless of whether we are building them top down, or from the top down, or from
// left to right. This grid abstracts away the physical dimensions of the bricks and allows us to concentrate
// on the core algorithm.
//
// In each orientation, a column is equal to the smallest width piece, and a row is equal to the smallest height
// piece. For instance, if we are looking top down at a mosaic, each row and column is of equal size.
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
	numRows, numCols int
	state            [][]State
}

// GridSolution encapsulates the original requested grid to solve, as well as the solution to that grid,
// mapping location to the brick that goes there.
type GridSolution struct {
	Original Grid
	Pieces   map[Location]Piece
}

// MakeGrid initializes an empty grid of size numRows by numCols.
func MakeGrid(numRows, numCols int) Grid {
	grid := make([][]State, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]State, numCols)
	}
	return Grid{
		numRows, numCols, grid,
	}
}

// MakeFilledGrid is a convenience function for creating
// a grid whose contents are entirely filled with state s.
func MakeFilledGrid(numRows, numCols int, s State) Grid {
	grid := MakeGrid(numRows, numCols)
	grid.Fill(s)
	return grid
}

// Fill mutates the given grid to be completely filled with state s.
func (g *Grid) Fill(s State) {
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			g.state[row][col] = s
		}
	}
}

// Find returns all of the locations that have the given state,
// in row major order.
func (g *Grid) Find(s State) []Location {
	locs := make([]Location, 0)
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			if g.state[row][col] == s {
				locs = append(locs, Location{row, col})
			}
		}
	}
	return locs
}

// Any determines whether any entry in the grid has state s.
func (g *Grid) Any(s State) bool {
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			if g.state[row][col] == s {
				return true
			}
		}
	}
	return false
}

// Get returns the state at row, col in the given grid. If the given
// row, col arguments are out of bouns, the method returns Empty.
func (g *Grid) Get(row, col int) State {
	if row < 0 || row >= g.numRows || col < 0 || col >= g.numCols {
		return Empty
	}
	return g.state[row][col]
}

// Set sets the state at the given (row, column) pair in the grid. If it's out of
// bounds, this does nothing.
func (g *Grid) Set(row, col int, state State) {
	if row < 0 || row >= g.numRows || col < 0 || col >= g.numCols {
		return
	}
	g.state[row][col] = state
}

// PieceFits determines if the given piece can fit at the desired location
// in the grid, where loc is the upper left hand corner of the piece.
func (g *Grid) PieceFits(piece Piece, loc Location) bool {
	relativeLocations := piece.Extent()
	for _, relLoc := range relativeLocations {
		absLoc := relLoc.Add(loc)
		if gridEntry := g.Get(absLoc.row, absLoc.col); gridEntry != ToBeFilled {
			return false
		}
	}
	return true
}

func (g *Grid) Clone() Grid {
	grid := MakeGrid(g.numRows, g.numCols)
	for row := 0; row < grid.numRows; row++ {
		for col := 0; col < grid.numCols; col++ {
			grid.state[row][col] = g.state[row][col]
		}
	}
	return grid
}

// Solve attempts to solve the grid by filling in the missing pieces.
// The pieces are considered in the order defined in the pieces list.
// They should be sorted accordingly, with the best entry first in the
// list (i.e.. least expensive). If the given pieces cannot exactly
// match the missing pieces, returns a non nil error
func (g *Grid) Solve(pieces []Piece) (GridSolution, error) {
	originalGrid := g.Clone()
	locs := make(map[Location]Piece)
	// Use a simple greedy strategy where we work
	// top to bottom, left to right
	for col := 0; col < g.numCols; col++ {
		for row := 0; row < g.numRows; row++ {

			//for row := 0; row < g.numRows; row++ {
			//	for col := 0; col < g.numCols; col++ {
			loc := Location{row, col}
			if g.Get(row, col) == ToBeFilled {
				for _, p := range pieces {
					// We found the best fit! Need to add it to the map, as
					// well as mark the internal state
					if g.PieceFits(p, loc) {
						locs[loc] = p
						for _, pieceLoc := range p.Extent() {
							absLoc := loc.Add(pieceLoc)
							g.state[absLoc.row][absLoc.col] = Filled
						}
					}
				}
			}
		}
	}
	if g.Any(ToBeFilled) {
		return GridSolution{originalGrid, locs}, fmt.Errorf("Following locations must still be filled: %v", g.Find(ToBeFilled))
	}
	return GridSolution{originalGrid, locs}, nil
}

// String returns a string representation of the grid, suitable for display in a terminal.
func (g Grid) String() string {
	result := "[\n"
	for _, row := range g.state {
		// Remove spaces
		result += strings.Replace(fmt.Sprintf("%v\n", row), " ", "", -1)
	}
	result += "\n]"
	return result
}

// TODO(ndunn): Remove, not very necessary given the mosaic rendering in svg
func (solution GridSolution) String() string {
	// TODO(ndunn): support more than 52 pieces
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	pieces := make([][]string, solution.Original.numRows)
	for row := 0; row < solution.Original.numRows; row++ {
		pieces[row] = make([]string, solution.Original.numCols)
		for col := 0; col < solution.Original.numCols; col++ {
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
			pieces[absLoc.row][absLoc.col] = pieceLetter
		}
	}

	result := "  ["
	// column headers
	for i := 0; i < solution.Original.numCols; i++ {
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
