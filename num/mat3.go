package num

var (
	Mat3Identity = NewMat3Identity()
)

type TMat3 [9]float64

func NewMat3Identity () *TMat3 {
	var mat = &TMat3 {}; mat.Identity(); return mat
}

func (me *TMat3) Identity () {
	me[0], me[3], me[6] = 1, 0, 0
	me[1], me[4], me[7] = 0, 1, 0
	me[2], me[5], me[8] = 0, 0, 1
}

func (me *TMat3) Transpose () {
	var a01, a02, a12 = me[1], me[2], me[5]
	me[1] = me[3]
	me[2] = me[6]
	me[3] = a01
	me[5] = me[7]
	me[6] = a02
	me[7] = a12
}
