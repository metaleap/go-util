package num

//	Represents a 3x3 matrix.
type Mat3 [9]float64

var (
	m3id Mat3
)

func init() {
	m3id[0], m3id[3], m3id[6] = 1, 0, 0
	m3id[1], m3id[4], m3id[7] = 0, 1, 0
	m3id[2], m3id[5], m3id[8] = 0, 0, 1
}

//	Sets this 3x3 matrix to the 3x3 identity matrix.
func (me *Mat3) Identity() {
	*me = m3id
}

//	Transposes this 3x3 matrix.
func (me *Mat3) Transpose() {
	// a01, a02, a12 := me[1], me[2], me[5]
	// me[1] = me[3]
	// me[2] = me[6]
	// me[3] = a01
	// me[5] = me[7]
	// me[6] = a02
	// me[7] = a12
	me[1], me[2], me[3], me[5], me[6], me[7] = me[3], me[6], me[1], me[7], me[2], me[5]
}

//	Calls the Identity() method on all specified mats.
func Mat3Identities(mats ...*Mat3) {
	for _, mat := range mats {
		mat.Identity()
	}
}

//	Returns a new 3x3 identity matrix.
func NewMat3Identity() (mat *Mat3) {
	mat = &Mat3{}
	mat.Identity()
	return
}
