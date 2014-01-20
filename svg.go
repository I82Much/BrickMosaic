package BrickMosaic

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ajstarks/svgo"
)

type SVGRenderer struct {}

func GetDimensionsForBlock(o ViewOrientation) (width, height int) {
	// Change aspect ratio
	switch o {
	case StudsUp:
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

// Upper left origin
func BoundingBox(p Piece, origin Location) (minRow, minCol, maxRow, maxCol int) {
	minRow = 9999999
	minCol = 99999999
	maxRow = -99999999
	maxCol = -99999999
	for _, loc := range p.Extent() {
		translated := origin.Add(loc)
		if translated.Row < minRow {
			minRow = translated.Row
		}
		if translated.Row > maxRow {
			maxRow = translated.Row
		}
		if translated.Col < minCol {
			minCol = translated.Col
		}
		if translated.Col > maxCol {
			maxCol = translated.Col
		}
	}
	return
}


func DoRender(p Plan, canvas *svg.SVG) {
	brickWidth, brickHeight := GetDimensionsForBlock(p.Orig().Orientation())
	imageWidth := brickWidth * p.Orig().NumCols()
	imageHeight := brickHeight * p.Orig().NumRows()

	canvas.Gid("blocks")
	bricksByColor := make(map[BrickColor][]PlacedBrick)
	for _, b := range p.Pieces() {
	  bricksByColor[b.Color] = append(bricksByColor[b.Color], b)
	}
	//Draw the blocks of color
	// Draw outlines around each piece
	for color, bricks := range bricksByColor {
		canvas.Gid(fmt.Sprintf("blocks-%v", color))
    for _, piece := range bricks {
      origin := piece.Origin
			for _, loc := range piece.Extent {
				translated := origin.Add(loc)
				startX := translated.Col * brickWidth
				startY := translated.Row * brickHeight
				// In range 0, 65535; need to convert to 0-255
				r, g, b, _ := color.RGBA()
				colorStr := canvas.RGB(int(r/255), int(g/255), int(b/255))
				canvas.Rect(startX, startY, brickWidth, brickHeight, colorStr)
			}
		}
		canvas.Gend()
	}
	canvas.Gend()

/*
	canvas.Gid("block_outlines")
	// Draw outlines around each piece
	for _, solution := range m.Solutions() {
		for loc, piece := range solution.Pieces {
			minRow, minCol, maxRow, maxCol := BoundingBox(piece, loc)

			// Offset by one because we draw to where it ends. e.g. if it takes up only one
			// row or column, we still need to draw it as if it went into right before the
			// next row or column.
			startX := minCol * brickWidth
			endX := (maxCol + 1) * brickWidth

			startY := minRow * brickHeight
			endY := (maxRow + 1) * brickHeight

			width := endX - startX
			height := endY - startY

			style := "fill='none' stroke='gray'"
			canvas.Rect(startX, startY, width, height, style)
		}
	}
	canvas.Gend()*/

  // FIXME(ndunn): this only works for studs right mosaic.
	canvas.Gid("gridlines")
	majorOpacity := 0.5
	minorOpacity := 0.2
	// Draw the grid lines
	for row := 0; row < p.Orig().NumRows()+1; row++ {
		y := int(row * brickHeight)
		// Every 4th row (corresponding to length of 2x4), draw it darker.
		alpha := minorOpacity
		if row > 0 && row%4 == 0 {
			alpha = majorOpacity
		}
		style := strings.Replace(canvas.RGBA(255, 0, 0, alpha), "fill", "stroke", -1)

		// Workaround for bug - need stroke- not file
		canvas.Line(0, y, imageWidth, y, style)
	}

	// Vertical grid lines
	for col := 0; col < p.Orig().NumCols()+1; col++ {
		x := int(col * brickWidth)
		// Every 3rd column (corresponding to 3 stacked plates), draw it darker.
		alpha := minorOpacity
		if col > 0 && col%3 == 0 {
			alpha = majorOpacity
		}
		style := strings.Replace(canvas.RGBA(255, 0, 0, alpha), "fill", "stroke", -1)
		canvas.Line(x, 0, x, imageHeight, style)
	}
	canvas.Gend()

}

func (r SVGRenderer) Render(p Plan) string {
  var buf bytes.Buffer
	canvas := svg.New(&buf)
	blockWidth, blockHeight := GetDimensionsForBlock(p.Orig().Orientation())
	width := blockWidth * p.Orig().NumCols()
	height := blockHeight * p.Orig().NumRows()
	canvas.Start(width, height)
	canvas.Title("Grid")
  DoRender(p, canvas)
	canvas.End()
	return buf.String()
}