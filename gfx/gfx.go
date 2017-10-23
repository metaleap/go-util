package ugfx

import (
	"image"
	"image/png"
	"math"
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

//	Converts the specified `vals` to a newly initialized `Rgba32` instance.
//
//	The first 4 `vals` are used for `R`, `G`, `B`, and `A` in that order, if present.
//	`A` is set to 1 if `vals[3]` is not present.
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

//	Converts the specified `vals` to a newly initialized `Rgba64` instance.
//
//	The first 4 `vals` are used for `R`, `G`, `B`, and `A` in that order, if present.
//	`A` is set to 1 if `vals[3]` is not present.
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

//	Converts the given value from gamma to linear color space.
func GammaToLinearSpace(f float64) float64 {
	if f > 0.0404482362771082 {
		return math.Pow((f+0.055)/1.055, 2.4)
	}
	return f / 12.92
}

//	If 2 dimensions are represented in a 1-dimensional linear array, this function provides a way to return a 1D index addressing the specified 2D coordinate.
func Index2D(x, y, ysize int) int {
	return (x * ysize) + y
}

//	If 3 dimensions are represented in a 1-dimensional linear array, this function provides a way to return a 1D index addressing the specified 3D coordinate.
func Index3D(x, y, z, xsize, ysize int) int {
	return (((z * xsize) + x) * ysize) + y
}

//	Converts the given value from linear to gamma color space.
func LinearToGammaSpace(f float64) float64 {
	if f > 0.00313066844250063 {
		return 1.055*math.Pow(f, 1/2.4) - 0.055
	}
	return f * 12.92
}

//	Saves any given `Image` as a local PNG file.
func SavePngImageFile(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err == nil {
		defer file.Close()
		err = png.Encode(file, img)
	}
	return err
}
