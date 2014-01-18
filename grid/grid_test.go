package grid

import (
	"reflect"
	"testing"
)

func TestAddDoesNotModify(t *testing.T) {
  a := Location{1, 5}
  b := Location{5, 6}
  a.Add(b)
  if a.Row != 1 || a.Col != 5 {
    t.Errorf("a was modified by addition. Expected {1, 5}, got %v", a)
  }
}

func TestAdd(t *testing.T) {
  tests := []struct {
    loc1, loc2 Location
    want Location
  } {
    {
      loc1: Location{0, 0}, 
      loc2: Location{1, 5},
      want: Location{1, 5},
    },
    {
      loc1: Location{0, 0}, 
      loc2: Location{0, 0},
      want: Location{0, 0},
    },
    {
      loc1: Location{1, 2}, 
      loc2: Location{4, 5},
      want: Location{5, 7},
    },
  }
  for _, test := range tests {
    if got := test.loc1.Add(test.loc2); !reflect.DeepEqual(got, test.want) {
      t.Errorf("%v + %v; wanted %v got %v", test.loc1, test.loc2, test.want, got)
    }
    // Addition should be commutative
    if got := test.loc2.Add(test.loc1); !reflect.DeepEqual(got, test.want) {
      t.Errorf("%v + %v; wanted %v got %v", test.loc2, test.loc1, test.want, got)
    }
  }

}

func TestGet(t *testing.T) {
	type getTest struct {
		name     string
		g        Grid
		row, col int
		want     State
	}
	various := New(5, 6)
	various.state[4][3] = ToBeFilled
	various.state[3][2] = Filled
	for _, test := range []getTest{
		{"empty", New(0, 0), 0, 0, Empty},
		{"negative row", New(0, 0), -1, 0, Empty},
		{"negative col", New(0, 0), 0, -1, Empty},
		{"out of bounds", New(5, 5), -1, 45, Empty},
		{"to be filled", various, 4, 3, ToBeFilled},
		{"filled", various, 3, 2, Filled},
	} {
		if got := test.g.Get(test.row, test.col); got != test.want {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got)
		}
	}
}

func TestPieceFits(t *testing.T) {
	type fitTest struct {
		name     string
		g        Grid
		p        Piece
		row, col int
		want     bool
	}
	oneByOne := RectPiece{1, 1}
	twoByTwo := RectPiece{2, 2}
	twoByFour := RectPiece{2, 4}

	various := New(5, 6)
	various.state[4][3] = ToBeFilled
	various.state[3][2] = Filled
	for _, test := range []fitTest{
		{"empty - no space", WithState(5, 5, Empty), oneByOne, 0, 0, false},
		{"filled case", WithState(5, 5, Filled), oneByOne, 0, 0, false},
		{"to be filled - with space (0, 0)", WithState(5, 5, ToBeFilled), oneByOne, 0, 0, true},
		{"to be filled - with space (4, 4)", WithState(5, 5, ToBeFilled), oneByOne, 4, 4, true},
		{"to be filled - with space, out of bounds (5, 4)", WithState(5, 5, ToBeFilled), oneByOne, 5, 4, false},
		{"2x2 doesn't fit in 1x1 spot", WithState(1, 1, ToBeFilled), twoByTwo, 0, 0, false},
		{"2x2 does fit in 2x2 spot", WithState(2, 2, ToBeFilled), twoByTwo, 0, 0, true},
		{"2x2 does fit in 2x2 spot, but not when offset", WithState(2, 2, ToBeFilled), twoByTwo, -1, 0, false},
		{"2x4 fits in horizontal grid", WithState(2, 4, ToBeFilled), twoByFour, 0, 0, true},
		{"2x4 does not fit in vertical grid", WithState(4, 2, ToBeFilled), twoByFour, 0, 0, false},
	} {
		if got := test.g.PieceFits(test.p, Location{test.row, test.col}); got != test.want {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got)
		}
	}
}

func TestSolve(t *testing.T) {
	type solveTest struct {
		name   string
		g      Grid
		p      []Piece
		want   map[Location]Piece
		hasErr bool
	}
	oneByOne := RectPiece{1, 1}
	oneByFour := RectPiece{1, 4}
	twoByTwo := RectPiece{2, 2}
	twoByFour := RectPiece{2, 4}
	fourByOne := RectPiece{4, 1}
	for _, test := range []solveTest{
		{
			"cannot be solved - no pieces",
			WithState(1, 1, ToBeFilled),
			[]Piece{},
			make(map[Location]Piece),
			true,
		},
		{
			"trivially solved - one piece",
			WithState(1, 1, ToBeFilled),
			[]Piece{oneByOne},
			map[Location]Piece{
				Location{0, 0}: oneByOne,
			},
			false,
		},
		{
			"trivially solved - one piece, 2x2",
			WithState(2, 2, ToBeFilled),
			[]Piece{twoByFour, twoByTwo, oneByOne},
			map[Location]Piece{
				Location{0, 0}: twoByTwo,
			},
			false,
		},
		{
			"5 x 5 grid - 2 2x4",
			WithState(5, 5, ToBeFilled),
			[]Piece{twoByFour, twoByTwo, oneByFour, fourByOne, oneByOne},
			map[Location]Piece{
				Location{0, 0}: twoByFour,
				Location{2, 0}: twoByFour,
				Location{0, 4}: fourByOne,
				Location{4, 0}: oneByFour,
				Location{4, 4}: oneByOne,
			},
			false,
		},
	} {
		got, err := test.g.Solve(test.p)
		if err != nil && !test.hasErr {
			t.Errorf("for %q wanted no error got %v", test.name, err)
		} else if err == nil && test.hasErr {
			t.Errorf("for %q should have had an error", test.name)
		}
		if !reflect.DeepEqual(got.Pieces, test.want) {
			t.Errorf("for %q wanted %v got %v", test.name, test.want, got.Pieces)
		}
	}
}

func TestClone(t *testing.T) {
  various := New(5, 6)
	various.state[4][3] = ToBeFilled
	various.state[3][2] = Filled
	
	cloned := various.Clone()
	if &cloned == &various {
	  t.Errorf("cloned value shares same pointer address")
	}
	
	if !reflect.DeepEqual(various, cloned) {
	  t.Errorf("cloned value was not correctly cloned")
	}
	various.state[0][0] = Filled
	if cloned.state[0][0] != Empty {
	  t.Errorf("somehow changing original value modifies cloned state")
	}
	cloned.state[4][3] = Empty
	if various.state[4][3] != ToBeFilled {
	  t.Errorf("somehow changing cloned value modifies cloned state")
	}
  
}