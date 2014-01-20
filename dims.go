package BrickMosaic

// The LDraw Co-ordinate System
// LDraw uses a right-handed co-ordinate system where -Y is "up".
// http://www.ldraw.org/article/218.html

// -Y - up
// +Z - into the screen
// +X - to the right

// All measurements are in terms of LDU
type LDU int

const (
	BrickWidth   LDU = 20
	BrickHeight  LDU = 24
	PlateHeight  LDU = 8
	StudDiameter LDU = 12
	StudHeight   LDU = 4
)

func ApproxSizeInch(size LDU) float32 {
	return float32(1.0 * size / 64.0)
}

func ApproxSizeMm(size LDU) float32 {
	return 0.4 * float32(size)
}
