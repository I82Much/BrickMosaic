BrickMosaic
===========

A program for converting images into LEGO mosaics.

# Installation
Get the svg library:

    go get github.com/ajstarks/svgo

Get the mosaic library:
    go get github.com/I82Much/BrickMosaic

Build the library:
    go build github.com/I82Much/BrickMosaic

Run the binary:
    go run main.go --rows=40 --cols=100 --path=/path/to/image.jpg --output_path=/path/to/destination.svg

Once the program is complete, you can view the svg in a web browser to see the instructions.

# Examples:

## Darth Vader
Here is the Darth Vader mosaic I created earlier this summer using a prior version of this program:

![Darth Vader finished][]
![Darth Vader in progress][]

## Iron Man
Here are three versions of Iron Man ([original image][Iron Man Original]):

Studs Up:
![Iron Man Studs Top][]

Studs Right:
![Iron Man Studs Right][]

Studs Out:
![Iron Man Studs Out][]

[Darth Vader finished]:https://lh5.googleusercontent.com/-mCzaiPeMcjc/UhzHhp9VYBI/AAAAAAAAQCo/mMZ3RbQDcp8/w620-h615-no/IMG_7020_cropped.jpg
[Darth Vader in progress]:https://lh3.googleusercontent.com/-6NUdcIC2hbc/UhzHgZZqPdI/AAAAAAAAQCY/SBcVMYzJPfg/w820-h615-no/IMG_6188.JPG
[Iron Man Original]:http://marvelwallpapers10.net/wp-content/uploads/images/3a/ironman-2.jpg
[Iron Man Studs Top]:https://rawgithub.com/I82Much/BrickMosaic/master/examples/iron_man_STUDS_TOP.svg
[Iron Man Studs Right]:https://rawgithub.com/I82Much/BrickMosaic/master/examples/iron_man_STUDS_RIGHT.svg
[Iron Man Studs Out]:https://rawgithub.com/I82Much/BrickMosaic/master/examples/iron_man_STUDS_OUT.svg
