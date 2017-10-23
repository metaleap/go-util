# unum

Go programming helpers for common maths needs; plus vectors, matrices and
quaternions.

## Usage

```go
const (
	Deg2Rad = math.Pi / 180
	Rad2Deg = 180 / math.Pi
)
```

```go
var (
	Infinity             = math.Inf(1)
	NegativeInfinity     = math.Inf(-1)
	Epsilon              = math.Nextafter(1, Infinity) - 1
	EpsilonEqFloat       = 1.121039E-44
	EpsilonEqFloatFactor = 1E-06
	EpsilonEqVec         = 9.99999944E-11
)
```

#### func  Clamp

```go
func Clamp(val, c0, c1 float64) float64
```
Clamps `val` between `c0` and `c1`.

#### func  Clamp01

```go
func Clamp01(v float64) float64
```
Clamps `v` between 0 and 1.

#### func  ClosestPowerOfTwo

```go
func ClosestPowerOfTwo(v uint32) uint32
```
Returns `v` if it is a power-of-two, or else the closest power-of-two.

#### func  DegToRad

```go
func DegToRad(degrees float64) float64
```
Converts the specified `degrees` to radians.

#### func  DeltaAngle

```go
func DeltaAngle(cur, target float64) float64
```
Calculates the shortest difference between two given angles.

#### func  Eq

```go
func Eq(a, b float64) (eq bool)
```
Compares two floating point values if they are approximately equivalent.

#### func  InvLerp

```go
func InvLerp(from, to, val float64) float64
```
Calculates the Lerp parameter between of two values.

#### func  IsPowerOfTwo

```go
func IsPowerOfTwo(x uint32) bool
```
Returns whether `x` is a power-of-two.

#### func  Lerp

```go
func Lerp(a, b, t float64) float64
```
Returns `a` if `t` is 0, or `b` if `t` is 1, or else the linear interpolation
from `a` to `b` according to `t`.

#### func  LerpAngle

```go
func LerpAngle(a, b, t float64) float64
```
Same as Lerp but makes sure the values interpolate correctly when they wrap
around 360 degrees.

#### func  Mat3Identities

```go
func Mat3Identities(mats ...*Mat3)
```
Calls the `Identity` method on all specified `mats`.

#### func  Mat4Identities

```go
func Mat4Identities(mats ...*Mat4)
```
Calls the `Identity` method on all specified `mats`.

#### func  NextPowerOfTwo

```go
func NextPowerOfTwo(v uint32) uint32
```
Returns `v` if it is a power-of-two, or else the next-highest power-of-two.

#### func  Percent

```go
func Percent(p, of float64) float64
```

#### func  PingPong

```go
func PingPong(t, l float64) float64
```
Ping-pongs the value `t`, so that it is never larger than `l` and never smaller
than 0.

#### func  RadToDeg

```go
func RadToDeg(radians float64) float64
```
Converts the specified `radians` to degrees.

#### func  Round

```go
func Round(v float64) (fint float64)
```
Returns the next-higher integer if fraction>0.5; if fraction<0.5 returns the
next-lower integer; if fraction==0.5, returns the next even integer.

#### func  Sign

```go
func Sign(v float64) (sign float64)
```
Returns -1 if `v` is negative, 1 if `v` is positive, or 0 if `v` is zero.

#### func  SmoothStep

```go
func SmoothStep(from, to, t float64) float64
```
Interpolates between `from` and `to` with smoothing at the limits.

#### func  SmootherStep

```go
func SmootherStep(from, to, t float64) float64
```
Interpolates between `from` and `to` with smoother smoothing at the limits.

#### func  SumFrom1To

```go
func SumFrom1To(to int) int
```
http://betterexplained.com/articles/techniques-for-adding-the-numbers-1-to-100/

#### func  SumFromTo

```go
func SumFromTo(from, to int) int
```
http://betterexplained.com/articles/techniques-for-adding-the-numbers-1-to-100/

#### type Mat3

```go
type Mat3 [9]float64
```

Represents a 3x3 matrix.

```go
var (
	//	The 3x3 identity matrix.
	Mat3Identity Mat3
)
```

#### func  NewMat3Identity

```go
func NewMat3Identity() (mat *Mat3)
```
Returns a new 3x3 identity matrix.

#### func (*Mat3) Identity

```go
func (me *Mat3) Identity()
```
Sets this 3x3 matrix to `Mat3Identity`.

#### func (*Mat3) Transpose

```go
func (me *Mat3) Transpose()
```
Transposes this 3x3 matrix.

#### type Mat4

```go
type Mat4 [16]float64
```

Represents a 4x4 column-major matrix.

```go
var (
	//	The 4x4 identity matrix.
	Mat4Identity Mat4
)
```

#### func  NewMat4Add

```go
func NewMat4Add(a, b *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` representing the result of adding `a` to `b`.

#### func  NewMat4Frustum

```go
func NewMat4Frustum(left, right, bottom, top, near, far float64) (mat *Mat4)
```
Returns a new `*Mat4` representing the specified frustum.

#### func  NewMat4Identity

```go
func NewMat4Identity() (mat *Mat4)
```
Returns a new `*Mat4` representing the identity matrix.

#### func  NewMat4Lookat

```go
func NewMat4Lookat(eyePos, lookTarget, upVec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` representing the "look-at matrix" computed from the
specified vectors.

#### func  NewMat4Mult1

```go
func NewMat4Mult1(m *Mat4, v float64) (mat *Mat4)
```
Returns a new `*Mat4` representing the result of multiplying all values in `m`
with `v`.

#### func  NewMat4Mult4

```go
func NewMat4Mult4(one, two *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents the result of multiplying `one` with
`two`.

#### func  NewMat4MultN

```go
func NewMat4MultN(mats ...*Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents the result of multiplying all specified
`mats` with one another.

#### func  NewMat4Orient

```go
func NewMat4Orient(lookTarget, worldUp *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` representing the "orientation matrix" computed from the
specified vectors.

#### func  NewMat4Perspective

```go
func NewMat4Perspective(fovY, aspect, near, far float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents the specified perspective-projection
matrix.

#### func  NewMat4RotationX

```go
func NewMat4RotationX(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the X
axis.

#### func  NewMat4RotationY

```go
func NewMat4RotationY(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the Y
axis.

#### func  NewMat4RotationZ

```go
func NewMat4RotationZ(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the Z
axis.

#### func  NewMat4Scaling

```go
func NewMat4Scaling(vec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` that represents a transformation of "scale by `vec`".

#### func  NewMat4Sub

```go
func NewMat4Sub(a, b *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents `a` minus `b`.

#### func  NewMat4Translation

```go
func NewMat4Translation(vec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` that represents a transformation of "translate by `vec`".

#### func (*Mat4) Abs

```go
func (me *Mat4) Abs() (abs *Mat4)
```
Returns a new `*Mat4` with each cell representing the `math.Abs` value of the
respective corresponding cell in `me`.

#### func (*Mat4) Add

```go
func (me *Mat4) Add(mat *Mat4)
```
Adds `mat` to `me`.

#### func (*Mat4) Clear

```go
func (me *Mat4) Clear()
```
Zeroes all cells in `me`.

#### func (*Mat4) Clone

```go
func (me *Mat4) Clone() (mat *Mat4)
```
Returns a new `*Mat` containing a copy of `me`.

#### func (*Mat4) CopyFrom

```go
func (me *Mat4) CopyFrom(mat *Mat4)
```
Copies all cells from `mat` to `me`.

#### func (*Mat4) CopyTo

```go
func (me *Mat4) CopyTo(mat *Mat4)
```
Copies all cells from `me` to `mat`.

#### func (*Mat4) Frustum

```go
func (me *Mat4) Frustum(left, right, bottom, top, near, far float64)
```
Sets `me` to represent the specified frustum.

#### func (*Mat4) Identity

```go
func (me *Mat4) Identity()
```
Copies all cells from `Mat4Identity` to `me`.

#### func (*Mat4) Lookat

```go
func (me *Mat4) Lookat(eyePos, lookTarget, upVec *Vec3)
```
Sets `me` to the "look-at matrix" computed from the specified vectors.

#### func (*Mat4) Mult1

```go
func (me *Mat4) Mult1(v float64)
```
Multiplies all cells in `me` with `v`.

#### func (*Mat4) Orient

```go
func (me *Mat4) Orient(lookTarget, worldUp *Vec3)
```
Sets `me` to the "orientation matrix" computed from the specified vectors.

#### func (*Mat4) Perspective

```go
func (me *Mat4) Perspective(fovYDeg, a, n, f float64) (fovYRadHalf float64)
```
Sets `me` to the specified perspective-projection matrix.

`fovYRad` -- vertical field-of-view angle in radians. `a` -- aspect ratio. `n`
-- near-plane. `f` -- far-plane.

#### func (*Mat4) RotationX

```go
func (me *Mat4) RotationX(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the X
axis".

#### func (*Mat4) RotationY

```go
func (me *Mat4) RotationY(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the Y
axis".

#### func (*Mat4) RotationZ

```go
func (me *Mat4) RotationZ(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the Z
axis".

#### func (*Mat4) Scaling

```go
func (me *Mat4) Scaling(vec *Vec3)
```
Sets `me` to a transformation matrix representing "scale by `vec`"

#### func (*Mat4) SetFromMult4

```go
func (me *Mat4) SetFromMult4(one, two *Mat4)
```
Sets `me` to the result of multiplying `one` times `two`.

#### func (*Mat4) SetFromMultN

```go
func (me *Mat4) SetFromMultN(mats ...*Mat4)
```
Sets `me` to the result of multiplying all the specified `mats` with one
another.

#### func (*Mat4) SetFromTransposeOf

```go
func (me *Mat4) SetFromTransposeOf(mat *Mat4)
```
Sets `me` to the transpose of `mat`.

#### func (*Mat4) Sub

```go
func (me *Mat4) Sub(mat *Mat4)
```
Subtracts `mat` from `me`.

#### func (*Mat4) Translation

```go
func (me *Mat4) Translation(vec *Vec3)
```
Sets `me` to a transformation matrix representing "translate by `vec`"

#### func (*Mat4) Transposed

```go
func (me *Mat4) Transposed() (mat *Mat4)
```
Returns the transpose of `me`.

#### type Quat

```go
type Quat struct {
	//	X, Y, Z, W
	Vec4
}
```

Quaternion

#### func  NewQuat

```go
func NewQuat(x, y, z, w float64) *Quat
```

#### func  Quat_Identity

```go
func Quat_Identity() (q Quat)
```

#### func (*Quat) AngleDeg

```go
func (me *Quat) AngleDeg(q *Quat) float64
```

#### func (*Quat) AngleRad

```go
func (me *Quat) AngleRad(q *Quat) float64
```

#### func (*Quat) Eq

```go
func (me *Quat) Eq(vec *Vec4) bool
```

#### func (*Quat) Mul

```go
func (me *Quat) Mul(q *Quat) *Quat
```

#### func (*Quat) MulVec3

```go
func (me *Quat) MulVec3(p *Vec3) *Vec3
```

#### type Vec2

```go
type Vec2 struct{ X, Y float64 }
```

A 2-dimensional vector.

#### func  Vec2_Lerp

```go
func Vec2_Lerp(from, to *Vec2, t float64) *Vec2
```

#### func  Vec2_Max

```go
func Vec2_Max(l, r *Vec2) *Vec2
```

#### func  Vec2_Min

```go
func Vec2_Min(l, r *Vec2) *Vec2
```

#### func  Vec2_One

```go
func Vec2_One() Vec2
```
Vec2{1, 1}

#### func  Vec2_Right

```go
func Vec2_Right() Vec2
```
Vec2{1, 0}

#### func  Vec2_Up

```go
func Vec2_Up() Vec2
```
Vec2{0, 1}

#### func  Vec2_Zero

```go
func Vec2_Zero() Vec2
```
Vec2{0, 0}

#### func (*Vec2) Add

```go
func (me *Vec2) Add(vec *Vec2)
```

#### func (*Vec2) AddedDiv

```go
func (me *Vec2) AddedDiv(a *Vec2, d float64) *Vec2
```

#### func (*Vec2) AngleDeg

```go
func (me *Vec2) AngleDeg(to *Vec2) float64
```

#### func (*Vec2) AngleRad

```go
func (me *Vec2) AngleRad(to *Vec2) float64
```

#### func (*Vec2) ClampMagnitude

```go
func (me *Vec2) ClampMagnitude(maxLength float64) *Vec2
```

#### func (*Vec2) Clear

```go
func (me *Vec2) Clear()
```

#### func (*Vec2) Distance

```go
func (me *Vec2) Distance(vec *Vec2) float64
```

#### func (*Vec2) Div

```go
func (me *Vec2) Div(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that is the result of dividing `me` by `vec` without
checking for division-by-0.

#### func (*Vec2) DivSafe

```go
func (me *Vec2) DivSafe(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that is the result of dividing `me` by `vec`, safely
checking for division-by-0.

#### func (*Vec2) Divide

```go
func (me *Vec2) Divide(d float64)
```

#### func (*Vec2) Divided

```go
func (me *Vec2) Divided(d float64) *Vec2
```

#### func (*Vec2) Dot

```go
func (me *Vec2) Dot(vec *Vec2) float64
```
Returns the dot product of `me` and `vec`.

#### func (*Vec2) Eq

```go
func (me *Vec2) Eq(vec *Vec2) bool
```

#### func (*Vec2) Length

```go
func (me *Vec2) Length() float64
```
Returns the 2D vector length of `me`.

#### func (*Vec2) Magnitude

```go
func (me *Vec2) Magnitude() float64
```
Returns the 2D vector magnitude of `me`.

#### func (*Vec2) MoveTowards

```go
func (me *Vec2) MoveTowards(target *Vec2, maxDistanceDelta float64) *Vec2
```

#### func (*Vec2) Mult

```go
func (me *Vec2) Mult(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that is the result of multiplying `me` with `vec`.

#### func (*Vec2) Negate

```go
func (me *Vec2) Negate() *Vec2
```

#### func (*Vec2) Normalize

```go
func (me *Vec2) Normalize()
```
Normalizes `me` in-place without checking for division-by-0.

#### func (*Vec2) NormalizeSafe

```go
func (me *Vec2) NormalizeSafe()
```
Normalizes `me` in-place, safely checking for division-by-0.

#### func (*Vec2) Normalized

```go
func (me *Vec2) Normalized() *Vec2
```
Returns a new `*Vec2` that is the normalized representation of `me` without
checking for division-by-0.

#### func (*Vec2) NormalizedSafe

```go
func (me *Vec2) NormalizedSafe() *Vec2
```
Returns a new `*Vec2` that is the normalized representation of `me`, safely
checking for division-by-0.

#### func (*Vec2) NormalizedScaled

```go
func (me *Vec2) NormalizedScaled(factor float64) *Vec2
```
Returns a new `*Vec2` that is the normalized representation of `me` scaled by
`factor` without checking for division-by-0.

#### func (*Vec2) NormalizedScaledSafe

```go
func (me *Vec2) NormalizedScaledSafe(factor float64) *Vec2
```
Returns a new `*Vec2` that is the normalized representation of `me` scaled by
`factor`, safely checking for division-by-0.

#### func (*Vec2) Scale

```go
func (me *Vec2) Scale(factor float64)
```
Multiplies all components in `me` with `factor`.

#### func (*Vec2) Scaled

```go
func (me *Vec2) Scaled(factor float64) *Vec2
```
Returns a new `*Vec2` that represents `me` scaled by `factor`.

#### func (*Vec2) Set

```go
func (me *Vec2) Set(x, y float64)
```

#### func (*Vec2) String

```go
func (me *Vec2) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

#### func (*Vec2) Sub

```go
func (me *Vec2) Sub(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that represents `me` minus `vec`.

#### func (*Vec2) Subtract

```go
func (me *Vec2) Subtract(vec *Vec2)
```
Subtracts `vec` from `me`.

#### type Vec3

```go
type Vec3 struct {
	X, Y, Z float64
}
```

Represents a 3-dimensional vector.

#### func  Vec3_Back

```go
func Vec3_Back() Vec3
```

#### func  Vec3_Down

```go
func Vec3_Down() Vec3
```

#### func  Vec3_Fwd

```go
func Vec3_Fwd() Vec3
```

#### func  Vec3_Left

```go
func Vec3_Left() Vec3
```

#### func  Vec3_Lerp

```go
func Vec3_Lerp(from, to *Vec3, t float64) *Vec3
```

#### func  Vec3_Max

```go
func Vec3_Max(l, r *Vec3) *Vec3
```

#### func  Vec3_Min

```go
func Vec3_Min(l, r *Vec3) *Vec3
```

#### func  Vec3_One

```go
func Vec3_One() Vec3
```

#### func  Vec3_Right

```go
func Vec3_Right() Vec3
```

#### func  Vec3_Up

```go
func Vec3_Up() Vec3
```

#### func  Vec3_Zero

```go
func Vec3_Zero() Vec3
```

#### func (*Vec3) Add

```go
func (me *Vec3) Add(vec *Vec3)
```
Adds `vec` to `me` in-place.

#### func (*Vec3) Add1

```go
func (me *Vec3) Add1(val float64)
```
Adds `val` to all 3 components of `me`.

#### func (*Vec3) Add3

```go
func (me *Vec3) Add3(x, y, z float64)
```
Adds the specified 3 components to the respective components in `me`.

#### func (*Vec3) Added

```go
func (me *Vec3) Added(vec *Vec3) *Vec3
```
Returns the sum of `me` and `vec`.

#### func (*Vec3) AllEq

```go
func (me *Vec3) AllEq(val float64) bool
```
Returns whether all 3 components in `me` are approximately equivalent to their
respective counterparts in `val`.

#### func (*Vec3) AllGEq

```go
func (me *Vec3) AllGEq(vec *Vec3) bool
```
Returns whether all 3 components in `me` are greater than (or approximately
equivalent to) their respective component counterparts in `vec`.

#### func (*Vec3) AllIn

```go
func (me *Vec3) AllIn(min, max *Vec3) bool
```
Returns whether all 3 components in `me` are greater than `min`, and also less
than `max`.

#### func (*Vec3) AllLEq

```go
func (me *Vec3) AllLEq(vec *Vec3) bool
```
Returns whether all 3 components in `me` are less than (or approximately
equivalent to) their respective component counterparts in `vec`.

#### func (*Vec3) AngleDeg

```go
func (me *Vec3) AngleDeg(to *Vec3) float64
```

#### func (*Vec3) AngleRad

```go
func (me *Vec3) AngleRad(to *Vec3) float64
```

#### func (*Vec3) Clamp

```go
func (me *Vec3) Clamp(min, max *Vec3)
```
Clamps each component in `me` between the respective corresponding counter-part
component in `min` and `max`.

#### func (*Vec3) Clamp01

```go
func (me *Vec3) Clamp01()
```
Clamps each component in `me` between 0 and 1.

#### func (*Vec3) ClampMagnitude

```go
func (me *Vec3) ClampMagnitude(maxLength float64) *Vec3
```

#### func (*Vec3) Clear

```go
func (me *Vec3) Clear()
```
Zeroes all 3 components in `me`.

#### func (*Vec3) Cross

```go
func (me *Vec3) Cross(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents the cross-product of `me` and `vec`.

#### func (*Vec3) CrossNormalized

```go
func (me *Vec3) CrossNormalized(vec *Vec3) (r *Vec3)
```
Returns a new `*Vec` that represents the cross-product of `me` and `vec`,
normalized.

#### func (*Vec3) Distance

```go
func (me *Vec3) Distance(vec *Vec3) float64
```
Returns the distance of `me` from `vec`.

#### func (*Vec3) DistanceManhattan

```go
func (me *Vec3) DistanceManhattan(vec *Vec3) float64
```
Returns the "manhattan distance" of `me` from `vec`.

#### func (*Vec3) Div

```go
func (me *Vec3) Div(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` divided by `vec`.

#### func (*Vec3) Divide

```go
func (me *Vec3) Divide(d float64)
```

#### func (*Vec3) Divided

```go
func (me *Vec3) Divided(d float64) *Vec3
```
Returns a new `*Vec3` that represents all 3 components in `me`, each divided by
`val`.

#### func (*Vec3) Dot

```go
func (me *Vec3) Dot(vec *Vec3) float64
```
Returns the dot-product of `me` and `vec`.

#### func (*Vec3) DotSub

```go
func (me *Vec3) DotSub(vec1, vec2 *Vec3) float64
```
Returns the dot-product of `me` and (`vec1` minus `vec2`).

#### func (*Vec3) Eq

```go
func (me *Vec3) Eq(vec *Vec3) bool
```

#### func (*Vec3) Length

```go
func (me *Vec3) Length() float64
```
Returns the 3D vector length of `me`.

#### func (*Vec3) Magnitude

```go
func (me *Vec3) Magnitude() float64
```
Returns the 3D vector magnitude of `me`.

#### func (*Vec3) Max

```go
func (me *Vec3) Max() float64
```
Returns the largest of the 3 components in `me`.

#### func (*Vec3) MaxAbs

```go
func (me *Vec3) MaxAbs() float64
```
Returns the `math.Max` of the `math.Abs` values of all 3 components in `me`.

#### func (*Vec3) Min

```go
func (me *Vec3) Min() float64
```
Returns the smallest of the 3 components in `me`.

#### func (*Vec3) Mult

```go
func (me *Vec3) Mult(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` multiplied with `vec`.

#### func (*Vec3) Mult3

```go
func (me *Vec3) Mult3(x, y, z float64) *Vec3
```
Returns a new `*Vec3` with each component in `me` multiplied by the respective
corresponding specified factor.

#### func (*Vec3) Negate

```go
func (me *Vec3) Negate()
```
Reverses the signs of all 3 vector components in `me`.

#### func (*Vec3) Negated

```go
func (me *Vec3) Negated() *Vec3
```
Returns a new `*Vec` with each component representing the negative (sign
inverted) corresponding component in `me`.

#### func (*Vec3) Normalize

```go
func (me *Vec3) Normalize()
```
Normalizes `me` in-place without checking for division-by-0.

#### func (*Vec3) NormalizeSafe

```go
func (me *Vec3) NormalizeSafe()
```
Normalizes `me` in-place, safely checking for division-by-0.

#### func (*Vec3) Normalized

```go
func (me *Vec3) Normalized() *Vec3
```
Returns a new `*Vec3` that represents `me`, normalized.

#### func (*Vec3) NormalizedScaled

```go
func (me *Vec3) NormalizedScaled(factor float64) (vec *Vec3)
```
Returns a new `*Vec3` that represents `me` normalized, then scaled by `factor`.

#### func (*Vec3) Rcp

```go
func (me *Vec3) Rcp() *Vec3
```
Returns a new `*Vec3` representing `1/me`.

#### func (*Vec3) RotateDeg

```go
func (me *Vec3) RotateDeg(angleDeg float64, axis *Vec3)
```
Rotates `me` `angleDeg` degrees around the specified `axis`.

#### func (*Vec3) RotateRad

```go
func (me *Vec3) RotateRad(angleRad float64, axis *Vec3)
```
Rotates `me` `angleRad` radians around the specified `axis`.

#### func (*Vec3) Scale

```go
func (me *Vec3) Scale(factor float64)
```
Scales `me` by `factor`.

#### func (*Vec3) ScaleAdd

```go
func (me *Vec3) ScaleAdd(factor, add *Vec3)
```
Scales `me` by `factor`, then adds `add`.

#### func (*Vec3) Scaled

```go
func (me *Vec3) Scaled(factor float64) *Vec3
```
Returns a new `*Vec3` that represents `me` scaled by `factor`.

#### func (*Vec3) ScaledAdded

```go
func (me *Vec3) ScaledAdded(factor float64, add *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` scaled by `factor`, then `add` added.

#### func (*Vec3) Set

```go
func (me *Vec3) Set(x, y, z float64)
```
Sets all 3 vector components in `me` to the corresponding respective specified
value.

#### func (*Vec3) SetFromAdd

```go
func (me *Vec3) SetFromAdd(vec1, vec2 *Vec3)
```
Sets `me` to the result of adding `vec1` and `vec2`.

#### func (*Vec3) SetFromAddAdd

```go
func (me *Vec3) SetFromAddAdd(a, b, c *Vec3)
```
`me = a + b + c`

#### func (*Vec3) SetFromAddScaled

```go
func (me *Vec3) SetFromAddScaled(vec1, vec2 *Vec3, mul float64)
```
`me = mul * vec2 + vec1`

#### func (*Vec3) SetFromAddSub

```go
func (me *Vec3) SetFromAddSub(a, b, c *Vec3)
```
`me = a + b - c`

#### func (*Vec3) SetFromCross

```go
func (me *Vec3) SetFromCross(vec *Vec3)
```
Sets `me` to the cross-product of `me` and `vec`.

#### func (*Vec3) SetFromCrossOf

```go
func (me *Vec3) SetFromCrossOf(one, two *Vec3)
```
Sets `me` to the cross-product of `one` and `two`.

#### func (*Vec3) SetFromDegToRad

```go
func (me *Vec3) SetFromDegToRad(deg *Vec3)
```
Sets each vector component in `me` to the radian equivalent of the degree angle
stored in the respective corresponding component of `vec`.

#### func (*Vec3) SetFromDivided

```go
func (me *Vec3) SetFromDivided(vec *Vec3, d float64)
```

#### func (*Vec3) SetFromMad

```go
func (me *Vec3) SetFromMad(mul1, mul2, add *Vec3)
```
`me = mul1 * mul2 + add`

#### func (*Vec3) SetFromMult

```go
func (me *Vec3) SetFromMult(v1, v2 *Vec3)
```
`me = v1 * v2`

#### func (*Vec3) SetFromNegated

```go
func (me *Vec3) SetFromNegated(vec *Vec3)
```
`me = -vec`

#### func (*Vec3) SetFromNormalized

```go
func (me *Vec3) SetFromNormalized(vec *Vec3)
```
Sets `me` to `vec` normalized.

#### func (*Vec3) SetFromRcp

```go
func (me *Vec3) SetFromRcp(vec *Vec3)
```
Sets `me` to the inverse of `vec`.

#### func (*Vec3) SetFromRotation

```go
func (me *Vec3) SetFromRotation(pos, rotCos, rotSin *Vec3)
```
Sets `me` to `pos` rotated as expressed in `rotCos` and `rotSin`.

#### func (*Vec3) SetFromScaled

```go
func (me *Vec3) SetFromScaled(vec *Vec3, mul float64)
```
`me = vec * mul`

#### func (*Vec3) SetFromScaledSub

```go
func (me *Vec3) SetFromScaledSub(vec1, vec2 *Vec3, mul float64)
```
`me = (vec1 - vec2) * mul`

#### func (*Vec3) SetFromSub

```go
func (me *Vec3) SetFromSub(vec1, vec2 *Vec3)
```
`me = vec1 - vec2`.

#### func (*Vec3) SetFromSubAdd

```go
func (me *Vec3) SetFromSubAdd(a, b, c *Vec3)
```
`me = a - b + c`

#### func (*Vec3) SetFromSubMult

```go
func (me *Vec3) SetFromSubMult(sub1, sub2, mul *Vec3)
```
`me = (sub1 - sub2) * mul`

#### func (*Vec3) SetFromSubScaled

```go
func (me *Vec3) SetFromSubScaled(v1, v2 *Vec3, v2Scale float64)
```
`me = v1 - v2 * v2Scale`

#### func (*Vec3) SetFromSubSub

```go
func (me *Vec3) SetFromSubSub(a, b, c *Vec3)
```
`me = a - b - c`

#### func (*Vec3) SetToMax

```go
func (me *Vec3) SetToMax()
```
Sets all 3 vector components in `me` to `math.MaxFloat64`.

#### func (*Vec3) SetToMin

```go
func (me *Vec3) SetToMin()
```
Sets all 3 vector components in `me` to `-math.MaxFloat64`.

#### func (*Vec3) Sign

```go
func (me *Vec3) Sign() *Vec3
```
Returns a new `*Vec3` with each vector component indicating the sign (-1, 1 or
0) of the respective corresponding component in `me`.

#### func (*Vec3) String

```go
func (me *Vec3) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

#### func (*Vec3) Sub

```go
func (me *Vec3) Sub(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` minus `vec`.

#### func (*Vec3) SubDivMult

```go
func (me *Vec3) SubDivMult(sub, div, mul *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `((me - sub) / div) * mul`.

#### func (*Vec3) SubFloorDivMult

```go
func (me *Vec3) SubFloorDivMult(div, mul float64) *Vec3
```
Returns a new `*Vec3` that represents `mul * math.Floor(me / div)`.

#### func (*Vec3) SubFrom

```go
func (me *Vec3) SubFrom(val float64) *Vec3
```
Returns a new `*Vec3` that represents `val` minus `me`.

#### func (*Vec3) SubScaled

```go
func (me *Vec3) SubScaled(vec *Vec3, val float64) *Vec3
```
Returns a new `*Vec3` that represents `(me - vec) * val`.

#### func (*Vec3) Subtract

```go
func (me *Vec3) Subtract(vec *Vec3)
```
Subtracts `vec` from `me`.

#### func (*Vec3) TransformCoord

```go
func (me *Vec3) TransformCoord(mat *Mat4)
```
Transform coordinate vector `me` according to the specified `*Mat4`.

#### func (*Vec3) TransformNormal

```go
func (me *Vec3) TransformNormal(mat *Mat4, absMat bool)
```
Transform normal vector `me` according to the specified `*Mat4`.

#### type Vec4

```go
type Vec4 struct {
	X, Y, Z, W float64
}
```

Represents an arbitrary 4-dimensional vector.

#### func  Vec4_Lerp

```go
func Vec4_Lerp(from, to *Vec4, t float64) *Vec4
```

#### func  Vec4_Max

```go
func Vec4_Max(l, r *Vec4) *Vec4
```

#### func  Vec4_Min

```go
func Vec4_Min(l, r *Vec4) *Vec4
```

#### func  Vec4_One

```go
func Vec4_One() Vec4
```

#### func  Vec4_Zero

```go
func Vec4_Zero() Vec4
```

#### func (*Vec4) AddedDiv

```go
func (me *Vec4) AddedDiv(a *Vec4, d float64) *Vec4
```

#### func (*Vec4) Clear

```go
func (me *Vec4) Clear()
```

#### func (*Vec4) Clone

```go
func (me *Vec4) Clone() (q *Vec4)
```
Returns a new `*Vec4` containing a copy of `me`.

#### func (*Vec4) Conjugate

```go
func (me *Vec4) Conjugate()
```
Negates the `X`, `Y`, `Z` components in `me`, but not `W`.

#### func (*Vec4) Conjugated

```go
func (me *Vec4) Conjugated() (v *Vec4)
```
Returns a new `*Vec4` that represents `me` conjugated.

#### func (*Vec4) Distance

```go
func (me *Vec4) Distance(vec *Vec4) float64
```

#### func (*Vec4) Divide

```go
func (me *Vec4) Divide(d float64)
```

#### func (*Vec4) Divided

```go
func (me *Vec4) Divided(d float64) *Vec4
```

#### func (*Vec4) Dot

```go
func (me *Vec4) Dot(vec *Vec4) float64
```

#### func (*Vec4) Eq

```go
func (me *Vec4) Eq(vec *Vec4) bool
```

#### func (*Vec4) Length

```go
func (me *Vec4) Length() float64
```
Returns the 4D vector length of `me`.

#### func (*Vec4) Magnitude

```go
func (me *Vec4) Magnitude() float64
```
Returns the 4D vector magnitude of `me`.

#### func (*Vec4) MoveTowards

```go
func (me *Vec4) MoveTowards(target *Vec4, maxDistanceDelta float64) *Vec4
```

#### func (*Vec4) MultMat4

```go
func (me *Vec4) MultMat4(mat *Mat4)
```
Sets `me` to the result of multiplying the specified `*Mat4` with `me`.

#### func (*Vec4) MultMat4Vec3

```go
func (me *Vec4) MultMat4Vec3(mat *Mat4, vec *Vec3)
```
Sets `me` to the result of multiplying the specified `*Mat4` with the specified
`*Vec3`.

#### func (*Vec4) MultMat4Vec4

```go
func (me *Vec4) MultMat4Vec4(mat *Mat4, vec *Vec4)
```
Sets `me` to the result of multiplying the specified `*Mat4` with the specified
`*Vec4`.

#### func (*Vec4) Negate

```go
func (me *Vec4) Negate()
```

#### func (*Vec4) Negated

```go
func (me *Vec4) Negated() *Vec4
```

#### func (*Vec4) Normalize

```go
func (me *Vec4) Normalize()
```
Normalizes `me` according to `me.Magnitude`.

#### func (*Vec4) NormalizeFrom

```go
func (me *Vec4) NormalizeFrom(magnitude float64)
```
Normalizes `me` according to the specified `magnitude`.

#### func (*Vec4) Normalized

```go
func (me *Vec4) Normalized() *Vec4
```
Returns a new `*Vec4` that represents `me` normalized according to
`me.Magnitude`.

#### func (*Vec4) Project

```go
func (me *Vec4) Project(vec *Vec4)
```

#### func (*Vec4) Projected

```go
func (me *Vec4) Projected(vec *Vec4) *Vec4
```

#### func (*Vec4) Scale

```go
func (me *Vec4) Scale(v float64)
```
Scales all 4 vector components in `me` by factor `v`.

#### func (*Vec4) Scaled

```go
func (me *Vec4) Scaled(v float64) *Vec4
```

#### func (*Vec4) SetFromConjugated

```go
func (me *Vec4) SetFromConjugated(c *Vec4)
```
Sets `me` to `c` conjugated.

#### func (*Vec4) SetFromMult

```go
func (me *Vec4) SetFromMult(l, r *Vec4)
```
Applies various 4D vector component computations of `l` and `r` to `me`, as
needed by the `Vec3.RotateRad` method.

#### func (*Vec4) SetFromMult3

```go
func (me *Vec4) SetFromMult3(q *Vec4, v *Vec3)
```
Applies various 4D vector component computations of `q` and `v` to `me`, as
needed by the `Vec3.RotateRad` method.

#### func (*Vec4) SetFromVec3

```go
func (me *Vec4) SetFromVec3(vec *Vec3)
```

#### func (*Vec4) String

```go
func (me *Vec4) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

#### func (*Vec4) Sub

```go
func (me *Vec4) Sub(vec *Vec4) *Vec4
```

#### func (*Vec4) Subtract

```go
func (me *Vec4) Subtract(vec *Vec4)
```
