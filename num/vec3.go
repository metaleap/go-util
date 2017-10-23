package unum

import (
	"math"

	"github.com/metaleap/go-util"
)

func Vec3_Back() Vec3 {
	return Vec3{0, 0, -1}
}

func Vec3_Down() Vec3 {
	return Vec3{0, -1, 0}
}

func Vec3_Fwd() Vec3 {
	return Vec3{0, 0, 1}
}

func Vec3_Left() Vec3 {
	return Vec3{-1, 0, 0}
}

func Vec3_One() Vec3 {
	return Vec3{1, 1, 1}
}

func Vec3_Right() Vec3 {
	return Vec3{1, 0, 0}
}

func Vec3_Up() Vec3 {
	return Vec3{0, 1, 0}
}

func Vec3_Zero() Vec3 {
	return Vec3{0, 0, 0}
}

func Vec3_Lerp(from, to *Vec3, t float64) *Vec3 {
	t = Clamp01(t)
	return &Vec3{t*(to.X-from.X) + from.X, t*(to.Y-from.Y) + from.Y, t*(to.Z-from.Z) + from.Z}
}

func Vec3_Max(l, r *Vec3) *Vec3 {
	return &Vec3{math.Max(l.X, r.X), math.Max(l.Y, r.Y), math.Max(l.Z, r.Z)}
}

func Vec3_Min(l, r *Vec3) *Vec3 {
	return &Vec3{math.Min(l.X, r.X), math.Min(l.Y, r.Y), math.Min(l.Z, r.Z)}
}

//	Represents a 3-dimensional vector.
type Vec3 struct {
	X, Y, Z float64
}

//	Adds `vec` to `me` in-place.
func (me *Vec3) Add(vec *Vec3) {
	me.X, me.Y, me.Z = me.X+vec.X, me.Y+vec.Y, me.Z+vec.Z
}

//	Returns the sum of `me` and `vec`.
func (me *Vec3) Added(vec *Vec3) *Vec3 {
	return &Vec3{me.X + vec.X, me.Y + vec.Y, me.Z + vec.Z}
}

//	Adds `val` to all 3 components of `me`.
func (me *Vec3) Add1(val float64) {
	me.X, me.Y, me.Z = me.X+val, me.Y+val, me.Z+val
}

//	Adds the specified 3 components to the respective components in `me`.
func (me *Vec3) Add3(x, y, z float64) {
	me.X, me.Y, me.Z = me.X+x, me.Y+y, me.Z+z
}

//	Returns whether all 3 components in `me` are approximately equivalent to their respective counterparts in `val`.
func (me *Vec3) AllEq(val float64) bool {
	return Eq(me.X, val) && Eq(me.Y, val) && Eq(me.Z, val)
}

//	Returns whether all 3 components in `me` are greater than (or approximately equivalent to) their respective component counterparts in `vec`.
func (me *Vec3) AllGEq(vec *Vec3) bool {
	return (me.X >= vec.X) && (me.Y >= vec.Y) && (me.Z >= vec.Z)
}

//	Returns whether all 3 components in `me` are greater than `min`, and also less than `max`.
func (me *Vec3) AllIn(min, max *Vec3) bool {
	return (me.X > min.X) && (me.X < max.X) && (me.Y > min.Y) && (me.Y < max.Y) && (me.Z > min.Z) && (me.Z < max.Z)
}

//	Returns whether all 3 components in `me` are less than (or approximately equivalent to) their respective component counterparts in `vec`.
func (me *Vec3) AllLEq(vec *Vec3) bool {
	return (me.X <= vec.X) && (me.Y <= vec.Y) && (me.Z <= vec.Z)
}

func (me *Vec3) AngleDeg(to *Vec3) float64 {
	return Rad2Deg * me.AngleRad(to)
}

func (me *Vec3) AngleRad(to *Vec3) float64 {
	return math.Acos(Clamp(me.Normalized().Dot(to.Normalized()), -1, 1))
}

//	Clamps each component in `me` between the respective corresponding counter-part component in `min` and `max`.
func (me *Vec3) Clamp(min, max *Vec3) {
	if me.X < min.X {
		me.X = min.X
	} else if me.X > max.X {
		me.X = max.X
	}
	if me.Y < min.Y {
		me.Y = min.Y
	} else if me.Y > max.Y {
		me.Y = max.Y
	}
	if me.Z < min.Z {
		me.Z = min.Z
	} else if me.Z > max.Z {
		me.Z = max.Z
	}
}

//	Clamps each component in `me` between 0 and 1.
func (me *Vec3) Clamp01() {
	me.X = Clamp01(me.X)
	me.Y = Clamp01(me.Y)
	me.Z = Clamp01(me.Z)
}

func (me *Vec3) ClampMagnitude(maxLength float64) *Vec3 {
	if l := me.Length(); l > maxLength*maxLength {
		return me.Scaled(maxLength * (1 / math.Sqrt(l)))
	}
	return me
}

//	Zeroes all 3 components in `me`.
func (me *Vec3) Clear() {
	me.X, me.Y, me.Z = 0, 0, 0
}

//	Returns a new `*Vec3` that represents the cross-product of `me` and `vec`.
func (me *Vec3) Cross(vec *Vec3) *Vec3 {
	return &Vec3{(me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X)}
}

//	Returns a new `*Vec` that represents the cross-product of `me` and `vec`, normalized.
func (me *Vec3) CrossNormalized(vec *Vec3) (r *Vec3) {
	r = me.Cross(vec)
	r.Normalize()
	return
}

//	Returns the distance of `me` from `vec`.
func (me *Vec3) Distance(vec *Vec3) float64 {
	return math.Sqrt(me.Sub(vec).Length())
}

//	Returns the "manhattan distance" of `me` from `vec`.
func (me *Vec3) DistanceManhattan(vec *Vec3) float64 {
	return math.Abs(vec.X-me.X) + math.Abs(vec.Y-me.Y) + math.Abs(vec.Z-me.Z)
}

//	Returns a new `*Vec3` that represents `me` divided by `vec`.
func (me *Vec3) Div(vec *Vec3) *Vec3 {
	return &Vec3{me.X / vec.X, me.Y / vec.Y, me.Z / vec.Z}
}

func (me *Vec3) Divide(d float64) {
	d = 1 / d
	me.X, me.Y, me.Z = me.X*d, me.Y*d, me.Z*d
}

//	Returns a new `*Vec3` that represents all 3 components in `me`, each divided by `val`.
func (me *Vec3) Divided(d float64) *Vec3 {
	d = 1 / d
	return &Vec3{me.X * d, me.Y * d, me.Z * d}
}

//	Returns the dot-product of `me` and `vec`.
func (me *Vec3) Dot(vec *Vec3) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y) + (me.Z * vec.Z)
}

//	Returns the dot-product of `me` and (`vec1` minus `vec2`).
func (me *Vec3) DotSub(vec1, vec2 *Vec3) float64 {
	return (me.X * (vec1.X - vec2.X)) + (me.Y * (vec1.Y - vec2.Y)) + (me.Z * (vec1.Z - vec2.Z))
}

func (me *Vec3) Eq(vec *Vec3) bool {
	return me.Sub(vec).Length() < EpsilonEqVec
}

//	Returns the 3D vector length of `me`.
func (me *Vec3) Length() float64 {
	return me.Dot(me)
}

//	Returns the 3D vector magnitude of `me`.
func (me *Vec3) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Returns the largest of the 3 components in `me`.
func (me *Vec3) Max() float64 {
	return math.Max(me.X, math.Max(me.Y, me.Z))
}

//	Returns the `math.Max` of the `math.Abs` values of all 3 components in `me`.
func (me *Vec3) MaxAbs() float64 {
	return math.Max(math.Abs(me.X), math.Max(math.Abs(me.Y), math.Abs(me.Z)))
}

//	Returns the smallest of the 3 components in `me`.
func (me *Vec3) Min() float64 {
	return math.Min(me.X, math.Min(me.Y, me.Z))
}

//	Returns a new `*Vec3` that represents `me` multiplied with `vec`.
func (me *Vec3) Mult(vec *Vec3) *Vec3 {
	return &Vec3{me.X * vec.X, me.Y * vec.Y, me.Z * vec.Z}
}

//	Returns a new `*Vec3` with each component in `me` multiplied by the respective corresponding specified factor.
func (me *Vec3) Mult3(x, y, z float64) *Vec3 {
	return &Vec3{me.X * x, me.Y * y, me.Z * z}
}

//	Reverses the signs of all 3 vector components in `me`.
func (me *Vec3) Negate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

//	Returns a new `*Vec` with each component representing the negative (sign inverted) corresponding component in `me`.
func (me *Vec3) Negated() *Vec3 {
	return &Vec3{-me.X, -me.Y, -me.Z}
}

//	Normalizes `me` in-place without checking for division-by-0.
func (me *Vec3) Normalize() {
	me.Divide(me.Magnitude())
}

//	Normalizes `me` in-place, safely checking for division-by-0.
func (me *Vec3) NormalizeSafe() {
	if mag := me.Magnitude(); mag > 0 {
		me.Divide(mag)
	} else {
		me.Clear()
	}
}

//	Returns a new `*Vec3` that represents `me`, normalized.
func (me *Vec3) Normalized() *Vec3 {
	return me.Divided(me.Magnitude())
}

//	Returns a new `*Vec3` that represents `me` normalized, then scaled by `factor`.
func (me *Vec3) NormalizedScaled(factor float64) (vec *Vec3) {
	return me.Normalized().Scaled(factor)
}

//	Returns a new `*Vec3` representing `1/me`.
func (me *Vec3) Rcp() *Vec3 {
	return &Vec3{1 / me.X, 1 / me.Y, 1 / me.Z}
}

//	Rotates `me` `angleDeg` degrees around the specified `axis`.
func (me *Vec3) RotateDeg(angleDeg float64, axis *Vec3) {
	me.RotateRad(DegToRad(angleDeg/2), axis)
}

//	Rotates `me` `angleRad` radians around the specified `axis`.
func (me *Vec3) RotateRad(angleRad float64, axis *Vec3) {
	var tmpQ, tmpQw, tmpQr, tmpQc Vec4
	sin, cos := math.Sincos(angleRad)
	tmpQr.X, tmpQr.Y, tmpQr.Z, tmpQr.W = axis.X*sin, axis.Y*sin, axis.Z*sin, cos
	tmpQc.SetFromConjugated(&tmpQr)
	tmpQ.SetFromMult3(&tmpQr, me)
	tmpQw.SetFromMult(&tmpQ, &tmpQc)
	me.X, me.Y, me.Z = tmpQw.X, tmpQw.Y, tmpQw.Z
}

//	Scales `me` by `factor`.
func (me *Vec3) Scale(factor float64) {
	me.X, me.Y, me.Z = me.X*factor, me.Y*factor, me.Z*factor
}

//	Scales `me` by `factor`, then adds `add`.
func (me *Vec3) ScaleAdd(factor, add *Vec3) {
	me.X, me.Y, me.Z = (me.X*factor.X)+add.X, (me.Y*factor.Y)+add.Y, (me.Z*factor.Z)+add.Z
}

//	Returns a new `*Vec3` that represents `me` scaled by `factor`.
func (me *Vec3) Scaled(factor float64) *Vec3 {
	return &Vec3{me.X * factor, me.Y * factor, me.Z * factor}
}

//	Returns a new `*Vec3` that represents `me` scaled by `factor`, then `add` added.
func (me *Vec3) ScaledAdded(factor float64, add *Vec3) *Vec3 {
	return &Vec3{(me.X * factor) + add.X, (me.Y * factor) + add.Y, (me.Z * factor) + add.Z}
}

//	Sets all 3 vector components in `me` to the corresponding respective specified value.
func (me *Vec3) Set(x, y, z float64) {
	me.X, me.Y, me.Z = x, y, z
}

//	Sets `me` to the result of adding `vec1` and `vec2`.
func (me *Vec3) SetFromAdd(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X+vec2.X, vec1.Y+vec2.Y, vec1.Z+vec2.Z
}

//	`me = a + b + c`
func (me *Vec3) SetFromAddAdd(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X+b.X+c.X, a.Y+b.Y+c.Y, a.Z+b.Z+c.Z
}

//	`me = mul * vec2 + vec1`
func (me *Vec3) SetFromAddScaled(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = mul*vec2.X+vec1.X, mul*vec2.Y+vec1.Y, mul*vec2.Z+vec1.Z
}

//	`me = a + b - c`
func (me *Vec3) SetFromAddSub(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X+b.X-c.X, a.Y+b.Y-c.Y, a.Z+b.Z-c.Z
}

//	Sets each vector component in `me` to the `math.Cos` of the respective corresponding component in `vec`.
func (me *Vec3) setFromCos(vec *Vec3) {
	me.X, me.Y, me.Z = math.Cos(vec.X), math.Cos(vec.Y), math.Cos(vec.Z)
}

//	Sets `me` to the cross-product of `me` and `vec`.
func (me *Vec3) SetFromCross(vec *Vec3) {
	me.X, me.Y, me.Z = (me.Y*vec.Z)-(me.Z*vec.Y), (me.Z*vec.X)-(me.X*vec.Z), (me.X*vec.Y)-(me.Y*vec.X)
}

//	Sets `me` to the cross-product of `one` and `two`.
func (me *Vec3) SetFromCrossOf(one, two *Vec3) {
	me.X, me.Y, me.Z = (one.Y*two.Z)-(one.Z*two.Y), (one.Z*two.X)-(one.X*two.Z), (one.X*two.Y)-(one.Y*two.X)
}

//	Sets each vector component in `me` to the radian equivalent of the degree angle stored in the respective corresponding component of `vec`.
func (me *Vec3) SetFromDegToRad(deg *Vec3) {
	me.X, me.Y, me.Z = DegToRad(deg.X), DegToRad(deg.Y), DegToRad(deg.Z)
}

//	`me = mul1 * mul2 + add`
func (me *Vec3) SetFromMad(mul1, mul2, add *Vec3) {
	me.X, me.Y, me.Z = mul1.X*mul2.X+add.X, mul1.Y*mul2.Y+add.Y, mul1.Z*mul2.Z+add.Z
}

func (me *Vec3) SetFromDivided(vec *Vec3, d float64) {
	d = 1 / d
	me.X, me.Y, me.Z = vec.X*d, vec.Y*d, vec.Z*d
}

//	`me = v1 * v2`
func (me *Vec3) SetFromMult(v1, v2 *Vec3) {
	me.X, me.Y, me.Z = v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z
}

//	`me = vec * mul`
func (me *Vec3) SetFromScaled(vec *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec.X*mul, vec.Y*mul, vec.Z*mul
}

//	`me = (vec1 - vec2) * mul`
func (me *Vec3) SetFromScaledSub(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = (vec1.X-vec2.X)*mul, (vec1.Y-vec2.Y)*mul, (vec1.Z-vec2.Z)*mul
}

//	`me = -vec`
func (me *Vec3) SetFromNegated(vec *Vec3) {
	me.X, me.Y, me.Z = -vec.X, -vec.Y, -vec.Z
}

//	Sets `me` to `vec` normalized.
func (me *Vec3) SetFromNormalized(vec *Vec3) {
	me.SetFromDivided(vec, vec.Magnitude())
}

//	Sets `me` to the inverse of `vec`.
func (me *Vec3) SetFromRcp(vec *Vec3) {
	me.X, me.Y, me.Z = 1/vec.X, 1/vec.Y, 1/vec.Z
}

//	Sets `me` to `pos` rotated as expressed in `rotCos` and `rotSin`.
func (me *Vec3) SetFromRotation(pos, rotCos, rotSin *Vec3) {
	tmpVal := ((pos.Y * rotSin.X) + (pos.Z * rotCos.X))
	me.X = (pos.X * rotCos.Y) + (tmpVal * rotSin.Y)
	me.Y = (pos.Y * rotCos.X) - (pos.Z * rotSin.X)
	me.Z = (-pos.X * rotSin.Y) + (tmpVal * rotCos.Y)
}

//	Sets each vector component in `me` to the `math.Sin` of the respective corresponding component in `vec`.
func (me *Vec3) setFromSin(vec *Vec3) {
	me.X, me.Y, me.Z = math.Sin(vec.X), math.Sin(vec.Y), math.Sin(vec.Z)
}

//	Component-wise, set `me` to `v0` if vec is less than `edge`, else `v1`.
func (me *Vec3) setFromStep(edge float64, vec, v0, v1 *Vec3) {
	me.X = umisc.IfF64(vec.X < edge, v0.X, v1.X)
	me.Y = umisc.IfF64(vec.Y < edge, v0.Y, v1.Y)
	me.Z = umisc.IfF64(vec.Z < edge, v0.Z, v1.Z)
}

//	`me = vec1 - vec2`.
func (me *Vec3) SetFromSub(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X-vec2.X, vec1.Y-vec2.Y, vec1.Z-vec2.Z
}

//	`me = a - b + c`
func (me *Vec3) SetFromSubAdd(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X-b.X+c.X, a.Y-b.Y+c.Y, a.Z-b.Z+c.Z
}

//	`me = v1 - v2 * v2Scale`
func (me *Vec3) SetFromSubScaled(v1, v2 *Vec3, v2Scale float64) {
	me.X, me.Y, me.Z = v1.X-v2.X*v2Scale, v1.Y-v2.Y*v2Scale, v1.Z-v2.Z*v2Scale
}

//	`me = a - b - c`
func (me *Vec3) SetFromSubSub(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X-b.X-c.X, a.Y-b.Y-c.Y, a.Z-b.Z-c.Z
}

//	`me = (sub1 - sub2) * mul`
func (me *Vec3) SetFromSubMult(sub1, sub2, mul *Vec3) {
	me.X, me.Y, me.Z = mul.X*(sub1.X-sub2.X), mul.Y*(sub1.Y-sub2.Y), mul.Z*(sub1.Z-sub2.Z)
}

//	Sets all 3 vector components in `me` to `math.MaxFloat64`.
func (me *Vec3) SetToMax() {
	me.X, me.Y, me.Z = math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
}

//	Sets all 3 vector components in `me` to `-math.MaxFloat64`.
func (me *Vec3) SetToMin() {
	me.X, me.Y, me.Z = -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
}

//	Returns a new `*Vec3` with each vector component indicating the sign (-1, 1 or 0) of the respective corresponding component in `me`.
func (me *Vec3) Sign() *Vec3 {
	return &Vec3{Sign(me.X), Sign(me.Y), Sign(me.Z)}
}

//	Returns a human-readable (imprecise) `string` representation of `me`.
func (me *Vec3) String() string {
	return strf("{X:%1.2f Y:%1.2f Z:%1.2f}", me.X, me.Y, me.Z)
}

//	Returns a new `*Vec3` that represents `me` minus `vec`.
func (me *Vec3) Sub(vec *Vec3) *Vec3 {
	return &Vec3{me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z}
}

//	Returns a new `*Vec3` that represents `((me - sub) / div) * mul`.
func (me *Vec3) SubDivMult(sub, div, mul *Vec3) *Vec3 {
	return &Vec3{mul.X * ((me.X - sub.X) / div.X), mul.Y * ((me.Y - sub.Y) / div.Y), mul.Z * ((me.Z - sub.Z) / div.Z)}
}

//	Returns a new `*Vec3` that represents `mul * math.Floor(me / div)`.
func (me *Vec3) SubFloorDivMult(div, mul float64) *Vec3 {
	div = 1 / div
	return me.Sub(&Vec3{mul * math.Floor(me.X*div), mul * math.Floor(me.Y*div), mul * math.Floor(me.Z*div)})
}

//	Returns a new `*Vec3` that represents `val` minus `me`.
func (me *Vec3) SubFrom(val float64) *Vec3 {
	return &Vec3{val - me.X, val - me.Y, val - me.Z}
}

//	Returns a new `*Vec3` that represents `(me - vec) * val`.
func (me *Vec3) SubScaled(vec *Vec3, val float64) *Vec3 {
	return &Vec3{val * (me.X - vec.X), val * (me.Y - vec.Y), val * (me.Z - vec.Z)}
}

//	Subtracts `vec` from `me`.
func (me *Vec3) Subtract(vec *Vec3) {
	me.X, me.Y, me.Z = me.X-vec.X, me.Y-vec.Y, me.Z-vec.Z
}

//	Transform coordinate vector `me` according to the specified `*Mat4`.
func (me *Vec3) TransformCoord(mat *Mat4) {
	var q Vec4
	q.MultMat4Vec3(mat, me)
	q.W = 1 / q.W
	me.X, me.Y, me.Z = q.X*q.W, q.Y*q.W, q.Z*q.W
}

//	Transform normal vector `me` according to the specified `*Mat4`.
func (me *Vec3) TransformNormal(mat *Mat4, absMat bool) {
	m11, m21, m31 := mat[0], mat[1], mat[2]
	m12, m22, m32 := mat[4], mat[5], mat[6]
	m13, m23, m33 := mat[8], mat[9], mat[10]
	if absMat {
		m11, m21, m31 = math.Abs(m11), math.Abs(m21), math.Abs(m31)
		m12, m22, m32 = math.Abs(m12), math.Abs(m22), math.Abs(m32)
		m13, m23, m33 = math.Abs(m13), math.Abs(m23), math.Abs(m33)
	}
	x := ((me.X * m11) + (me.Y * m21)) + (me.Z * m31)
	y := ((me.X * m12) + (me.Y * m22)) + (me.Z * m32)
	z := ((me.X * m13) + (me.Y * m23)) + (me.Z * m33)
	me.X, me.Y, me.Z = x, y, z
}
