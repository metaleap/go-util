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
	patterns     []matcherPattern
	hasWildcards bool
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
		if len(patts[i].contains) > 0 || len(patts[i].prefix) > 0 || len(patts[i].suffix) > 0 {
			me.hasWildcards = true
		}
	}
	me.patterns = append(me.patterns, patts...)
}

//	Returns whether any of the patterns specified for this Matcher contains a *-wildcard.
func (me *Matcher) HasWildcardPatterns() bool {
	return me.hasWildcards
}

//	Matches s against all patterns in me.
func (me *Matcher) IsMatch(s string) bool {
	for i := 0; i < len(me.patterns); i++ {
		if s == me.patterns[i].pattern {
			return true
		}
		if me.hasWildcards {
			if len(me.patterns[i].prefix) > 0 && strings.HasPrefix(s, me.patterns[i].prefix) {
				return true
			}
			if len(me.patterns[i].suffix) > 0 && strings.HasSuffix(s, me.patterns[i].suffix) {
				return true
			}
			if len(me.patterns[i].contains) > 0 && strings.Contains(s, me.patterns[i].contains) {
				return true
			}
		}
	}
	return false
}
