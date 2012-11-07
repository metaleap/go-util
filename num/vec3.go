package num

import (
	"fmt"
	"math"
)

var (
	tmpVal, tmpCos, tmpSin float64
	tmpQ, tmpQw, tmpQr, tmpQc = &Quat {}, &Quat {}, &Quat {}, &Quat {}
)

type Vec3 struct {
	X, Y, Z float64
}

func (me *Vec3) Add (vec *Vec3) {
	me.X, me.Y, me.Z = me.X + vec.X, me.Y + vec.Y, me.Z + vec.Z
}

func (me *Vec3) Add1 (add float64)  {
	me.X, me.Y, me.Z = me.X + add, me.Y + add, me.Z + add
}

func (me *Vec3) AllEqual (val float64) bool {
	return (me.X == val) && (me.Y == val) && (me.Z == val)
}

func (me *Vec3) AllGreaterOrEqual (test *Vec3) bool {
	return (me.X >= test.X) && (me.Y >= test.Y) && (me.Z >= test.Z)
}

func (me *Vec3) AllInRange (min, max float64) bool {
	return (me.X >= min) && (me.X < max) && (me.Y >= min) && (me.Y < max) && (me.Z >= min) && (me.Z < max)
}

func (me *Vec3) AllInside (min, max *Vec3) bool {
	return (me.X > min.X) && (me.X < max.X) && (me.Y > min.Y) && (me.Y < max.Y) && (me.Z > min.Z) && (me.Z < max.Z)
}

func (me *Vec3) AllLessOrEqual (test *Vec3) bool {
	return (me.X <= test.X) && (me.Y <= test.Y) && (me.Z <= test.Z)
}

func (me *Vec3) Cross (vec *Vec3) *Vec3 {
	return &Vec3 { (me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X) }
}

func (me *Vec3) CrossNormalized (vec *Vec3) *Vec3 {
	var r = &Vec3 { (me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X) }
	r.Normalize()
	return r
}

func (me *Vec3) Div (vec *Vec3) *Vec3 {
	return &Vec3 { me.X / vec.X, me.Y / vec.Y, me.Z / vec.Z }
}

func (me *Vec3) Div1 (val float64) *Vec3 {
	return &Vec3 { me.X / val, me.Y / val, me.Z / val }
}

func (me *Vec3) Dot (vec *Vec3) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y) + (me.Z * vec.Z)
}

func (me *Vec3) DotSub (vec1, vec2 *Vec3) float64 {
	return (me.X * (vec1.X - vec2.X)) + (me.Y * (vec1.Y - vec2.Y)) + (me.Z * (vec1.Z - vec2.Z))
}

func (me *Vec3) Equals (vec *Vec3) bool {
	return (me.X == vec.X) && (me.Y == vec.Y) && (me.Z == vec.Z)
}

func (me *Vec3) Inv () *Vec3 {
	return &Vec3 { 1 / me.X, 1 / me.Y, 1 / me.Z }
}

func (me *Vec3) Length () float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z)
}

func (me *Vec3) Magnitude () float64 {
	return math.Sqrt(me.Length())
}

func (me *Vec3) MakeFinite (v *Vec3) {
	if math.IsInf(me.X, 0) { me.X = v.X }
	if math.IsInf(me.Y, 0) { me.Y = v.Y }
	if math.IsInf(me.Z, 0) { me.Z = v.Z }
}

func (me *Vec3) Max () float64 {
	return math.Max(me.X, math.Max(me.Y, me.Z))
}

func (me *Vec3) Min () float64 {
	return math.Min(me.X, math.Min(me.Y, me.Z))
}

func (me *Vec3) Mult (vec *Vec3) *Vec3 {
	return &Vec3 { me.X * vec.X, me.Y * vec.Y, me.Z * vec.Z }
}

func (me *Vec3) Mult1 (val float64) *Vec3 {
	return &Vec3 { me.X * val, me.Y * val, me.Z * val }
}

func (me *Vec3) Normalize () {
	tmpVal = me.Magnitude()
	if tmpVal == 0 {
		me.X, me.Y, me.Z = tmpVal, tmpVal, tmpVal
	} else {
		tmpVal = 1 / tmpVal
		me.X, me.Y, me.Z = me.X * tmpVal, me.Y * tmpVal, me.Z * tmpVal
	}
}

func (me *Vec3) Normalized () *Vec3 {
	tmpVal = me.Magnitude()
	if tmpVal == 0 { return &Vec3 { tmpVal, tmpVal, tmpVal } }
	tmpVal = 1 / tmpVal
	return &Vec3 { me.X * tmpVal, me.Y * tmpVal, me.Z * tmpVal }
}

func (me *Vec3) NormalizedScaled (mul float64) *Vec3 {
	var vec = me.Normalized(); vec.Scale(mul); return vec
}

func (me *Vec3) RotateDeg (angleDeg float64, axis *Vec3) {
	tmpCos = math.Cos(DegToRad(angleDeg / 2))
	tmpSin = math.Sin(DegToRad(angleDeg / 2))
	tmpQr.X, tmpQr.Y, tmpQr.Z, tmpQr.W = axis.X * tmpSin, axis.Y * tmpSin, axis.Z * tmpSin, tmpCos
	tmpQc.SetFromConjugated(tmpQr)
	tmpQ.SetFromMult3(tmpQr, me)
	tmpQw.SetFromMult(tmpQ, tmpQc)
	me.X, me.Y, me.Z = tmpQw.X, tmpQw.Y, tmpQw.Z
}

func (me *Vec3) Scale (mul float64) {
	me.X, me.Y, me.Z = me.X * mul, me.Y * mul, me.Z * mul
}

func (me *Vec3) ScaleAdd (mul, add *Vec3) {
	me.X, me.Y, me.Z = (me.X * mul.X) + add.X, (me.Y * mul.Y) + add.Y, (me.Z * mul.Z) + add.Z
}

func (me *Vec3) Scaled (by float64) *Vec3 {
	return &Vec3 { me.X * by, me.Y * by, me.Z * by }
}

func (me *Vec3) ScaledAdded (mul float64, add *Vec3) *Vec3 {
	return &Vec3 { (me.X * mul) + add.X, (me.Y * mul) + add.Y, (me.Z * mul) + add.Z }
}

func (me *Vec3) Set (x, y, z float64) {
	me.X, me.Y, me.Z = x, y, z
}

func (me *Vec3) SetFrom (vec *Vec3) {
	me.X, me.Y, me.Z = vec.X, vec.Y, vec.Z
}

func (me *Vec3) SetFromAdd (vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X + vec2.X, vec1.Y + vec2.Y, vec1.Z + vec2.Z
}

func (me *Vec3) SetFromAddMult (add, mul1, mul2 *Vec3) {
	me.X, me.Y, me.Z = add.X + (mul1.X * mul2.X), add.Y + (mul1.Y * mul2.Y), add.Z + (mul1.Z * mul2.Z)
}

func (me *Vec3) SetFromAddMult1 (vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec1.X + (vec2.X * mul), vec1.Y + (vec2.Y * mul), vec1.Z + (vec2.Z * mul)
}

func (me *Vec3) SetFromCos (vec *Vec3) {
	me.X, me.Y, me.Z = math.Cos(vec.X), math.Cos(vec.Y), math.Cos(vec.Z)
}

func (me *Vec3) SetFromDegToRad (deg *Vec3) {
	me.X, me.Y, me.Z = DegToRad(deg.X), DegToRad(deg.Y), DegToRad(deg.Z)
}

func (me *Vec3) SetFromEpsilon () {
	if math.Abs(me.X) < Epsilon { me.X = Epsilon }
	if math.Abs(me.Y) < Epsilon { me.Y = Epsilon }
	if math.Abs(me.Z) < Epsilon { me.Z = Epsilon }
}

func (me *Vec3) SetFromInv (vec *Vec3) {
	me.X, me.Y, me.Z = 1 / vec.X, 1 / vec.Y, 1 / vec.Z
}

func (me *Vec3) SetFromMult (v1, v2 *Vec3) {
	me.X, me.Y, me.Z = v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z
}

func (me *Vec3) SetFromMult1 (vec *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec.X * mul, vec.Y * mul, vec.Z * mul
}

func (me *Vec3) SetFromMult1Sub (vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = (vec1.X - vec2.X) * mul, (vec1.Y - vec2.Y) * mul, (vec1.Z - vec2.Z) * mul
}

func (me *Vec3) SetFromNeg (vec *Vec3) {
	me.X, me.Y, me.Z = -vec.X, -vec.Y, -vec.Z
}

func (me *Vec3) SetFromRotation (pos, rotCos, rotSin *Vec3) {
	tmpVal = ((pos.Y * rotSin.X) + (pos.Z * rotCos.X))
	me.X = (pos.X * rotCos.Y) + (tmpVal * rotSin.Y)
	me.Y = (pos.Y * rotCos.X) - (pos.Z * rotSin.X)
	me.Z = (-pos.X * rotSin.Y) + (tmpVal * rotCos.Y)
}

func (me *Vec3) SetFromSin (vec *Vec3) {
	me.X, me.Y, me.Z = math.Sin(vec.X), math.Sin(vec.Y), math.Sin(vec.Z)
}

func (me *Vec3) SetFromStep1 (edge float64, vec, zero, one *Vec3) {
	if vec.X < edge { me.X = zero.X } else { me.X = one.X }
	if vec.Y < edge { me.Y = zero.Y } else { me.Y = one.Y }
	if vec.Z < edge { me.Z = zero.Z } else { me.Z = one.Z }
}

func (me *Vec3) SetFromSub (vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X - vec2.X, vec1.Y - vec2.Y, vec1.Z - vec2.Z
}

func (me *Vec3) SetFromSubMult (sub1, sub2, mul *Vec3) {
	me.X, me.Y, me.Z = (sub1.X - sub2.X) * mul.X, (sub1.Y - sub2.Y) * mul.Y, (sub1.Z - sub2.Z) * mul.Z
}

func (me *Vec3) SetFromSubMult1 (vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec1.X - (vec2.X * mul), vec1.Y - (vec2.Y * mul), vec1.Z - (vec2.Z * mul)
}

func (me *Vec3) Sign () *Vec3 {
	return &Vec3 { Sign(me.X), Sign(me.Y), Sign(me.Z) }
}

func (me *Vec3) String () string {
	return fmt.Sprintf("{X: %6.2f Y:%6.2f Z:%6.2f}", me.X, me.Y, me.Z)
}

func (me *Vec3) Sub (vec *Vec3) *Vec3 {
	return &Vec3 { me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z }
}

func (me *Vec3) SubDivMult (sub, div, mul *Vec3) *Vec3 {
	return &Vec3 { ((me.X - sub.X) / div.X) * mul.X, ((me.Y - sub.Y) / div.Y) * mul.Y, ((me.Z - sub.Z) / div.Z) * mul.Z }
}

func (me *Vec3) SubDot (vec *Vec3) float64 {
	return ((me.X - vec.X) * (me.X - vec.X)) + ((me.Y - vec.Y) * (me.Y - vec.Y)) + ((me.Z - vec.Z) * (me.Z - vec.Z))
}

func (me *Vec3) SubFloorDivMult (floorDiv, mul float64) *Vec3 {
	return me.Sub(&Vec3 { math.Floor(me.X / floorDiv) * mul, math.Floor(me.Y / floorDiv) * mul, math.Floor(me.Z / floorDiv) * mul })
}

func (me *Vec3) SubFrom (val float64) *Vec3 {
	return &Vec3 { val - me.X, val - me.Y, val - me.Z }
}

func (me *Vec3) SubMult (vec *Vec3, val float64) *Vec3 {
	return &Vec3 { (me.X - vec.X) * val, (me.Y - vec.Y) * val, (me.Z - vec.Z) * val }
}

func (me *Vec3) SubVec (vec *Vec3) {
	me.X, me.Y, me.Z = me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z
}

func (me *Vec3) SwapSigns () {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

func (me *Vec3) ToDegInts () []int {
	return []int { int(RadToDeg(me.X)), int(RadToDeg(me.Y)), int(RadToDeg(me.Z)) }
}

func (me *Vec3) ToInts () []int {
	return []int { int(me.X), int(me.Y), int(me.Z) }
}
