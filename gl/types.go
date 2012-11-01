package glutil

import (
	gl "github.com/chsc/gogl/gl42"

	numutil "github.com/go3d/go-util/num"
)

type TGlVec4 [4]gl.Float

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
		SizeOfGlFloat gl.Sizeiptr = 4
		SizeOfGlUint gl.Sizeiptr = 4
	)

func ClampF (val, min, max gl.Float) gl.Float {
	return MinF(MaxF(val, min), max)
}

func IfE (cond bool, ifTrue, ifFalse gl.Enum) gl.Enum {
	if cond { return ifTrue }
	return ifFalse
}

func IfI (cond bool, ifTrue, ifFalse gl.Int) gl.Int {
	if cond { return ifTrue }
	return ifFalse
}

func MaxF (one, two gl.Float) gl.Float {
	if two > one { return two }
	return one
}

func MinF (one, two gl.Float) gl.Float {
	if two < one { return two }
	return one
}
