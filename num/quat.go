package num

import (
	"math"
)

type TQuat struct {
	X, Y, Z, W float64
}

func (me *TQuat) Conjugate () {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

func (me *TQuat) Conjugated () *TQuat {
	return &TQuat { -me.X, -me.Y, -me.Z, me.W }
}

func (me *TQuat) Length () float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z) + (me.W * me.W)
}

func (me *TQuat) Magnitude () float64 {
	return math.Sqrt(me.Length())
}

func (me *TQuat) Normalize () {
	tfQ = me.Magnitude()
	if tfQ == 0 {
		me.X, me.Y, me.Z = tfQ, tfQ, tfQ
	} else {
		tfQ = 1 / tfQ
		me.X, me.Y, me.Z, me.W = me.X * tfQ, me.Y * tfQ, me.Z * tfQ, me.W * tfQ
	}
}

func (me *TQuat) Normalized () *TQuat {
	tfQ = me.Magnitude()
	if tfQ == 0 { return &TQuat { tfQ, tfQ, tfQ, tfQ } }
	tfQ = 1 / tfQ
	return &TQuat { me.X * tfQ, me.Y * tfQ, me.Z * tfQ, me.W * tfQ }
}

func (me *TQuat) SetFromConjugated (c *TQuat) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

func (me *TQuat) SetFromMult (l, r *TQuat) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

func (me *TQuat) SetFromMult1 (v float64) {
	me.X, me.Y, me.Z, me.W = me.X * v, me.Y * v, me.Z * v, me.W * v
}

func (me *TQuat) SetFromMult3 (q *TQuat, v *TVec3) {
	me.W = - (q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

func (me *TQuat) SetFromMult4 (mat *TMat4) {
	*me = *mat.Mult4(me)
}
