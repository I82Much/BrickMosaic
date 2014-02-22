package BrickMosaic

// Brick represents a prototypical piece, not bound to any specific orientation or color.
type Brick interface {
	Name() string
	Id() string
	Width() int
	Length() int
	Height() int
	// TODO(ndunn): This somehow has to take into account color
	// Cost in cents.
	ApproximateCost() int
}

// brick represents a prototypical piece, not bound to any specific orientation or color.
type brick struct {
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

	// Cost in cents
	cost int
}

func (b brick) Name() string {
	return b.name
}

func (b brick) Id() string {
	return b.id
}

func (b brick) Width() int {
	return b.width
}

func (b brick) Length() int {
	return b.length
}

func (b brick) Height() int {
	return b.height
}

func (b brick) ApproximateCost() int {
	return b.cost
}

var (
	// OneByEight represents a 1 x 8 brick. See http://lego.wikia.com/wiki/Part_3008.
	OneByEight = brick{
		name:   "1x8 brick",
		id:     "3008",
		width:  1,
		length: 8,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=326&sz=10&searchSort=P
		cost: 30,
	}

	// OneBySix represents a 1 x 6 brick. See http://lego.wikia.com/wiki/Part_3009.
	OneBySix = brick{
		name:   "1x6 brick",
		id:     "3009",
		width:  1,
		length: 6,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=316&sz=10&searchSort=P
		cost: 18,
	}

	// OneByFour represents a 1 x 4 brick. See http://lego.wikia.com/wiki/Part_3010.
	OneByFour = brick{
		name:   "1x4 brick",
		id:     "3010",
		width:  1,
		length: 4,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=326&sz=10&searchSort=P
		cost: 7,
	}

	// OneByThree represents a 1 x 3 brick. See http://lego.wikia.com/wiki/Part_3622.
	OneByThree = brick{
		name:   "1x3 brick",
		id:     "3622",
		width:  1,
		length: 3,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=632&sz=10&searchSort=P
		cost: 7,
	}
	// OneByTwo represents a 1 x 2 brick. See http://lego.wikia.com/wiki/Part_3004.
	OneByTwo = brick{
		name:   "1x2 brick",
		id:     "3004",
		width:  1,
		length: 2,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=280&sz=10&searchSort=P
		cost: 2,
	}
	// OneByOne represents a 1 x 1 brick. See http://lego.wikia.com/wiki/Part_3005.
	OneByOne = brick{
		name:   "1x1 brick",
		id:     "3005",
		width:  1,
		length: 1,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=299&sz=10&searchSort=P
		cost: 4,
	}

	// TwoByEight represents a 2 x 8 brick. See http://lego.wikia.com/wiki/Part_3007.
	TwoByEight = brick{
		name:   "2x8 brick",
		id:     "3007",
		width:  2,
		length: 8,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=306&sz=10&searchSort=P
		cost: 27,
	}

	// TwoBySix represents a 2 x 6 brick. See http://lego.wikia.com/wiki/Part_2456.
	TwoBySix = brick{
		name:   "2x6 brick",
		id:     "2456", // AKA 44237
		width:  2,
		length: 6,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=110&sz=10&searchSort=P
		cost: 25,
	}

	// TwoByFour represents a 2 x 4 brick. See http://lego.wikia.com/wiki/Part_3001.
	TwoByFour = brick{
		name:   "2x4 brick",
		id:     "3001",
		width:  2,
		length: 4,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=280&sz=10&searchSort=P
		cost: 14,
	}
	// TwoByThree represents a 2 x 3 brick. See http://lego.wikia.com/wiki/Part_3002.
	TwoByThree = brick{
		name:   "2x3 brick",
		id:     "3002",
		width:  2,
		length: 3,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=266&sz=10&searchSort=P
		cost: 9,
	}
	// TwoByTwo represents a 2 x 2 brick. See http://lego.wikia.com/wiki/Part_3003.
	TwoByTwo = brick{
		name:   "2x2 brick",
		id:     "3003",
		width:  2,
		length: 2,
		height: 3,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=272&sz=10&searchSort=P
		cost: 4,
	}

	// Plates

	// OneByPlate represents a 1 x 1 plate. See http://lego.wikia.com/wiki/Part_3024.
	OneByOnePlate = brick{
		name:   "1x1 plate",
		id:     "3024",
		width:  1,
		length: 1,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=381&sz=10&searchSort=P
		cost: 5,
	}
	// OneByTwoPlate represents a 1 x 2 plate. See http://lego.wikia.com/wiki/Part_3023.
	OneByTwoPlate = brick{
		name:   "1x2 plate",
		id:     "3023",
		width:  1,
		length: 2,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=378&sz=10&searchSort=P
		cost: 2,
	}
	// OneByThreePlate represents a 1 x 3 plate. See http://lego.wikia.com/wiki/Part_3623.
	OneByThreePlate = brick{
		name:   "1x3 plate",
		id:     "3623",
		width:  1,
		length: 3,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=634&sz=10&searchSort=P
		cost: 5,
	}
	// OneByFourPlate represents a 1 x 4 plate. See http://lego.wikia.com/wiki/Part_3710.
	OneByFourPlate = brick{
		name:   "1x4 plate",
		id:     "3710",
		width:  1,
		length: 4,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=683&sz=10&searchSort=P
		cost: 3,
	}
	// OneBySixPlate represents a 1 x 6 plate. See http://lego.wikia.com/wiki/Part_3666.
	OneBySixPlate = brick{
		name:   "1x6 plate",
		id:     "3666",
		width:  1,
		length: 6,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=659&sz=10&searchSort=P
		cost: 4,
	}
	// OneByEightPlate represents a 1 x 8 plate. See http://brickowl.com/catalog/lego-plate-1-x-8-3460.
	OneByEightPlate = brick{
		name:   "1x8 plate",
		id:     "3460",
		width:  1,
		length: 8,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=604&sz=10&searchSort=P
		cost: 8,
	}
	// OneByTenPlate represents a 1 x 10 plate. See http://brickowl.com/catalog/lego-plate-1-x-10-4477.
	OneByTenPlate = brick{
		name:   "1x10 plate",
		id:     "4477",
		width:  1,
		length: 10,
		height: 1,
		// http://www.bricklink.com/search.asp?pg=1&colorID=11&itemID=908&sz=10&searchSort=P
		cost: 10,
	}

	// Bricks represents a slice of all of the bricks (full height, not plates). They are listed in descending
	// order of area.
	Bricks = []Brick{
		TwoByEight,
		TwoBySix,
		TwoByFour,
		TwoByThree,
		TwoByTwo,

		OneByEight,
		OneBySix,
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

// MosaicPiece represents a given physical brick in a certain orientation, which determines
// its extent in the 2d grid.
type MosaicPiece interface {
	Brick
	Piece
	Rows() int
	Cols() int
}

type mosaicPiece struct {
	Brick Brick
	// In whatever orientation the mosaic is facing. e.g. a 2x4 brick when viewed above has size 2x4.
	// When viewed from the side, it has size 3x4 (3 plates high, 4 bricks wide)
	Rect RectPiece
}

// Extent() fulfills Extent interface
func (r mosaicPiece) Extent() []Location {
	return r.Rect.Extent()
}

func (r mosaicPiece) Name() string {
	return r.Brick.Name()
}

func (r mosaicPiece) Id() string {
	return r.Brick.Id()
}

func (r mosaicPiece) Width() int {
	return r.Brick.Width()
}

func (r mosaicPiece) Length() int {
	return r.Brick.Length()
}

func (r mosaicPiece) Height() int {
	return r.Brick.Height()
}

func (r mosaicPiece) Rows() int {
	return r.Rect.NumRows
}

func (r mosaicPiece) Cols() int {
	return r.Rect.NumCols
}

func (r mosaicPiece) ApproximateCost() int {
	return r.Brick.ApproximateCost()
}

// TODO(ndunn): This could either be facing horizontally or vertically. This is not taking
// that into consideration.
func StudsOutPiece(piece Brick) MosaicPiece {
	// Studs up, so rows = width, cols = length
	r := RectPiece{piece.Width(), piece.Length()}
	return mosaicPiece{
		Brick: piece,
		Rect:  r,
	}
}

func StudsTopPiece(piece Brick) MosaicPiece {
	// Studs to the top on side, so rows = height, cols = length
	r := RectPiece{piece.Height(), piece.Length()}
	return mosaicPiece{
		piece,
		r,
	}
}

func StudsRightPiece(piece Brick) MosaicPiece {
	// Studs to the right on its side, so rows = length, cols = height
	r := RectPiece{piece.Length(), piece.Height()}
	return mosaicPiece{
		piece,
		r,
	}
}

func PiecesForOrientation(o ViewOrientation, pieces []Brick) []MosaicPiece {
	result := make([]MosaicPiece, len(pieces))
	switch o {
	case StudsOut:
		for i, p := range pieces {
			result[i] = StudsOutPiece(p)
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
