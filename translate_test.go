package BrickMosaic

import (
  "reflect"
  "testing"
)

func loc(row, col int) Location {
  return Location{Row: row, Col: col}
}

func TestTranslate(t *testing.T) {
  tests := []struct {
    name string
    locs []Location
    point AnchorPoint
    want []Location
  } {
    {
      name: "empty",
      locs: nil,
      point: UpperRight,
      want: nil,
    },
    {
      name: "upper left",
      locs: []Location{loc(5,2), loc(5,3)},
      point: UpperLeft,
      want: []Location{loc(5,2), loc(5,3)},
    },
    {
      name: "upper right",
      locs: []Location{loc(5,2), loc(5,3)},
      point: UpperRight,
      want: []Location{loc(5,-2), loc(5,-3)},
    },
    {
      name: "lower right",
      locs: []Location{loc(5,2), loc(5,3)},
      point: LowerRight,
      want: []Location{loc(-5,-2), loc(-5,-3)},
    },
    {
      name: "lower left",
      locs: []Location{loc(5,2), loc(5,3)},
      point: LowerLeft,
      want: []Location{loc(-5,2), loc(-5,3)},
    },
  }
  for _, test := range tests {
    if got := Translate(test.locs, test.point); !reflect.DeepEqual(got, test.want) {
      t.Errorf("Translate(%q): got %v want %v", test.name, got, test.want)
    }
  }
}