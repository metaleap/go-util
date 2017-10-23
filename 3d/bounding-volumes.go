package u3d


type Bounds struct {
	Sphere float64
	AaBox  AaBb
}

func (me *Bounds) Clear() {
	me.Sphere = 0
	me.AaBox.Clear()
}

func (me *Bounds) Reset() {
	me.Sphere = 0
	me.AaBox.ResetMinMax()
}
