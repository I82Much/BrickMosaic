BrickMosaic
===========

A program for converting images into Lego mosaics

# Installation
Get the svg library:

    go get github.com/ajstarks/svgo

Get the mosaic library:
    go get github.com/I82Much/BrickMosaic

Build the library:
    go build

Run the binary:
    go run main.go --rows=40 --cols=100 --path=/path/to/image.jpg --output_path=/path/to/destination.svg

Once the program is complete, you can view the svg in a web browser to see the instructions.