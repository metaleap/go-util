package unum

import (
	"math"
)

type Vec4 struct {
	Vec3
	W float64
}

func (me *Vec4) Clone() (q *Vec4) {
	*q = *me
	return
}

func (me *Vec4) Conjugate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

func (me *Vec4) Conjugated() (v *Vec4) {
	v = new(Vec4)
	v.X, v.Y, v.Z, v.W = -me.X, -me.Y, -me.Z, me.W
	return
}

func (me *Vec4) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z) + (me.W * me.W)
}

func (me *Vec4) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

func (me *Vec4) Normalize() {
	if mag := me.Magnitude(); mag == 0 {
		me.X, me.Y, me.Z, me.W = 0, 0, 0, 0
	} else {
		mag = 1 / mag
		me.X, me.Y, me.Z, me.W = me.X*mag, me.Y*mag, me.Z*mag, me.W*mag
	}
}

func (me *Vec4) NormalizeFrom(magnitude float64) {
	magnitude = 1 / magnitude
	me.X, me.Y, me.Z, me.W = me.X*magnitude, me.Y*magnitude, me.Z*magnitude, me.W*magnitude
}

func (me *Vec4) Normalized() *Vec4 {
	var q Vec4
	if mag := me.Magnitude(); mag != 0 {
		mag = 1 / mag
		q.X, q.Y, q.Z, q.W = me.X*mag, me.Y*mag, me.Z*mag, me.W*mag
	}
	return &q
}

func (me *Vec4) SetFromConjugated(c *Vec4) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

func (me *Vec4) SetFromMult(l, r *Vec4) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

func (me *Vec4) Scale(v float64) {
	me.X, me.Y, me.Z, me.W = me.X*v, me.Y*v, me.Z*v, me.W*v
}

func (me *Vec4) MultMat4Vec3(mat *Mat4, vec *Vec3) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * 1)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * 1)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * 1)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * 1)
}

func (me *Vec4) MultMat4Vec4(mat *Mat4, vec *Vec4) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * vec.W)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * vec.W)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * vec.W)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * vec.W)
}

func (me *Vec4) SetFromMult3(q *Vec4, v *Vec3) {
	me.W = -(q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

func (me *Vec4) SetFromMultMat4(mat *Mat4) {
	me.MultMat4Vec4(mat, me.Clone())
}

//	Returns a string representation of this 3D vector.
func (me *Vec4) String() string {
	return strf("{X:%1.2f Y:%1.2f Z:%1.2f W:%1.2f}", me.X, me.Y, me.Z, me.W)
}
