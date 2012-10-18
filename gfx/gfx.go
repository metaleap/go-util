package gfx

import (
	"image"
	"image/png"
	"os"
)

func Index3D (x, y, z, xsize, ysize int) int {
	return (((z * xsize) + x) * ysize) + y
}

func SaveImageFile (img image.Image, filePath string) error {
	var file, err = os.Create(filePath)
	if err == nil {
		defer file.Close()
		png.Encode(file, img)
	}
	return err
}
