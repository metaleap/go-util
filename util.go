package util

import (
	"reflect"
)

type FAnyToAny func (src interface{}) interface{}

func Ifb (cond, ifTrue, ifFalse bool) bool {
	if cond { return ifTrue }
	return ifFalse
}

func Ifd (cond bool, ifTrue, ifFalse float64) float64 {
	if cond { return ifTrue }
	return ifFalse
}

func Ifi (cond bool, ifTrue, ifFalse int) int {
	if cond { return ifTrue }
	return ifFalse
}

func Ifi16 (cond bool, ifTrue, ifFalse int16) int16 {
	if cond { return ifTrue }
	return ifFalse
}

func Ifi32 (cond bool, ifTrue, ifFalse int32) int32 {
	if cond { return ifTrue }
	return ifFalse
}

func Ifi64 (cond bool, ifTrue, ifFalse int64) int64 {
	if cond { return ifTrue }
	return ifFalse
}

func Ifs (cond bool, ifTrue string, ifFalse string) string {
	if cond { return ifTrue }
	return ifFalse
}

func Ifu32 (cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond { return ifTrue }
	return ifFalse
}

func IsFloat64s (any interface{}) bool {
	if any != nil {
		if _, isT := any.([]float64); isT {
			return true
		}
	}
	return false
}

func PtrVal (ptr interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(ptr)).Interface()
}
