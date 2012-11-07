package num

import (
	"math"
)

type Quat struct {
	X, Y, Z, W float64
}

func (me *Quat) Conjugate () {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

func (me *Quat) Conjugated () *Quat {
	return &Quat { -me.X, -me.Y, -me.Z, me.W }
}

func (me *Quat) Length () float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z) + (me.W * me.W)
}

func (me *Quat) Magnitude () float64 {
	return math.Sqrt(me.Length())
}

func (me *Quat) Normalize () {
	tfQ = me.Magnitude()
	if tfQ == 0 {
		me.X, me.Y, me.Z = tfQ, tfQ, tfQ
	} else {
		tfQ = 1 / tfQ
		me.X, me.Y, me.Z, me.W = me.X * tfQ, me.Y * tfQ, me.Z * tfQ, me.W * tfQ
	}
}

func (me *Quat) Normalized () *Quat {
	tfQ = me.Magnitude()
	if tfQ == 0 { return &Quat { tfQ, tfQ, tfQ, tfQ } }
	tfQ = 1 / tfQ
	return &Quat { me.X * tfQ, me.Y * tfQ, me.Z * tfQ, me.W * tfQ }
}

func (me *Quat) SetFromConjugated (c *Quat) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

func (me *Quat) SetFromMult (l, r *Quat) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

func (me *Quat) SetFromMult1 (v float64) {
	me.X, me.Y, me.Z, me.W = me.X * v, me.Y * v, me.Z * v, me.W * v
}

func (me *Quat) SetFromMult3 (q *Quat, v *Vec3) {
	me.W = - (q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

func (me *Quat) SetFromMult4 (mat *Mat4) {
	*me = *mat.Mult4(me)
}
