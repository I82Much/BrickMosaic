// Package render is concerned with rendering Plans.
package render

import (
  	"github.com/I82Much/BrickMosaic"
)

// Renderers somehow convert the plan into a form that's easy for humans to build. For instance it
// might render the plan as an SVG file embedded in a webpage, or print it to standard out,
// or render it as LDRAW instructions.
type Renderer interface {
  func Render(p BrickMosaic.Plan)
}