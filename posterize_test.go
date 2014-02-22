package BrickMosaic

import (
	"image"
	"image/color"
	"reflect"
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

func TestAddError(t *testing.T) {
  tests := []struct {
    name string
    col color.Color
    err QuantizationError
    want color.Color
  } {
    {
      name: "Adding zero does nothing",
      col: color.RGBA{R:100, G:120, B:140},
      err: QuantizationError{},
      want: color.RGBA{R:100, G:120, B:140},
    },
    {
      name: "Underflow and overflow is avoided",
      col: color.RGBA{R:100, G:120, B:140},
      err: QuantizationError{
        r: -101,
        g: 136,
        b: 20,
      },
      want: color.RGBA{R:0, G:255, B:160},
    },
  }
  for _, test := range tests {
    if got := AddError(test.col, test.err); got != test.want {
      t.Errorf("AddError(%q): got %v want %v", test.name, got, test.want)
    }
    
  }
}

func TestError(t *testing.T) {
  tests := []struct {
    c0, c1 color.Color
    want QuantizationError
  } {
    {
      c0: color.RGBA{0,255,50,0},
      c1: color.RGBA{255,0,100,0},
      want: QuantizationError {
        r: -255,
        g: 255,
        b: -50,
      },
    },
  }  
  for _, test := range tests {
    if got := Error(test.c0, test.c1); !reflect.DeepEqual(test.want, got) {
      t.Errorf("Error(%v, %v): got %v wanted %v", test.c0, test.c1, got, test.want)
    }
  }
}

func TestScale(t *testing.T) {
  tests := []struct {
    orig QuantizationError
    scaleFactor float32
    want QuantizationError
  } {
    {
      orig: QuantizationError {
        r: -22,
        g: -29,
        b: -50,
        a: 0,
      },
      scaleFactor: 1.0,
      want: QuantizationError {
        r: -22,
        g: -29,
        b: -50,
        a: 0,
      },
    },
    {
      orig: QuantizationError {
        r: -22,
        g: -29,
        b: -50,
        a: 0,
      },
      scaleFactor: 2.0,
      want: QuantizationError {
        r: -44,
        g: -58,
        b: -100,
        a: 0,
      },
    },
    {
        orig: QuantizationError {
          r: -22,
          g: -29,
          b: -50,
          a: 0,
        },
        scaleFactor: 0.5,
        want: QuantizationError {
          r: -11,
          g: -14,
          b: -25,
          a: 0,
        },
      },
    }
  for _, test := range tests {
    if got := test.orig.Scale(test.scaleFactor); !reflect.DeepEqual(got, test.want) {
      t.Errorf("%v scale by %v: got %v want %v", test.orig, test.scaleFactor, got, test.want)
    }
  }
}