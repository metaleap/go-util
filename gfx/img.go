package ugfx

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"

	unum "github.com/metaleap/go-util/num"
)

//	The interface that the image package was missing...
type Picture interface {
	image.Image

	//	Set pixel at x,y to the specified color.
	Set(int, int, color.Color)
}

//	Creates and returns a copy of src.
//	If copyPixels is true, pixels in src are copied to dst, otherwise dst will be an
//	empty/black image of the same dimensions, color format, stride/offset/etc as src.
func CloneImage(src image.Image, copyPixels bool) (dst Picture) {
	makePix := func(pix []uint8) (cp []uint8) {
		if cp = make([]uint8, len(pix)); copyPixels {
			copy(cp, pix)
		}
		return
	}
	switch pic := src.(type) {
	case *image.Alpha:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.Alpha16:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.Gray:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.Gray16:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.NRGBA:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.NRGBA64:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.RGBA:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	case *image.RGBA64:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst = &clone
	default:
		rect := src.Bounds()
		tmpImg := image.NewRGBA(rect)
		if copyPixels {
			draw.Draw(tmpImg, rect, src, rect.Min, draw.Src)
		}
		dst = tmpImg
	}
	return
}

func PreprocessImage(src image.Image, dst Picture, flipY, toBgra, toLinear bool) {
	const preprocessParallel = true
	var wg sync.WaitGroup
	srgbToLinear := func(c *uint8) {
		var f float64
		if f = unum.Din1(float64(*c), 255); f > 0.04045 {
			f = math.Pow((f+0.055)/1.055, 2.4)
		} else {
			f = f / 12.92
		}
		*c = uint8(f * 255)
	}
	dstY, width, height := -1, src.Bounds().Dx(), src.Bounds().Dy()
	workRow := func(y, dy int) {
		if preprocessParallel {
			defer wg.Done()
		}
		for x := 0; x < width; x++ {
			col := src.At(x, y)
			if toLinear || toBgra {
				rgba := color.RGBAModel.Convert(col).(color.RGBA)
				if toBgra {
					rgba.R, rgba.B = rgba.B, rgba.R
				}
				if toLinear {
					srgbToLinear(&rgba.R)
					srgbToLinear(&rgba.G)
					srgbToLinear(&rgba.B)
				}
				col = rgba
			}
			dst.Set(x, dy, col)
		}
	}
	for y := height - 1; y >= 0; y-- {
		if flipY {
			dstY++
		} else {
			dstY = y
		}
		if preprocessParallel {
			wg.Add(1)
			go workRow(y, dstY)
		} else {
			workRow(y, dstY)
		}
	}
	if preprocessParallel {
		wg.Wait()
	}
}
