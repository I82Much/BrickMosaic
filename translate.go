package BrickMosaic

// This package is responsible for translating the Extent ([]Location) of pieces relative to
// different anchor points. E.g. by default the extent is relative to 'upper left' corner
// of brick. But if we're placing it such that lower right corner is the origin, we need
// to translate the upper left locations to match.

type AnchorPoint int
const (
  UpperLeft AnchorPoint = iota
  UpperRight AnchorPoint = iota
  LowerRight AnchorPoint = iota
  LowerLeft AnchorPoint = iota
)

func Translate(locs []Location, pt AnchorPoint) []Location {
  if pt == UpperLeft {
    return locs
  }
  // All of the x (col) values need to become negative
  if pt == UpperRight {
    var points []Location
    for _, p := range locs {
      points = append(points, Location{Row:p.Row, Col:-p.Col})
    }
    return points
  }
  //All of the x values need to become negative, and all of the y values as well 
  if pt == LowerRight {
    var points []Location
    for _, p := range locs {
      points = append(points, Location{Row:-p.Row, Col:-p.Col})
    }
    return points
  }
  // All of the y (row) values need to become negative
  if pt == LowerLeft {
    var points []Location
    for _, p := range locs {
      points = append(points, Location{Row:-p.Row, Col:p.Col})
    }
    return points
  }
  panic("Shouldn't reach here")
}