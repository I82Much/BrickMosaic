//
// Inventory is a convenience helper for determining how many pieces of which color are used in the model.
// This makes it easier to actually procure the pieces necessary to build the physical version of the model.
package BrickMosaic

import (
	"sort"
)

type Inventory struct {
	pieces map[BrickColor][]BrickPiece
}

type Usage struct {
	NumPieces int
}

type ColorUsage struct {
	color BrickColor
	usage Usage
}

type ColorUsages []ColorUsage

func (c ColorUsages) Len() int {
	return len(c)
}

func (c ColorUsages) Less(i, j int) bool {
	return c[i].usage.NumPieces < c[j].usage.NumPieces
}

func (c ColorUsages) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (inv Inventory) PiecesForColor(c BrickColor) []BrickPiece {
	return inv.pieces[c]
}

func (inventory Inventory) UsageForColorMap() map[BrickColor]Usage {
	usageMap := make(map[BrickColor]Usage)
	for color, pieces := range inventory.pieces {
	  usageMap[color] = Usage{len(pieces)}
	}
	return usageMap
}

func (inventory Inventory) DescendingUsage() []ColorUsage {
	usages := make([]ColorUsage, 0)
	for color, usage := range inventory.UsageForColorMap() {
		usages = append(usages, ColorUsage{color, usage})
	}
	sort.Sort(ColorUsages(usages))
	return usages
}

func MakeInventory() Inventory {
	return Inventory{make(map[BrickColor][]BrickPiece)}
}

func (inventory *Inventory) Add(c BrickColor, p BrickPiece) {
	inventory.pieces[c] = append(inventory.pieces[c], p)
}
