package unum

import (
	"math"
)

//	Represents a 4x4 column-major matrix.
type Mat4 [16]float64

var (
	//	The 4x4 identity matrix.
	Mat4Identity Mat4

	m4z Mat4
)

func init() {
	Mat4Identity[0], Mat4Identity[4], Mat4Identity[8], Mat4Identity[12] = 1, 0, 0, 0
	Mat4Identity[1], Mat4Identity[5], Mat4Identity[9], Mat4Identity[13] = 0, 1, 0, 0
	Mat4Identity[2], Mat4Identity[6], Mat4Identity[10], Mat4Identity[14] = 0, 0, 1, 0
	Mat4Identity[3], Mat4Identity[7], Mat4Identity[11], Mat4Identity[15] = 0, 0, 0, 1
}

//	Adds mat to this 4x4 matrix.
func (me *Mat4) Add(mat *Mat4) {
	me[0], me[4], me[8], me[12] = me[0]+mat[0], me[4]+mat[4], me[8]+mat[8], me[12]+mat[12]
	me[1], me[5], me[9], me[13] = me[1]+mat[1], me[5]+mat[5], me[9]+mat[9], me[13]+mat[13]
	me[2], me[6], me[10], me[14] = me[2]+mat[2], me[6]+mat[6], me[10]+mat[10], me[14]+mat[14]
	me[3], me[7], me[11], me[15] = me[3]+mat[3], me[7]+mat[7], me[11]+mat[11], me[15]+mat[15]
}

//	Zeroes this 4x4 matrix.
func (me *Mat4) Clear() {
	*me = m4z
}

//	Returns a pointer to a newly allocated copy of me.
func (me *Mat4) Clone() (mat *Mat4) {
	mat = &Mat4{}
	me.CopyTo(mat)
	return
}

//	Sets this 4x4 matrix to the same values as mat.
func (me *Mat4) CopyFrom(mat *Mat4) {
	*me = *mat
}

//	Sets mat to the same values as this 4x4 matrix.
func (me *Mat4) CopyTo(mat *Mat4) {
	*mat = *me
}

//	Sets this 4x4 matrix to represent the specified frustum.
func (me *Mat4) Frustum(left, right, bottom, top, near, far float64) {
	me[0], me[4], me[8], me[12] = ((near * 2) / (right - left)), 0, ((right + left) / (right - left)), 0
	me[1], me[5], me[9], me[13] = 0, ((near * 2) / (top - bottom)), ((top + bottom) / (top - bottom)), 0
	me[2], me[6], me[10], me[14] = 0, 0, -(far+near)/(far-near), (-(far * near * 2) / (far - near))
	me[3], me[7], me[11], me[15] = 0, 0, -1, 0
}

//	Sets this 4x4 matrix to Mat4Identity.
func (me *Mat4) Identity() {
	*me = Mat4Identity
}

//	Sets this 4x4 matrix to the specified look-at matrix.
func (me *Mat4) LookAt(bag *Bag, lookTarget, worldUp *Vec3) {
	/*
		var aFwd, aSide, aUp = &Vec3 { eyePos.X - lookTarget.X, eyePos.Y - lookTarget.Y, eyePos.Z - lookTarget.Z }, &Vec3 {}, &Vec3 {}
		aFwd.Normalize()
		aSide.X, aSide.Y, aSide.Z = (worldUp.Y * aFwd.Z) - (worldUp.Z * aFwd.Y), (worldUp.Z * aFwd.X) - (worldUp.X * aFwd.Z), (worldUp.X * aFwd.Y) - (worldUp.Y * aFwd.X)
		aSide.Normalize()
		aUp.X, aUp.Y, aUp.Z = (aFwd.Y * aSide.Z) - (aFwd.Z * aSide.Y), (aFwd.Z * aSide.X) - (aFwd.X * aSide.Z), (aFwd.X * aSide.Y) - (aFwd.Y * aSide.X)
		aUp.Normalize()
		me[0], me[4], me[8], me[12] = aSide.X, aSide.Y, aSide.Z, -((aSide.X * eyePos.X) + (aSide.Y * eyePos.Y) + (aSide.Z * eyePos.Z))
		me[1], me[5], me[9], me[13] = aUp.X, aUp.Y, aUp.Z, -((aUp.X * eyePos.X) + (aUp.Y * eyePos.Y) + (aUp.Z * eyePos.Z))
		me[2], me[6], me[10], me[14] = aFwd.X, aFwd.Y, aFwd.Z, -((aFwd.X * eyePos.X) + (aFwd.Y * eyePos.Y) + (aFwd.Z * eyePos.Z))
		me[3], me[7], me[11], me[15] = 0, 0, 0, 1
	*/
	// bag.m4l.tvN = lookTarget.Normalized()
	// bag.m4l.tvU = worldUp.Normalized().Cross(lookTarget)
	// bag.m4l.tvV = bag.m4l.tvN.Cross(bag.m4l.tvU)
	bag.m4l.tvN.SetFromNormalized(lookTarget)
	bag.m4l.tvU.SetFromCrossOf(worldUp.Normalized(), lookTarget)
	bag.m4l.tvV.SetFromCrossOf(&bag.m4l.tvN, &bag.m4l.tvU)
	me[0], me[4], me[8], me[12] = bag.m4l.tvU.X, bag.m4l.tvU.Y, bag.m4l.tvU.Z, 0
	me[1], me[5], me[9], me[13] = bag.m4l.tvV.X, bag.m4l.tvV.Y, bag.m4l.tvV.Z, 0
	me[2], me[6], me[10], me[14] = bag.m4l.tvN.X, bag.m4l.tvN.Y, bag.m4l.tvN.Z, 0
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Multiplies all values in this 4x4 matrix with v.
func (me *Mat4) Mult1(v float64) {
	me[0], me[4], me[8], me[12] = me[0]*v, me[4]*v, me[8]*v, me[12]*v
	me[1], me[5], me[9], me[13] = me[1]*v, me[5]*v, me[9]*v, me[13]*v
	me[2], me[6], me[10], me[14] = me[2]*v, me[6]*v, me[10]*v, me[14]*v
	me[3], me[7], me[11], me[15] = me[3]*v, me[7]*v, me[11]*v, me[15]*v
}

//	Sets this 4x4 matrix to the specified perspective-projection matrix.
func (me *Mat4) Perspective(bag *Bag, fovY, aspect, near, far float64) {
	bag.m4p.tfY = near * math.Tan(fovY*math.Pi/360)
	bag.m4p.tfX = bag.m4p.tfY * aspect
	me.Frustum(-bag.m4p.tfX, bag.m4p.tfX, -bag.m4p.tfY, bag.m4p.tfY, near, far)
}

/*
func (me *Mat4) Rotation (rad float64, axes *Vec3) {
	var cos, sin = math.Cos(rad), math.Sin(rad)
	var x, y, z = axes.X, axes.Y, axes.Z
	var xx, yy, zz, xy, xz, yz = x * x, y * y, z * z, x * y, x * z, y * z
	me[0], me[4], me[8], me[12] = (xx + (1 - xx) * cos),		(xy * (1 - cos) - z * sin),	(xz * (1 - cos) + y * sin),	0
	me[1], me[5], me[9], me[13] = (xy * (1 - cos) + z * sin),	(yy + (1 - yy) * cos),		(yz * (1 - cos) - x * sin),	0
	me[2], me[6], me[10], me[14] = (xz * (1 - cos) - y * sin),	(yz * (1 - cos) + x * sin),	(zz + (1 - zz) * cos),		0
	me[3], me[7], me[11], me[15] = 0,							0,							0,							1
}
*/

//	Sets this 4x4 matrix to a rotation matrix representing "rotate rad radians around the X asis".
func (me *Mat4) RotationX(bag *Bag, rad float64) {
	bag.m4p.tfX, bag.m4p.tfY = math.Cos(rad), math.Sin(rad)
	me[0], me[4], me[8], me[12] = 1, 0, 0, 0
	me[1], me[5], me[9], me[13] = 0, bag.m4p.tfX, -bag.m4p.tfY, 0
	me[2], me[6], me[10], me[14] = 0, bag.m4p.tfY, bag.m4p.tfX, 0
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Sets this 4x4 matrix to a rotation matrix representing "rotate rad radians around the Y asis".
func (me *Mat4) RotationY(bag *Bag, rad float64) {
	bag.m4p.tfX, bag.m4p.tfY = math.Cos(rad), math.Sin(rad)
	me[0], me[4], me[8], me[12] = bag.m4p.tfX, 0, bag.m4p.tfY, 0
	me[1], me[5], me[9], me[13] = 0, 1, 0, 0
	me[2], me[6], me[10], me[14] = -bag.m4p.tfY, 0, bag.m4p.tfX, 0
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Sets this 4x4 matrix to a rotation matrix representing "rotate rad radians around the Z asis".
func (me *Mat4) RotationZ(bag *Bag, rad float64) {
	bag.m4p.tfX, bag.m4p.tfY = math.Cos(rad), math.Sin(rad)
	me[0], me[4], me[8], me[12] = bag.m4p.tfX, -bag.m4p.tfY, 0, 0
	me[1], me[5], me[9], me[13] = bag.m4p.tfY, bag.m4p.tfX, 0, 0
	me[2], me[6], me[10], me[14] = 0, 0, 1, 0
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Sets this 4x4 matrix to a transformation matrix representing "scale by vec"
func (me *Mat4) Scaling(vec *Vec3) {
	me[0], me[4], me[8], me[12] = vec.X, 0, 0, 0
	me[1], me[5], me[9], me[13] = 0, vec.Y, 0, 0
	me[2], me[6], me[10], me[14] = 0, 0, vec.Z, 0
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Sets this 4x4 matrix to the result of multiplying one with two.
func (me *Mat4) SetFromMult4(one, two *Mat4) {
	me[0], me[4], me[8], me[12] = (one[0]*two[0])+(one[4]*two[1])+(one[8]*two[2])+(one[12]*two[3]), (one[0]*two[4])+(one[4]*two[5])+(one[8]*two[6])+(one[12]*two[7]), (one[0]*two[8])+(one[4]*two[9])+(one[8]*two[10])+(one[12]*two[11]), (one[0]*two[12])+(one[4]*two[13])+(one[8]*two[14])+(one[12]*two[15])
	me[1], me[5], me[9], me[13] = (one[1]*two[0])+(one[5]*two[1])+(one[9]*two[2])+(one[13]*two[3]), (one[1]*two[4])+(one[5]*two[5])+(one[9]*two[6])+(one[13]*two[7]), (one[1]*two[8])+(one[5]*two[9])+(one[9]*two[10])+(one[13]*two[11]), (one[1]*two[12])+(one[5]*two[13])+(one[9]*two[14])+(one[13]*two[15])
	me[2], me[6], me[10], me[14] = (one[2]*two[0])+(one[6]*two[1])+(one[10]*two[2])+(one[14]*two[3]), (one[2]*two[4])+(one[6]*two[5])+(one[10]*two[6])+(one[14]*two[7]), (one[2]*two[8])+(one[6]*two[9])+(one[10]*two[10])+(one[14]*two[11]), (one[2]*two[12])+(one[6]*two[13])+(one[10]*two[14])+(one[14]*two[15])
	me[3], me[7], me[11], me[15] = (one[3]*two[0])+(one[7]*two[1])+(one[11]*two[2])+(one[15]*two[3]), (one[3]*two[4])+(one[7]*two[5])+(one[11]*two[6])+(one[15]*two[7]), (one[3]*two[8])+(one[7]*two[9])+(one[11]*two[10])+(one[15]*two[11]), (one[3]*two[12])+(one[7]*two[13])+(one[11]*two[14])+(one[15]*two[15])
}

//	Sets this 4x4 matrix to the result of multiplying all the specified mats with one another.
func (me *Mat4) SetFromMultN(bag *Bag, mats ...*Mat4) {
	bag.m4n.m0.Clear()
	bag.m4n.m1 = mats[0]
	for bag.m4n.i = 1; bag.m4n.i < len(mats); bag.m4n.i++ {
		if bag.m4n.m2 = mats[bag.m4n.i]; bag.m4n.m2 != nil {
			me[0], me[4], me[8], me[12] = (bag.m4n.m1[0]*bag.m4n.m2[0])+(bag.m4n.m1[4]*bag.m4n.m2[1])+(bag.m4n.m1[8]*bag.m4n.m2[2])+(bag.m4n.m1[12]*bag.m4n.m2[3]), (bag.m4n.m1[0]*bag.m4n.m2[4])+(bag.m4n.m1[4]*bag.m4n.m2[5])+(bag.m4n.m1[8]*bag.m4n.m2[6])+(bag.m4n.m1[12]*bag.m4n.m2[7]), (bag.m4n.m1[0]*bag.m4n.m2[8])+(bag.m4n.m1[4]*bag.m4n.m2[9])+(bag.m4n.m1[8]*bag.m4n.m2[10])+(bag.m4n.m1[12]*bag.m4n.m2[11]), (bag.m4n.m1[0]*bag.m4n.m2[12])+(bag.m4n.m1[4]*bag.m4n.m2[13])+(bag.m4n.m1[8]*bag.m4n.m2[14])+(bag.m4n.m1[12]*bag.m4n.m2[15])
			me[1], me[5], me[9], me[13] = (bag.m4n.m1[1]*bag.m4n.m2[0])+(bag.m4n.m1[5]*bag.m4n.m2[1])+(bag.m4n.m1[9]*bag.m4n.m2[2])+(bag.m4n.m1[13]*bag.m4n.m2[3]), (bag.m4n.m1[1]*bag.m4n.m2[4])+(bag.m4n.m1[5]*bag.m4n.m2[5])+(bag.m4n.m1[9]*bag.m4n.m2[6])+(bag.m4n.m1[13]*bag.m4n.m2[7]), (bag.m4n.m1[1]*bag.m4n.m2[8])+(bag.m4n.m1[5]*bag.m4n.m2[9])+(bag.m4n.m1[9]*bag.m4n.m2[10])+(bag.m4n.m1[13]*bag.m4n.m2[11]), (bag.m4n.m1[1]*bag.m4n.m2[12])+(bag.m4n.m1[5]*bag.m4n.m2[13])+(bag.m4n.m1[9]*bag.m4n.m2[14])+(bag.m4n.m1[13]*bag.m4n.m2[15])
			me[2], me[6], me[10], me[14] = (bag.m4n.m1[2]*bag.m4n.m2[0])+(bag.m4n.m1[6]*bag.m4n.m2[1])+(bag.m4n.m1[10]*bag.m4n.m2[2])+(bag.m4n.m1[14]*bag.m4n.m2[3]), (bag.m4n.m1[2]*bag.m4n.m2[4])+(bag.m4n.m1[6]*bag.m4n.m2[5])+(bag.m4n.m1[10]*bag.m4n.m2[6])+(bag.m4n.m1[14]*bag.m4n.m2[7]), (bag.m4n.m1[2]*bag.m4n.m2[8])+(bag.m4n.m1[6]*bag.m4n.m2[9])+(bag.m4n.m1[10]*bag.m4n.m2[10])+(bag.m4n.m1[14]*bag.m4n.m2[11]), (bag.m4n.m1[2]*bag.m4n.m2[12])+(bag.m4n.m1[6]*bag.m4n.m2[13])+(bag.m4n.m1[10]*bag.m4n.m2[14])+(bag.m4n.m1[14]*bag.m4n.m2[15])
			me[3], me[7], me[11], me[15] = (bag.m4n.m1[3]*bag.m4n.m2[0])+(bag.m4n.m1[7]*bag.m4n.m2[1])+(bag.m4n.m1[11]*bag.m4n.m2[2])+(bag.m4n.m1[15]*bag.m4n.m2[3]), (bag.m4n.m1[3]*bag.m4n.m2[4])+(bag.m4n.m1[7]*bag.m4n.m2[5])+(bag.m4n.m1[11]*bag.m4n.m2[6])+(bag.m4n.m1[15]*bag.m4n.m2[7]), (bag.m4n.m1[3]*bag.m4n.m2[8])+(bag.m4n.m1[7]*bag.m4n.m2[9])+(bag.m4n.m1[11]*bag.m4n.m2[10])+(bag.m4n.m1[15]*bag.m4n.m2[11]), (bag.m4n.m1[3]*bag.m4n.m2[12])+(bag.m4n.m1[7]*bag.m4n.m2[13])+(bag.m4n.m1[11]*bag.m4n.m2[14])+(bag.m4n.m1[15]*bag.m4n.m2[15])
			bag.m4n.m0 = *me
			bag.m4n.m1 = &bag.m4n.m0
		}
	}
}

//	Subtracts mat from this 4x4 matrix.
func (me *Mat4) Sub(mat *Mat4) {
	me[0], me[4], me[8], me[12] = me[0]-mat[0], me[4]-mat[4], me[8]-mat[8], me[12]-mat[12]
	me[1], me[5], me[9], me[13] = me[1]-mat[1], me[5]-mat[5], me[9]-mat[9], me[13]-mat[13]
	me[2], me[6], me[10], me[14] = me[2]-mat[2], me[6]-mat[6], me[10]-mat[10], me[14]-mat[14]
	me[3], me[7], me[11], me[15] = me[3]-mat[3], me[7]-mat[7], me[11]-mat[11], me[15]-mat[15]
}

//	Sets the specified 3x3 matrix to the inverse of me.
//	This method is currently in "not needed right now and not sure if actually correct" limbo.
func (me *Mat4) ToInverseMat3(mat *Mat3) {
	a00, a01, a02 := me[0], me[1], me[2]
	a10, a11, a12 := me[4], me[5], me[6]
	a20, a21, a22 := me[8], me[9], me[10]
	b01 := a22*a11 - a12*a21
	b11 := -a22*a10 + a12*a20
	b21 := a21*a10 - a11*a20
	d := a00*b01 + a01*b11 + a02*b21
	dInv := 1 / d

	mat[0], mat[3], mat[6] = b01*dInv, b11*dInv, b21*dInv
	mat[1], mat[4], mat[7] = (-a22*a01+a02*a21)*dInv, (a22*a00-a02*a20)*dInv, (-a21*a00+a01*a20)*dInv
	mat[2], mat[5], mat[8] = (a12*a01-a02*a11)*dInv, (-a12*a00+a02*a10)*dInv, (a11*a00-a01*a10)*dInv
}

//	Sets this 4x4 matrix to a transformation matrix representing "translate by vec"
func (me *Mat4) Translation(vec *Vec3) {
	me[0], me[4], me[8], me[12] = 1, 0, 0, vec.X
	me[1], me[5], me[9], me[13] = 0, 1, 0, vec.Y
	me[2], me[6], me[10], me[14] = 0, 0, 1, vec.Z
	me[3], me[7], me[11], me[15] = 0, 0, 0, 1
}

//	Calls the Identity() method on all specified mats.
func Mat4Identities(mats ...*Mat4) {
	for _, mat := range mats {
		mat.Identity()
	}
}

//	Returns a new 4x4 matrix representing the result of adding a to b.
func NewMat4Add(a, b *Mat4) (mat *Mat4) {
	mat = &Mat4{}
	mat[0], mat[4], mat[8], mat[12] = a[0]+b[0], a[4]+b[4], a[8]+b[8], a[12]+b[12]
	mat[1], mat[5], mat[9], mat[13] = a[1]+b[1], a[5]+b[5], a[9]+b[9], a[13]+b[13]
	mat[2], mat[6], mat[10], mat[14] = a[2]+b[2], a[6]+b[6], a[10]+b[10], a[14]+b[14]
	mat[3], mat[7], mat[11], mat[15] = a[3]+b[3], a[7]+b[7], a[11]+b[11], a[15]+b[15]
	return
}

//	Returns a new 4x4 matrix representing the specified frustum.
func NewMat4Frustum(left, right, bottom, top, near, far float64) (mat *Mat4) {
	mat = &Mat4{}
	mat.Frustum(left, right, bottom, top, near, far)
	return
}

//	Returns a new 4x4 matrix representing the identity matrix.
func NewMat4Identity() (mat *Mat4) {
	mat = &Mat4{}
	mat.Identity()
	return
}

//	Returns a new 4x4 matrix representing the specified look-at matrix.
func NewMat4LookAt(bag *Bag, lookTarget, worldUp *Vec3) (mat *Mat4) {
	mat = &Mat4{}
	mat.LookAt(bag, lookTarget, worldUp)
	return
}

//	Returns a new 4x4 matrix representing the result of multiplying all values in m with v.
func NewMat4Mult1(m *Mat4, v float64) (mat *Mat4) {
	mat = &Mat4{}
	mat[0], mat[4], mat[8], mat[12] = m[0]*v, m[4]*v, m[8]*v, m[12]*v
	mat[1], mat[5], mat[9], mat[13] = m[1]*v, m[5]*v, m[9]*v, m[13]*v
	mat[2], mat[6], mat[10], mat[14] = m[2]*v, m[6]*v, m[10]*v, m[14]*v
	mat[3], mat[7], mat[11], mat[15] = m[3]*v, m[7]*v, m[11]*v, m[15]*v
	return
}

//	Returns a new 4x4 matrix that represents the result of multiplying one with two.
func NewMat4Mult4(one, two *Mat4) (mat *Mat4) {
	mat = &Mat4{}
	mat.SetFromMult4(one, two)
	return
}

//	Returns a new 4x4 matrix that represents the result of multiplying all mats with one another.
func NewMat4MultN(bag *Bag, mats ...*Mat4) (mat *Mat4) {
	mat = &Mat4{}
	mat.SetFromMultN(bag, mats...)
	return
}

//	Returns a new 4x4 matrix that represents the specified perspective-projection matrix.
func NewMat4Perspective(bag *Bag, fovY, aspect, near, far float64) (mat *Mat4) {
	mat = &Mat4{}
	mat.Perspective(bag, fovY, aspect, near, far)
	return
}

/*
func NewMat4Rotation (rad float64, axes *Vec3) *Mat4 {
	var mat = &Mat4 {}; mat.Rotation(rad, axes); return mat
}
*/

//	Returns a new 4x4 matrix that representing a rotation of rad radians around the X asis.
func NewMat4RotationX(bag *Bag, rad float64) (mat *Mat4) {
	mat = &Mat4{}
	mat.RotationX(bag, rad)
	return
}

//	Returns a new 4x4 matrix that representing a rotation of rad radians around the Y asis.
func NewMat4RotationY(bag *Bag, rad float64) (mat *Mat4) {
	mat = &Mat4{}
	mat.RotationY(bag, rad)
	return
}

//	Returns a new 4x4 matrix that representing a rotation of rad radians around the Z asis.
func NewMat4RotationZ(bag *Bag, rad float64) (mat *Mat4) {
	mat = &Mat4{}
	mat.RotationZ(bag, rad)
	return
}

//	Returns a new 4x4 matrix that represents a transformation of "scale by vec".
func NewMat4Scaling(vec *Vec3) (mat *Mat4) {
	mat = &Mat4{}
	mat.Scaling(vec)
	return
}

//	Returns a new 4x4 matrix that represents a minus b.
func NewMat4Sub(a, b *Mat4) (mat *Mat4) {
	mat = &Mat4{}
	mat[0], mat[4], mat[8], mat[12] = a[0]-b[0], a[4]-b[4], a[8]-b[8], a[12]-b[12]
	mat[1], mat[5], mat[9], mat[13] = a[1]-b[1], a[5]-b[5], a[9]-b[9], a[13]-b[13]
	mat[2], mat[6], mat[10], mat[14] = a[2]-b[2], a[6]-b[6], a[10]-b[10], a[14]-b[14]
	mat[3], mat[7], mat[11], mat[15] = a[3]-b[3], a[7]-b[7], a[11]-b[11], a[15]-b[15]
	return
}

//	Returns a new 4x4 matrix that represents a transformation of "translate by vec".
func NewMat4Translation(vec *Vec3) (mat *Mat4) {
	mat = &Mat4{}
	mat.Translation(vec)
	return
}
