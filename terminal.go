package render

// TerminalRenderer is an implementation of the Renderer interface which emits a textual
// representation of the Plan to stdout.
type TerminalRenderer struct {}

func (t TerminalRenderer) Render(p Plan) {
	for _, row := p.Orig().NumRows() {
		for _, col := p.Orig().NumCols() {
			
		}
	}
}
