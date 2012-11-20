package gfx

import (
	"image"
	"image/png"
	"os"
)

//	If 3-dimensions are represented in a 1-dimensional linear array, this function provides one way to return a 1D index addressing a 3D coordinate...
func Index3D (x, y, z, xsize, ysize int) int {
	return (((z * xsize) + x) * ysize) + y
}

//	Saves any given image as a local PNG file.
func SavePngImageFile (img image.Image, filePath string) error {
	var file, err = os.Create(filePath)
	if err == nil {
		defer file.Close()
		png.Encode(file, img)
	}
	return err
}
