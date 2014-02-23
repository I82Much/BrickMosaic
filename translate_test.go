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
		name  string
		locs  []Location
		point AnchorPoint
		want  []Location
	}{
		{
			name:  "empty",
			locs:  nil,
			point: UpperRight,
			want:  nil,
		},
		{
			name:  "upper left",
			locs:  []Location{loc(5, 2), loc(5, 3)},
			point: UpperLeft,
			want:  []Location{loc(5, 2), loc(5, 3)},
		},
		{
			name:  "upper right",
			locs:  []Location{loc(5, 2), loc(5, 3)},
			point: UpperRight,
			want:  []Location{loc(5, -2), loc(5, -3)},
		},
		{
			name:  "lower right",
			locs:  []Location{loc(5, 2), loc(5, 3)},
			point: LowerRight,
			want:  []Location{loc(-5, -2), loc(-5, -3)},
		},
		{
			name:  "lower left",
			locs:  []Location{loc(5, 2), loc(5, 3)},
			point: LowerLeft,
			want:  []Location{loc(-5, 2), loc(-5, 3)},
		},
	}
	for _, test := range tests {
		if got := Translate(test.locs, test.point); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Translate(%q): got %v want %v", test.name, got, test.want)
		}
	}
}

func TestTranslateAbsoluteOrigin(t *testing.T) {
	tests := []struct {
		name   string
		p      MosaicPiece
		absLoc Location
		pt     AnchorPoint
		want   Location
	}{
		{
			name:   "lower right",
			p:      StudsTopPiece(TwoByFour),
			absLoc: loc(5, 6),
			pt:     LowerRight,
			// Shift up by 3 plates, left by 4
			want: loc(3, 3),
		},
		{
			name:   "lower left",
			p:      StudsTopPiece(TwoByFour),
			absLoc: loc(5, 6),
			pt:     LowerLeft,
			// Shift up by 3 plates
			want: loc(3, 6),
		},
		{
			name:   "upper right",
			p:      StudsTopPiece(TwoByFour),
			absLoc: loc(5, 6),
			pt:     UpperRight,
			want: loc(5, 3),
		},
		{
			name:   "upper left",
			p:      StudsTopPiece(TwoByFour),
			absLoc: loc(5, 6),
			pt:     UpperLeft,
			want:   loc(5, 6),
		},
	}
	for _, test := range tests {
		if got := TranslateAbsoluteOrigin(test.absLoc, test.p, test.pt); got != test.want {
			t.Errorf("TranslateAbsoluteOrigin(%q): got %v want %v", test.name, got, test.want)
		}
	}
}
