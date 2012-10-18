package glutil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"

	gl "github.com/chsc/gogl/gl42"
)

const (
	TEXTURE_MAX_ANISOTROPY_EXT = 0x84FE
	MAX_TEXTURE_MAX_ANISOTROPY_EXT = 0x84FF
)

var (
	TexFilter gl.Int = gl.NEAREST
	TexWrap gl.Int = gl.CLAMP_TO_BORDER

	maxTexAnisotropy gl.Float = -100
	maxTexBufSize, maxTexSize1D, maxTexSize2D, maxTexSize3D gl.Int = 0, 0, 0, 0
)

func FindTexSize2D (maxSize, numTexels, minSize float64) (float64, float64) {
	var wh float64
	if math.Floor(numTexels) != numTexels { log.Panicf("AAAAH %v", numTexels) }
	if numTexels <= maxSize { return numTexels, 1 }
	for h := 2.0; h < maxSize; h ++ {
		for w := 2.0; w < maxSize; w ++ {
			wh = w * h
			if wh == numTexels {
				if minSize > 0 { for (math.Mod(w, 2) == 0) && (math.Mod(h, 2) == 0) && ((w / 2) >= minSize) && ((h / 2) >= minSize) { w, h = w / 2, h / 2 } }
				for ((h * 2) < w) && (math.Mod(w, 2) == 0) { w, h = w / 2, h * 2 }
				if minSize > 0 { for (math.Mod(w, 2) == 0) && (math.Mod(h, 2) == 0) && ((w / 2) >= minSize) && ((h / 2) >= minSize) { w, h = w / 2, h / 2 } }
				return w, h
			} else if wh > numTexels { break }
		}
	}
	return 0, 0
}

func MakeTexture (glPtr *gl.Uint, dimensions uint8, texFormat gl.Enum, width, height, depth gl.Sizei) error {
	return MakeTextureForTarget(glPtr, dimensions, width, height, depth, 0, texFormat, false)
}

func MakeTextureForTarget (glPtr *gl.Uint, dimensions uint8, width, height, depth gl.Sizei, texTarget gl.Enum, texFormat gl.Enum, reuseGlPtr bool) error {
	var is3d, is2d = (dimensions == 3), (dimensions == 2)
	if texTarget == 0 { texTarget = Ife(is3d, gl.TEXTURE_3D, Ife(is2d, gl.TEXTURE_2D, gl.TEXTURE_1D)) }
	if texFormat == 0 { texFormat = gl.RGBA8 }
	if width == 0 { return errors.New("MakeTextureForTarget() needs at least width") }
	if height == 0 { height = width }
	if depth == 0 { depth = height }
	if (!reuseGlPtr) || (*glPtr == 0) { gl.GenTextures(1, glPtr) }
	gl.BindTexture(texTarget, *glPtr)
	gl.TexParameteri(texTarget, gl.TEXTURE_MAG_FILTER, TexFilter)
	gl.TexParameteri(texTarget, gl.TEXTURE_MIN_FILTER, TexFilter)
	gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_S, TexWrap)
	if is2d || is3d { gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_T, TexWrap) }
	if is3d { gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_R, TexWrap) }
	if is3d {
		gl.TexStorage3D(texTarget, 1, texFormat, width, height, depth)
	} else if is2d {
		gl.TexStorage2D(texTarget, 1, texFormat, width, height)
	} else {
		gl.TexStorage1D(texTarget, 1, texFormat, width)
	}
	gl.BindTexture(texTarget, 0)
	return LastError("MakeTextureForTarget(dim=%v w=%v h=%v d=%v)", dimensions, width, height, depth)
}

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
			gl.TexStorage2D(gl.TEXTURE_2D, 1, gl.RGBA8, sw, sh)
			gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, sw, sh, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&rgba.Pix[0]))
			if (mipMaps) { gl.GenerateMipmap(gl.TEXTURE_2D) }
			gl.BindTexture(gl.TEXTURE_2D, 0)
		}
	}
	return tex
}

func MakeTextureFromImageFloatsFile (filePath string, w, h int) (gl.Uint, error) {
	var file, err = os.Open(filePath)
	var tex gl.Uint
	var pix = make([]gl.Float, w * h * 3)
	var fVal float32
	var raw []uint8
	var buf *bytes.Buffer
	var i int
	if err != nil { return tex, err }
	defer file.Close()
	raw, err = ioutil.ReadAll(file)
	if err != nil { return tex, err }
	buf = bytes.NewBuffer(raw)
	for i = 0; (err == nil) && (i < len(pix)); i++ {
		if err = binary.Read(buf, binary.LittleEndian, &fVal); err == io.EOF {
			err = nil; break
		} else if err == nil {
			pix[i] = gl.Float(fVal)
		}
	}
	if err != nil { return tex, err }
	sw, sh := gl.Sizei(w), gl.Sizei(h)
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexStorage2D(gl.TEXTURE_2D, 1, gl.RGB16F, sw, sh)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, sw, sh, gl.RGB, gl.FLOAT, gl.Pointer(&pix[0]))
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return tex, LastError("MakeTextureFromImageFloatsFile(%v)", filePath)
}

func MakeTextures (num gl.Sizei, glPtr []gl.Uint, dimensions uint8, texFormat gl.Enum, width, height, depth gl.Sizei) error {
	var err error
	var is3d, is2d = (dimensions == 3), (dimensions == 2)
	var texTarget gl.Enum = Ife(is3d, gl.TEXTURE_3D, Ife(is2d, gl.TEXTURE_2D, gl.TEXTURE_1D))
	gl.GenTextures(num, &glPtr[0])
	if err = LastError("MakeTextures.GenTextures(num=%v)", num); err != nil { return err }
	if width == 0 { return errors.New("MakeTextures() needs at least width") }
	if height == 0 { height = width }
	if depth == 0 { depth = height }
	for i := 0; i < len(glPtr); i++ {
		gl.BindTexture(texTarget, glPtr[i])
		gl.TexParameteri(texTarget, gl.TEXTURE_MAG_FILTER, TexFilter)
		gl.TexParameteri(texTarget, gl.TEXTURE_MIN_FILTER, TexFilter)
		gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_R, TexWrap)
		gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_S, TexWrap)
		gl.TexParameteri(texTarget, gl.TEXTURE_WRAP_T, TexWrap)
		if is3d {
			gl.TexStorage3D(texTarget, 1, texFormat, width, height, depth)
		} else if is2d {
			gl.TexStorage2D(texTarget, 1, texFormat, width, height)
		} else {
			gl.TexStorage1D(texTarget, 1, texFormat, width)
		}
		gl.BindTexture(texTarget, 0)
		if err = LastError("MakeTextures.Loop(i=%v dim=%v w=%v h=%v d=%v)", i, dimensions, width, height, depth); err != nil { return err }
	}
	return err
}

func MaxTextureAnisotropy () gl.Float {
	if maxTexAnisotropy == -100 {
		maxTexAnisotropy = 0
		if Extension("texture_filter_anisotropic") { gl.GetFloatv(MAX_TEXTURE_MAX_ANISOTROPY_EXT, &maxTexAnisotropy) }
	}
	return maxTexAnisotropy
}

func MaxTextureBufferSize () (uint64, error) {
	var err error
	if maxTexBufSize == 0 {
		gl.GetIntegerv(gl.MAX_TEXTURE_BUFFER_SIZE, &maxTexBufSize)
		err = LastError("MaxTextureBufferSize()")
	}
	return uint64(maxTexBufSize), err
}
