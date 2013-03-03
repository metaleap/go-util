package uslice

//#begin-gt -gen.gt N:Bool T:bool

//	Appends v to sl only if sl does not already contain v.
func BoolAppendUnique(ref *[]bool, v bool) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Returns the position of val in slice.
func BoolAt(slice []bool, val bool) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func BoolEnsureCap(ref *[]bool, capacity int) {
	if cap(*ref) < capacity {
		BoolSetCap(ref, capacity)
	}
}

func BoolEnsureLen(ref *[]bool, length int) {
	if len(*ref) < length {
		BoolSetLen(ref, length)
	}
}

//	Returns true if one and two only contain identical values, regardless of ordering.
func BoolEquivalent(one, two []bool) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if BoolAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns true if val is in slice.
func BoolHas(slice []bool, val bool) bool {
	return BoolAt(slice, val) >= 0
}

//	Returns whether one of the specified vals is contained in slice.
func BoolHasAny(slice []bool, vals ...bool) bool {
	for _, v1 := range vals {
		for _, v2 := range slice {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func BoolRemove(ref *[]bool, v bool, all bool) {
	for i := 0; i < len(*ref); i++ {
		if (*ref)[i] == v {
			before, after := (*ref)[:i], (*ref)[i+1:]
			*ref = append(before, after...)
			if !all {
				break
			}
		}
	}
}

func BoolSetCap(ref *[]bool, capacity int) {
	nu := make([]bool, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

func BoolSetLen(ref *[]bool, length int) {
	nu := make([]bool, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all withoutVals from slice.
func BoolWithout(slice []bool, keepOrder bool, withoutVals ...bool) []bool {
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos := BoolAt(slice, w); pos >= 0; pos = BoolAt(slice, w) {
				if keepOrder {
					slice = append(slice[:pos], slice[pos+1:]...)
				} else {
					slice[pos] = slice[len(slice)-1]
					slice = slice[:len(slice)-1]
				}
			}
		}
	}
	return slice
}

//#end-gt
