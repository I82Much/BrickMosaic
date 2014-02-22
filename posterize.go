// Package image is responsible for manipulating input images into a format that is
// compatible with the brick colors we have. For instance, it converts from an image
// of potentially millions of colors into a much smaller color palette.
//
// According to Wikipedia:
//   "Posterization of an image entails conversion of a continuous gradation of tone to several regions of fewer tones, with abrupt changes from one tone to another." (http://en.wikipedia.org/wiki/Posterization, retrieved 2014/01/19)
//
// This package is responsible for converting from raw images to the Ideal.

// This implements Floyd Steinberg's algorithm for dithering. See  http://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering

package BrickMosaic

import (
	"fmt"
	"image"
	"image/color"
)

const (
  // Each pixel in image is blown up by this factor in paletted image
  scaleFactor = 5
)


// IdealImage is an object that implements both the Image interface and the Ideal interface
type IdealImage interface {
  image.Image
  Ideal
}

// Posterize is the interface for converting from images into DesiredMosaic objects.
type Posterize func(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) IdealImage

// BrickImage is an implementation of Ideal interface. 
type BrickImage struct {
	img        image.Image
	palette    color.Palette
	rows, cols int

	colors map[Location]color.Color
	
	// Maps each grid cell to its color
	avgColors   map[Location]BrickColor
	orientation ViewOrientation

	// 0.0 = no dithering at all
	// 1.0 = standard amount of dithering. Scales the quantization error that is propagated
	// through the image.
	errorScalingFactor float32

	// Frames are snapshots of the process of creating the final image, for debuggin
	// purposes
	Frames []*image.Paletted
}

// AverageColor determines the 'average' color of the subimage whose coordinates are contained in the
// given bounds. The average is an arithmetic average in RGB color space.
// TODO(ndunn): try different color spaces.
func AverageColor(si image.Image, bounds image.Rectangle) color.Color {
	R, G, B, A := uint64(0), uint64(0), uint64(0), uint64(0)
	numPixels := uint64(0)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := si.At(x, y)
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

// NumRows returns the number of rows in the piece.
func (si *BrickImage) NumRows() int {
	return si.rows
}

// NumRows returns the number of columns in the piece.
func (si *BrickImage) NumCols() int {
	return si.cols
}

// Orientation returns the way in which the image is oriented.
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



func (si *BrickImage) ColorModel() color.Model {
	return si.palette
}

// FIXME ndunn base it on the orientation
func (si *BrickImage) Bounds() image.Rectangle {
  return image.Rectangle{image.Pt(0, 0), image.Pt(scaleFactor*si.cols, scaleFactor*si.rows)}
}

// Color returns the best palette.BrickColor for the given row/column
// in the image based on the palette this image was instantiated with.
func (si *BrickImage) Color(row, col int) BrickColor {
	loc := Location{row, col}
	if c, ok := si.avgColors[loc]; ok {
		return c
	}
	avgColor := si.IdealColor(row, col)
	bestMatch := si.palette.Convert(avgColor).(BrickColor)
	si.avgColors[loc] = bestMatch
	return bestMatch
}

// IdealColor returns the color that ideally we would use for the row / col combination
// if we had pieces of every color. Later this will be quantized into the nearest
// neighbor, into a BrickColor.
func (si *BrickImage) IdealColor(row, col int) color.Color {
  loc := Location{row, col}
  if c, ok := si.colors[loc]; ok {
    return c
  }

  // Convert rows/columns into x/y coordinates in the image
	y1 := si.rowToY(row)
	y2 := si.rowToY(row + 1)
	
	x1 := si.colToX(col)
	x2 := si.colToX(col + 1)

	bounds := image.Rect(x1, y1, x2, y2)
	avgColor := AverageColor(si.img, bounds)
	return avgColor
}

// Paletted renders the current state of the image as a Paletted image. Useful for debugging
// purposes
func (si *BrickImage) Paletted() *image.Paletted {
  p := image.NewPaletted(si.Bounds(), si.palette)
  for row := 0; row < si.NumRows(); row++ {
    for col := 0; col < si.NumCols(); col++ {
      for x := 0; x < scaleFactor; x++ {
        for y := 0; y < scaleFactor; y++ {
          x1 := col * scaleFactor
          y1 := row * scaleFactor
          p.Set(x1 + x, y1 + y, si.IdealColor(row, col))
        } 
      }
    }
  }
  return p
}

// QuantizationError represents an error between a desired color and the best possible
// color that we can use to represent it.
type QuantizationError struct {
  // amount of error in r, g, b, a channels. Assumes 8 bit color.
  r, g, b, a int
}

// Error returns how much error is there from c1 relative to c0? High numbers means c1 has higher in that channel.
func Error(oldC, newC color.Color) QuantizationError {
	  r0, g0, b0, a0 := oldC.RGBA()
	  r1, g1, b1, a1 := newC.RGBA()

    // Avoid underflow when subtracting uints
    rdiff := int(int(r0) - int(r1))
    gdiff := int(int(g0) - int(g1))
    bdiff := int(int(b0) - int(b1))
    adiff := int(int(a0) - int(a1))
    return QuantizationError{
      r: rdiff / 256,
      g: gdiff / 256,
      b: bdiff / 256,
      a: adiff / 256,
    }
}

// Scale scales the given error by the given factor. For instance, Scale(2.0) doubles the
// error, while Scale(.5) halves it. This returns a new object.
func (e QuantizationError) Scale(factor float32) QuantizationError {
  return QuantizationError {
    r: int(float32(e.r) * factor),
    g: int(float32(e.g) * factor),
    b: int(float32(e.b) * factor),
    //a: int8(float32(e.a) * factor),
  }
}

// AddError adds the given amount of error to the given color, transforming it into a new color. For instance,
// if c is {R:100, G:100, B:100} and error is {r:-50,g:50,b:-50}, then the final result is
// {R:50, G:150, B:50}
func AddError(c color.Color, err QuantizationError) color.Color {
  // These are in 16 bit color; later one we convert back to 8 bit color through integer
  // division.
  r0, g0, b0, a0 := c.RGBA()
  
  // Avoid under and overflow.
  clamp := func(x int) uint8 {
    if x < 0 { 
      return uint8(0)
    } else if x > 255 {
      return uint8(255)
    }
    return uint8(x)
  }
  
  return color.RGBA {
    R: clamp(int(r0/256) + err.r),
    G: clamp(int(g0/256) + err.g),
    B: clamp(int(b0/256) + err.b),
    A: clamp(int(a0/256) + err.a),
  }
}

// DitherPosterize converts the given image into an Ideal form using a standard amount of
// dithering (error propagation).
func DitherPosterize(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) Ideal {
	return NewBrickImage(img, rows, cols, p, o, 1.0)
}

// EucPosterize returns an Ideal representation of the image with no dithering.
func EucPosterize(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) Ideal {
  return NewBrickImage(img, rows, cols, p, o, 0.0)
}

// NewBrickImage returns a BrickImage based on the given inputs.
func NewBrickImage(img image.Image, rows, cols int, palette color.Palette, o ViewOrientation, errorScalingFactor float32) *BrickImage {
	brickImage := &BrickImage{
	  img        : img,
  	palette    : palette,
  	rows: rows,
  	cols: cols,
  	colors: make(map[Location]color.Color),
	  avgColors: make(map[Location]BrickColor),
	  orientation: o,
	  errorScalingFactor: 1.0,
	  Frames: nil,
	}
  
// Initialize the color map
 	for row := 0; row < rows; row++ {
 		for col := 0; col < cols; col++ {
      oldPixel := brickImage.IdealColor(row, col)
      brickImage.colors[Location{row, col}] = oldPixel
      
      if row == 0 && col < 10 {
        fmt.Printf("row %d col %d color: %v\n", row, col, oldPixel)
      }
    } 
	}
  //brickImage.Frames = append(brickImage.Frames, brickImage.Paletted())
  

	// Initialize the color map
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
		  oldPixel := brickImage.IdealColor(row, col)
		  bestMatch := brickImage.Color(row, col)
		  err := Error(oldPixel, bestMatch)

      if col != cols - 1 {
        // To the right
        brickImage.colors[Location{row,col+1}] = AddError(brickImage.colors[Location{row,col+1}], err.Scale(errorScalingFactor*7.0/16.0))
      }

      if col != 0 && row != rows - 1 {
        // To Left, below
        brickImage.colors[Location{row+1,col-1}] = AddError(brickImage.colors[Location{row+1,col-1}], err.Scale(errorScalingFactor*3.0/16.0))
      }

      if row != rows - 1 {
        // Center, below
        brickImage.colors[Location{row+1,col}] = AddError(brickImage.colors[Location{row+1,col}], err.Scale(errorScalingFactor*5.0/16.0))
      }

      if row != rows -1 && col != cols -1 {
        // To right, below
        brickImage.colors[Location{row+1,col+1}] = AddError(brickImage.colors[Location{row+1,col+1}], err.Scale(errorScalingFactor*1.0/16.0))
      }
      //brickImage.Frames = append(brickImage.Frames, brickImage.Paletted())
    }
  }
  // Final version
  brickImage.Frames = append(brickImage.Frames, brickImage.Paletted())
	return brickImage
}
