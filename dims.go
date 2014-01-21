package BrickMosaic

import (
  "fmt"
)

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

func GetDimensionsForBlock(o ViewOrientation) (width, height int) {
	// Change aspect ratio
	switch o {
	case StudsOut:
		height = int(BrickWidth)
		width = int(BrickWidth)
	case StudsTop:
		height = int(PlateHeight)
		width = int(BrickWidth)
	case StudsRight:
		height = int(BrickWidth)
		width = int(PlateHeight)
	}
	return
}

func CalculateRowsAndColumns(width, height, maxStuds int, orientation ViewOrientation) (rows, cols int) {
  if width <= 0 {
    panic(fmt.Sprintf("width must be > 0; was %d", width))
  }
  if height <= 0 {
    panic(fmt.Sprintf("height must be > 0; was %d", height))
  }
  // Do the conversion in LDU rather than studs for more precision
  var heightDim LDU
  var widthDim LDU
  // width / height ratio = aspect ratio.
  aspectRatio := float64(width) / float64(height)
  // Wider than tall
  if aspectRatio > 1.0 {
    widthDim = BrickWidth * LDU(maxStuds)
    heightDim = LDU((float64)(BrickWidth * LDU(maxStuds)) / aspectRatio)
  } else {
    // Taller than wide, or equally tall
    heightDim = BrickWidth * LDU(maxStuds)
    widthDim = LDU((float64)(BrickWidth * LDU(maxStuds)) * aspectRatio)
  }

  // How wide and tall is the base brick in the requested orientation?
  brickWidth, brickHeight := GetDimensionsForBlock(orientation)
  rows = int(heightDim / LDU(brickHeight))
  cols = int(widthDim / LDU(brickWidth))
  
  fmt.Printf("width %d height %d max studs %d orientation %v height dim %d width dim %d rows %d cols %d\n",
  width, height, maxStuds, orientation, heightDim, widthDim, rows, cols)
  return rows, cols
} 
