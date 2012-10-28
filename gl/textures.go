package glutil

import (
	"fmt"
	"image"
	"math"
	"reflect"

	gl "github.com/chsc/gogl/gl42"
)

const (
	TEXTURE_MAX_ANISOTROPY_EXT = 0x84FE
	MAX_TEXTURE_MAX_ANISOTROPY_EXT = 0x84FF
)

var (
	maxTexAnisotropy gl.Float = -100
)

func ImageTextureProperties (img image.Image, width, height, numLevels *gl.Sizei, sizedInternalFormat, pixelDataFormat, pixelDataType *gl.Enum) (pixPtr gl.Pointer) {
	*pixelDataType = gl.UNSIGNED_BYTE
	*numLevels = gl.Sizei(int(math.Log2(math.Max(float64(img.Bounds().Dx()), float64(img.Bounds().Dy()))) + 1))
	*width, *height = gl.Sizei(img.Bounds().Dx()), gl.Sizei(img.Bounds().Dy())
	switch pic := img.(type) {
	case *image.Alpha:
		*sizedInternalFormat = gl.R8
		*pixelDataFormat = gl.RED
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.Alpha16:
		*sizedInternalFormat = gl.R16
		*pixelDataFormat = gl.RED
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.Gray:
		*sizedInternalFormat = gl.R8
		*pixelDataFormat = gl.RED
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.Gray16:
		*sizedInternalFormat = gl.R16
		*pixelDataFormat = gl.RED
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.NRGBA:
		*sizedInternalFormat = gl.RGBA8
		*pixelDataFormat = gl.RGBA
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.NRGBA64:
		*sizedInternalFormat = gl.RGBA16
		*pixelDataFormat = gl.RGBA
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.RGBA:
		*sizedInternalFormat = gl.RGBA8
		*pixelDataFormat = gl.RGBA
		pixPtr = gl.Pointer(&pic.Pix[0])
	case *image.RGBA64:
		*sizedInternalFormat = gl.RGBA16
		*pixelDataFormat = gl.RGBA
		pixPtr = gl.Pointer(&pic.Pix[0])
	default:
		panic(fmt.Errorf("Unsupported image.Image type (%v) for use as OpenGL texture", reflect.TypeOf(pic)))
	}
	return
}

/*
func MakeTextureFromImageFile (filePath string, wrapS, wrapT, minFilter, magFilter gl.Int, mipMaps bool) gl.Uint {
	var file, err = os.Open(filePath)
	var img image.Image
	var tex gl.Uint
	if err == nil {
		defer file.Close()
		img, _, err = image.Decode(file)
		if err == nil {
			w, h := img.Bounds().Dx(), img.Bounds().Dy()
			sw, sh := gl.Sizei(w), gl.Sizei(h)
			// GL_BGRA, GL_UNSIGNED_INT_8_8_8_8_REV
			rgba := image.NewRGBA(image.Rect(0, 0, w, h))
			for x := 0; x < w; x++ { for y := 0; y < h; y++ { rgba.Set(x, y, img.At(x, y)) } }
			gl.GenTextures(1, &tex)
			gl.BindTexture(gl.TEXTURE_2D, tex)
			if MaxTextureAnisotropy() > 0 {
				gl.TexParameterf(gl.TEXTURE_2D, TEXTURE_MAX_ANISOTROPY_EXT, maxTexAnisotropy)
			}
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, magFilter)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, minFilter)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrapS)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrapT)
			if IsGl42 {
				gl.TexStorage2D(gl.TEXTURE_2D, 1, gl.RGBA8, sw, sh)
				gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, sw, sh, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&rgba.Pix[0]))
			} else {
				gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, sw, sh, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&rgba.Pix[0]))
			}
			if (mipMaps) { gl.GenerateMipmap(gl.TEXTURE_2D) }
			gl.BindTexture(gl.TEXTURE_2D, 0)
		}
	}
	return tex
}
*/

func MaxTextureAnisotropy () gl.Float {
	if maxTexAnisotropy == -100 {
		maxTexAnisotropy = 0
		if Extension("texture_filter_anisotropic") {
			gl.GetFloatv(MAX_TEXTURE_MAX_ANISOTROPY_EXT, &maxTexAnisotropy)
		}
	}
	return maxTexAnisotropy
}
