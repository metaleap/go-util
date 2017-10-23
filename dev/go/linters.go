package udevgo

import (
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/str"
)

var (
	GolintIgnoreSubstrings = []string{
		" should have comment ",
		"if block ends with a return statement, so drop this else and outdent its block",
		"ALL_CAPS",
		"underscore",
		"CamelCase",
		"should be of the form \"",
		"should omit 2nd value from range; this loop is equivalent to ",
		"don't use generic names",
	}
)

func LintCheck(cmdname string, pkgimppath string) (msgs udev.SrcMsgs) {
	reline := func(ln string) string {
		if strings.HasPrefix(ln, pkgimppath+": ") {
			return udev.LnRelify(ln[len(pkgimppath)+2:])
		}
		return ""
	}
	for _, srcref := range udev.CmdExecOnSrc(false, reline, cmdname, pkgimppath) {
		if strings.HasPrefix(srcref.Msg, pkgimppath+".") {
			srcref.Msg = srcref.Msg[len(pkgimppath)+1:]
		}
		if cmdname != "aligncheck" {
			srcref.Msg = "unused & unexported: " + srcref.Msg
		}
		msgs = append(msgs, srcref)
	}
	return
}

func LintIneffAssign(dirrelpath string) (msgs udev.SrcMsgs) {
	msgs = append(msgs, udev.CmdExecOnSrc(false, udev.LnRelify, "ineffassign", "-n", dirrelpath)...)
	return
}

func LintMDempsky(cmdname string, pkgimppath string) (msgs udev.SrcMsgs) {
	msgs = append(msgs, udev.CmdExecOnSrc(false, udev.LnRelify, cmdname, pkgimppath)...)
	return
}

func LintMvDan(cmdname string, pkgimppath string) udev.SrcMsgs {
	reline := func(ln string) string {
		if rln := udev.LnRelify(ln); len(rln) > 0 {
			return rln
		}
		return ln
	}
	return udev.CmdExecOnSrc(false, reline, cmdname, pkgimppath)
}

func LintHonnef(cmdname string, pkgimppath string) (msgs udev.SrcMsgs) {
	msgs = append(msgs, udev.CmdExecOnSrc(false, nil, cmdname, pkgimppath)...)
	return
}

func LintErrcheck(pkgimppath string) (msgs udev.SrcMsgs) {
	for _, m := range udev.CmdExecOnSrc(false, nil, "errcheck", pkgimppath) {
		m.Msg = "Ignores a returned `error`: " + m.Msg
		msgs = append(msgs, m)
	}
	return
}

func LintGolint(dirrelpath string) (msgs udev.SrcMsgs) {
	censored := func(msg string) bool {
		for _, s := range GolintIgnoreSubstrings {
			if ustr.Has(msg, s) {
				return true
			}
		}
		return false
	}
	for _, msg := range udev.CmdExecOnSrc(false, nil, "golint", dirrelpath) {
		if !censored(msg.Msg) {
			msgs = append(msgs, msg)
		}
	}
	return
}

func LintGoVet(dirrelpath string) udev.SrcMsgs {
	reline := func(ln string) string {
		if strings.HasPrefix(ln, "vet: ") {
			return ""
		} else {
			return ln
		}
	}
	return udev.CmdExecOnSrc(true, reline, "go", "tool", "vet", "-shadow=true", "-shadowstrict", "-all", dirrelpath)
}
