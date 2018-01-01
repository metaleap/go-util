package udev

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type SrcMsg struct {
	Flag   int                    `json:",omitempty"`
	Ref    string                 `json:",omitempty"`
	Msg    string                 `json:",omitempty"`
	Misc   string                 `json:",omitempty"`
	Pos1Ln int                    `json:",omitempty"`
	Pos1Ch int                    `json:",omitempty"`
	Pos2Ln int                    `json:",omitempty"`
	Pos2Ch int                    `json:",omitempty"`
	Data   map[string]interface{} `json:",omitempty"`
}

type SrcMsgs []*SrcMsg

func (me SrcMsgs) Len() int           { return len(me) }
func (me SrcMsgs) Swap(i, j int)      { me[i], me[j] = me[j], me[i] }
func (me SrcMsgs) Less(i, j int) bool { return me[i].Msg < me[j].Msg }

var (
	SrcDir string
)

func LnRelify(ln string) string {
	if ufs.PathPrefix(ln, SrcDir) {
		return strings.TrimLeft(ln[len(SrcDir):], "/\\")
	}
	return ""
}

func CmdExecOnSrcIn(dir string, inclstderr bool, reline func(string) string, cmdname string, cmdargs ...string) SrcMsgs {
	var output []byte
	cmd := exec.Command(cmdname, cmdargs...)
	cmd.Dir = dir
	if inclstderr {
		output, _ = cmd.CombinedOutput()
	} else {
		output, _ = cmd.Output()
	}
	cmdout := strings.TrimSpace(string(output))
	msgs := SrcMsgsFromLns(uslice.StrMap(ustr.Split(cmdout, "\n"), reline))
	if len(msgs) == 0 && cmdout != "" && dir == "" && inclstderr && reline == nil {
		msgs = append(msgs, &SrcMsg{Msg: cmdout, Pos1Ch: 1, Pos1Ln: 1})
	}
	return msgs
}

func CmdExecOnStdIn(stdin string, dir string, reline func(string) string, cmdname string, cmdargs ...string) (SrcMsgs, error) {
	cmdout, cmderr, err := urun.CmdExecStdin(stdin, dir, cmdname, cmdargs...)
	if len(cmderr) > 0 {
		err = umisc.E(cmderr)
	}
	return SrcMsgsFromLns(uslice.StrMap(ustr.Split(strings.TrimSpace(cmdout), "\n"), reline)), err
}

func CmdExecOnSrc(inclstderr bool, reline func(string) string, cmdname string, cmdargs ...string) SrcMsgs {
	return CmdExecOnSrcIn("", inclstderr, reline, cmdname, cmdargs...)
}

func SrcMsgFromLn(ln string) (item *SrcMsg) {
	if lnbits := ustr.Split(ln, ":"); len(lnbits) >= 3 {
		if msgfrom, filepath := 3, lnbits[0]; len(filepath) > 0 {
			pos1ln, errln := strconv.Atoi(lnbits[1])
			if errln != nil {
				pos1ln = 1
			}
			pos1ch, errch := strconv.Atoi(lnbits[2])
			if errch != nil {
				pos1ch = 1
			}
			if errln != nil {
				msgfrom = 1
			} else if errch != nil {
				msgfrom = 2
			}
			if msg := strings.TrimSpace(strings.Join(lnbits[msgfrom:], ":")); len(msg) > 0 || true { // not sure yet, for godef we need to allow "no msg"
				item = &SrcMsg{Msg: msg, Ref: filepath}
				item.Pos1Ln, item.Pos1Ch = pos1ln, pos1ch
			}
		}
	}
	return
}

func SrcMsgsFromLns(lines []string) (msgs SrcMsgs) {
	for i, _ := range lines {
		if item := SrcMsgFromLn(lines[i]); item != nil {
			msgs = append(msgs, item)
		} else if l := len(msgs); l > 0 {
			msgs[l-1].Msg += "\n" + lines[i]
		}
	}
	return
}
