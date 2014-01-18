package BrickMosaic

import (
	"sort"

	//"github.com/I82Much/BrickMosaic/palette"
)

type Inventory struct {
	pieces map[BrickColor][]MosaicPiece
}

type Usage struct {
	NumPieces int
	// Area in terms of units used in the mosaic.
	Area int
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
	return c[i].usage.Area < c[j].usage.Area
}

func (c ColorUsages) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (inventory Inventory) mosaicPiecesForColor(c BrickColor) []MosaicPiece {
	return inventory.pieces[c]
}

func (inventory Inventory) PiecesForColor(c BrickColor) []BrickPiece {
	mosaics := inventory.mosaicPiecesForColor(c)
	coloredPieces := make([]BrickPiece, 0)
	for _, p := range mosaics {
		coloredPieces = append(coloredPieces, p.Brick)
	}
	return coloredPieces
}

func (inventory Inventory) UsageForColorMap() map[BrickColor]Usage {
	usageMap := make(map[BrickColor]Usage)
	for color := range inventory.pieces {
		coloredPieces := inventory.mosaicPiecesForColor(color)
		area := 0
		for _, piece := range coloredPieces {
			area += len(piece.locs)
		}
		// At this point we've calculated the area and number of bricks.
		usageMap[color] = Usage{
			NumPieces: len(coloredPieces),
			Area:      area,
		}
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
	return Inventory{make(map[BrickColor][]MosaicPiece)}
}

func (inventory *Inventory) Add(c BrickColor, p MosaicPiece) {
	inventory.pieces[c] = append(inventory.pieces[c], p)
}
