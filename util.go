package util

import (
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type AnyToAny func(src interface{}) interface{}

func BaseCodePath (gitHubName string, subDirNames ... string) string {
	return filepath.Join(append([]string { os.Getenv("GOPATH"), "src", "github.com", gitHubName }, subDirNames ...) ...)
}

func Ifb(cond, ifTrue, ifFalse bool) bool {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifd(cond bool, ifTrue, ifFalse float64) float64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifi(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifi16(cond bool, ifTrue, ifFalse int16) int16 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifi32(cond bool, ifTrue, ifFalse int32) int32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifi64(cond bool, ifTrue, ifFalse int64) int64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifs(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifu32(cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func IsFloat64s(any interface{}) bool {
	if any != nil {
		if _, isT := any.([]float64); isT {
			return true
		}
	}
	return false
}

func ParseVersion(vString string) []int {
	//	3.3.0 - Build 8.15.10.2761.
	var nums = []int{}
	var i uint64
	var pos int
	var err error
	var parts = split(vString, ".")
	for _, p := range parts {
		if pos = strings.Index(p, " "); pos > 0 { p = p[ : pos] }
		if i, err = strconv.ParseUint(p, 10, 8); err == nil {
			nums = append(nums, int(i))
		} else {
			break
		}
	}
	return nums
}

func PtrVal(ptr interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(ptr)).Interface()
}

func split (v, s string) (sl []string) {
	if len(v) > 0 { sl = strings.Split(v, s) }; return
}
