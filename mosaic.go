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
	placedBricks map[Location]PlacedBrick
}

func (g* gridBasedPlan) Orig() Ideal {
  return g.img
}

func (g* gridBasedPlan) Pieces() []PlacedBrick {
  var bricks []PlacedBrick
  for _, b := range g.placedBricks {
    bricks = append(bricks, b)
  }
  return bricks
}

func (g *gridBasedPlan) Piece(row, col int) PlacedBrick {
  return g.placedBricks[Location{row, col}]
}

func (g *gridBasedPlan) Inventory() Inventory {
  i := MakeInventory()
  for _, p := range g.Pieces() {
    i.Add(p.Color, p.Shape)
  }
  return i
}

func CreateGridMosaic(m Ideal) Plan {
  grids := makeGrids(m)
  
  // TODO(ndunn): how do I inject which pieces are allowed?
  allPieces := PiecesForOrientation(m.Orientation(), allBrickPieces())
  // cast all pieces to Piece rather than MosaicPiece
  rawPieces := make([]Piece, len(allPieces))
  for i, d := range allPieces {
    rawPieces[i] = d
  }
	solutions := make(map[BrickColor]Solution)
	placedBricks := make(map[Location]PlacedBrick)
	for color, grid := range grids {
		solution, _ := grid.Solve(rawPieces)
		solutions[color] = solution
		
		// Now we know where each piece goes. Create PlacedBrick representations of the pieces.
		counter := 0
  	for loc, piece := range solution.Pieces {
  	  // We know that each entry is not just a Piece but a MosaicPiece
  	  // TODO(ndunn): do we really need BrickPiece, Piece, MosaicPiece, and PlacedBrick?
  	  mp := piece.(MosaicPiece)
  	  pb := PlacedBrick {
      	Id: counter,
      	Origin: loc,
      	Extent: mp.Extent(),
      	Color: color,
      	Shape: mp.Brick,
      	Orientation: m.Orientation(),
  	  }
  	  placedBricks[loc] = pb
  	  counter++
  	}
	}
	return &gridBasedPlan{
		m,
		grids,
		m.Orientation(),
		solutions,
		placedBricks,
	}
}

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