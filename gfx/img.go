package ugfx

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"

	unum "github.com/metaleap/go-util/num"
)

//	The "missing interface" from the `image` package:
//	`Set(x, y, color)` is implemented by most (but not all) `image` types that also implement `Image`.
type Picture interface {
	image.Image

	//	Set pixel at `x, y` to the specified `Color`.
	Set(int, int, color.Color)
}

//	Creates and returns a `Picture` just like `src`:
//
//	If `copyPixels` is `true`, pixels in `src` are copied to `dst`, otherwise `dst` will be an
//	empty/black `Picture` of the exact same dimensions, color format, stride/offset/etc as `src`.
//
//	The resulting `dst` will be of the same type as `src` if `src` is an `*image.Alpha`, `*image.Alpha16`, `*image.Gray`, `*image.Gray16`, `*image.NRGBA`, `*image.NRGBA16`, or `*image.RGBA64` --- otherwise, `dst` will be an `*image.RGBA`.
func CreateLike(src image.Image, copyPixels bool) (dst Picture, pix []byte) {
	makePix := func(pix []byte) (cp []byte) {
		if cp = make([]byte, len(pix)); copyPixels {
			copy(cp, pix)
		}
		return
	}
	switch pic := src.(type) {
	case *image.Alpha:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.Alpha16:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.Gray:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.Gray16:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.NRGBA:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.NRGBA64:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.RGBA:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	case *image.RGBA64:
		clone := *pic
		clone.Pix = makePix(pic.Pix)
		dst, pix = &clone, clone.Pix
	default:
		rect := src.Bounds()
		tmpImg := image.NewRGBA(rect)
		if copyPixels {
			draw.Draw(tmpImg, rect, src, rect.Min, draw.Src)
		}
		dst, pix = tmpImg, tmpImg.Pix
	}
	return
}

//	Processes the specified `Image` and writes the result to the specified `Picture`:
//
//	If `flipY` is `true`, all pixel rows are inverted (`dst` becomes `src` vertically mirrored).
//
//	If `toBgra` is `true`, all pixels' red and blue components are swapped.
//
//	If `toLinear` is `true`, all pixels are converted from gamma/sRGB to linear space --
//	only use this if you're certain that `src` is not already in linear space.
//
//	`dst` and `src` may point to the same `Image` object ONLY if `flipY` is `false`.
func PreprocessImage(src image.Image, dst Picture, flipY, toBgra, toLinear bool) {
	const preprocessParallel = true
	var wg sync.WaitGroup
	srgbToLinear := func(c *byte) {
		var f float64
		if f = unum.Din1(float64(*c), 255); f > 0.04045 {
			f = math.Pow((f+0.055)/1.055, 2.4)
		} else {
			f = f / 12.92
		}
		*c = byte(f * 255)
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
