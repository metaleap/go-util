package gfx

import (
	"image"
	"image/png"
	"os"
)

type Rgba32 struct {
	R, G, B, A float32
}

func NewRgba32(vals []float64) *Rgba32 {
	return &Rgba32{R: float32(vals[0]), G: float32(vals[1]), B: float32(vals[2]), A: float32(vals[3])}
}

type Rgba64 struct {
	R, G, B, A float64
}

//	If 3-dimensions are represented in a 1-dimensional linear array, this function provides one way to return a 1D index addressing a 3D coordinate...
func Index3D(x, y, z, xsize, ysize int) int {
	return (((z * xsize) + x) * ysize) + y
}

//	Saves any given image as a local PNG file.
func SavePngImageFile(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err == nil {
		defer file.Close()
		png.Encode(file, img)
	}
	return err
}
