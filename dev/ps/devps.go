package udevps

import (
	"strings"
)

var (
	NotImplErr        func(string, string, interface{}) error
	StrReplUnsanitize *strings.Replacer
)
