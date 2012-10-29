package str

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	util "github.com/go3d/go-util"
)

func Concat (vals ... string) string {
	return strings.Join(vals, "")
}

func ContainsOnce (str1, str2 string) bool {
	var first, last = strings.Index(str1, str2), strings.LastIndex(str1, str2)
	if (first >= 0) && (first == last) {
		return true
	}
	return false
}

func Distance (s1, s2 string) int {
	var cost, min1, min2, min3, i, j int
	var d = make([][]int, len(s1) + 1)
	for i = 0; i < len(d); i++ {
		d[i] = make([]int, len(s2) + 1)
		d[i][0] = i
	}
	for i = 0; i < len(d[0]); i++ {
		d[0][i] = i
	}
	for i = 1; i < len(d); i++ {
		for j = 1; j < len(d[0]); j++ {
			cost = util.Ifi(s1[i - 1] == s2[j - 1], 0, 1)
			min1 = d[i - 1][j] + 1
			min2 = d[i][j - 1] + 1
			min3 = d[i - 1][j - 1] + cost
			d[i][j] = int(math.Min(math.Min(float64(min1), float64(min2)), float64(min3)))
		}
	}
	return d[len(s1)][len(s2)]
}

func First (fun func (s string) bool, step int, vals ... string) string {
	var l = len(vals)
	var reverse = step < 0
	for i := util.Ifi(reverse, l - 1, 0); util.Ifb(reverse, i >= 0, i < l); i += step {
		if fun(vals[i]) {
			return vals[i]
		}
	}
	return ""
}

func FirstNonEmpty (step int, vals ... string) string {
	return First(func (s string) bool {
		return len(s) > 0
	}, step, vals ...)
}

func ForEach (fun func (i int, s string), vals ... string) {
	for i, s := range vals {
		fun(i, s)
	}
}

func InSliceAt (vals []string, val string) int {
	for i, v := range vals { if v == val { return i } }
	return -1
}

func InSliceAtIgnoreCase (vals []string, val string) int {
	var lv = strings.ToLower(val)
	for i, v := range vals {
		if (v == val) || (strings.ToLower(v) == lv) {
			return i
		}
	}
	return -1
}

func IsAscii (str string) bool {
	for _, c := range str {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func IsInSlice (vals []string, val string) bool {
	return InSliceAt(vals, val) >= 0
}

func IsInSliceIgnoreCase (vals []string, val string) bool {
	return InSliceAtIgnoreCase(vals, val) >= 0
}

func NonEmpties (breakAtFirstEmpty bool, vals ... string) []string {
	var slice = []string {}
	for _, s := range vals {
		if (len(s) > 0) {
			slice = append(slice, s)
		} else if breakAtFirstEmpty {
			break
		}
	}
	return slice
}

func Split (str string, sep string) []string {
	var spl []string = nil
	if len(str) > 0 {
		spl = strings.Split(str, sep)
		/*
		ls := len(sep)
		for si, sv := range spl {
			if len(sv) >= ls {
				for sv[0:ls] == sep {
					sv = sv[ls:]
				}
				for sv[len(sv) - ls:] == sep {
					sv = sv[:len(sv) - ls]
				}
				spl[si] = sv
			}
		}
		return Without(spl, true, sep)
		*/
	}
	return spl
}

func Title (str string) string {
	return strings.Title(strings.ToLower(str))
}

func ToFloat32 (str string) float32 {
	var f, err = strconv.ParseFloat(str, 32)
	if err == nil {
		return float32(f)
	}
	return 0.0
}

func ToFloat64 (str string) float64 {
	var f, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return f
	}
	return 0.0
}

func ToFloat64s (strs ... string) []float64 {
	var f = make([]float64, len(strs))
	for i, s := range strs {
		f[i] = ToFloat64(s)
	}
	return f
}

func ToInt (str string) int {
	var i, err = strconv.Atoi(str)
	if err == nil {
		return i
	}
	return 0
}

func ToString (any interface{}, nilVal string) string {
	if any == nil {
		return nilVal
	}
	if s, isS := any.(string); isS {
		return s
	}
	if f, isF := any.(fmt.Stringer); isF {
		return f.String()
	}
	return fmt.Sprintf("%v", any)
}

func ToStrings (any interface{}) []string {
	if sl, isSl := any.([]string); isSl {
		return sl
	}
	return nil
}

func TrimLeft (val, trim string) string {
	if strings.Index(val, trim) == 0 {
		return val[len(trim):]
	}
	return val
}

func Without (slice []string, keepOrder bool, withoutVals ... string) []string {
	var pos int
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos = InSliceAt(slice, w); pos >= 0; pos = InSliceAt(slice, w) {
				if keepOrder {
					slice = append(slice[:pos], slice[pos + 1:] ...)
				} else {
					slice[pos] = slice[len(slice) - 1]
					slice = slice[:len(slice) - 1]
				}
			}
		}
	}
	return slice
}
