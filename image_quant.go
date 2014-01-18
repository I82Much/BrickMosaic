// Package BrickMosaic contains functions for image quantization
// and resampling large images into a format suitable for 
package BrickMosaic

import (
	"fmt"
	"image"
	"image/color"

	brickpalette "github.com/I82Much/BrickMosaic/palette"
)

type Location struct {
	row, col int
}

func (loc Location) Add(loc2 Location) Location {
	return Location{loc.row + loc2.row, loc.col + loc2.col}
}

// TODO(ndunn): top down vs side on
type BrickImage struct {
	img        image.Image
	palette    color.Palette
	rows, cols uint
	// Maps each grid cell to its color
	avgColors map[Location]brickpalette.BrickColor
}

// AverageColor determines the 'average' color of the subimage whose coordinates are contained in the
// given bounds. The average is an arithmetic average in RGB color space.
// TODO(ndunn): use a different color space
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

func (si *BrickImage) ColorModel() color.Model {
	return si.img.ColorModel()
}

func (si *BrickImage) Bounds() image.Rectangle {
	return si.img.Bounds()
}

// ColorAt returns the best palette.BrickColor for the given row/column
// in the image based on the palette this image was instantiated with.
func (si *BrickImage) ColorAt(row, col int) brickpalette.BrickColor {
	loc := Location{row, col}
	if c, ok := si.avgColors[loc]; ok {
		return c
	}
	w := si.img.Bounds().Max.X - si.img.Bounds().Min.X
	h := si.img.Bounds().Max.Y - si.img.Bounds().Min.Y
	colWidth := w / int(si.cols)
	rowHeight := h / int(si.rows)

	x1 := col * colWidth
	x2 := (col + 1) * colWidth
	y1 := row * rowHeight
	y2 := (row + 1) * rowHeight

	bounds := image.Rect(x1, y1, x2, y2)
	avgColor := AverageColor(&si.img, bounds)
	bestMatch := si.palette.Convert(avgColor).(brickpalette.BrickColor)
	si.avgColors[loc] = bestMatch
	return bestMatch
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
	if uint(gridRow) >= si.rows {
		panic(fmt.Sprintf("Too many rows; was rendering row %d; max of %d rows", gridRow, si.rows))
	}

	// Grid line 
	if x%colWidth == 0 || y%rowHeight == 0 {
		return brickpalette.Red
	}
	return si.ColorAt(gridRow, gridCol)
}

func NewBrickImage(img image.Image, rows, cols int, palette color.Palette) image.Image {
	brickImage := &BrickImage{img, palette, uint(rows), uint(cols), make(map[Location]brickpalette.BrickColor)}
	// Initialize the color map
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			_ = brickImage.ColorAt(row, col)
		}
	}

	return brickImage
}
