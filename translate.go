package BrickMosaic

// This package is responsible for translating the Extent ([]Location) of pieces relative to
// different anchor points. E.g. by default the extent is relative to 'upper left' corner
// of brick. But if we're placing it such that lower right corner is the origin, we need
// to translate the upper left locations to match.

type AnchorPoint int

const (
	UpperLeft AnchorPoint = iota
	UpperRight
	LowerRight
	LowerLeft
)

func (a AnchorPoint) String() string {
	switch a {
	case UpperLeft:
		return "UpperLeft"
	case UpperRight:
		return "UpperRight"
	case LowerRight:
		return "LowerRight"
	case LowerLeft:
		return "LowerLeft"
	}
	panic("shouldn't reach here")
}

func Translate(locs []Location, pt AnchorPoint) []Location {
	if pt == UpperLeft {
		return locs
	}
	// All of the x (col) values need to become negative
	if pt == UpperRight {
		var points []Location
		for _, p := range locs {
			points = append(points, Location{Row: p.Row, Col: -p.Col})
		}
		return points
	}
	//All of the x values need to become negative, and all of the y values as well
	if pt == LowerRight {
		var points []Location
		for _, p := range locs {
			points = append(points, Location{Row: -p.Row, Col: -p.Col})
		}
		return points
	}
	// All of the y (row) values need to become negative
	if pt == LowerLeft {
		var points []Location
		for _, p := range locs {
			points = append(points, Location{Row: -p.Row, Col: p.Col})
		}
		return points
	}
	panic("Shouldn't reach here")
}

func TranslateAbsoluteOrigin(absLoc Location, p MosaicPiece, pt AnchorPoint) Location {
	if pt == UpperLeft {
		return absLoc
	} else if pt == UpperRight {
		// Need to translate the point LEFT by the width of the brick
		return absLoc.Add(Location{Col: -p.Cols()+1})
	} else if pt == LowerRight {
		// Tranlsate LEFT and UP
		return absLoc.Add(Location{Row: -p.Rows()+1, Col: -p.Cols()+1})
	} else if pt == LowerLeft {
		// Translate UP
		return absLoc.Add(Location{Row: -p.Rows()+1})
	}
	panic("Shouldn't reach here")
}
