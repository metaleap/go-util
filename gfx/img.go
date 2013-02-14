package ugfx

import (
	"image"
	"image/color"
	"math"

	unum "github.com/metaleap/go-util/num"
)

//	The interface that the image package was missing...
type Picture interface {
	image.Image

	//	Set pixel at x,y to the specified color.
	Set(int, int, color.Color)
}

//	Creates and returns a copy of src.
//	If copyPixels is true, pixels in src are copied to dst, otherwise dst will be an empty/black image of the same dimensions, color format, stride/offset/etc as src.
func CloneImage(src image.Image, copyPixels bool) (dst Picture) {
	switch pic := src.(type) {
	case *image.Alpha:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.Alpha16:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.Gray:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.Gray16:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.NRGBA:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.NRGBA64:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.RGBA:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	case *image.RGBA64:
		clone := *pic
		if clone.Pix = make([]uint8, len(pic.Pix)); copyPixels {
			copy(clone.Pix, pic.Pix)
		}
		dst = &clone
	default:
	}
	return
}

//	Sets dst to the vertically-inverted picture of src.
//	dst must have at least the same width and height as src and shouldn't be the same pointer as src.
func FlipVertical(src image.Image, dst Picture) {
	var (
		col  color.Color
		dstY int
	)
	width, height := src.Bounds().Dx(), src.Bounds().Dy()
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			col = src.At(x, y)
			dst.Set(x, dstY, col)
		}
		dstY++
	}
}

func SrgbToLinear(src image.Image, dst Picture) {
	var (
		col color.RGBA
		f   float64
	)
	conv := func(c *uint8) {
		if f = unum.Din1(float64(*c), 255); f > 0.04045 {
			f = math.Pow((f+0.055)/1.055, 2.4)
		} else {
			f = f / 12.92
		}
		*c = uint8(f * 255)
	}
	width, height := src.Bounds().Dx(), src.Bounds().Dy()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			col = color.RGBAModel.Convert(src.At(x, y)).(color.RGBA)
			conv(&col.R)
			conv(&col.G)
			conv(&col.B)
			dst.Set(x, y, col)
		}
	}
}
