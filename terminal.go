package BrickMosaic

import (
  "fmt"
  "io"
)
// TerminalRenderer is an implementation of the Renderer interface which emits a textual
// representation of the Plan to stdout.
type WriterRenderer struct {
    w io.Writer
}

func (t WriterRenderer) Render(p Plan) {
	for row := 0; row < p.Orig().NumRows(); row++ {
		for col := 0; col < p.Orig().NumCols(); col++ {
		  piece := p.Piece(row, col)
		  _, err := io.WriteString(t.w, fmt.Sprintf("%03d ", piece.Id))
		  if err != nil {
		    panic("couldn't write string")
		  }
		}
		_, err := io.WriteString(t.w, "\n")
		if err != nil {
		  panic("couldn't write string")
		}
	}
}
