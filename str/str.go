package str

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	util "github.com/metaleap/go-util"
)

func Concat (vals ... string) string {
	return strings.Join(vals, "")
}

func ContainsOnce (str1, str2 string) bool {
	var first, last = strings.Index(str1, str2), strings.LastIndex(str1, str2)
	if (first >= 0) && (first == last) { return true }
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
		if fun(vals[i]) { return vals[i] }
	}
	return ""
}

func FirstNonEmpty (step int, vals ... string) string {
	return First(func (s string) bool { return len(s) > 0 }, step, vals ...)
}

func ForEach (fun func (i int, s string), vals ... string) {
	for i, s := range vals { fun(i, s) }
}

func InSliceAt (vals []string, val string) int {
	for i, v := range vals { if v == val { return i } }
	return -1
}

func InSliceAtIgnoreCase (vals []string, val string) int {
	var lv = strings.ToLower(val)
	for i, v := range vals { if (v == val) || (strings.ToLower(v) == lv) { return i } }
	return -1
}

func IsAscii (str string) bool {
	for _, c := range str { if c > unicode.MaxASCII { return false } }
	return true
}

func IsInSlice (vals []string, val string) bool {
	return InSliceAt(vals, val) >= 0
}

func IsInSliceIgnoreCase (vals []string, val string) bool {
	return InSliceAtIgnoreCase(vals, val) >= 0
}

func IsLower (s string) bool {
	for _, r := range s { if unicode.IsLetter(r) && !unicode.IsLower(r) { return false } }
	return true
}

func IsUpper (s string) bool {
	for _, r := range s { if unicode.IsLetter(r) && !unicode.IsUpper(r) { return false } }
	return true
}

func LettersOnly (s string) (ret string) {
	for _, r := range s { if unicode.IsLetter(r) { ret += string(r) } }
	return
}

func NonEmpties (breakAtFirstEmpty bool, vals ... string) []string {
	var slice = []string {}
	for _, s := range vals {
		if (len(s) > 0) { slice = append(slice, s) } else if breakAtFirstEmpty { break }
	}
	return slice
}

func PrefixWithSep (prefix, sep, v string) string {
	if len(prefix) > 0 { return prefix + sep + v }
	return v
}

func PrependIf (s, p string) string {
	if strings.HasPrefix(s, p) { return s }
	return p + s
}

func Replace (str string, repls map[string]string) string {
	for k, v := range repls { str = strings.Replace(str, k, v, -1) }
	return str
}

func RuneAt (str string, pos int) rune {
	for i, r := range str { if i == pos { return r } }
	return 0
}

func SafeIdentifier (s string) (ret string) {
	var words []string
	var isL, isD, last bool
	for i, r := range s {
		if isL, isD = unicode.IsLetter(r), unicode.IsDigit(r); isL || isD || ((r == '_') && (i == 0)) {
			if (i > 0) && (isL != last) { ret += " " }
			ret += string(r)
		} else {
			ret += " "
		}
		last = isL
	}
	words = Split(strings.Title(ret), " ")
	for i, w := range words {
		if (len(w) > 1) && IsUpper(w) { words[i] = strings.Title(strings.ToLower(w)) }
	}
	ret = strings.Join(words, "") // if !unicode.IsLetter(RuneAt(ret, 0)) { ret = safePrefix + ret }
	return
}

func Split (v, s string) (sl []string) {
	if len(v) > 0 { sl = strings.Split(v, s) }; return
}

func StripPrefix (val, prefix string) string {
	for strings.HasPrefix(val, prefix) { val = val[len(prefix) :] }
	return val
}

func StripSuffix (val, suffix string) string {
	for strings.HasSuffix(val, suffix) { val = val[: len(val) - len(suffix)] }
	return val
}

func ToFloat32 (str string) float32 {
	var f, err = strconv.ParseFloat(str, 32)
	if err == nil { return float32(f) }
	return 0.0
}

func ToFloat64 (str string) float64 {
	var f, err = strconv.ParseFloat(str, 64)
	if err == nil { return f }
	return 0.0
}

func ToFloat64s (strs ... string) []float64 {
	var f = make([]float64, len(strs))
	for i, s := range strs { f[i] = ToFloat64(s) }
	return f
}

func ToInt (str string) int {
	var i, err = strconv.Atoi(str)
	if err == nil { return i }
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

func ToLowerIfUpper (s string) string {
	if IsUpper(s) { return strings.ToLower(s) }; return s
}

func ToUpperIfLower (s string) string {
	if IsLower(s) { return strings.ToUpper(s) }; return s
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
