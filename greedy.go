package BrickMosaic

import (
  "fmt"
)

// Greedy implements a greedy algorithm for placing bricks within the mosaic. 
// It works its way through the grid looking at pieces in descending size order. As soon as
// it finds a piece that fits, it places it and moves on.
// TODO(ndunn): this is the piece that should be an interface

// Solve attempts to solve the grid by filling in the missing pieces.
// The pieces are considered in the order defined in the pieces list.
// They should be sorted accordingly, with the best entry first in the
// list (i.e.. least expensive). If the given pieces cannot exactly
// match the missing pieces, returns a non nil error
func GreedySolve(g *Grid, pieces []MosaicPiece) (Solution, error) {
	originalGrid := g.Clone()
	locs := make(map[Location]MosaicPiece)
	// Use a simple greedy strategy where we work
	// top to bottom, left to right
	for col := 0; col < g.Cols; col++ {
		for row := 0; row < g.Rows; row++ {

			//for row := 0; row < g.numRows; row++ {
			//	for col := 0; col < g.Cols; col++ {
			loc := Location{row, col}
			if g.Get(row, col) == ToBeFilled {
				for _, p := range pieces {
					// We found the best fit! Need to add it to the map, as
					// well as mark the internal state
					if g.PieceFits(p, loc) {
						locs[loc] = p
						for _, pieceLoc := range p.Extent() {
							absLoc := loc.Add(pieceLoc)
							g.State[absLoc.Row][absLoc.Col] = Filled
						}
					}
				}
			}
		}
	}
	if g.Any(ToBeFilled) {
		return Solution{originalGrid, locs}, fmt.Errorf("Following locations must still be filled: %v", g.Find(ToBeFilled))
	}
	return Solution{originalGrid, locs}, nil
}