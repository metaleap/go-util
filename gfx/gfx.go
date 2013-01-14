package gfx

import (
	"image"
	"image/png"
	"os"
)

//	Describes a literal color using four 32-bit floating-point numbers in RGBA order.
type Rgba32 struct {
	//	Red component
	R float32
	//	Green component
	G float32
	//	Blue component
	B float32
	//	Alpha component
	A float32
}

//	Converts the specified vals to a newly initialized Rgba32 instance.
//	The first 4 vals are used for R, G, B, and A in that order, if present.
//	A is set to 1 if vals[3] is not present.
func NewRgba32(vals ...float64) (me *Rgba32) {
	me = &Rgba32{}
	if len(vals) > 0 {
		if me.R = float32(vals[0]); len(vals) > 1 {
			if me.G = float32(vals[1]); len(vals) > 2 {
				if me.B = float32(vals[2]); len(vals) > 3 {
					me.A = float32(vals[3])
				} else {
					me.A = 1
				}
			}
		}
	}
	return
}

//	Describes a literal color using four 64-bit floating-point numbers in RGBA order.
type Rgba64 struct {
	//	Red component
	R float64
	//	Green component
	G float64
	//	Blue component
	B float64
	//	Alpha component
	A float64
}

//	Converts the specified vals to a newly initialized Rgba64 instance.
//	The first 4 vals are used for R, G, B, and A in that order, if present.
//	A is set to 1 if vals[3] is not present.
func NewRgba64(vals ...float64) (me *Rgba64) {
	me = &Rgba64{}
	if len(vals) > 0 {
		if me.R = vals[0]; len(vals) > 1 {
			if me.G = vals[1]; len(vals) > 2 {
				if me.B = vals[2]; len(vals) > 3 {
					me.A = vals[3]
				} else {
					me.A = 1
				}
			}
		}
	}
	return
}

//	If 2-dimensions are represented in a 1-dimensional linear array, this function provides one way to return a 1D index addressing a 2D coordinate...
func Index2D(x, y, ysize int) int {
	return (x * ysize) + y
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
