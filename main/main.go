// main runs the mosaic program.
package main

import (
	"flag"
	"fmt"

	"image"
	"image/color"
	// Support reading both jpeg and png
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/I82Much/BrickMosaic"
)

var (
	maxSizeStuds = flag.Int("studs", 40, "number of studs to have on maximum length side. The number of rows and columns will be automatically calculated")
	rows         = flag.Int("rows", -1, "number of rows. If set, will be used in preference to --studs")
	cols         = flag.Int("cols", -1, "number of columns. If set, will be used in preference to --studs")
	orientation  = flag.String("orientation", "STUDS_RIGHT", "how the grid should be oriented. Either STUDS_RIGHT, STUDS_OUT, or STUDS_TOP")
	inputPath    = flag.String("path", "", "path to input file")
	outputPath   = flag.String("output_path", "", "path to output svg file")
	palette      = flag.String("palette", "full", "comma separated list of color names, or predefined color palette name")

	orientationMap = map[string]BrickMosaic.ViewOrientation{
		"STUDS_RIGHT": BrickMosaic.StudsRight,
		"STUDS_OUT":   BrickMosaic.StudsOut,
		"STUDS_TOP":   BrickMosaic.StudsTop,
	}

	paletteMap = map[string]color.Palette{
		"gray":      BrickMosaic.GrayScalePalette,
		"gray_plus": BrickMosaic.GrayPlusPalette,
		"basic":     BrickMosaic.LimitedPalette,
		"full":      BrickMosaic.FullPalette,
		"primary":   BrickMosaic.Primary,
	}
)

func paletteFromArg(p string) color.Palette {
	if palette, ok := paletteMap[p]; ok {
		return palette
	}
	// Treat this as comma separated list
	colorStrings := strings.Split(p, ",")
	var colors []color.Color
	for _, c := range colorStrings {
		if color := BrickMosaic.ColorForName(c); color != nil {
			colors = append(colors, *color)
		} else {
			fmt.Printf("Could not find color matching %v", c)
		}
	}
	return color.Palette(colors)
}

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
	palette := paletteFromArg(*palette)
	if len(palette) == 0 {
		panic("no color palette set")
	}
	viewOrientation := orientationMap[*orientation]

	var numRows, numCols int
	// Use the command line arguments directly
	if *rows > 0 && *cols > 0 {
		numRows = *rows
		numCols = *cols
	} else if *maxSizeStuds > 0 {
		imgWidth := img.Bounds().Size().X
		imgHeight := img.Bounds().Size().Y
		numRows, numCols = BrickMosaic.CalculateRowsAndColumns(imgWidth, imgHeight, *maxSizeStuds, viewOrientation)
	} else {
		panic("must set (--rows and --colors) or --studs")
	}

	// What is the ideal representation of the mosaic? Handles downsampling from many colors to few.
	ideal := BrickMosaic.EucPosterize(img, palette, numRows, numCols, viewOrientation)
	// How are we going to build this mosaic?
	plan := BrickMosaic.CreateGridMosaic(ideal)
	inventory := plan.Inventory()
	fmt.Printf("%v", inventory.DescendingUsage())

	renderer := BrickMosaic.SVGRenderer{}
	if _, err := outputFile.Write([]byte(renderer.Render(plan))); err != nil {
		panic(err)
	}
}
