package num

import (
	"math"
)

const (
	PiDiv180 = math.Pi / 180
	PiHalfDiv = 0.5 / math.Pi
)

var (
	Epsilon float64 = 0
	EpsilonMax float64 = 0
	Infinity float64
	NegInfinity float64
)

func AllEqual (test float64, vals ... float64) bool {
	for i := 0; i < len(vals); i++ { if vals[i] != test { return false } }
	return true
}

func Clamp (val, c0, c1 float64) float64 {
	return math.Min(math.Max(val, c0), c1);
}

func DegToRad (deg float64) float64 {
	return PiDiv180 * deg
}

func Din1 (val, max float64) float64 {
	return 1 / (max / val)
}

func Fin1 (val, max float32) float32 {
	return 1 / (max / val)
}

func Iin1 (val, max int) int {
	return 1 / (max / val)
}

func IsEveni (val int) bool {
	return (math.Mod(float64(val), 2) == 0)
}

func IsInt (val float64) bool {
	_, f := math.Modf(val)
	return f == 0
}

func IsMod0 (v, m int) bool {
	return math.Mod(float64(v), float64(m)) == 0
}

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

func Mini (v1, v2 int) int {
	if v1 < v2 { return v1 }
	return v2
}

func Mix (x, y, a float64) float64 {
	return (x * y) + ((1 - y) * a)
}

func RadToDeg (rad float64) float64 {
	return rad * PiDiv180
}

func Round (v float64) float64 {
	var frac float64
	if _, frac = math.Modf(v); frac >= 0.5 { return math.Ceil(v) }
	return math.Floor(v)
}

func Sign (v float64) float64 {
	if v == 0 { return 0 }
	return v / math.Abs(v)
}

func Step (edge, x float64) int {
	if x < edge { return 0 }
	return 1
}

func init () {
	var eps, i float64
	Infinity, NegInfinity = math.Inf(1), math.Inf(-1)
	for i = 0; i <= 8192; i++ {
		if eps = math.Pow(2, -i); eps == 0 { break }
		if i == 23 {
			Epsilon = eps
		} else {
			EpsilonMax = eps
		}
	}
}
