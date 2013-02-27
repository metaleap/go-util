package ustr

import (
	"strings"
)

type matcherPattern struct {
	pattern, prefix, suffix, contains string
}

//	Matches a string against "simple-patterns".
//	Simple-patterns are strings that can have *-wildcards only at the beginning, at the end, or both.
type Matcher struct {
	patterns []matcherPattern
}

//	Adds the specified pattern to me.
func (me *Matcher) AddPatterns(patterns ...string) {
	var s string
	patts := make([]matcherPattern, len(patterns))
	for i := 0; i < len(patterns); i++ {
		s = patterns[i]
		patts[i].pattern = s
		if strings.HasPrefix(s, "*") && strings.HasSuffix(s, "*") {
			patts[i].contains = s[1 : len(s)-1]
		} else if strings.HasPrefix(s, "*") {
			patts[i].suffix = s[1:]
		} else if strings.HasSuffix(s, "*") {
			patts[i].prefix = s[:len(s)-1]
		}
	}
	me.patterns = append(me.patterns, patts...)
}

//	Returns whether any of the patterns specified for this Matcher contains a *-wildcard.
func (me *Matcher) HasWildcardPatterns() bool {
	for i := 0; i < len(me.patterns); i++ {
		if len(me.patterns[i].contains) > 0 || len(me.patterns[i].prefix) > 0 || len(me.patterns[i].suffix) > 0 {
			return true
		}
	}
	return false
}

//	Convenience short-hand that informs whether Match(s) returns a non-empty result value.
func (me *Matcher) IsMatch(s string) bool {
	return len(me.Match(s)) > 0
}

//	Matches s against all patterns in me.
//	Returns the first pattern (minus wildcards) that matches s, or "" if there is no match.
func (me *Matcher) Match(s string) string {
	for i := 0; i < len(me.patterns); i++ {
		if s == me.patterns[i].pattern {
			return me.patterns[i].pattern
		}
		if len(me.patterns[i].prefix) > 0 && strings.HasPrefix(s, me.patterns[i].prefix) {
			return me.patterns[i].prefix
		}
		if len(me.patterns[i].suffix) > 0 && strings.HasSuffix(s, me.patterns[i].suffix) {
			return me.patterns[i].suffix
		}
		if len(me.patterns[i].contains) > 0 && strings.Contains(s, me.patterns[i].contains) {
			return me.patterns[i].contains
		}
	}
	return ""
}
