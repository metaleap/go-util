package unum

import (
	"math"
)

//	Vec2{1, 1}
func Vec2_One() Vec2 {
	return Vec2{1, 1}
}

//	Vec2{1, 0}
func Vec2_Right() Vec2 {
	return Vec2{1, 0}
}

//	Vec2{0, 1}
func Vec2_Up() Vec2 {
	return Vec2{0, 1}
}

//	Vec2{0, 0}
func Vec2_Zero() Vec2 {
	return Vec2{0, 0}
}

func Vec2_Lerp(from, to *Vec2, t float64) *Vec2 {
	t = Clamp01(t)
	return &Vec2{t*(to.X-from.X) + from.X, t*(to.Y-from.Y) + from.Y}
}

func Vec2_Max(l, r *Vec2) *Vec2 {
	return &Vec2{math.Max(l.X, r.X), math.Max(l.Y, r.Y)}
}

func Vec2_Min(l, r *Vec2) *Vec2 {
	return &Vec2{math.Min(l.X, r.X), math.Min(l.Y, r.Y)}
}

//	A 2-dimensional vector.
type Vec2 struct{ X, Y float64 }

func (me *Vec2) Add(vec *Vec2) {
	me.X, me.Y = me.X+vec.X, me.Y+vec.Y
}

func (me *Vec2) AddedDiv(a *Vec2, d float64) *Vec2 {
	d = 1 / d
	return &Vec2{a.X*d + me.X, a.Y*d + me.Y}
}

func (me *Vec2) AngleDeg(to *Vec2) float64 {
	return Rad2Deg * me.AngleRad(to)
}

func (me *Vec2) AngleRad(to *Vec2) float64 {
	return math.Acos(Clamp(me.Normalized().Dot(to.Normalized()), -1, 1))
}

func (me *Vec2) ClampMagnitude(maxLength float64) *Vec2 {
	if l := me.Length(); l > maxLength*maxLength {
		return me.Scaled(maxLength * (1 / math.Sqrt(l)))
	}
	return me
}

func (me *Vec2) Clear() {
	me.X, me.Y = 0, 0
}

func (me *Vec2) Distance(vec *Vec2) float64 {
	return me.Sub(vec).Magnitude()
}

//	Returns a new `*Vec2` that is the result of dividing `me` by `vec` without checking for division-by-0.
func (me *Vec2) Div(vec *Vec2) *Vec2 {
	return &Vec2{me.X / vec.X, me.Y / vec.Y}
}

func (me *Vec2) Divide(d float64) {
	d = 1 / d
	me.X, me.Y = me.X*d, me.Y*d
}

func (me *Vec2) Divided(d float64) *Vec2 {
	d = 1 / d
	return &Vec2{me.X * d, me.Y * d}
}

//	Returns a new `*Vec2` that is the result of dividing `me` by `vec`, safely checking for division-by-0.
func (me *Vec2) DivSafe(vec *Vec2) *Vec2 {
	r := Vec2{}
	if vec.X != 0 {
		r.X = me.X / vec.X
	}
	if vec.Y != 0 {
		r.Y = me.Y / vec.Y
	}
	return &r
}

//	Returns the dot product of `me` and `vec`.
func (me *Vec2) Dot(vec *Vec2) float64 {
	return me.X*vec.X + me.Y*vec.Y
}

func (me *Vec2) Eq(vec *Vec2) bool {
	return me.Sub(vec).Length() < EpsilonEqVec
}

//	Returns the 2D vector length of `me`.
func (me *Vec2) Length() float64 {
	return me.Dot(me)
}

//	Returns the 2D vector magnitude of `me`.
func (me *Vec2) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

func (me *Vec2) MoveTowards(target *Vec2, maxDistanceDelta float64) *Vec2 {
	a := target.Sub(me)
	m := a.Magnitude()
	if m <= maxDistanceDelta || m == 0 {
		return target
	}
	return me.AddedDiv(a, m*maxDistanceDelta)
}

//	Returns a new `*Vec2` that is the result of multiplying `me` with `vec`.
func (me *Vec2) Mult(vec *Vec2) *Vec2 {
	return &Vec2{me.X * vec.X, me.Y * vec.Y}
}

func (me *Vec2) Negate() *Vec2 {
	return &Vec2{-me.X, -me.Y}
}

//	Normalizes `me` in-place without checking for division-by-0.
func (me *Vec2) Normalize() {
	me.Divide(me.Magnitude())
}

//	Normalizes `me` in-place, safely checking for division-by-0.
func (me *Vec2) NormalizeSafe() {
	if mag := me.Magnitude(); mag > 0 {
		me.Divide(mag)
	} else {
		me.Clear()
	}
}

//	Returns a new `*Vec2` that is the normalized representation of `me` without checking for division-by-0.
func (me *Vec2) Normalized() *Vec2 {
	return me.Divided(me.Magnitude())
}

//	Returns a new `*Vec2` that is the normalized representation of `me`, safely checking for division-by-0.
func (me *Vec2) NormalizedSafe() *Vec2 {
	if mag := me.Magnitude(); mag > 0 {
		return me.Divided(mag)
	}
	return &Vec2{0, 0}
}

//	Returns a new `*Vec2` that is the normalized representation of `me` scaled by `factor` without checking for division-by-0.
func (me *Vec2) NormalizedScaled(factor float64) *Vec2 {
	return me.Normalized().Scaled(factor)
}

//	Returns a new `*Vec2` that is the normalized representation of `me` scaled by `factor`, safely checking for division-by-0.
func (me *Vec2) NormalizedScaledSafe(factor float64) *Vec2 {
	return me.NormalizedSafe().Scaled(factor)
}

//	Multiplies all components in `me` with `factor`.
func (me *Vec2) Scale(factor float64) {
	me.X, me.Y = me.X*factor, me.Y*factor
}

//	Returns a new `*Vec2` that represents `me` scaled by `factor`.
func (me *Vec2) Scaled(factor float64) *Vec2 {
	return &Vec2{me.X * factor, me.Y * factor}
}

func (me *Vec2) Set(x, y float64) {
	me.X, me.Y = x, y
}

//	Returns a human-readable (imprecise) `string` representation of `me`.
func (me *Vec2) String() string {
	return strf("{X:%1.2f Y:%1.2f}", me.X, me.Y)
}

//	Returns a new `*Vec2` that represents `me` minus `vec`.
func (me *Vec2) Sub(vec *Vec2) *Vec2 {
	return &Vec2{me.X - vec.X, me.Y - vec.Y}
}

//	Subtracts `vec` from `me`.
func (me *Vec2) Subtract(vec *Vec2) {
	me.X, me.Y = me.X-vec.X, me.Y-vec.Y
}
