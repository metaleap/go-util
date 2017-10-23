package unum

import (
	"math"
)

func NewQuat(x, y, z, w float64) *Quat {
	var q Quat
	q.X, q.Y, q.Z, q.W = x, y, z, w
	return &q
}

func Quat_Identity() (q Quat) {
	q.Vec4.W = 1
	return
}

//	Quaternion
type Quat struct {
	//	X, Y, Z, W
	Vec4
}

func (me *Quat) AngleDeg(q *Quat) float64 {
	return Rad2Deg * me.AngleRad(q)
}

func (me *Quat) AngleRad(q *Quat) float64 {
	return 2 * math.Acos(math.Min(1, math.Abs(me.Dot(&q.Vec4))))
}

func (me *Quat) Eq(vec *Vec4) bool {
	return me.Dot(vec) > 0.999999
}

func (me *Quat) Mul(q *Quat) *Quat {
	return NewQuat(me.W*q.X+me.X*q.W+me.Y*q.Z-me.Z*q.Y, me.W*q.Y+me.Y*q.W+me.Z*q.X-me.X*q.Z, me.W*q.Z+me.Z*q.W+me.X*q.Y-me.Y*q.X, me.W*q.W-me.X*q.X-me.Y*q.Y-me.Z*q.Z)
}

func (me *Quat) MulVec3(p *Vec3) *Vec3 {
	r := Vec3{me.X * 2, me.Y * 2, me.Z * 2}
	mr := me.MulVec3(&r)
	mrc := Vec3{me.X * r.Y, me.X * r.Z, me.Y * r.Z}
	mrw := r.Scaled(me.W)
	r.X = (1-(mr.Y+mr.Z))*p.X + (mrc.X-mrw.Z)*p.Y + (mrc.Y+mrw.Y)*p.Z
	r.Y = (mrc.X+mrw.Z)*p.X + (1-(mr.X+mr.Z))*p.Y + (mrc.Z-mrw.X)*p.Z
	r.Z = (mrc.Y-mrw.Y)*p.X + (mrc.Z+mrw.X)*p.Y + (1-(mr.X+mr.Y))*p.Z
	return &r
}
