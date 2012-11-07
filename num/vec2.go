package num

import (
	"math"
	"strconv"
)

type Vec2 struct { X, Y float64 }

func (me Vec2) Div (vec Vec2) Vec2 {
	return Vec2 { me.X / vec.X, me.Y / vec.Y }
}

func (me Vec2) Dot (vec Vec2) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y)
}

func (me Vec2) Length () float64 {
	return (me.X * me.X) + (me.Y * me.Y)
}

func (me Vec2) Magnitude () float64 {
	return math.Sqrt(me.Length())
}

func (me Vec2) Mult (vec Vec2) Vec2 {
	return Vec2 { me.X * vec.X, me.Y * vec.Y }
}

func (me *Vec2) Normalize () {
	var l = 1 / me.Magnitude()
	me.X *= l
	me.Y *= l
}

func (me Vec2) Normalized () Vec2 {
	var l = 1 / me.Magnitude()
	return Vec2 { me.X * l, me.Y * l }
}

func (me Vec2) NormalizedScaled (by float64) Vec2 {
	var l = 1 / me.Magnitude()
	return Vec2 { me.X * l * by, me.Y * l * by }
}

func (me Vec2) Scaled (by float64) Vec2 {
	return Vec2 { me.X * by, me.Y * by }
}

func (me Vec2) Sub (vec Vec2) Vec2 {
	return Vec2 { me.X - vec.X, me.Y - vec.Y }
}

func IsVec2 (any interface{}) bool {
	if any != nil {
		if _, isT := any.(Vec2); isT {
			return true
		}
	}
	return false
}

func NewVec2 (vals ... string) (Vec2, error) {
	var err error
	var f Vec2
	if f.X, err = strconv.ParseFloat(vals[0], 64); err == nil {
		f.Y, err = strconv.ParseFloat(vals[1], 64)
	}
	return f, err
}
