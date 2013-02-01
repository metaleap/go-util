package unum

import (
	"math"
)

//	Represents a quaternion.
type Quat struct {
	X, Y, Z, W float64
}

func (me *Quat) Clone() (q *Quat) {
	*q = *me
	return
}

//	Conjugates this quaternion: swaps all signs for the X, Y and Z components only.
func (me *Quat) Conjugate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

//	Returns a new quaternion that is the conjugated representation of this quaternion.
func (me *Quat) Conjugated() *Quat {
	return &Quat{-me.X, -me.Y, -me.Z, me.W}
}

//	Returns the length of this quaternion.
func (me *Quat) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z) + (me.W * me.W)
}

//	Returns the magnitude of this quaternion.
func (me *Quat) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Normalizes this quaternion.
func (me *Quat) Normalize(bag *Bag) {
	bag.qfv = me.Magnitude()
	if bag.qfv == 0 {
		me.X, me.Y, me.Z, me.W = 0, 0, 0, 0
	} else {
		bag.qfv = 1 / bag.qfv
		me.X, me.Y, me.Z, me.W = me.X*bag.qfv, me.Y*bag.qfv, me.Z*bag.qfv, me.W*bag.qfv
	}
}

//	Returns a new quaternion that is the normalized representation of this quaternion.
func (me *Quat) Normalized(bag *Bag) (q *Quat) {
	q = &Quat{}
	bag.qfv = me.Magnitude()
	if bag.qfv != 0 {
		bag.qfv = 1 / bag.qfv
		q.X, q.Y, q.Z, q.W = me.X*bag.qfv, me.Y*bag.qfv, me.Z*bag.qfv, me.W*bag.qfv
	}
	return
}

//	Sets this quaternion to the conjugated representation of c.
func (me *Quat) SetFromConjugated(c *Quat) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

//	Sets this quaternion to the result of multiplying l with r.
func (me *Quat) SetFromMult(l, r *Quat) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

//	Multiplies this quaternion's components with v.
func (me *Quat) Scale(v float64) {
	me.X, me.Y, me.Z, me.W = me.X*v, me.Y*v, me.Z*v, me.W*v
}

//	Sets this quaternion to the result of multiplying all values in the specified 4x4 matrix with the specified 3D vector.
func (me *Quat) MultMat4Vec3(mat *Mat4, vec *Vec3) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * 1)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * 1)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * 1)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * 1)
}

//	Sets this quaternion to the result of multiplying all values in the specified 4x4 matrix with the specified quaternion.
func (me *Quat) MultMat4Vec4(mat *Mat4, vec *Quat) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * vec.Y)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * vec.Y)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * vec.Y)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * vec.Y)
}

//	Sets this quaternion to the result of multiplying quaternion q with vector v.
func (me *Quat) SetFromMult3(q *Quat, v *Vec3) {
	me.W = -(q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

//	Multiplies this quaternion with the specified matrix.
func (me *Quat) SetFromMultMat4(mat *Mat4) {
	me.MultMat4Vec4(mat, me.Clone())
}
