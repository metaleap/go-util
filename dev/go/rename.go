package udevgo

import (
	"fmt"
	"strings"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

func Gorename(cmdname string, filepath string, offset int, newname string, eol string) (fileedits udev.SrcMsgs, err error) {
	cmdargs := []string{"-d", "-to", newname, "-offset", fmt.Sprintf("%s:#%d", filepath, offset)}
	var renout, renerr string
	if len(cmdname) == 0 {
		cmdname = "gorename"
	}
	if renout, renerr, err = urun.CmdExec(cmdname, cmdargs...); err != nil {
		return
	} else if renout = strings.TrimSpace(renout); renerr != "" && renout == "" {
		if join, re, msgs := " — ", "", udev.SrcMsgsFromLns(strings.Split(renerr, "\n")); len(msgs) > 0 {
			for _, m := range msgs {
				if m.Ref != "" && ufs.FileExists(m.Ref) {
					re += m.Msg + join
				}
			}
			if re != "" {
				renerr = re[:len(re)-len(join)]
			}
		}
		err = umisc.E(renerr)
		return
	}
	i := ustr.Idx(renout, "--- ")
	if i < 0 {
		err = umisc.E(renout)
		return
	}
	renout = renout[i+4:]
	rendiffs := uslice.StrMap(ustr.Split(renout, "--- "), strings.TrimSpace)
	if len(rendiffs) == 0 {
		return nil, umisc.E("Renaming aborted: no diffs could be obtained.")
	}

	for _, rendiff := range rendiffs {
		if i = ustr.Idx(rendiff, "\t"); i <= 0 {
			return nil, umisc.E("Renaming aborted: could not detect file path in diffs.")
		}
		if ffp := rendiff[:i]; !ufs.FileExists(ffp) {
			return nil, umisc.E("Renaming aborted: bad absolute file path `" + ffp + "` in diffs.")
		} else {
			if i = ustr.Idx(rendiff, "@@ -"); i <= 0 {
				return nil, umisc.E("Renaming aborted: `@@ -` expected.")
			} else {
				for _, hunkchunk := range ustr.Split(rendiff[i+4:], "@@ -") {
					if lns := ustr.Split(hunkchunk, "\n"); len(lns) > 0 {
						i = ustr.Idx(lns[0], ",")
						lb := ustr.ToInt(lns[0][:i])
						s := lns[0][i+1:]
						ll := ustr.ToInt(s[:ustr.Idx(s, " +")])
						if lb == 0 || ll == 0 {
							return nil, umisc.E("Renaming aborted: diffs contained invalid or unparsable line hints.")
						} else {
							fed := &udev.SrcMsg{Ref: ffp, Pos1Ln: lb - 1, Pos1Ch: 0, Pos2Ln: lb - 1 + ll, Pos2Ch: 0}
							for _, ln := range lns[1:] {
								if ustr.Pref(ln, " ") || ustr.Pref(ln, "+") {
									fed.Msg = fed.Msg + ln[1:] + eol
								}
							}
							fileedits = append(fileedits, fed)
						}
					} else {
						return nil, umisc.E("Renaming aborted: expected something between one `@@ -` and the next.")
					}
				}
				if len(fileedits) == 0 {
					err = umisc.E("Renaming aborted: a diff without effective edits.")
				}
			}
		}
	}
	return
}
