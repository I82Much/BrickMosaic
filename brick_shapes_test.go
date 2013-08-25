package BrickMosaic

import (
  "reflect"
  "testing"
)

// Make sure all bricks are accounted for
func TestAllBricks(t *testing.T) {
  allPieces := make(map[BrickPiece]bool)
  for _, b := range Bricks {
    allPieces[b] = true
  }
  for _, p := range Plates {
    allPieces[p] = true
  }
  foundPieces := make(map[BrickPiece]bool)
  for _, b := range Pieces {
    foundPieces[b] = true
  }
  if len(foundPieces) != len(allPieces) {
    t.Errorf("got %d pieces wanted %d pieces", len(foundPieces), len(allPieces))
  }
  if !reflect.DeepEqual(foundPieces, allPieces) {
    t.Errorf("got %v wanted %v", foundPieces, allPieces)
  }
}