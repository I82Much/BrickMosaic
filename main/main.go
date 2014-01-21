// main runs the mosaic program.
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
)

var (
	rows       = flag.Int("rows", 10, "number of rows")
	cols       = flag.Int("cols", 25, "number of columns")
	orientation = flag.String("orientation", "STUDS_UP", "how the grid should be oriented. Either STUDS_RIGHT, STUDS_UP, or STUDS_TOP")
	inputPath  = flag.String("path", "", "path to input file")
	outputPath = flag.String("output_path", "", "path to output svg file")

	orientationMap = map[string]BrickMosaic.ViewOrientation {
	  "STUDS_RIGHT": BrickMosaic.StudsRight,
	  "STUDS_UP": BrickMosaic.StudsUp,
	  "STUDS_TOP": BrickMosaic.StudsTop,
	}
)


func main() {
  // Flag handling; fail fast if anything is amiss
	flag.Parse()
	if _, ok := orientationMap[*orientation]; !ok {
	  panic("Must set --orientation to one of STUDS_RIGHT, STUDS_UP, or STUDS_TOP")
	}
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

  // TODO(ndunn): Expose the Palette as a command line option
	//palette := BrickMosaic.GrayPlusPalette
	palette := BrickMosaic.FullPalette
	viewOrientation := orientationMap[*orientation]
	
	// What is the ideal representation of the mosaic? Handles downsampling from many colors to few.
	ideal := BrickMosaic.EucPosterize(img, palette, *rows, *cols, viewOrientation)
  // How are we going to build this mosaic?
  plan := BrickMosaic.CreateGridMosaic(ideal)
	inventory := plan.Inventory()
	fmt.Printf("%v", inventory.DescendingUsage())
	
	renderer := BrickMosaic.SVGRenderer{}
	if _, err := outputFile.Write([]byte(renderer.Render(plan))); err != nil {
		panic(err)
	}
}
