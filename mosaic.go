package BrickMosaic

import (
	"github.com/I82Much/BrickMosaic/palette"
	"github.com/I82Much/BrickMosaic/grid"
)

// ViewOrientation represents the orientation of each brick in the mosaic.
type ViewOrientation int

const (
	// StudsUp is a top down view, studs on top. Rows and columns refer to equal distances.
	StudsUp ViewOrientation = iota
	// StudsTop indicates a view from the side - pieces build on top of each other. Rows refer to plate height, 
	// columns are 1x1 width.
	StudsTop
	// StudsRight indicates a view from the side, where the top of a piece faces to the right. Rows refer to 
	// piece width, columns are plate height.
	StudsRight
)

type Mosaic struct {
	img         *BrickImage
	colorGrid   map[palette.BrickColor]grid.Grid
	orientation ViewOrientation
	solutions   map[palette.BrickColor]grid.GridSolution
}

func makeGrids(numRows, numCols uint, colorMap map[grid.Location]palette.BrickColor) map[palette.BrickColor]grid.Grid {
	grids := make(map[palette.BrickColor]grid.Grid)
	for _, color := range colorMap {
		// New color - initialize the grid
		if _, ok := grids[color]; !ok {
			grids[color] = grid.MakeGrid(int(numRows), int(numCols))
		}
	}
	// Set all of the 'to be filled' bits for each color. Every thing else is
	// 'empty' so it won't be filled in with this color.
	for loc, color := range colorMap {
		g := grids[color]
		g.Set(loc.Row, loc.Col, grid.ToBeFilled)
	}
	return grids
}

func MakeMosaic(img *BrickImage, orientation ViewOrientation, pieces []grid.Piece) Mosaic {
	grids := makeGrids(img.rows, img.cols, img.avgColors)
	solutions := make(map[palette.BrickColor]grid.GridSolution)
	for color, grid := range grids {
		solution, _ := grid.Solve(pieces)
		solutions[color] = solution
	}
	return Mosaic{
		img,
		grids,
		orientation,
		solutions,
	}
}

func (m *Mosaic) Grids() map[palette.BrickColor]grid.Grid {
	return m.colorGrid
}

func (m *Mosaic) Solutions() map[palette.BrickColor]grid.GridSolution {
	return m.solutions
}
