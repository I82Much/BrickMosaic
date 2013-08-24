package BrickMosaic

import (
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
  if foundPieces != allPieces {
    t.Errorf("got %v wanted %v", foundPieces, allPieces)
  }
}