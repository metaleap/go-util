package unum

import (
	"math"
)

//	A 2-dimensional vector.
type Vec2 struct{ X, Y float64 }

//	Returns a new 2D vector that is the result of dividing this 2D vector by vec.
func (me *Vec2) Div(vec *Vec2) *Vec2 {
	return &Vec2{me.X / vec.X, me.Y / vec.Y}
}

//	Returns the dot product of this 2D vector and vec.
func (me *Vec2) Dot(vec *Vec2) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y)
}

//	Returns the length of this 2D vector.
func (me *Vec2) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y)
}

//	Returns the magnitude of this 2D vector.
func (me *Vec2) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Returns a new 2D vector that is the result of multiplying this 2D vector with vec.
func (me *Vec2) Mult(vec *Vec2) *Vec2 {
	return &Vec2{me.X * vec.X, me.Y * vec.Y}
}

//	Normalizes this 2D vector.
func (me *Vec2) Normalize() {
	// l := 1 / me.Magnitude()
	// me.X, me.Y = me.X*l, me.Y*l
	me.Scale(1 / me.Magnitude())
}

//	Returns a new 2D vector that is the normalized representation of this 2D vector.
func (me *Vec2) Normalized() *Vec2 {
	return me.Scaled(1 / me.Magnitude())
	// l := 1 / me.Magnitude()
	// return Vec2{me.X * l, me.Y * l}
}

//	Returns a new 2D vector that is the normalized representation of this 2D vector scaled by factor.
func (me *Vec2) NormalizedScaled(factor float64) (vec *Vec2) {
	vec = me.Normalized()
	vec.Scale(factor)
	return
	// l := 1 / me.Magnitude()
	// return Vec2{me.X * l * factor, me.Y * l * factor}
}

//	Multiplies all components in me with factor.
func (me *Vec2) Scale(factor float64) {
	me.X, me.Y = me.X*factor, me.Y*factor
}

//	Returns a new 2D vector that represents this 2D vector scaled by factor.
func (me *Vec2) Scaled(factor float64) *Vec2 {
	return &Vec2{me.X * factor, me.Y * factor}
}

func (me *Vec2) String() string {
	return strf("{X:%1.2f Y:%1.2f}", me.X, me.Y)
}

//	Returns a new 2D vector that represents this 2D vector with vec subtracted.
func (me *Vec2) Sub(vec *Vec2) *Vec2 {
	return &Vec2{me.X - vec.X, me.Y - vec.Y}
}

/*
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
*/
