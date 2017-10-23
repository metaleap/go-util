package unum

import (
	"fmt"
	"math"
)

const (
	Deg2Rad = math.Pi / 180
	Rad2Deg = 180 / math.Pi
)

var (
	Infinity             = math.Inf(1)
	NegativeInfinity     = math.Inf(-1)
	Epsilon              = math.Nextafter(1, Infinity) - 1
	EpsilonEqFloat       = 1.121039E-44
	EpsilonEqFloatFactor = 1E-06
	EpsilonEqVec         = 9.99999944E-11
)

//	Clamps `val` between `c0` and `c1`.
func Clamp(val, c0, c1 float64) float64 {
	switch {
	case val < c0:
		return c0
	case val > c1:
		return c1
	}
	return val
}

//	Clamps `v` between 0 and 1.
func Clamp01(v float64) float64 {
	return Clamp(v, 0, 1)
}

//	Returns `v` if it is a power-of-two, or else the closest power-of-two.
func ClosestPowerOfTwo(v uint32) uint32 {
	next := NextPowerOfTwo(v)
	if prev := next / 2; (v - prev) < (next - v) {
		next = prev
	}
	return next
}

//	Converts the specified `degrees` to radians.
func DegToRad(degrees float64) float64 {
	return degrees * Deg2Rad
}

//	Calculates the shortest difference between two given angles.
func DeltaAngle(cur, target float64) float64 {
	sin, cos := math.Sincos(target - cur)
	return math.Atan2(sin, cos) // HACK: could be atan2(cos,sin) instead of atan2(sin,cos) ...
}

//	Compares two floating point values if they are approximately equivalent.
func Eq(a, b float64) (eq bool) {
	if eq = (a == b); !eq {
		diff := math.Abs(b - a)
		eq = diff <= Epsilon || diff < math.Max(EpsilonEqFloat, EpsilonEqFloatFactor*math.Max(math.Abs(a), math.Abs(b)))
	}
	return
}

//	Calculates the Lerp parameter between of two values.
func InvLerp(from, to, val float64) float64 {
	return (val - from) / (to - from)
}

//	Returns whether `x` is a power-of-two.
func IsPowerOfTwo(x uint32) bool {
	return x == (x & ^(x & (x - 1)))
}

//	Returns `a` if `t` is 0, or `b` if `t` is 1, or else the linear interpolation from `a` to `b` according to `t`.
func Lerp(a, b, t float64) float64 {
	return ((b-a)*t)+a
}

//	Same as Lerp but makes sure the values interpolate correctly when they wrap around 360 degrees.
func LerpAngle(a, b, t float64) float64 {
	return t * (math.Mod(math.Mod(b-a, 360)+540, 360) - 180)
}

//	Returns `v` if it is a power-of-two, or else the next-highest power-of-two.
func NextPowerOfTwo(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func Percent(p, of float64) float64 {
	return p * of * 0.01
}

func PingPong(t, l float64) float64 {
	l = l * 2
	return l - math.Mod(t, l)
}

//	Converts the specified `radians` to degrees.
func RadToDeg(radians float64) float64 {
	return radians * Rad2Deg
}

//	Returns the next-higher integer if fraction>0.5; if fraction<0.5 returns the next-lower integer; if fraction==0.5, returns the next even integer.
func Round(v float64) (fint float64) {
	var frac float64
	if fint, frac = math.Modf(v); frac > 0.5 || (frac == 0.5 && math.Mod(fint, 2) != 0) {
		fint++
	}
	return
}

//	Returns -1 if `v` is negative, 1 if `v` is positive, or 0 if `v` is zero.
func Sign(v float64) (sign float64) {
	if v > 0 {
		sign = 1
	} else if v < 0 {
		sign = -1
	}
	return
	// return v / math.Abs(v)
}

//	Interpolates between `from` and `to` with smoothing at the limits.
func SmoothStep(from, to, t float64) float64 {
	t = Clamp01((t - from) / (to - from))
	return (t * t) * (3 - 2*t)
}

//	Interpolates between `from` and `to` with smoother smoothing at the limits.
func SmootherStep(from, to, t float64) float64 {
	t = Clamp01((t - from) / (to - from))
	return t * t * t * (t*(t*6-15) + 10)
}

//	http://betterexplained.com/articles/techniques-for-adding-the-numbers-1-to-100/
func SumFromTo(from, to int) int {
	return ((to * (to + 1)) / 2) - ((from - 1) * from / 2)
}

//	http://betterexplained.com/articles/techniques-for-adding-the-numbers-1-to-100/
func SumFrom1To(to int) int {
	return (to * (to + 1)) / 2
}

func strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

/*

//	Returns whether `val` is even.
func IsEven(val int) bool {
	return (math.Mod(float64(val), 2) == 0)
}

//	Returns whether `val` represents an integer.
func IsInt(val float64) bool {
	_, f := math.Modf(val)
	return (f == 0)
}

//	Returns whether `math.Mod(v, m)` is 0.
func IsMod0(v, m int) bool {
	return (math.Mod(float64(v), float64(m)) == 0)
}

func Absi (v int32) int32 {
	return v - (v ^ (v >> 31))
}

func Absl (v int64) int64 {
	return v - (v ^ (v >> 63))
}

func Max (x, y float64) float64 {
	return 0.5 * (x + y + math.Abs(x - y))
}

func Min (x, y float64) float64 {
	return 0.5 * (x + y - math.Abs(x - y))
}

//	Returns the smaller of two `int` values.
func Mini(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
*/
