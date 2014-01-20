package BrickMosaic

import (
//	"github.com/I82Much/BrickMosaic/palette"
	//"github.com/I82Much/BrickMosaic/grid"
//	"github.com/I82Much/BrickMosaic/image"
)

// Ideal is the idealized grid of how the mosaic should look. Basically a 2d grid of color.
type Ideal interface  {
	Orientation() ViewOrientation
	NumRows() int
	NumCols() int
	Color(row, col int) BrickColor
}

// PlacedBrick represents a physical brick placed within the mosaic, at a certain location,
// with a certain color, orientation, and shape.
type PlacedBrick struct {
	// Unique identifier for this brick within the mosaic
	Id int
	// Upper left corner of the piece
	Origin Location
	// The relative locations of how big this brick is. Add to origin to get absolute
	// location
	Extent []Location
	// What color is this brick?
	Color BrickColor
	// Characteristics of the brick - 2x4, etc
	Shape BrickPiece 
	// Orientation represents how the brick is placed in the mosaic
	Orientation ViewOrientation 
}


// Plan represents how to build the mosaic. The resulting plan may not match the
// ideal DesiredMosaic perfectly; for instance, an implementation might decide
// to depart slightly from the desired colors if it leads to enhanced rigidity
// in the structure.
type Plan interface {
	Orig() Ideal
	Pieces() []PlacedBrick
	Piece(row, col int) PlacedBrick
	Inventory() Inventory
}

// Creator is the interface by which we convert DesiredMosaic objects into a plan
// for building it. As discussed in Plan, different Creators might build Plans
// that do not perfectly match the DesiredMosaic.
type Creator interface {
	Create(i Ideal) Plan
}


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

// Create is the interface by which we convert Ideal mosaics into a plan
// for building it. As discussed in Plan, different Creators might build Plans
// that do not perfectly match the DesiredMosaic.
type Create func(i Ideal) Plan

// gridBasedPlan is a basic implementation of the Plan interface
type gridBasedPlan struct {
	img         Ideal
	colorGrid   map[BrickColor]Grid
	orientation ViewOrientation
	solutions   map[BrickColor]Solution
}

/*
func (g* gridBasedPlan) Orig() Ideal {
  return g.img
}

func (g* gridBasedPlan) Pieces() []PlacedBrick {

}

func (g *gridBasedPlan) Piece(row, col int) PlacedBrick {
  
}
*/

/*Plan interface {
	Orig() Ideal
	Pieces() []PlacedBrick
	Piece(row, col int) PlacedBrick
	Inventory() Inventory
}*/

//func CreateGridMosaic() {}

// makeGrids is the core piece of the algorithm. For each color in the ideal image, we create a grid whose
// 'TO_BE_FILLED' cells are set to the places in the ideal location for that color. In other words, say we have
// a square image whose upper left quadrant is red, upper right is blue, lower left is black, lower right
// is gray.
// Red grid would be
// 1 1 0 0
// 0 0 0 0
//
// Blue is
// 0 0 1 1
// 0 0 0 0
//
// Black is
// 0 0 0 0
// 1 1 0 0
//
// Gray is
// 0 0 0 0
// 0 0 1 1
//
// Where the 0's indicate Empty and 1 represents ToBeFilled.
func makeGrids(i Ideal) map[BrickColor]Grid {
  grids := make(map[BrickColor]Grid)
  for row := 0; row < i.NumRows(); row++ {
    for col := 0; col < i.NumCols(); col++ {
      color := i.Color(row, col)
      // New color - initialize the grid
  		if _, ok := grids[color]; !ok {
  			grids[color] = NewGrid(i.NumRows(), i.NumCols())
  		}
  		colorGrid := grids[color]
  		// Set all of the 'to be filled' bits. Every thing else is 'empty' so
  		// it won't be filled.
  		colorGrid.Set(row, col, ToBeFilled)
    }
  }
	return grids
}

/*
func MakeMosaic(img *image.BrickImage, orientation ViewOrientation, pieces []grid.Piece) Mosaic {
	grids := makeGrids(img.rows, img.cols, img.avgColors)
	solutions := make(map[BrickColor]grid.Solution)
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

func (m *Mosaic) Grids() map[BrickColor]grid.Grid {
	return m.colorGrid
}

func (m *Mosaic) Solutions() map[BrickColor]grid.Solution {
	return m.solutions
}
*/