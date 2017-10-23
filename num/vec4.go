package unum

import (
	"math"
)

func Vec4_One() Vec4 {
	return Vec4{1, 1, 1, 1}
}

func Vec4_Zero() Vec4 {
	return Vec4{0, 0, 0, 0}
}

func Vec4_Lerp(from, to *Vec4, t float64) *Vec4 {
	t = Clamp01(t)
	return &Vec4{t*(to.X-from.X) + from.X, t*(to.Y-from.Y) + from.Y, t*(to.Z-from.Z) + from.Z, t*(to.W-from.W) + from.W}
}

func Vec4_Max(l, r *Vec4) *Vec4 {
	return &Vec4{math.Max(l.X, r.X), math.Max(l.Y, r.Y), math.Max(l.Z, r.Z), math.Max(l.W, r.W)}
}

func Vec4_Min(l, r *Vec4) *Vec4 {
	return &Vec4{math.Min(l.X, r.X), math.Min(l.Y, r.Y), math.Min(l.Z, r.Z), math.Min(l.W, r.W)}
}

//	Represents an arbitrary 4-dimensional vector.
type Vec4 struct {
	X, Y, Z, W float64
}

func (me *Vec4) AddedDiv(a *Vec4, d float64) *Vec4 {
	d = 1 / d
	return &Vec4{a.X*d + me.X, a.Y*d + me.Y, a.Z*d + me.Z, a.W*d + me.W}
}

func (me *Vec4) Clear() {
	me.X, me.Y, me.Z, me.W = 0, 0, 0, 0
}

//	Returns a new `*Vec4` containing a copy of `me`.
func (me *Vec4) Clone() (q *Vec4) {
	*q = *me
	return
}

//	Negates the `X`, `Y`, `Z` components in `me`, but not `W`.
func (me *Vec4) Conjugate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

//	Returns a new `*Vec4` that represents `me` conjugated.
func (me *Vec4) Conjugated() (v *Vec4) {
	v = new(Vec4)
	v.X, v.Y, v.Z, v.W = -me.X, -me.Y, -me.Z, me.W
	return
}

func (me *Vec4) Distance(vec *Vec4) float64 {
	return me.Sub(vec).Magnitude()
}

func (me *Vec4) Divide(d float64) {
	d = 1 / d
	me.X, me.Y, me.Z, me.W = me.X*d, me.Y*d, me.Z*d, me.W*d
}

func (me *Vec4) Divided(d float64) *Vec4 {
	d = 1 / d
	return &Vec4{me.X * d, me.Y * d, me.Z * d, me.W * d}
}

func (me *Vec4) Dot(vec *Vec4) float64 {
	return me.X*vec.X + me.Y*vec.Y + me.Z*vec.Z + me.W*vec.W
}

func (me *Vec4) Eq(vec *Vec4) bool {
	return me.Sub(vec).Length() < EpsilonEqVec
}

//	Returns the 4D vector length of `me`.
func (me *Vec4) Length() float64 {
	return me.Dot(me)
}

//	Returns the 4D vector magnitude of `me`.
func (me *Vec4) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

func (me *Vec4) MoveTowards(target *Vec4, maxDistanceDelta float64) *Vec4 {
	a := target.Sub(me)
	m := a.Magnitude()
	if m <= maxDistanceDelta || m == 0 {
		return target
	}
	return me.AddedDiv(a, m*maxDistanceDelta)
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with `me`.
func (me *Vec4) MultMat4(mat *Mat4) {
	me.MultMat4Vec4(mat, me.Clone())
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with the specified `*Vec3`.
func (me *Vec4) MultMat4Vec3(mat *Mat4, vec *Vec3) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * 1)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * 1)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * 1)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * 1)
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with the specified `*Vec4`.
func (me *Vec4) MultMat4Vec4(mat *Mat4, vec *Vec4) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * vec.W)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * vec.W)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * vec.W)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * vec.W)
}

func (me *Vec4) Negate() {
	me.X, me.Y, me.Z, me.W = -me.X, -me.Y, -me.Z, -me.W
}

func (me *Vec4) Negated() *Vec4 {
	return &Vec4{-me.X, -me.Y, -me.Z, -me.W}
}

//	Normalizes `me` according to `me.Magnitude`.
func (me *Vec4) Normalize() {
	me.NormalizeFrom(me.Magnitude())
}

//	Normalizes `me` according to the specified `magnitude`.
func (me *Vec4) NormalizeFrom(magnitude float64) {
	if magnitude > 0 {
		me.Divide(magnitude)
	} else {
		me.Clear()
	}
}

//	Returns a new `*Vec4` that represents `me` normalized according to `me.Magnitude`.
func (me *Vec4) Normalized() *Vec4 {
	if mag := me.Magnitude(); mag > 0 {
		return me.Divided(mag)
	} else {
		return &Vec4{0, 0, 0, 0}
	}
}

func (me *Vec4) Project(vec *Vec4) {
	me.Scale(me.Dot(vec) / vec.Length())
}

func (me *Vec4) Projected(vec *Vec4) *Vec4 {
	return vec.Scaled(me.Dot(vec) / vec.Length())
}

//	Sets `me` to `c` conjugated.
func (me *Vec4) SetFromConjugated(c *Vec4) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

//	Applies various 4D vector component computations of `l` and `r` to `me`, as needed by the `Vec3.RotateRad` method.
func (me *Vec4) SetFromMult(l, r *Vec4) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

//	Scales all 4 vector components in `me` by factor `v`.
func (me *Vec4) Scale(v float64) {
	me.X, me.Y, me.Z, me.W = me.X*v, me.Y*v, me.Z*v, me.W*v
}

func (me *Vec4) Scaled(v float64) *Vec4 {
	return &Vec4{me.X * v, me.Y * v, me.Z * v, me.W * v}
}

//	Applies various 4D vector component computations of `q` and `v` to `me`, as needed by the `Vec3.RotateRad` method.
func (me *Vec4) SetFromMult3(q *Vec4, v *Vec3) {
	me.W = -(q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

func (me *Vec4) SetFromVec3(vec *Vec3) {
	me.X, me.Y, me.Z = vec.X, vec.Y, vec.Z
}

//	Returns a human-readable (imprecise) `string` representation of `me`.
func (me *Vec4) String() string {
	return strf("{X:%1.2f Y:%1.2f Z:%1.2f W:%1.2f}", me.X, me.Y, me.Z, me.W)
}

func (me *Vec4) Sub(vec *Vec4) *Vec4 {
	return &Vec4{me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z, me.W - vec.W}
}

func (me *Vec4) Subtract(vec *Vec4) {
	me.X, me.Y, me.Z, me.W = me.X-vec.X, me.Y-vec.Y, me.Z-vec.Z, me.W-vec.W
}
