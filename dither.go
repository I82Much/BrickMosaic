package BrickMosaic

// Package image is responsible for manipulating input images into a format that is
// compatible with the brick colors we have. For instance, it converts from an image
// of potentially millions of colors into a much smaller color palette.
//
// According to Wikipedia:
//   "Posterization of an image entails conversion of a continuous gradation of tone to several regions of fewer tones, with abrupt changes from one tone to another." (http://en.wikipedia.org/wiki/Posterization, retrieved 2014/01/19)
//
// This package is responsible for converting from raw images to the Ideal.


// This library does 
// Dither implements Floyd Steinberg's algorithm for dithering. See  http://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering

// TODO(ndunn): reduce duplication.

import (
	"fmt"
	"image"
	"image/color"
)

// DitheredBrickImage is an implementation of DesiredMosaic interface.
type DitheredBrickImage struct {
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
	
	Frames []*image.Paletted
}

const (
  // Each pixel in image is blown up by this factor in paletted image
  scaleFactor = 5
)

func (si *DitheredBrickImage) NumRows() int {
	return si.rows
}

func (si *DitheredBrickImage) NumCols() int {
	return si.cols
}

func (si *DitheredBrickImage) Orientation() ViewOrientation {
	return si.orientation
}

func (si *DitheredBrickImage) ColorModel() color.Model {
	return si.palette
}

// FIXME ndunn base it on the orientation
func (si *DitheredBrickImage) Bounds() image.Rectangle {
  return image.Rectangle{image.Pt(0, 0), image.Pt(scaleFactor*si.cols, scaleFactor*si.rows)}
	//return image.Rectangle{image.Pt(0, 0), image.Pt(si.cols, si.rows)}
}

func (si *DitheredBrickImage) rowToY(row int) int {
  return int(doMap(float64(row), 0.0, float64(si.rows), float64(si.img.Bounds().Min.Y), float64(si.img.Bounds().Max.Y)))
}

func (si *DitheredBrickImage) colToX(col int) int {
  return int(doMap(float64(col), 0.0, float64(si.cols), float64(si.img.Bounds().Min.X), float64(si.img.Bounds().Max.X)))
}

// Color returns the best palette.BrickColor for the given row/column
// in the image based on the palette this image was instantiated with.
func (si *DitheredBrickImage) Color(row, col int) BrickColor {
	loc := Location{row, col}
	if c, ok := si.avgColors[loc]; ok {
		return c
	}
	avgColor := si.IdealColor(row, col)
	bestMatch := si.palette.Convert(avgColor).(BrickColor)
	si.avgColors[loc] = bestMatch
	return bestMatch
}

func (si *DitheredBrickImage) IdealColor(row, col int) color.Color {
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
func (si *DitheredBrickImage) Paletted() *image.Paletted {
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
	return NewDitheredBrickImage(img, rows, cols, p, o, 1.0)
}

// Posterize returns an Ideal representation of the image with no dithering.
func Posterize(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) Ideal {
  return NewDitheredBrickImage(img, rows, cols, p, o, 0.0)
}

// NewDitheredBrickImage returns a DitheredBrickImage based on the given inputs.
func NewDitheredBrickImage(img image.Image, rows, cols int, palette color.Palette, o ViewOrientation, errorScalingFactor float32) *DitheredBrickImage {
	brickImage := &DitheredBrickImage{
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
