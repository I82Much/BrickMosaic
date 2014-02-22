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
			loc := Location{row, col}
			if g.Get(row, col) == ToBeFilled {
				for _, p := range pieces {
					// We found the best fit! Need to add it to the map, as
					// well as mark the internal state
					if g.PieceFits(p.Extent(), loc) {
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

func SymmetricalGreedySolve(g *Grid, pieces []MosaicPiece) (Solution, error) {
	originalGrid := g.Clone()
	locs := make(map[Location]MosaicPiece)
	// Alternate going left and right working towards the middle.
	// Alternate rows from top to bottom
	for colIndex := 0; colIndex < g.Cols; colIndex++ {
		for rowIndex := 0; rowIndex < g.Rows; rowIndex++ {
		  left := colIndex % 2 == 0
		  top := rowIndex % 2 == 0
		  var anchorPoint AnchorPoint
		  if left && top {
		    anchorPoint = UpperLeft
		  } else if left && !top {
		    anchorPoint = LowerLeft		    
		  } else if !left && top {
		    anchorPoint = UpperRight		    
		  } else if !left && !top {
		    anchorPoint = LowerRight		    
		  } else {
		    panic("shouldn't reach here")
		  }

			// On even calls we'll go left. On odd calls we'll go right
			colOffset := colIndex / 2
			var col int
			if colIndex%2 == 0 {
				col = colOffset
			} else {
				col = (g.Cols - colOffset) - 1
			}

			// {0} - works (0 / 2 = 0)

			// {0, 1} - works
			// 0 / 2 = 0. even, 0
			// 1 / 2 = 0. odd, 2 - 0 - 1 = 1. works

			// {0, 1, 2}
			// 0 / 2 = 0. even, 0
			// 1 / 2 = 0. odd, 3 - 0 - 1 = 2. right
			// 2 / 2 = 1. even. 1. right

			// {0, 1, 2, 3}
			// 0 / 2 = 0. even, 0
			// 1 /2 = 0. odd, 4 - 0 - 1 = 3. right
			// 2 / 2 = 1. even. 1. right.
			// 3 / 2 = 1. odd. 4 - 1 - 1 = 2. right

			rowOffset := rowIndex / 2
			var row int
			if rowIndex % 2 == 0 {
				row = rowOffset
			} else {
				row = (g.Rows - rowOffset) - 1
			}
			loc := Location{row, col}
			if g.Get(row, col) == ToBeFilled {
				for _, p := range pieces {
				  // Based on anchor point, translate the extent accordingly
				  translated := Translate(p.Extent(), anchorPoint)
				  
					// We found the best fit! Need to add it to the map, as
					// well as mark the internal state
					if g.PieceFits(translated, loc) {
					  // Translate the absolute location here as to where the absolute location of the
					  // upper left corner of the piece is located.
					  upperLeft := TranslateAbsoluteOrigin(loc, p, anchorPoint)
						locs[upperLeft] = p
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
