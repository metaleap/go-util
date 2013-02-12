package ustr

import (
	"math"
	"strings"
	"unicode"

	ugo "github.com/metaleap/go-util"
)

//	Appends s to sl only if sl does not already contain s.
func AppendUnique(sl *[]string, s string) {
	for _, str := range *sl {
		if str == s {
			return
		}
	}
	*sl = append(*sl, s)
}

//	Does a strings.Join() on the specified string values.
func Concat(vals ...string) string {
	return strings.Join(vals, "")
}

//	Returns true if str2 is contained in str1 exactly once.
func ContainsOnce(str1, str2 string) bool {
	first, last := strings.Index(str1, str2), strings.LastIndex(str1, str2)
	if (first >= 0) && (first == last) {
		return true
	}
	return false
}

//	A simple string-similarity algorithm.
func Distance(s1, s2 string) int {
	var (
		cost, min1, min2, min3, i, j int
		d                            = make([][]int, len(s1)+1)
	)
	for i = 0; i < len(d); i++ {
		d[i] = make([]int, len(s2)+1)
		d[i][0] = i
	}
	for i = 0; i < len(d[0]); i++ {
		d[0][i] = i
	}
	for i = 1; i < len(d); i++ {
		for j = 1; j < len(d[0]); j++ {
			cost = ugo.Ifi(s1[i-1] == s2[j-1], 0, 1)
			min1 = d[i-1][j] + 1
			min2 = d[i][j-1] + 1
			min3 = d[i-1][j-1] + cost
			d[i][j] = int(math.Min(math.Min(float64(min1), float64(min2)), float64(min3)))
		}
	}
	return d[len(s1)][len(s2)]
}

//	Returns true if one and two contain the same strings, regardless of their respective slice positions.
func Equivalent(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	if len(one) > 0 {
		for _, v := range one {
			if InSliceAt(two, v) < 0 {
				return false
			}
		}
	}
	return true
}

//	Extracts all identifiers (no duplicates, ordered by occurrence) starting with prefix occurring in src.
func ExtractAllIdentifiers(src, prefix string) (identifiers []string) {
	minPos := 0
	id := ExtractFirstIdentifier(src, prefix, minPos)
	for len(id) > 0 {
		if minPos = strings.Index(src, id) + 1; !IsInSlice(identifiers, id) {
			identifiers = append(identifiers, id)
		}
		id = ExtractFirstIdentifier(src, prefix, minPos)
	}
	return
}

//	Extracts the first occurrence (at or after minPos) of an identifier starting with prefix in src.
func ExtractFirstIdentifier(src, prefix string, minPos int) (identifier string) {
	sub := src[minPos:]
	pos := strings.Index(sub, prefix)
	if pos >= 0 {
		for i, r := range sub[pos:] {
			if !(unicode.IsNumber(r) || unicode.IsLetter(r) || r == '_') {
				identifier = sub[pos : pos+i]
				break
			}
		}
	}
	return
}

//	Returns the first string in vals to match the specified predicate.
//	step: 1 to test all values. A higher value to skip n values after each test. Negative for reverse slice traversal. Or use 0 to get stuck in an infinite loop.
func First(predicate func(s string) bool, step int, vals ...string) string {
	l := len(vals)
	reverse := step < 0
	for i := ugo.Ifi(reverse, l-1, 0); ugo.Ifb(reverse, i >= 0, i < l); i += step {
		if predicate(vals[i]) {
			return vals[i]
		}
	}
	return ""
}

//	Returns the first non-empty string in vals.
func FirstNonEmpty(vals ...string) (val string) {
	// return First(func(s string) bool { return len(s) > 0 }, step, vals...)
	for _, val = range vals {
		if len(val) > 0 {
			return
		}
	}
	return
}

//	Returns true if s starts with any one of the specified prefixes.
func HasAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

//	Returns true if s ends with any one of the specified suffixes.
func HasAnySuffix(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifm(cond bool, ifTrue, ifFalse map[string]string) map[string]string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifs(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	For all seps, computes the index of first occurrence in s, then returns the smallest index.
func IndexAny(s string, seps ...string) (pos int) {
	pos = -1
	for index, sep := range seps {
		if index = strings.Index(s, sep); pos < 0 || (index >= 0 && index < pos) {
			pos = index
		}
	}
	return
}

//	Returns the position of val in vals.
func InSliceAt(vals []string, val string) int {
	for i, v := range vals {
		if v == val {
			return i
		}
	}
	return -1
}

//	Returns the position of lower-case val in lower-case vals.
func InSliceAtIgnoreCase(vals []string, val string) int {
	lv := strings.ToLower(val)
	for i, v := range vals {
		if (v == val) || (strings.ToLower(v) == lv) {
			return i
		}
	}
	return -1
}

//	Returns whether one of the specified vals is contained in slice.
func IsAnyInSlice(slice []string, vals ...string) bool {
	var big, small []string
	if len(slice) > len(vals) {
		big, small = slice, vals
	} else {
		big, small = vals, slice
	}
	for _, s1 := range big {
		for _, s2 := range small {
			if s1 == s2 {
				return true
			}
		}
	}
	return false
}

//	Returns true if str is ASCII-compatible.
func IsAscii(str string) bool {
	for _, c := range str {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

//	Returns true if val is in vals.
func IsInSlice(vals []string, val string) bool {
	return InSliceAt(vals, val) >= 0
}

//	Returns true if lower-case val is in lower-case vals.
func IsInSliceIgnoreCase(vals []string, val string) bool {
	return InSliceAtIgnoreCase(vals, val) >= 0
}

//	Returns true if all Letter-runes in s are lower-case.
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

//	Returns true if MatchSimplePattern(s, patterns...) returns a match.
func IsSimplePatternMatch(s string, patterns ...string) bool {
	return len(MatchSimplePattern(s, patterns...)) > 0
}

//	Returns true if s is in all.
func IsOneOf(s string, all ...string) bool {
	for _, a := range all {
		if s == a {
			return true
		}
	}
	return false
}

//	Returns true if all Letter-runes in s are upper-case.
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

//	Returns a string representation of s with all non-Letter-runes removed.
func LettersOnly(s string) (ret string) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			ret += string(r)
		}
	}
	return
}

//	Checks s against the specified simple-patterns and returns the first matching pattern encountered, or "" if there is no match.
//	(A "simple-pattern" is a string that can optionally have one leading or trailing (or both) asterisk ('*') wildcard.)
func MatchSimplePattern(s string, patterns ...string) string {
	for _, p := range patterns {
		if (s == p) || (strings.HasPrefix(p, "*") && strings.HasSuffix(p, "*") && strings.Contains(s, p[1:len(p)-1])) || (strings.HasPrefix(p, "*") && strings.HasSuffix(s, p[1:])) || (strings.HasSuffix(p, "*") && strings.HasPrefix(s, p[:len(p)-1])) {
			return p
		}
	}
	return ""
}

//	Returns a slice that contains the non-empty strings in vals.
func NonEmpties(breakAtFirstEmpty bool, vals ...string) (slice []string) {
	for _, s := range vals {
		if len(s) > 0 {
			slice = append(slice, s)
		} else if breakAtFirstEmpty {
			break
		}
	}
	return
}

//	A most simplistic (not linguistically-correct) English-language pluralizer that may be useful for code or doc generation.
//	If s ends with "s", only appends "es": bus -> buses, mess -> messes
//	If s ends with "y" (but not "ay", "ey", "oy", "uy" or "iy"), removes "y" and appends "ies": autonomy -> autonomies, dictionary -> dictionaries etc.
//	Otherwise, appends "s".
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	if (len(s) > 1) && strings.HasSuffix(s, "y") && !IsOneOf(s[(len(s)-2):], "ay", "ey", "oy", "uy", "iy") {
		return s[0:(len(s)-1)] + "ies"
	}
	return s + "s"
}

//	Prepends prefix + sep to v only if prefix isn't empty.
func PrefixWithSep(prefix, sep, v string) string {
	if len(prefix) > 0 {
		return prefix + sep + v
	}
	return v
}

//	Prepends p to s only if s doesn't already have that prefix.
func PrependIf(s, p string) string {
	if strings.HasPrefix(s, p) {
		return s
	}
	return p + s
}

//	Replaces in str all occurrences of all repls map keys with their associated (mapped) value.
func Replace(str string, repls map[string]string) string {
	for k, v := range repls {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

//	Creates a Pascal-cased "identifier" version of the specified string.
func SafeIdentifier(s string) (ret string) {
	var isL, isD, last bool
	for i, r := range s {
		if isL, isD = unicode.IsLetter(r), unicode.IsDigit(r); isL || isD || ((r == '_') && (i == 0)) {
			if (i > 0) && (isL != last) {
				ret += " "
			}
			ret += string(r)
		} else {
			ret += " "
		}
		last = isL
	}
	words := Split(strings.Title(ret), " ")
	for i, w := range words {
		if (len(w) > 1) && IsUpper(w) {
			words[i] = strings.Title(strings.ToLower(w))
		}
	}
	ret = strings.Join(words, "")
	return
}

//	Returns an empty slice is v is emtpy, otherwise like strings.Split()
func Split(v, s string) (sl []string) {
	if len(v) > 0 {
		sl = strings.Split(v, s)
	}
	return
}

//	Strips prefix off val if possible.
func StripPrefix(val, prefix string) string {
	for strings.HasPrefix(val, prefix) {
		val = val[len(prefix):]
	}
	return val
}

//	Strips suffix off val if possible.
func StripSuffix(val, suffix string) string {
	for strings.HasSuffix(val, suffix) {
		val = val[:len(val)-len(suffix)]
	}
	return val
}

/*
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
*/

//	Returns the lower-case representation of s only if it is currently fully upper-case as per IsUpper().
func ToLowerIfUpper(s string) string {
	if IsUpper(s) {
		return strings.ToLower(s)
	}
	return s
}

//	Returns the upper-case representation of s only if it is currently fully lower-case as per IsLower().
func ToUpperIfLower(s string) string {
	if IsLower(s) {
		return strings.ToUpper(s)
	}
	return s
}

//	Removes all withoutVals from slice.
func Without(slice []string, keepOrder bool, withoutVals ...string) []string {
	if len(withoutVals) > 0 {
		for _, w := range withoutVals {
			for pos := InSliceAt(slice, w); pos >= 0; pos = InSliceAt(slice, w) {
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
