package BrickMosaic

import (
	"image"
	"image/color"
	"testing"
)

// LimitedUniform is a finite-sized Image of uniform color. Most of implementation is copied from
// http://golang.org/src/pkg/image/names.go?s=668:706#L14
type LimitedUniform struct {
	C color.Color
	B image.Rectangle
}

func (c *LimitedUniform) RGBA() (r, g, b, a uint32) {
	return c.C.RGBA()
}

func (c *LimitedUniform) ColorModel() color.Model {
	return c
}

func (c *LimitedUniform) Convert(color.Color) color.Color {
	return c.C
}

func (c *LimitedUniform) Bounds() image.Rectangle { return c.B }

func (c *LimitedUniform) At(x, y int) color.Color { return c.C }

func NewUniform(c color.Color, b image.Rectangle) *LimitedUniform {
	return &LimitedUniform{c, b}
}

func TestColorChosenCorrectly(t *testing.T) {
	tests := []struct {
		name string
		c    color.Color
		p    color.Palette
		want color.Color
	}{
		{
			name: "Black",
			c:    Black,
			p:    []color.Color{Black, White, Red},
			want: Black,
		},
		{
			name: "White",
			c:    White,
			p:    []color.Color{Black, White, Red},
			want: White,
		},
		{
			name: "Gray",
			c:    color.Gray{128},
			p:    []color.Color{Black, White, DarkGrey, Red},
			want: DarkGrey,
		},
	}
	for _, test := range tests {
		img := NewUniform(test.c, image.Rectangle{image.Point{0, 0}, image.Point{10, 10}})
		ideal := EucPosterize(img, test.p, 1, 1, StudsOut)
		if got := ideal.Color(0, 0); got != test.want {
			t.Errorf("Posterize(%q): got %v want %v\n", test.name, got, test.want)
		}
	}
}
