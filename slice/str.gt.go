package uslice

import (
	"strings"
)

//	Returns the position of lower-case val in lower-case vals.
func StrAtIgnoreCase(vals []string, val string) int {
	lv := strings.ToLower(val)
	for i, v := range vals {
		if (v == val) || (strings.ToLower(v) == lv) {
			return i
		}
	}
	return -1
}

//	Returns true if lower-case val is in lower-case vals.
func StrHasIgnoreCase(vals []string, val string) bool {
	return StrAtIgnoreCase(vals, val) >= 0
}

//#begin-gt -gen.gt N:Str T:string

//	Appends v to sl only if sl does not already contain v.
func StrAppendUnique(ref *[]string, v string) {
	for _, sv := range *ref {
		if sv == v {
			return
		}
	}
	*ref = append(*ref, v)
}

//	Returns the position of val in slice.
func StrAt(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Returns true if one and two only contain identical values, regardless of ordering.
func StrEquivalent(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if StrAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Returns true if val is in slice.
func StrHas(slice []string, val string) bool {
	return StrAt(slice, val) >= 0
}

//	Returns whether one of the specified vals is contained in slice.
func StrHasAny(slice []string, vals ...string) bool {
	for _, v1 := range vals {
		for _, v2 := range slice {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func StrRemove(ref *[]string, v string, all bool) {
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

//	Removes all withoutVals from slice.
func StrWithout(slice []string, keepOrder bool, withoutVals ...string) []string {
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos := StrAt(slice, w); pos >= 0; pos = StrAt(slice, w) {
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
