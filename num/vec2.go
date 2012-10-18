package num

import (
	"math"
	"strconv"
)

type TVec2 struct { X, Y float64 }

func (me TVec2) Div (vec TVec2) TVec2 {
	return TVec2 { me.X / vec.X, me.Y / vec.Y }
}

func (me TVec2) Dot (vec TVec2) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y)
}

func (me TVec2) Length () float64 {
	return (me.X * me.X) + (me.Y * me.Y)
}

func (me TVec2) Magnitude () float64 {
	return math.Sqrt(me.Length())
}

func (me TVec2) Mult (vec TVec2) TVec2 {
	return TVec2 { me.X * vec.X, me.Y * vec.Y }
}

func (me *TVec2) Normalize () {
	var l = 1 / me.Magnitude()
	me.X *= l
	me.Y *= l
}

func (me TVec2) Normalized () TVec2 {
	var l = 1 / me.Magnitude()
	return TVec2 { me.X * l, me.Y * l }
}

func (me TVec2) NormalizedScaled (by float64) TVec2 {
	var l = 1 / me.Magnitude()
	return TVec2 { me.X * l * by, me.Y * l * by }
}

func (me TVec2) Scaled (by float64) TVec2 {
	return TVec2 { me.X * by, me.Y * by }
}

func (me TVec2) Sub (vec TVec2) TVec2 {
	return TVec2 { me.X - vec.X, me.Y - vec.Y }
}

func IsVec2 (any interface{}) bool {
	if any != nil {
		if _, isT := any.(TVec2); isT {
			return true
		}
	}
	return false
}

func NewVec2 (vals ... string) (TVec2, error) {
	var err error
	var f TVec2
	if f.X, err = strconv.ParseFloat(vals[0], 64); err == nil {
		f.Y, err = strconv.ParseFloat(vals[1], 64)
	}
	return f, err
}
