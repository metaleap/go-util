package ugfx

import (
	"image"
	"image/color"
)

type Picture interface {
	Set(int, int, color.Color)
}

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
