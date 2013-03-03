package uslice

//#begin-gt -gen.gt N:Int T:int

//	Appends v to sl only if sl does not already contain v.
func IntAppendUnique(ref *[]int, v int) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Returns the position of val in slice.
func IntAt(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func IntEnsureCap(ref *[]int, capacity int) {
	if cap(*ref) < capacity {
		IntSetCap(ref, capacity)
	}
}

func IntEnsureLen(ref *[]int, length int) {
	if len(*ref) < length {
		IntSetLen(ref, length)
	}
}

//	Returns true if one and two only contain identical values, regardless of ordering.
func IntEquivalent(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if IntAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns true if val is in slice.
func IntHas(slice []int, val int) bool {
	return IntAt(slice, val) >= 0
}

//	Returns whether one of the specified vals is contained in slice.
func IntHasAny(slice []int, vals ...int) bool {
	for _, v1 := range vals {
		for _, v2 := range slice {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func IntRemove(ref *[]int, v int, all bool) {
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

func IntSetCap(ref *[]int, capacity int) {
	nu := make([]int, len(*ref), capacity)
	copy(nu, *ref)
	*ref = nu
}

func IntSetLen(ref *[]int, length int) {
	nu := make([]int, length)
	copy(nu, *ref)
	*ref = nu
}

//	Removes all withoutVals from slice.
func IntWithout(slice []int, keepOrder bool, withoutVals ...int) []int {
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos := IntAt(slice, w); pos >= 0; pos = IntAt(slice, w) {
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
