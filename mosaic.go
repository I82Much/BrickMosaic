package BrickMosaic

import (
/*  "image"
"image/color"*/
)

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
	colorGrid   map[BrickColor]Grid
	orientation ViewOrientation
	solutions   map[BrickColor]GridSolution
}

func makeGrids(numRows, numCols uint, colorMap map[Location]BrickColor) map[BrickColor]Grid {
	grids := make(map[BrickColor]Grid)
	for _, color := range colorMap {
		// New color - initialize the grid
		if _, ok := grids[color]; !ok {
			grids[color] = MakeGrid(int(numRows), int(numCols))
		}
	}
	// Set all of the 'to be filled' bits for each color. Every thing else is
	// 'empty' so it won't be filled in with this color.
	for loc, color := range colorMap {
		grid := grids[color]
		grid.Set(loc.row, loc.col, ToBeFilled)
	}
	return grids
}

func MakeMosaic(img *BrickImage, orientation ViewOrientation, pieces []Piece) Mosaic {
	grids := makeGrids(img.rows, img.cols, img.avgColors)
	solutions := make(map[BrickColor]GridSolution)
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

func (m *Mosaic) Grids() map[BrickColor]Grid {
	return m.colorGrid
}

func (m *Mosaic) Solutions() map[BrickColor]GridSolution {
	return m.solutions
}
