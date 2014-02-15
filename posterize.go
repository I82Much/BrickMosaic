// Package image is responsible for manipulating input images into a format that is
// compatible with the brick colors we have. For instance, it converts from an image
// of potentially millions of colors into a much smaller color palette.
//
// According to Wikipedia:
//   "Posterization of an image entails conversion of a continuous gradation of tone to several regions of fewer tones, with abrupt changes from one tone to another." (http://en.wikipedia.org/wiki/Posterization, retrieved 2014/01/19)
//
// This package is responsible for converting from raw images to the Ideal.
package BrickMosaic

import (
	"fmt"
	"image"
	"image/color"
)

// Posterize is the interface for converting from images into DesiredMosaic objects.
type Posterize func(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) Ideal

// EucPosterize is a posterization process that uses Euclidean distance.
func EucPosterize(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) Ideal {
	return NewBrickImage(img, rows, cols, p, o)
}

// BrickImage is an implementation of DesiredMosaic interface. It also implements the image.Image interface
// so that it can be rendered for debugging purposes.
type BrickImage struct {
	img        image.Image
	palette    color.Palette
	rows, cols int
	// Maps each grid cell to its color
	avgColors   map[Location]BrickColor
	orientation ViewOrientation
}

// AverageColor determines the 'average' color of the subimage whose coordinates are contained in the
// given bounds. The average is an arithmetic average in RGB color space.
// TODO(ndunn): try different color spaces.
func AverageColor(si *image.Image, bounds image.Rectangle) color.Color {
	R, G, B, A := uint64(0), uint64(0), uint64(0), uint64(0)
	numPixels := uint64(0)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := (*si).At(x, y)
			r, g, b, a := c.RGBA()
			R += uint64(r)
			G += uint64(g)
			B += uint64(b)
			A += uint64(a)
			numPixels++
		}
	}
	// Convert from range 0 ... 65535 -> 0 ... 255
	R /= (numPixels * 256)
	G /= (numPixels * 256)
	B /= (numPixels * 256)
	A /= (numPixels * 256)
	return color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
}

func (si *BrickImage) NumRows() int {
	return si.rows
}

func (si *BrickImage) NumCols() int {
	return si.cols
}

func (si *BrickImage) Orientation() ViewOrientation {
	return si.orientation
}

// doMap converts a value from one range [low1, high1] into another [low2, high2].
func doMap(v, low1, high1, low2, high2 float64) float64 {
  diff := v - low1
  proportion := diff / (high1 - low1)
  return lerp(low2, high2, proportion)
}

// lerp performs linear interpolation between v1 and v2. If amt == 0, v1 is used. If amt == 1.0,
// v2 is used. At 0.5, the average of v1, v2 is used, and so on and so forth.
func lerp(v1, v2, amt float64) float64 {
  return ((v2 - v1) * amt) + v1
}

func (si *BrickImage) rowToY(row int) int {
  return int(doMap(float64(row), 0.0, float64(si.rows), float64(si.img.Bounds().Min.Y), float64(si.img.Bounds().Max.Y)))
}

func (si *BrickImage) colToX(col int) int {
  return int(doMap(float64(col), 0.0, float64(si.cols), float64(si.img.Bounds().Min.X), float64(si.img.Bounds().Max.X)))
}

// Color returns the best palette.BrickColor for the given row/column
// in the image based on the palette this image was instantiated with.
func (si *BrickImage) Color(row, col int) BrickColor {
	loc := Location{row, col}
	if c, ok := si.avgColors[loc]; ok {
		return c
	}
	
	// Map from row/col coordinates to image coordinates
	y1 := si.rowToY(row)
	y2 := si.rowToY(row + 1)
	
	x1 := si.colToX(col)
	x2 := si.colToX(col + 1)
	
	bounds := image.Rect(x1, y1, x2, y2)
	avgColor := AverageColor(&si.img, bounds)
	bestMatch := si.palette.Convert(avgColor).(BrickColor)
	si.avgColors[loc] = bestMatch
	return bestMatch
}

func (si *BrickImage) ColorModel() color.Model {
	return si.img.ColorModel()
}

// image.Image implementation follows

func (si *BrickImage) Bounds() image.Rectangle {
	return si.img.Bounds()
}

// At returns what color should be rendered at this x, y coordinate.
func (si *BrickImage) At(x, y int) color.Color {
	w := si.img.Bounds().Max.X - si.img.Bounds().Min.X
	h := si.img.Bounds().Max.Y - si.img.Bounds().Min.Y
	colWidth := w / int(si.cols)
	rowHeight := h / int(si.rows)
	gridCol := x / colWidth
	gridRow := y / rowHeight

	// FIXME(ndunn): this only works with perfectly oriented pictures.
	if gridRow >= si.rows {
		panic(fmt.Sprintf("Too many rows; was rendering row %d; max of %d rows", gridRow, si.rows))
	}

	// Grid line
	if x%colWidth == 0 || y%rowHeight == 0 {
		return Red
	}
	return si.Color(gridRow, gridCol)
}

func NewBrickImage(img image.Image, rows, cols int, palette color.Palette, o ViewOrientation) *BrickImage {
	brickImage := &BrickImage{img, palette, rows, cols, make(map[Location]BrickColor), o}
	// Initialize the color map
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			_ = brickImage.Color(row, col)
		}
	}
	return brickImage
}
