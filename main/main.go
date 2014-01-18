// main runs the mosaic program.
// +build !appengine
package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	//"image/png"
	"os"
	//"strings"

	"github.com/I82Much/BrickMosaic"
_	"github.com/I82Much/BrickMosaic/grid/grid"
)

var (
	rows       = flag.Int("rows", 10, "number of rows")
	cols       = flag.Int("cols", 25, "number of columns")
	inputPath  = flag.String("path", "", "path to input file")
	outputPath = flag.String("output_path", "", "path to output svg file")
)

func main() {
	flag.Parse()
	if *inputPath == "" {
		panic("Must set --path, path to the input file")
	}
	if *outputPath == "" {
		panic("Must set --output_path, path to the output file")
	}

	path := *inputPath
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Couldn't load file path %q: %v", path, err))
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		panic(fmt.Sprintf("Couldn't decode file %v: %v", path, err))
	}

	fmt.Printf("Image format %v\n", format)

	outputFile, err := os.Create(*outputPath)
	if err != nil {
		panic("Couldn't create output file")
	}
	// close the output file on exit and check for its returned error
	defer func() {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}()

	//palette := BrickMosaic.GrayPlusPalette
	palette := BrickMosaic.FullPalette
	legoImg := BrickMosaic.NewBrickImage(img, *rows, *cols, palette)
	orientation := BrickMosaic.StudsRight

	pieces := BrickMosaic.Pieces
	//pieces := BrickMosaic.Bricks
	//pieces = append(pieces, BrickMosaic.OneByTenPlate)
	studsRightPieces := BrickMosaic.PiecesForOrientation(orientation, pieces)
	fmt.Printf("pieces %v", studsRightPieces)

	mosaic := BrickMosaic.MakeMosaic(legoImg.(*BrickMosaic.BrickImage), orientation, studsRightPieces)
	fmt.Printf("%v\n", mosaic)

	inventory := BrickMosaic.MakeInventory()
	for color, grid := range mosaic.Grids() {
		fmt.Printf("%v\n", color)
		fmt.Printf("%v\n", grid)
		solution, _ := grid.Solve(studsRightPieces)

		fmt.Printf("%d pieces", len(solution.Pieces))
		piecesNeeded := make(map[BrickMosaic.BrickPiece]int)
		for _, piece := range solution.Pieces {
			mosaicPiece := piece.(BrickMosaic.MosaicPiece)
			piecesNeeded[mosaicPiece.Brick]++
			inventory.Add(color, mosaicPiece)
		}
		fmt.Printf("%v\n", piecesNeeded)
		fmt.Printf("%v\n", solution)
	}
	fmt.Printf("%v", inventory.DescendingUsage())

	if _, err := outputFile.Write(BrickMosaic.MakeSvgInstructions(mosaic)); err != nil {
		panic(err)
	}
}
