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
	patterns []*matcherPattern
}

//	Initializes and returns a new Matcher with the specified patterns.
func NewMatcher(patterns ...string) (me *Matcher) {
	me = &Matcher{}
	for _, s := range patterns {
		me.AddPattern(s)
	}
	return
}

//	Adds the specified pattern to me.
func (me *Matcher) AddPattern(s string) {
	mp := &matcherPattern{pattern: s}
	if strings.HasPrefix(s, "*") && strings.HasSuffix(s, "*") {
		mp.contains = s[1 : len(s)-1]
	} else if strings.HasPrefix(s, "*") {
		mp.suffix = s[1:]
	} else if strings.HasSuffix(s, "*") {
		mp.prefix = s[:len(s)-1]
	}
	me.patterns = append(me.patterns, mp)
}

//	Returns whether any of the patterns specified for this Matcher contains a *-wildcard.
func (me *Matcher) HasWildcardPatterns() bool {
	for _, mp := range me.patterns {
		if len(mp.contains) > 0 || len(mp.prefix) > 0 || len(mp.suffix) > 0 {
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
	for _, mp := range me.patterns {
		if s == mp.pattern {
			return mp.pattern
		}
		if len(mp.prefix) > 0 && strings.HasPrefix(s, mp.prefix) {
			return mp.prefix
		}
		if len(mp.suffix) > 0 && strings.HasSuffix(s, mp.suffix) {
			return mp.suffix
		}
		if len(mp.contains) > 0 && strings.Contains(s, mp.contains) {
			return mp.contains
		}
	}
	return ""
}
