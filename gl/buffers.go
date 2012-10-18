package glutil

import (
	"fmt"

	gl "github.com/chsc/gogl/gl42"
)

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
)

var (
	BufMode gl.Enum = gl.STATIC_DRAW
	FillWithZeroes = false
)

func MakeArrayBuffer (glPtr *gl.Uint, size uint64, sl interface{}, isLen, makeTex bool) (gl.Uint, error) {
	var err error
	var ptr = gl.Pointer(nil)
	var glTex gl.Uint = 0
	var glTexFormat gl.Enum = gl.R8UI
	var sizeFactor, sizeTotal, sizeMax uint64 = 1, 0, 0
	var tm = false
	var handle = func (sf uint64, glPtr gl.Pointer, le int, tf gl.Enum) {
		tm = true; if le > 1 { ptr = glPtr }; if size == 0 { size = uint64(le); isLen = true }; if isLen { sizeFactor = sf }; if tf != 0 { glTexFormat = tf }
	}
	if (sl == nil) && FillWithZeroes { sl = make([]uint8, size) }
	gl.GenBuffers(1, glPtr)
	gl.BindBuffer(gl.ARRAY_BUFFER, *glPtr)
	if sl != nil {
		if tv, tb := sl.([]uint8); tb { handle(1, gl.Pointer(&tv[0]), len(tv), gl.R8UI) }
		if tv, tb := sl.([]uint16); tb { handle(2, gl.Pointer(&tv[0]), len(tv), gl.R16UI) }
		if tv, tb := sl.([]uint32); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32UI) }
		if tv, tb := sl.([]uint64); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32UI) }
		if tv, tb := sl.([]int8); tb { handle(1, gl.Pointer(&tv[0]), len(tv), gl.R8I) }
		if tv, tb := sl.([]int16); tb { handle(2, gl.Pointer(&tv[0]), len(tv), gl.R16I) }
		if tv, tb := sl.([]int32); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32I) }
		if tv, tb := sl.([]int64); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32I) }
		if tv, tb := sl.([]float32); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32F) }
		if tv, tb := sl.([]float64); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32F) }
		if tv, tb := sl.([]gl.Bitfield); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R8UI) }
		if tv, tb := sl.([]gl.Byte); tb { handle(1, gl.Pointer(&tv[0]), len(tv), gl.R8I) }
		if tv, tb := sl.([]gl.Ubyte); tb { handle(1, gl.Pointer(&tv[0]), len(tv), gl.R8UI) }
		if tv, tb := sl.([]gl.Ushort); tb { handle(2, gl.Pointer(&tv[0]), len(tv), gl.R16UI) }
		if tv, tb := sl.([]gl.Short); tb { handle(2, gl.Pointer(&tv[0]), len(tv), gl.R16I) }
		if tv, tb := sl.([]gl.Uint); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32UI) }
		if tv, tb := sl.([]gl.Uint64); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32UI) }
		if tv, tb := sl.([]gl.Int); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32I) }
		if tv, tb := sl.([]gl.Int64); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32I) }
		if tv, tb := sl.([]gl.Clampd); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32F) }
		if tv, tb := sl.([]gl.Clampf); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32F) }
		if tv, tb := sl.([]gl.Float); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32F) }
		if tv, tb := sl.([]gl.Half); tb { handle(2, gl.Pointer(&tv[0]), len(tv), gl.R16F) }
		if tv, tb := sl.([]gl.Double); tb { handle(8, gl.Pointer(&tv[0]), len(tv), gl.RG32F) }
		if tv, tb := sl.([]gl.Enum); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32I) }
		if tv, tb := sl.([]gl.Sizei); tb { handle(4, gl.Pointer(&tv[0]), len(tv), gl.R32UI) }
		if tv, tb := sl.([]gl.Char); tb { handle(1, gl.Pointer(&tv[0]), len(tv), gl.R8UI) }
		if !tm { err = fmt.Errorf("MakeArrayBuffer() -- slice type unsupported:\n%+v", sl) }
	}
	if (err == nil) {
		sizeTotal = size * sizeFactor
		gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(sizeTotal), ptr, BufMode)
		if makeTex {
			if sizeMax, err = MaxTextureBufferSize(); err == nil {
				if  sizeTotal > sizeMax {
					err = fmt.Errorf("Texture buffer size (%vMB) would exceed your GPU's maximum texture buffer size (%vMB)", sizeTotal / MB, maxTexBufSize / MB)
				} else {
					gl.GenTextures(1, &glTex)
					gl.BindTexture(gl.TEXTURE_BUFFER, glTex)
					gl.TexBuffer(gl.TEXTURE_BUFFER, glTexFormat, *glPtr)
					gl.BindTexture(gl.TEXTURE_2D, 0)
				}
			}
		}
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		err = LastError("MakeArrayBuffer()")
	}
	return glTex, err
}

func MakeArrayBuffer32 (glPtr *gl.Uint, sliceOrLen interface{}, makeTex bool) (gl.Uint, error) {
	if le, isLe := sliceOrLen.(uint64); isLe {
		return MakeArrayBuffer(glPtr, le * 4, nil, false, makeTex)
	}
	return MakeArrayBuffer(glPtr, 0, sliceOrLen, true, makeTex)
}

func MakeAtomicCounters (glPtr *gl.Uint, num gl.Sizei) {
	gl.GenBuffers(1, glPtr)
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, *glPtr)
	gl.BufferData(gl.ATOMIC_COUNTER_BUFFER, gl.Sizeiptr(4 * num), gl.Pointer(nil), BufMode)
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, 0)
}

func MakeRttFramebuffer (texFormat gl.Enum, width, height, mipLevels gl.Sizei, anisoFiltering gl.Int) (gl.Uint, gl.Uint, error) {
	var glPtrTex, glPtrFrameBuf gl.Uint
	var glMagFilter, glMinFilter gl.Int = gl.LINEAR, gl.LINEAR
	if anisoFiltering < 1 { glMagFilter, glMinFilter = gl.NEAREST, gl.NEAREST }
	if mipLevels <= 0 { mipLevels = 1; for size := Mins(width, height); size > 1; size /= 2 { mipLevels++ } }
	if mipLevels > 1 { if anisoFiltering < 1 { glMinFilter = gl.NEAREST_MIPMAP_NEAREST } else { glMinFilter = gl.LINEAR_MIPMAP_LINEAR } }
	gl.GenTextures(1, &glPtrTex)
	gl.BindTexture(gl.TEXTURE_2D, glPtrTex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, glMagFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, glMinFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)
	if anisoFiltering > 1 { gl.TexParameteri(gl.TEXTURE_2D, 0x84FE, anisoFiltering) }  // max 16
	gl.TexStorage2D(gl.TEXTURE_2D, mipLevels, texFormat, width, height)
	gl.GenFramebuffers(1, &glPtrFrameBuf)
	gl.BindFramebuffer(gl.FRAMEBUFFER, glPtrFrameBuf)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, glPtrTex, 0)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return glPtrFrameBuf, glPtrTex, LastError("MakeRttFramebuffer")
}
