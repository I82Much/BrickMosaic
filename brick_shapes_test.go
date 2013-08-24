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
  if len(Pieces) != len(allPieces) {
    t.Errorf("Pieces: got length %d expected %d", len(Pieces), len(allPieces))
  }
}