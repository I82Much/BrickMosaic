// TODO(ndunn): move to a more appropriate package
package BrickMosaic

// Brick represents a prototypical piece, not bound to any specific orientation or color.
type Brick struct {
	// name is the human readable name for this brick.
	name string
	// id is the the LDRAW identifier for this Brick.
	id string
	// width is measured in studs.
	width int
	// length is measured in studs.
	length int
	// height is measured in terms of plates - a standard brick is 3 plates high.
	height int
}

var (
	// OneByFour represents a 1 x 4 brick. See http://lego.wikia.com/wiki/Part_3010.
	OneByFour = Brick{
		name:   "1x4 brick",
		id:     "3010",
		width:  1,
		length: 4,
		height: 3,
	}
	// OneByThree represents a 1 x 3 brick. See http://lego.wikia.com/wiki/Part_3622.
	OneByThree = Brick{
		name:   "1x3 brick",
		id:     "3622",
		width:  1,
		length: 3,
		height: 3,
	}
	// OneByTwo represents a 1 x 2 brick. See http://lego.wikia.com/wiki/Part_3004.
	OneByTwo = Brick{
		name:   "1x2 brick",
		id:     "3004",
		width:  1,
		length: 2,
		height: 3,
	}
	// OneByOne represents a 1 x 1 brick. See http://lego.wikia.com/wiki/Part_3005.
	OneByOne = Brick{
		name:   "1x1 brick",
		id:     "3005",
		width:  1,
		length: 1,
		height: 3,
	}
	// TwoByFour represents a 2 x 4 brick. See http://lego.wikia.com/wiki/Part_3001.
	TwoByFour = Brick{
		name:   "2x4 brick",
		id:     "3001",
		width:  2,
		length: 4,
		height: 3,
	}
	// TwoByThree represents a 2 x 3 brick. See http://lego.wikia.com/wiki/Part_3002.
	TwoByThree = Brick{
		name:   "2x3 brick",
		id:     "3002",
		width:  2,
		length: 3,
		height: 3,
	}
	// TwoByTwo represents a 2 x 2 brick. See http://lego.wikia.com/wiki/Part_3003.
	TwoByTwo = Brick{
		name:   "2x2 brick",
		id:     "3003",
		width:  2,
		length: 2,
		height: 3,
	}

	// Plates

	// OneByPlate represents a 1 x 1 plate. See http://lego.wikia.com/wiki/Part_3024.
	OneByOnePlate = Brick{
		name:   "1x1 plate",
		id:     "3024",
		width:  1,
		length: 1,
		height: 1,
	}
	// OneByTwoPlate represents a 1 x 2 plate. See http://lego.wikia.com/wiki/Part_3023.
	OneByTwoPlate = Brick{
		name:   "1x2 plate",
		id:     "3023",
		width:  1,
		length: 2,
		height: 1,
	}
	// OneByThreePlate represents a 1 x 3 plate. See http://lego.wikia.com/wiki/Part_3623.
	OneByThreePlate = Brick{
		name:   "1x3 plate",
		id:     "3623",
		width:  1,
		length: 3,
		height: 1,
	}
	// OneByFourPlate represents a 1 x 4 plate. See http://lego.wikia.com/wiki/Part_3710.
	OneByFourPlate = Brick{
		name:   "1x4 plate",
		id:     "3710",
		width:  1,
		length: 4,
		height: 1,
	}
	// OneBySixPlate represents a 1 x 6 plate. See http://lego.wikia.com/wiki/Part_3666.
	OneBySixPlate = Brick{
		name:   "1x6 plate",
		id:     "3666",
		width:  1,
		length: 6,
		height: 1,
	}
	// OneByEightPlate represents a 1 x 8 plate. See http://brickowl.com/catalog/lego-plate-1-x-8-3460.
	OneByEightPlate = Brick{
		name:   "1x8 plate",
		id:     "3460",
		width:  1,
		length: 8,
		height: 1,
	}
	// OneByTenPlate represents a 1 x 10 plate. See http://brickowl.com/catalog/lego-plate-1-x-10-4477.
	OneByTenPlate = Brick{
		name:   "1x10 plate",
		id:     "4477",
		width:  1,
		length: 10,
		height: 1,
	}

	// Bricks represents a slice of all of the bricks (full height, not plates). They are listed in descending
	// order of area.
	Bricks = []Brick{
		TwoByFour,
		TwoByThree,
		TwoByTwo,

		OneByFour,
		OneByThree,
		OneByTwo,
		OneByOne,
	}

	// Plates represents a slice of all of the standard plates (thinner than bricks). They are listed in
	// descending order of area.
	Plates = []Brick{
		OneByTenPlate,
		OneByEightPlate,
		OneBySixPlate,
		OneByFourPlate,
		OneByThreePlate,
		OneByTwoPlate,
		OneByOnePlate,
	}

	// Pieces represents a slice of all of the standard Bricks; the concatenation of Bricks and Plates.
	Pieces = allBricks()
)

func allBricks() []Brick {
	result := make([]Brick, len(Bricks)+len(Plates))
	copy(result, Bricks)
	copy(result[len(Bricks):], Plates)
	return result
}

// MosaicPiece represents a given physical brick in a certain orientation, which determines its extent
// in the 2d grid.
type MosaicPiece struct {
	Brick Brick
	r     RectPiece
	// In whatever orientation the mosaic is facing. e.g. a 2x4 brick when viewed above has size 2x4.
	// When viewed from the side, it has size 3x4 (3 plates high, 4 bricks wide)
	locs []Location
}

func (r MosaicPiece) Extent() []Location {
	return r.locs
}

func StudsUpPiece(piece Brick) MosaicPiece {
	// Studs up, so rows = width, cols = length
	r := RectPiece{piece.width, piece.length}
	return MosaicPiece{
		piece,
		r,
		r.Extent(),
	}
}

func StudsTopPiece(piece Brick) MosaicPiece {
	// Studs to the top on side, so rows = height, cols = length
	r := RectPiece{piece.height, piece.length}
	return MosaicPiece{
		piece,
		r,
		r.Extent(),
	}
}

func StudsRightPiece(piece Brick) MosaicPiece {
	// Studs to the right on its side, so rows = length, cols = height
	r := RectPiece{piece.length, piece.height}
	return MosaicPiece{
		piece,
		r,
		r.Extent(),
	}
}

func PiecesForOrientation(o ViewOrientation, pieces []Brick) []MosaicPiece {
	result := make([]MosaicPiece, len(pieces))
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
