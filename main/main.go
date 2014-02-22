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
	//	"path/filepath"
	//	"image/gif"
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
	dither       = flag.Bool("dither", true, "If true, use dithering when converting the imagine into a mosaic")

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
		"bw":        BrickMosaic.BlackAndWhite,
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
		panic("Must set --orientation to one of STUDS_RIGHT, STUDS_OUT, or STUDS_TOP")
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
		panic("must set (--rows and --cols) or --studs")
	}

	// What is the ideal representation of the mosaic? Handles downsampling from many colors to few.
	var ideal BrickMosaic.Ideal
	if *dither {
		ideal = BrickMosaic.DitherPosterize(img, palette, numRows, numCols, viewOrientation)
	} else {
		ideal = BrickMosaic.EucPosterize(img, palette, numRows, numCols, viewOrientation)
	}

	// How are we going to build this mosaic?
	plan := BrickMosaic.CreateGridMosaic(ideal, BrickMosaic.GreedySolve)
	inventory := plan.Inventory()
	fmt.Printf("%v", inventory.DescendingUsage())
	fmt.Printf("Will cost approximately %d dollars to build", inventory.ApproximateCost() / 100)

	renderer := BrickMosaic.SVGRenderer{}
	if _, err := outputFile.Write([]byte(renderer.Render(plan))); err != nil {
		panic(err)
	}

	/*
			// TODO handle this more gracefully

			for _, scalingFactor := range []float32{0.0, 0.25, 0.5, 1.0, 2.0} {
			  baseName := filepath.Base(*inputPath)
			  ditheredImg := BrickMosaic.NewDitheredBrickImage(img, numRows, numCols, palette, viewOrientation, scalingFactor)
			  gifFile, err := os.Create(fmt.Sprintf("%v_%v_%d_rows_%d_cols_%v.gif", baseName, viewOrientation, numRows, numCols, scalingFactor))
		  	if err != nil {
		  		panic("Couldn't create output file")
		  	}
		  	// close the output file on exit and check for its returned error
		  	defer func() {
		  		if err := gifFile.Close(); err != nil {
		  			panic(err)
		  		}
		  	}()
		  	// TODO fixme
		  	frames := ditheredImg.Frames
		  	var delay []int
		  	for _ = range frames {
		  	  delay = append(delay, 100)
		  	}

		    g := &gif.GIF {
		      Image: frames,
		      Delay: delay,
		      LoopCount: -1,
		    }
		    err = gif.EncodeAll(gifFile, g)
		    if err != nil {
		      panic(fmt.Sprintf("Couldn't encode gif: %v", err))
		    }
			}*/
}
