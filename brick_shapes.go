// TODO(ndunn): move to a more appropriate package
package BrickMosaic

type BrickPiece struct {
  // name is the human readable name for this brick.
	name string
	// id is the the LDRAW identifier for this piece.
	id string
	// width is measured in studs.
	width int
	// length is measured in studs.
	length int
	// height is measured in terms of plates - a standard brick is 3 plates high.
	height int
}

var (
	OneByFour = BrickPiece{
		name:   "1x4 brick",
		id:     "3010",
		width:  1,
		length: 4,
		height: 3,
	}
	OneByThree = BrickPiece{
		name:   "1x3 brick",
		id:     "3622",
		width:  1,
		length: 3,
		height: 3,
	}
	OneByTwo = BrickPiece{
		name:   "1x2 brick",
		id:     "3004",
		width:  1,
		length: 2,
		height: 3,
	}
	OneByOne = BrickPiece{
		name:   "1x1 brick",
		id:     "3005",
		width:  1,
		length: 1,
		height: 3,
	}
	TwoByFour = BrickPiece{
		name:   "2x4 brick",
		id:     "3001",
		width:  2,
		length: 4,
		height: 3,
	}
	TwoByThree = BrickPiece{
		name:   "2x3 brick",
		id:     "3002",
		width:  2,
		length: 3,
		height: 3,
	}
	TwoByTwo = BrickPiece{
		name:   "2x2 brick",
		id:     "3003",
		width:  2,
		length: 2,
		height: 3,
	}

	// Plates
	OneByOnePlate = BrickPiece{
		name:   "1x1 plate",
		id:     "3024",
		width:  1,
		length: 1,
		height: 1,
	}
	OneByTwoPlate = BrickPiece{
		name:   "1x2 plate",
		id:     "3023",
		width:  1,
		length: 2,
		height: 1,
	}
	OneByThreePlate = BrickPiece{
		name:   "1x3 plate",
		id:     "3623",
		width:  1,
		length: 3,
		height: 1,
	}
	OneByFourPlate = BrickPiece{
		name:   "1x4 plate",
		id:     "3710",
		width:  1,
		length: 4,
		height: 1,
	}
	OneBySixPlate = BrickPiece{
		name:   "1x6 plate",
		id:     "3666",
		width:  1,
		length: 6,
		height: 1,
	}
	OneByEightPlate = BrickPiece{
		name:   "1x8 plate",
		id:     "3460",
		width:  1,
		length: 8,
		height: 1,
	}
	OneByTenPlate = BrickPiece{
		name:   "1x10 plate",
		id:     "4477",
		width:  1,
		length: 10,
		height: 1,
	}
	Bricks = []BrickPiece{
	  TwoByFour,
		TwoByThree,
		TwoByTwo,
		
		OneByFour,
		OneByThree,
		OneByTwo,
		OneByOne,
	}
	Plates = []BrickPiece{
	  OneByTenPlate,
	  OneByEightPlate,
	  OneBySixPlate,
		OneByFourPlate,
		OneByThreePlate,
		OneByTwoPlate,	  
		OneByOnePlate,
	}
	Pieces = allPieces()//append(make([]BrickPiece, 0), Bricks..., Plates...)
)

func allPieces() []BrickPiece {
	result := make([]BrickPiece, len(Bricks)+len(Plates))
	copy(result, Bricks)
	copy(result[len(Bricks):], Plates)
	return result
}

// The LDraw Co-ordinate System
// LDraw uses a right-handed co-ordinate system where -Y is "up".
// http://www.ldraw.org/article/218.html

// -Y - up
// +Z - into the screen
// +X - to the right

// All measurements are in terms of LDU
type LDU int

const (
	BrickWidth   LDU = 20
	BrickHeight  LDU = 24
	PlateHeight  LDU = 8
	StudDiameter LDU = 12
	StudHeight   LDU = 4
)

func ApproxSizeInch(size LDU) float32 {
	return float32(1.0 * size / 64.0)
}

func ApproxSizeMm(size LDU) float32 {
	return 0.4 * float32(size)
}
