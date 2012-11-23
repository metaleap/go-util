package num

import (
	"math"
)

//	Represents a quaternion.
type Quat struct {
	X, Y, Z, W float64
}

//	Conjugates this quaternion.
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
func (me *Quat) Normalize() {
	tfQ = me.Magnitude()
	if tfQ == 0 {
		me.X, me.Y, me.Z = tfQ, tfQ, tfQ
	} else {
		tfQ = 1 / tfQ
		me.X, me.Y, me.Z, me.W = me.X*tfQ, me.Y*tfQ, me.Z*tfQ, me.W*tfQ
	}
}

//	Returns a new quaternion that is the normalized representation of this quaternion.
func (me *Quat) Normalized() *Quat {
	tfQ = me.Magnitude()
	if tfQ == 0 {
		return &Quat{tfQ, tfQ, tfQ, tfQ}
	}
	tfQ = 1 / tfQ
	return &Quat{me.X * tfQ, me.Y * tfQ, me.Z * tfQ, me.W * tfQ}
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
func (me *Quat) SetFromMult1(v float64) {
	me.X, me.Y, me.Z, me.W = me.X*v, me.Y*v, me.Z*v, me.W*v
}

//	Sets this quaternion to the result of multiplying quaternion q with vector v.
func (me *Quat) SetFromMult3(q *Quat, v *Vec3) {
	me.W = -(q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

//	Multiplies this quaternion with the specified matrix.
func (me *Quat) SetFromMult4(mat *Mat4) {
	*me = *(mat.Mult4(me))
}
