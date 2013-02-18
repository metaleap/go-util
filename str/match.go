package ustr

import (
	"strings"
)

type matcherPattern struct {
	pattern, prefix, suffix, contains string
}

type Matcher struct {
	patterns []*matcherPattern
}

func NewMatcher(patterns ...string) (me *Matcher) {
	me = &Matcher{}
	for _, s := range patterns {
		me.AddPattern(s)
	}
	return
}

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

func (me *Matcher) HasWildcardPatterns() bool {
	for _, mp := range me.patterns {
		if len(mp.contains) > 0 || len(mp.prefix) > 0 || len(mp.suffix) > 0 {
			return true
		}
	}
	return false
}

func (me *Matcher) IsMatch(s string) bool {
	return len(me.Match(s)) > 0
}

func (me *Matcher) Match(s string) string {
	for _, mp := range me.patterns {
		if len(mp.prefix) > 0 && strings.HasPrefix(s, mp.prefix) {
			return mp.prefix
		}
		if len(mp.suffix) > 0 && strings.HasSuffix(s, mp.suffix) {
			return mp.suffix
		}
		if len(mp.contains) > 0 && strings.Contains(s, mp.contains) {
			return mp.contains
		}
		if s == mp.pattern {
			return mp.pattern
		}
	}
	return ""
}
