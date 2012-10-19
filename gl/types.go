package glutil

import (
	gl "github.com/chsc/gogl/gl42"

	numutil "github.com/go-ngine/go-util/num"
)

type TGlMat3 [9]gl.Float

func NewGlMat3 (mat *numutil.TMat3) *TGlMat3 {
	var glMat = &TGlMat3 {}; glMat.Load(mat); return glMat
}

func (me *TGlMat3) Load (mat *numutil.TMat3) {
	me[0], me[3], me[6] = gl.Float(mat[0]), gl.Float(mat[3]), gl.Float(mat[6])
	me[1], me[4], me[7] = gl.Float(mat[1]), gl.Float(mat[4]), gl.Float(mat[7])
	me[2], me[5], me[8] = gl.Float(mat[2]), gl.Float(mat[5]), gl.Float(mat[8])
}

type TGlMat4 [16]gl.Float

func NewGlMat4 (mat *numutil.TMat4) *TGlMat4 {
	var glMat = &TGlMat4 {}; glMat.Load(mat); return glMat
}

func (me *TGlMat4) Load (mat *numutil.TMat4) {
	me[0], me[4], me[8], me[12] = gl.Float(mat[0]), gl.Float(mat[4]), gl.Float(mat[8]), gl.Float(mat[12])
	me[1], me[5], me[9], me[13] = gl.Float(mat[1]), gl.Float(mat[5]), gl.Float(mat[9]), gl.Float(mat[13])
	me[2], me[6], me[10], me[14] = gl.Float(mat[2]), gl.Float(mat[6]), gl.Float(mat[10]), gl.Float(mat[14])
	me[3], me[7], me[11], me[15] = gl.Float(mat[3]), gl.Float(mat[7]), gl.Float(mat[11]), gl.Float(mat[15])
}

var (
	SizeOfGlUint gl.Sizeiptr = 4
)

func Fin1 (val, max gl.Float) gl.Float {
	return 1 / (max / val)
}

func Ife (cond bool, ifTrue, ifFalse gl.Enum) gl.Enum {
	if cond { return ifTrue }
	return ifFalse
}

func Ifi (cond bool, ifTrue, ifFalse gl.Int) gl.Int {
	if cond { return ifTrue }
	return ifFalse
}

func Ifui (cond bool, ifTrue, ifFalse gl.Uint) gl.Uint {
	if cond { return ifTrue }
	return ifFalse
}

func InSliceAt (vals []gl.Enum, val gl.Enum) int {
	for i, v := range vals {
		if v == val {
			return i
		}
	}
	return -1
}

func Mins (s1, s2 gl.Sizei) gl.Sizei {
	if s2 < s1 { return s2 }; return s1;
}

func OffsetIntPtr (ptr gl.Pointer, offset gl.Sizei) gl.Intptr {
	return gl.Intptr(uintptr(ptr) + uintptr(offset))
}

func OffsetPointer (ptr gl.Pointer, offset uint) gl.Pointer {
	return gl.Pointer(uintptr(ptr) + uintptr(offset))
}

func WriteFloats4 (vecX, vecY, vecZ, vecW float64, index int, glFloats []gl.Float) {
	glFloats[index + 0], glFloats[index + 1], glFloats[index + 2], glFloats[index + 3] = gl.Float(vecX), gl.Float(vecY), gl.Float(vecZ), gl.Float(vecW)
}

func WriteFloatsVec3 (vec numutil.TVec3, index int, glFloats []gl.Float) {
	glFloats[index + 0], glFloats[index + 1], glFloats[index + 2] = gl.Float(vec.X), gl.Float(vec.Y), gl.Float(vec.Z)
}

func WriteFloatsVec4 (vec numutil.TVec3, vecW float64, index int, glFloats []gl.Float) {
	glFloats[index + 0], glFloats[index + 1], glFloats[index + 2], glFloats[index + 3] = gl.Float(vec.X), gl.Float(vec.Y), gl.Float(vec.Z), gl.Float(vecW)
}
