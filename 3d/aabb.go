package u3d

import (
	"math"

	"github.com/metaleap/go-util/num"
)

type AaBb struct {
	Min, Max, Center, Extent unum.Vec3
}

func (me *AaBb) BoundingSphere(center *unum.Vec3) (radius float64) {
	return math.Max(me.Min.Distance(center), me.Max.Distance(center))
}

func (me *AaBb) Clear() {
	me.Min.Clear()
	me.Max.Clear()
	me.Center.Clear()
	me.Extent.Clear()
}

func (me *AaBb) ResetMinMax() {
	me.Max.SetToMin()
	me.Min.SetToMax()
}

func (me *AaBb) SetCenterExtent() {
	me.Center.SetFromAdd(&me.Max, &me.Min)
	me.Center.Scale(0.5)
	me.Extent.SetFromSub(&me.Max, &me.Min)
	me.Extent.Scale(0.5)
}

func (me *AaBb) SetMinMax() {
	me.Min.SetFromSub(&me.Center, &me.Extent)
	me.Max.SetFromAdd(&me.Center, &me.Extent)
}

func (me *AaBb) Transform(mat *unum.Mat4) {
	me.Center.TransformCoord(mat)
	me.Extent.TransformNormal(mat, true)
	me.SetMinMax()
}

func (me *AaBb) UpdateMinMax(vec *unum.Vec3) {
	vec.Clamp(&me.Min, &me.Max)
}

func (me *AaBb) UpdateMinMaxFrom(aabb *AaBb) {
	me.UpdateMinMax(&aabb.Min)
	me.UpdateMinMax(&aabb.Max)
}
