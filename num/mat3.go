package num

var (
	//	Represents the 3x3 identity matrix.
	Mat3Identity = NewMat3Identity()
)

//	Represents a 3x3 matrix.
type Mat3 [9]float64

//	Returns a new 3x3 identity matrix.
func NewMat3Identity () *Mat3 {
	var mat = &Mat3 {}; mat.Identity(); return mat
}

//	Sets this 3x3 matrix to the 3x3 identity matrix.
func (me *Mat3) Identity () {
	me[0], me[3], me[6] = 1, 0, 0
	me[1], me[4], me[7] = 0, 1, 0
	me[2], me[5], me[8] = 0, 0, 1
}

//	Transposes this 3x3 matrix.
func (me *Mat3) Transpose () {
	var a01, a02, a12 = me[1], me[2], me[5]
	me[1] = me[3]
	me[2] = me[6]
	me[3] = a01
	me[5] = me[7]
	me[6] = a02
	me[7] = a12
}
