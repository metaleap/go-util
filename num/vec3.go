package num

import (
	"fmt"
	"math"
)

var (
	tmpQ, tmpQw, tmpQr, tmpQc = Quat{}, Quat{}, Quat{}, Quat{}
)

//	Represented a 3-dimensional vector.
type Vec3 struct {
	X, Y, Z float64
}

//	Adds vec to this 3D vector.
func (me *Vec3) Add(vec *Vec3) {
	me.X, me.Y, me.Z = me.X+vec.X, me.Y+vec.Y, me.Z+vec.Z
}

//	Adds val to all components of this 3D vector.
func (me *Vec3) Add1(val float64) {
	me.X, me.Y, me.Z = me.X+val, me.Y+val, me.Z+val
}

//	Returns true if all components of this 3D vector equal val.
func (me *Vec3) AllEqual(val float64) bool {
	return (me.X == val) && (me.Y == val) && (me.Z == val)
}

//	Returns true if this 3D vector is greater than or equal to vec.
func (me *Vec3) AllGreaterOrEqual(vec *Vec3) bool {
	return (me.X >= vec.X) && (me.Y >= vec.Y) && (me.Z >= vec.Z)
}

//	Returns true if this 3D vector is greater than or equal to min, and less than max.
func (me *Vec3) AllInRange(min, max float64) bool {
	return (me.X >= min) && (me.X < max) && (me.Y >= min) && (me.Y < max) && (me.Z >= min) && (me.Z < max)
}

//	Returns true if this 3D vector is greater than min and less than max.
func (me *Vec3) AllInside(min, max *Vec3) bool {
	return (me.X > min.X) && (me.X < max.X) && (me.Y > min.Y) && (me.Y < max.Y) && (me.Z > min.Z) && (me.Z < max.Z)
}

//	Returns true if this 3D vector is less than or equal to vec.
func (me *Vec3) AllLessOrEqual(vec *Vec3) bool {
	return (me.X <= vec.X) && (me.Y <= vec.Y) && (me.Z <= vec.Z)
}

//	Returns a new 3D vector that represents the cross-product of this 3D vector and vec.
func (me *Vec3) Cross(vec *Vec3) *Vec3 {
	return &Vec3{(me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X)}
}

//	Returns a new 3D vector that represents the cross-product of this 3D vector and vec, normalized.
func (me *Vec3) CrossNormalized(vec *Vec3) (r *Vec3) {
	r = &Vec3{(me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X)}
	r.Normalize()
	return
}

//	Returns a new 3D vector that represents this 3D vector divided by vec.
func (me *Vec3) Div(vec *Vec3) *Vec3 {
	return &Vec3{me.X / vec.X, me.Y / vec.Y, me.Z / vec.Z}
}

//	Returns a new 3D vector that represents this 3D vector's components all divided by val.
func (me *Vec3) Div1(val float64) *Vec3 {
	return &Vec3{me.X / val, me.Y / val, me.Z / val}
}

//	Returns the dot product of this 3D vector and vec.
func (me *Vec3) Dot(vec *Vec3) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y) + (me.Z * vec.Z)
}

//	Returns the dot product of this 3D vector and (vec1 minus vec2).
func (me *Vec3) DotSub(vec1, vec2 *Vec3) float64 {
	return (me.X * (vec1.X - vec2.X)) + (me.Y * (vec1.Y - vec2.Y)) + (me.Z * (vec1.Z - vec2.Z))
}

//	Returns true if this 3D vector equals vec.
func (me *Vec3) Equals(vec *Vec3) bool {
	return (me.X == vec.X) && (me.Y == vec.Y) && (me.Z == vec.Z)
}

//	Returns the inverse of this 3D vector.
func (me *Vec3) Inv() *Vec3 {
	return &Vec3{1 / me.X, 1 / me.Y, 1 / me.Z}
}

//	Returns the length of this 3D vector.
func (me *Vec3) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z)
}

//	Returns the magnitude of this 3D vector.
func (me *Vec3) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Sets each component of this 3D vector to the corresponding component in vec if the former is infinity.
func (me *Vec3) MakeFinite(vec *Vec3) {
	if math.IsInf(me.X, 0) {
		me.X = vec.X
	}
	if math.IsInf(me.Y, 0) {
		me.Y = vec.Y
	}
	if math.IsInf(me.Z, 0) {
		me.Z = vec.Z
	}
}

//	Returns the biggest component of this 3D vector.
func (me *Vec3) Max() float64 {
	return math.Max(me.X, math.Max(me.Y, me.Z))
}

//	Returns the smallest component of this 3D vector.
func (me *Vec3) Min() float64 {
	return math.Min(me.X, math.Min(me.Y, me.Z))
}

//	Returns a new 3D vector that represents this 3D vector multiplied with vec.
func (me *Vec3) Mult(vec *Vec3) *Vec3 {
	return &Vec3{me.X * vec.X, me.Y * vec.Y, me.Z * vec.Z}
}

//	Returns a new 3D vector that represents this 3D vector's components each multipled with val.
func (me *Vec3) Mult1(val float64) *Vec3 {
	return &Vec3{me.X * val, me.Y * val, me.Z * val}
}

//	Normalizes this 3D vector.
func (me *Vec3) Normalize() {
	tmpVal := me.Magnitude()
	if tmpVal == 0 {
		me.X, me.Y, me.Z = tmpVal, tmpVal, tmpVal
	} else {
		tmpVal = 1 / tmpVal
		me.X, me.Y, me.Z = me.X*tmpVal, me.Y*tmpVal, me.Z*tmpVal
	}
}

//	Returns a new 3D vector that represents this 3D vector, normalized.
func (me *Vec3) Normalized() *Vec3 {
	tmpVal := me.Magnitude()
	if tmpVal == 0 {
		return &Vec3{tmpVal, tmpVal, tmpVal}
	}
	tmpVal = 1 / tmpVal
	return &Vec3{me.X * tmpVal, me.Y * tmpVal, me.Z * tmpVal}
}

//	Returns a new 3D vector that represents this 3D vector, normalized, then scaled by factor.
func (me *Vec3) NormalizedScaled(factor float64) (vec *Vec3) {
	vec = me.Normalized()
	vec.Scale(factor)
	return
}

//	Rotates this 3D vector angleDeg degrees around the specified axis.
func (me *Vec3) RotateDeg(angleDeg float64, axis *Vec3) {
	me.RotateRad(DegToRad(angleDeg/2), axis)
}

func (me *Vec3) RotateRad(angleRad float64, axis *Vec3) {
	tmpCos := math.Cos(angleRad)
	tmpSin := math.Sin(angleRad)
	tmpQr.X, tmpQr.Y, tmpQr.Z, tmpQr.W = axis.X*tmpSin, axis.Y*tmpSin, axis.Z*tmpSin, tmpCos
	tmpQc.SetFromConjugated(&tmpQr)
	tmpQ.SetFromMult3(&tmpQr, me)
	tmpQw.SetFromMult(&tmpQ, &tmpQc)
	me.X, me.Y, me.Z = tmpQw.X, tmpQw.Y, tmpQw.Z
}

//	Scales this 3D vector by factor.
func (me *Vec3) Scale(factor float64) {
	me.X, me.Y, me.Z = me.X*factor, me.Y*factor, me.Z*factor
}

//	Scales this 3D vector by factor, then adds add.
func (me *Vec3) ScaleAdd(factor, add *Vec3) {
	me.X, me.Y, me.Z = (me.X*factor.X)+add.X, (me.Y*factor.Y)+add.Y, (me.Z*factor.Z)+add.Z
}

//	Returns a new 3D vector that represents this 3D vector scaled by factor.
func (me *Vec3) Scaled(factor float64) *Vec3 {
	return &Vec3{me.X * factor, me.Y * factor, me.Z * factor}
}

//	Returns a new 3D vector that represents this 3D vector scaled by factor, then add added.
func (me *Vec3) ScaledAdded(factor float64, add *Vec3) *Vec3 {
	return &Vec3{(me.X * factor) + add.X, (me.Y * factor) + add.Y, (me.Z * factor) + add.Z}
}

//	Sets the components of this 3D vector to the specified value.
func (me *Vec3) Set(x, y, z float64) {
	me.X, me.Y, me.Z = x, y, z
}

//	Sets the components of this 3D vector to the same values as the components of vec.
func (me *Vec3) SetFrom(vec *Vec3) {
	me.X, me.Y, me.Z = vec.X, vec.Y, vec.Z
}

//	Sets this 3D vector to the result of adding vec1 and vec2.
func (me *Vec3) SetFromAdd(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X+vec2.X, vec1.Y+vec2.Y, vec1.Z+vec2.Z
}

//	Sets this 3D vector to the result of multiplying mul1 with mul2, then adding add.
func (me *Vec3) SetFromAddMult(add, mul1, mul2 *Vec3) {
	me.X, me.Y, me.Z = add.X+(mul1.X*mul2.X), add.Y+(mul1.Y*mul2.Y), add.Z+(mul1.Z*mul2.Z)
}

//	Sets this 3D vector to the result of scaling vec2 by mul, then adding vec1.
func (me *Vec3) SetFromAddMult1(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec1.X+(vec2.X*mul), vec1.Y+(vec2.Y*mul), vec1.Z+(vec2.Z*mul)
}

//	Sets each component of this 3D vector to the cosine of the corresponding component in vec.
func (me *Vec3) SetFromCos(vec *Vec3) {
	me.X, me.Y, me.Z = math.Cos(vec.X), math.Cos(vec.Y), math.Cos(vec.Z)
}

//	Sets each component of this 3D vector to the radian equivalent of the degree angle stored in the corresponding component in vec.
func (me *Vec3) SetFromDegToRad(deg *Vec3) {
	me.X, me.Y, me.Z = DegToRad(deg.X), DegToRad(deg.Y), DegToRad(deg.Z)
}

//	Sets each component of this 3D vector to Epsilon32 if it is 0 or greater but smaller than Epsilon32.
func (me *Vec3) SetFromEpsilon32() {
	if (me.X >= 0) && (me.X < Epsilon32) {
		me.X = Epsilon32
	}
	if (me.Y >= 0) && (me.Y < Epsilon32) {
		me.Y = Epsilon32
	}
	if (me.Z >= 0) && (me.Z < Epsilon32) {
		me.Z = Epsilon32
	}
}

//	Sets each component of this 3D vector to Epsilon64 if it is 0 or greater but smaller than Epsilon64.
func (me *Vec3) SetFromEpsilon64() {
	if (me.X >= 0) && (me.X < Epsilon64) {
		me.X = Epsilon64
	}
	if (me.Y >= 0) && (me.Y < Epsilon64) {
		me.Y = Epsilon64
	}
	if (me.Z >= 0) && (me.Z < Epsilon64) {
		me.Z = Epsilon64
	}
}

//	Sets this 3D vector to the inverse of vec.
func (me *Vec3) SetFromInv(vec *Vec3) {
	me.X, me.Y, me.Z = 1/vec.X, 1/vec.Y, 1/vec.Z
}

//	Sets this 3D vector to the result of v1 multiplied with v2.
func (me *Vec3) SetFromMult(v1, v2 *Vec3) {
	me.X, me.Y, me.Z = v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z
}

//	Sets this 3D vector to the result of vec scaled by mul.
func (me *Vec3) SetFromMult1(vec *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec.X*mul, vec.Y*mul, vec.Z*mul
}

//	Sets this 3D vector to the result of (vec1 minus vec2), scaled by mul.
func (me *Vec3) SetFromMult1Sub(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = (vec1.X-vec2.X)*mul, (vec1.Y-vec2.Y)*mul, (vec1.Z-vec2.Z)*mul
}

//	Sets this 3D vector to vec with each component's sign reversed.
func (me *Vec3) SetFromNeg(vec *Vec3) {
	me.X, me.Y, me.Z = -vec.X, -vec.Y, -vec.Z
}

//	Sets this 3D vector to pos rotated as expressed in rotCos and rotSin.
func (me *Vec3) SetFromRotation(pos, rotCos, rotSin *Vec3) {
	tmpVal := ((pos.Y * rotSin.X) + (pos.Z * rotCos.X))
	me.X = (pos.X * rotCos.Y) + (tmpVal * rotSin.Y)
	me.Y = (pos.Y * rotCos.X) - (pos.Z * rotSin.X)
	me.Z = (-pos.X * rotSin.Y) + (tmpVal * rotCos.Y)
}

//	Sets this 3D vector to the sine of vec.
func (me *Vec3) SetFromSin(vec *Vec3) {
	me.X, me.Y, me.Z = math.Sin(vec.X), math.Sin(vec.Y), math.Sin(vec.Z)
}

//	Sets each component of this 3D vector to the corresponding component in v0 if it is less than edge, else to the corresponding component in v1.
func (me *Vec3) SetFromStep(edge float64, vec, v0, v1 *Vec3) {
	if vec.X < edge {
		me.X = v0.X
	} else {
		me.X = v1.X
	}
	if vec.Y < edge {
		me.Y = v0.Y
	} else {
		me.Y = v1.Y
	}
	if vec.Z < edge {
		me.Z = v0.Z
	} else {
		me.Z = v1.Z
	}
}

//	Sets this 3D vector to the result of vec1 minus vec2.
func (me *Vec3) SetFromSub(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X-vec2.X, vec1.Y-vec2.Y, vec1.Z-vec2.Z
}

//	Sets this 3D vector to the result of (sub1 minus sub2), scaled by mul.
func (me *Vec3) SetFromSubMult(sub1, sub2, mul *Vec3) {
	me.X, me.Y, me.Z = (sub1.X-sub2.X)*mul.X, (sub1.Y-sub2.Y)*mul.Y, (sub1.Z-sub2.Z)*mul.Z
}

//	Sets this 3D vector to the result of sub1 minus (vec2 scaled by mul).
func (me *Vec3) SetFromSubMult1(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec1.X-(vec2.X*mul), vec1.Y-(vec2.Y*mul), vec1.Z-(vec2.Z*mul)
}

//	Returns a 3D vector where each component indicates the sign of this 3D vector's corresponding component.
func (me *Vec3) Sign() *Vec3 {
	return &Vec3{Sign(me.X), Sign(me.Y), Sign(me.Z)}
}

//	Returns a string representation of this 3D vector.
func (me *Vec3) String() string {
	return fmt.Sprintf("{X: %6.2f Y:%6.2f Z:%6.2f}", me.X, me.Y, me.Z)
}

//	Returns a new 3D vector that represents (this 3D vector minus vec).
func (me *Vec3) Sub(vec *Vec3) *Vec3 {
	return &Vec3{me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z}
}

//	Returns a new 3D vector that represents (this 3D vector minus sub) divided by div, multiplied with mul.
func (me *Vec3) SubDivMult(sub, div, mul *Vec3) *Vec3 {
	return &Vec3{((me.X - sub.X) / div.X) * mul.X, ((me.Y - sub.Y) / div.Y) * mul.Y, ((me.Z - sub.Z) / div.Z) * mul.Z}
}

//	Returns the dot product of (this 3D vector minus vec)
func (me *Vec3) SubDot(vec *Vec3) float64 {
	return ((me.X - vec.X) * (me.X - vec.X)) + ((me.Y - vec.Y) * (me.Y - vec.Y)) + ((me.Z - vec.Z) * (me.Z - vec.Z))
}

//	Returns a new 3D vector that represents the result of (this 3D vector divided by floorDiv) floored, scaled by mul.
func (me *Vec3) SubFloorDivMult(floorDiv, mul float64) *Vec3 {
	return me.Sub(&Vec3{math.Floor(me.X/floorDiv) * mul, math.Floor(me.Y/floorDiv) * mul, math.Floor(me.Z/floorDiv) * mul})
}

//	Returns a new 3D vector that represents (val minus this 3D vector).
func (me *Vec3) SubFrom(val float64) *Vec3 {
	return &Vec3{val - me.X, val - me.Y, val - me.Z}
}

//	Returns a new 3D vector that represents (this 3D vector minus vec), scaled by val.
func (me *Vec3) SubMult(vec *Vec3, val float64) *Vec3 {
	return &Vec3{(me.X - vec.X) * val, (me.Y - vec.Y) * val, (me.Z - vec.Z) * val}
}

//	Subtracts vec from this 3D vector.
func (me *Vec3) SubVec(vec *Vec3) {
	me.X, me.Y, me.Z = me.X-vec.X, me.Y-vec.Y, me.Z-vec.Z
}

//	Reverses the sign of each of this 3D vector's components.
func (me *Vec3) SwapSigns() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

/*
func (me *Vec3) ToDegInts () []int {
	return []int { int(RadToDeg(me.X)), int(RadToDeg(me.Y)), int(RadToDeg(me.Z)) }
}

func (me *Vec3) ToInts () []int {
	return []int { int(me.X), int(me.Y), int(me.Z) }
}
*/
