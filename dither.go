package BrickMosaic

// Dither implements Floyd Steinberg's algorithm for dithering. See  http://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering

// TODO(ndunn): reduce duplication.

import (
	"fmt"
	"image"
	"image/color"
)

// BrickImage is an implementation of DesiredMosaic interface. It also implements the image.Image interface
// so that it can be rendered for debugging purposes.
type DitheredBrickImage struct {
	img        image.Image
	palette    color.Palette
	rows, cols int
	
	colors map[Location]color.Color
	
	// Maps each grid cell to its color
	avgColors   map[Location]BrickColor
	orientation ViewOrientation
	
	Frames []*image.Paletted
}

const (
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

// image.Image implementation follows

// FIXME ndunn base it on the orientation
func (si *DitheredBrickImage) Bounds() image.Rectangle {
  return image.Rectangle{image.Pt(0, 0), image.Pt(scaleFactor*si.cols, scaleFactor*si.rows)}
	//return image.Rectangle{image.Pt(0, 0), image.Pt(si.cols, si.rows)}
}

// At returns what color should be rendered at this x, y coordinate.
func (si *DitheredBrickImage) At(x, y int) color.Color {
  return si.Color(y, x)
  /*
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
	return si.Color(gridRow, gridCol)*/
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

// ColorIndexAt returns the palette index of the pixel at (x, y).
func (si *DitheredBrickImage) ColorIndexAt(x, y int) uint8 {
  avg := si.At(x, y)
  for i, c := range si.palette {
    if c == avg {
      return uint8(i)
    }
  }
  return uint8(0)
}

// FIXME ndunn remove
func (si *DitheredBrickImage) OriginalIdealColor(row, col int) color.Color {
  // Convert rows/columns into x/y coordinates in the image
	y1 := si.rowToY(row)
	y2 := si.rowToY(row + 1)
	
	x1 := si.colToX(col)
	x2 := si.colToX(col + 1)

	bounds := image.Rect(x1, y1, x2, y2)
	avgColor := AverageColor(si.img, bounds)
	return avgColor
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

// Paletted renders the current state of the image as a Paletted image.
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
  //return QuantizationError{}
   
	  r0, g0, b0, a0 := oldC.RGBA()
	  r1, g1, b1, a1 := newC.RGBA()

    // Avoid underflow when subtracting uints
    rdiff := int(int(r0) - int(r1))
    gdiff := int(int(g0) - int(g1))
    bdiff := int(int(b0) - int(b1))
    adiff := int(int(a0) - int(a1))
    //adiff := int8(a1 - a0)
    return QuantizationError{
      r: rdiff / 256,
      g: gdiff / 256,
      b: bdiff / 256,
      a: adiff / 256,
    }
}
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

// DitherPosterize is a posterization process that uses Euclidean distance.
func DitherPosterize(img image.Image, p color.Palette, rows int, cols int, o ViewOrientation) IdealImage {
	return NewDitheredBrickImage(img, rows, cols, p, o)
}

// NewDitheredBrickImage returns a DitheredBrickImage based on the given inputs.
func NewDitheredBrickImage(img image.Image, rows, cols int, palette color.Palette, o ViewOrientation) *DitheredBrickImage {
	brickImage := &DitheredBrickImage{
	  img        : img,
  	palette    : palette,
  	rows: rows,
  	cols: cols,
  	colors: make(map[Location]color.Color),
	  avgColors: make(map[Location]BrickColor),
	  orientation: o,
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
  brickImage.Frames = append(brickImage.Frames, brickImage.Paletted())
  

	// Initialize the color map
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
		  origIdeal := brickImage.OriginalIdealColor(row, col)
		  
		  oldPixel := brickImage.IdealColor(row, col)
		  bestMatch := brickImage.Color(row, col)
		  err := Error(oldPixel, bestMatch)

      // row 9 col 42 originalIdeal {44 46 49 255} old {112 15 37 0} best match {21 BrightRed {196 40 27 0}} err {-84 -25 10 0}
      
      /*
      if bestMatch == BrightRed {
      //if row == 0 && col < 10 {
        fmt.Printf("row %d col %d originalIdeal %v old %v best match %v err %v\n", row, col, origIdeal, oldPixel, bestMatch, err)
      }*/
      
      // Look at the 3x3 neighborhood around the pixel in question
      if row >= 7 && row < 9 && col >= 41 && col < 43 {
        fmt.Printf("row %d col %d originalIdeal %v old %v best match %v err %v\n", row, col, origIdeal, oldPixel, bestMatch, err)
      }


      if col != cols - 1 {
        //existRight := brickImage.colors[Location{row,col+1}]
        //withErr := AddError(existRight, err.Scale(7.0 / 16.0))
        
        //priorBestMatch := brickImage.palette.Convert(existRight).(BrickColor)
        //newBestMatch := brickImage.palette.Convert(withErr).(BrickColor)
        
        /*if row < 10 && col < 10 {
          fmt.Printf("right: %v with err %v prior best %v new best match %v\n", existRight, withErr, priorBestMatch, newBestMatch)
        }*/
        
        // To the right
        brickImage.colors[Location{row,col+1}] = AddError(brickImage.colors[Location{row,col+1}], err.Scale(7.0/16.0))
      }

      if col != 0 && row != rows - 1 {
        // To Left, below
        brickImage.colors[Location{row+1,col-1}] = AddError(brickImage.colors[Location{row+1,col-1}], err.Scale(3.0/16.0))
      }

      if row != rows - 1 {
        // Center, below
        brickImage.colors[Location{row+1,col}] = AddError(brickImage.colors[Location{row+1,col}], err.Scale(5.0/16.0))
      }

      if row != rows -1 && col != cols -1 {
        // To right, below
        brickImage.colors[Location{row+1,col+1}] = AddError(brickImage.colors[Location{row+1,col+1}], err.Scale(1.0/16.0))
      }
      
      brickImage.Frames = append(brickImage.Frames, brickImage.Paletted())
    }
  }
	return brickImage
}
