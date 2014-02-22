package BrickMosaic

import (
  "reflect"
  "testing"
)

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
  }
  for _, test := range tests {
    if got := Translate(test.locs, test.point); !reflect.DeepEqual(got, test.want) {
      t.Errorf("Translate(%q): got %v want %v", test.name, got, test.want)
    }
  }
}