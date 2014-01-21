package BrickMosaic

import (
  "testing"
)

func TestCalculateRowsAndColumns(t *testing.T) {
  tests := []struct {
    name string
    width int
    height int
    maxStuds int
    orientation ViewOrientation
    
    // Wanted
    rows, cols int
  } {
    {
      name: "square - out",
      width: 1000,
      height: 1000,
      maxStuds: 50,
      orientation: StudsOut,
      rows: 50,
      cols: 50,
    },
    
    {
      name: "square - top",
      width: 1000,
      height: 1000,
      maxStuds: 50,
      orientation: StudsTop,
      rows: 125,
      cols: 50,
    },
    {
      name: "square - on side",
      width: 1000,
      height: 1000,
      maxStuds: 50,
      orientation: StudsRight,
      rows: 50,
      // Ratio of plate height to brick width is 8 / 20. 2.5 * 50 = 125
      cols: 125,
    },
    {
      name: "rectangle - wider than tall",
      width: 1000,
      height: 500,
      maxStuds: 50,
      orientation: StudsOut,
      rows: 25,
      cols: 50,
    },
    {
      name: "rectangle - taller than wide",
      width: 500,
      height: 1000,
      maxStuds: 50,
      orientation: StudsOut,
      rows: 50,
      cols: 25,
    },
    {
      name: "rectangle - taller than wide, studs up",
      width: 500,
      height: 1000,
      maxStuds: 50,
      orientation: StudsTop,
      rows: 125,
      cols: 25,
    },
    {
      name: "rectangle - taller than wide, studs right",
      width: 500,
      height: 1000,
      maxStuds: 50,
      orientation: StudsRight,
      rows: 50,
      cols: 62,
    },
  }
  for _, test := range tests {
    gotRows, gotCols := CalculateRowsAndColumns(test.width, test.height, test.maxStuds, test.orientation)
    if gotRows != test.rows {
      t.Errorf("CalculateRowsAndColumns(%q): got rows %d want %d", test.name, gotRows, test.rows)
    } 
    if gotCols != test.cols {
      t.Errorf("CalculateRowsAndColumns(%q): got cols %d want %d", test.name, gotCols, test.cols)
    }
  }
} 
