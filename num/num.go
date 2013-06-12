package unum

import (
	"fmt"
	"math"
)

const (
	PiDiv180  = math.Pi / 180
	PiDiv360  = math.Pi / 360
	PiHalfDiv = 0.5 / math.Pi
)

var (
	Infinity    = math.Inf(1)
	NegInfinity = math.Inf(-1)

	Epsilon32 = Nextafter32(1, Infinity) - 1
	Epsilon64 = math.Nextafter(1, Infinity) - 1
)

//	Returns true if all vals equal test.
func AllEqual(test float64, vals ...float64) bool {
	for i := 0; i < len(vals); i++ {
		if vals[i] != test {
			return false
		}
	}
	return true
}

//	Clamps val between c0 and c1
func Clamp(val, c0, c1 float64) float64 {
	// return math.Min(math.Max(val, c0), c1)
	if val < c0 {
		return c0
	}
	if val > c1 {
		return c1
	}
	return val
}

//	Converts the specified degrees to radians.
func DegToRad(deg float64) float64 {
	return PiDiv180 * deg
}

//	Returns the "normalized ratio" of val to max.
//	Example: for max = 900 and val = 300, returns 0.33333.
func Din1(val, max float64) float64 {
	return 1 / (max / val)
}

//	Returns the "normalized ratio" of val to max.
//	Example: for max = 900 and val = 300, returns 0.33333.
func Fin1(val, max float32) float32 {
	return 1 / (max / val)
}

/*
func Hash3 (one, two, three uint) uint {
	var rshift = func (x, y uint) uint {
		return x >> y
	}
	one = one - two;  one = one - three;  one = one ^ (rshift(three, 13));
	two = two - three;  two = two - one;  two = two ^ (one << 8);
	three = three - one;  three = three - two;  three = three ^ (rshift(two, 13));
	one = one - two;  one = one - three;  one = one ^ (rshift(three, 12));
	two = two - three;  two = two - one;  two = two ^ (one << 16);
	three = three - one;  three = three - two;  three = three ^ (rshift(two, 5));
	one = one - two;  one = one - three;  one = one ^ (rshift(three, 3));
	two = two - three;  two = two - one;  two = two ^ (one << 10);
	three = three - one;  three = three - two;  three = three ^ (rshift(two, 15));
	return three;
}

func Iin1 (val, max int) int {
	return 1 / (max / val)
}
*/

//	Returns true if val is even
func IsEven(val int) bool {
	return (math.Mod(float64(val), 2) == 0)
}

//	Returns true if val represents an integer
func IsInt(val float64) bool {
	_, f := math.Modf(val)
	return (f == 0)
}

//	Returns true if math.Mod(v, m) is zero
func IsMod0(v, m int) bool {
	return (math.Mod(float64(v), float64(m)) == 0)
}

/*
func Lin1 (val, max int64) int64 {
	return 1 / (max / val)
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
*/

//	Returns the smaller of two values.
func Mini(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

//	Returns x if a is 0, y if a is 1, or a corresponding mix of both if a is between 0 and 1.
func Mix(x, y, a float64) float64 {
	return (x * y) + ((1 - y) * a)
}

//	https://groups.google.com/d/topic/golang-nuts/dVtKN8QLUNM/discussion
func Nextafter32(x, y float64) (r float64) {
	switch {
	case math.IsNaN(x) || math.IsNaN(y):
		r = math.NaN()
	case x == y:
		r = x
	case x == 0:
		// const sign = 1 << 31
		// r = math.Float32frombits(1 | (math.Float32bits(y) & sign))
		r = math.Copysign(float64(math.Float32frombits(1)), float64(y))
	case (y > x) == (x > 0):
		r = float64(math.Float32frombits(math.Float32bits(float32(x)) + 1))
	default:
		r = float64(math.Float32frombits(math.Float32bits(float32(x)) - 1))
	}
	return
}

//	Converts the specified radians to degrees.
func RadToDeg(rad float64) float64 {
	return rad * PiDiv180
}

//	Returns (the equivalent of) math.Ceil(v) if fraction >= 0.5, otherwise returns (the equivalent of) math.Floor(v).
func Round(v float64) (fint float64) {
	var frac float64
	if fint, frac = math.Modf(v); frac >= 0.5 {
		fint++
	}
	return
}

//	Clamps v between 0 and 1.
func Saturate(v float64) float64 {
	return Clamp(v, 0, 1)
}

//	Returns -1 if v is negative, 1 if v is positive, or 0 if v is zero.
func Sign(v float64) (sign float64) {
	if v > 0 {
		sign = 1
	} else if v < 0 {
		sign = -1
	}
	return
	// return v / math.Abs(v)
}

//	Returns 0 if x < edge, otherwise returns 1.
func Step(edge, x float64) (step int) {
	if edge >= x {
		step = 1
	}
	return
	// if x < edge {
	// 	return 0
	// }
	// return 1
}

func strf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
