package BrickMosaic

import (
	"fmt"
	"testing"
)

func TestSVGRender(t *testing.T) {
	p := &fakePlan{}
	svg := SVGRenderer{}
	res := svg.Render(p)
	fmt.Println(res)
}
