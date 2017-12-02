package udevgo

import (
	"encoding/json"
	"path/filepath"
	"strings"

	gurujson "golang.org/x/tools/cmd/guru/serial"

	"github.com/metaleap/go-util"
	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type Gogetdoc struct {
	Name   string `json:"name,omitempty"`
	ImpP   string `json:"import,omitempty"`
	Decl   string `json:"decl,omitempty"`
	Doc    string `json:"doc,omitempty"`
	DocUrl string `json:",omitempty"`
	Pos    string `json:"pos,omitempty"`
	Pkg    string `json:"pkg,omitempty"`
	ImpS   string `json:",omitempty"`
	ImpN   string `json:",omitempty"`

	Err     string `json:",omitempty"`
	ErrMsgs string `json:",omitempty"`
}

var (
	GuruScopeExclPkgs []string
)

func queryGuru(gurucmd string, fullsrcfilepath string, srcin string, bpos1 string, bpos2 string, singlevar interface{}, multnextvar func() (interface{}, func(interface{}))) (allok bool, err error) {
	cmdargs := []string{"-json", gurucmd, fullsrcfilepath + ":#" + bpos1}
	if len(bpos2) > 0 {
		cmdargs[len(cmdargs)-1] = cmdargs[len(cmdargs)-1] + ",#" + bpos2
	}
	if len(SnipImp) > 0 {
		cmdargs = append([]string{"-scope", SnipImp + "..."}, cmdargs...)
		for _, exclpkg := range GuruScopeExclPkgs {
			cmdargs[1] = cmdargs[1] + ",-" + exclpkg
		}
	}
	var jsonerr error
	cmdout, cmderr, _ := urun.CmdExecStdin(srcin, "", "guru", cmdargs...)
	if len(cmderr) > 0 {
		if ustr.Has(cmderr, "guru: ") {
			cmderr = ustr.Join(uslice.StrFiltered(ustr.Split(cmderr, "\n"), func(ln string) bool { return ustr.Pref(ln, "guru: ") }), "\n")
		}
		err = umisc.E(cmderr)
	}
	if cmdout = ustr.Trim(cmdout); len(cmdout) > 0 {
		if singlevar != nil {
			jsonerr = json.Unmarshal([]byte(cmdout), singlevar)
			if allok = (jsonerr == nil); !allok {
				err = jsonerr
			}
		} else {
			allok = true
			for _, subjson := range ustr.Split(cmdout[1:len(cmdout)-1], "}\n{") {
				jsonvar, ondecoded := multnextvar()
				jsonerr = json.Unmarshal([]byte("{"+subjson+"}"), jsonvar)
				if jsonerr == nil {
					ondecoded(jsonvar)
				} else {
					allok = false
					err = jsonerr
				}
			}
		}
	}
	return
}

func QueryDef_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.Definition {
	var gr gurujson.Definition
	if ok, _ := queryGuru("definition", fullsrcfilepath, srcin, bytepos, "", &gr, nil); ok {
		return &gr
	}
	return nil
}

func QueryDesc_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.Describe {
	var gr gurujson.Describe
	if ok, _ := queryGuru("describe", fullsrcfilepath, srcin, bytepos, "", &gr, nil); ok {
		return &gr
	}
	return nil
}

func QueryImpl_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.Implements {
	var gr gurujson.Implements
	if ok, _ := queryGuru("implements", fullsrcfilepath, srcin, bytepos, "", &gr, nil); ok {
		return &gr
	}
	return nil
}

func QueryWhat_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.What {
	var gr gurujson.What
	if ok, _ := queryGuru("what", fullsrcfilepath, srcin, bytepos, "", &gr, nil); ok {
		return &gr
	}
	return nil
}

func QueryRefs_Guru(fullsrcfilepath string, srcin string, bytepos string) (refs []gurujson.Ref) {
	onrefpkg := func(mayberefpkg interface{}) {
		if refpkg, ok := mayberefpkg.(*gurujson.ReferrersPackage); ok && refpkg != nil {
			refs = append(refs, refpkg.Refs...)
		}
	}
	queryGuru("referrers", fullsrcfilepath, srcin, bytepos, "", nil, func() (interface{}, func(interface{})) {
		return &gurujson.ReferrersPackage{}, onrefpkg
	})
	return
}

func QueryCallees_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gc *gurujson.Callees, err error) {
	var gr gurujson.Callees
	var ok bool
	if ok, err = queryGuru("callees", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil); ok {
		gc = &gr
	}
	return
}

func QueryCallers_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) ([]gurujson.Caller, error) {
	var gr []gurujson.Caller
	_, err := queryGuru("callers", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil)
	return gr, err
}

func QueryCallstack_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gcs *gurujson.CallStack, err error) {
	var gr gurujson.CallStack
	var ok bool
	if ok, err = queryGuru("callstack", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil); ok {
		gcs = &gr
	}
	return
}

func QueryWhicherrs_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gwe *gurujson.WhichErrs, err error) {
	var gr gurujson.WhichErrs
	var ok bool
	if ok, err = queryGuru("whicherrs", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil); ok {
		gwe = &gr
	}
	return
}

func QueryPeers_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gp *gurujson.Peers, err error) {
	var gr gurujson.Peers
	var ok bool
	if ok, err = queryGuru("peers", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil); ok {
		gp = &gr
	}
	return
}

func QueryPointsto_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) ([]gurujson.PointsTo, error) {
	var gr []gurujson.PointsTo
	_, err := queryGuru("pointsto", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil)
	return gr, err
}

func QueryFreevars_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gfvs []*gurujson.FreeVar, err error) {
	ongfv := func(maybegfv interface{}) {
		if gfv, ok := maybegfv.(*gurujson.FreeVar); ok && gfv != nil {
			gfvs = append(gfvs, gfv)
		}
	}
	_, err = queryGuru("freevars", fullsrcfilepath, srcin, bytepos1, bytepos2, nil, func() (interface{}, func(interface{})) {
		return &gurujson.FreeVar{}, ongfv
	})
	return
}

func QueryCmplSugg_Gocode(fullsrcfilepath string, srcin string, pos string) (cmpls []map[string]string) {
	var args []string
	if len(srcin) == 0 {
		args = []string{"-in=" + fullsrcfilepath}
	}
	args = append(args, "-f=json", "autocomplete", fullsrcfilepath, pos)
	if cmdout, _, _ := urun.CmdExecStdin(srcin, filepath.Dir(fullsrcfilepath), "gocode", args...); len(cmdout) > 0 {
		if i := ustr.Idx(cmdout, "[{"); i > 0 {
			cmdout = cmdout[:len(cmdout)-1][i:]
		}
		json.Unmarshal([]byte(cmdout), &cmpls)
	}
	return
}

func cmdArgs_Godef(fullsrcfilepath string, srcin string, bytepos string) (args []string) {
	args = []string{"-f", fullsrcfilepath, "-o", bytepos}
	if len(srcin) > 0 {
		args = append(args, "-i")
	}
	return
}

func QueryDefLoc_Godef(fullsrcfilepath string, srcin string, bytepos string) *udev.SrcMsg {
	var refs udev.SrcMsgs
	args := cmdArgs_Godef(fullsrcfilepath, srcin, bytepos)
	refs, _ = udev.CmdExecOnStdIn(srcin, filepath.Dir(fullsrcfilepath), nil, "godef", args...)
	for _, srcmsg := range refs {
		if isfile := ufs.FileExists(srcmsg.Ref); isfile {
			return srcmsg
		}
	}
	return nil
}

func QueryDefDecl_GoDef(fullsrcfilepath string, srcin string, bytepos string) (defdecl string) {
	foreachln := func(ln string) string {
		if l, _ := ustr.BreakOn(ln, ":"); l == "godef" || ufs.FileExists(l) || len(l) == 1 || ln == "-" {
			return ""
		}
		return ln
	}
	args := append(cmdArgs_Godef(fullsrcfilepath, srcin, bytepos), "-t")
	if cmdout, _, err := urun.CmdExecStdin(srcin, filepath.Dir(fullsrcfilepath), "godef", args...); err == nil {
		if cmdout = ustr.Trim(cmdout); cmdout != "" {
			defdecl = strings.TrimSpace(ustr.Join(uslice.StrWithout(uslice.StrMap(ustr.Split(cmdout, "\n"), foreachln), true, ""), "\n"))
		}
	}
	return
}

func QueryDefLoc_Gogetdoc(fullsrcfilepath string, srcin string, bytepos string) *udev.SrcMsg {
	if ggd := Query_Gogetdoc(fullsrcfilepath, srcin, bytepos); ggd != nil && len(ggd.Pos) > 0 {
		if srcref := udev.SrcMsgFromLn(ggd.Pos); srcref != nil {
			return srcref
		}
	}
	return nil
}

func Query_Gogetdoc(fullsrcfilepath string, srcin string, bytepos string) *Gogetdoc {
	var ggd Gogetdoc
	cmdargs := []string{"-json", "-u", "-linelength", "50", "-pos", fullsrcfilepath + ":#" + bytepos}
	if len(srcin) > 0 {
		cmdargs = append(cmdargs, "-modified")
		srcin = fullsrcfilepath + "\n" + umisc.Str(len([]byte(srcin))) + "\n" + srcin
	}
	cmdout, cmderr, err := urun.CmdExecStdin(srcin, "", "gogetdoc", cmdargs...)
	if cmdout, cmderr = ustr.Trim(cmdout), ustr.Trim(cmderr); err == nil && len(cmdout) > 0 {
		if err = json.Unmarshal([]byte(cmdout), &ggd); err == nil {
			ggd.ImpS = ShortImpP(ggd.ImpP)
			ggd.Decl = ShortenImPs(ggd.Decl)
			ggd.DocUrl = ggd.ImpP + "#"
			dochash := ggd.Name
			var tname string
			if ustr.Pref(ggd.Decl, "func (") {
				tname = ggd.Decl[6:ustr.Idx(ggd.Decl, ")")]
				_, tname = ustr.BreakOn(tname, " ")
				_, tname = ustr.BreakOnLast(tname, ".")
			} else if ustr.Pref(ggd.Decl, "field ") && Has_guru {
				if gw := QueryWhat_Guru(fullsrcfilepath, srcin, bytepos); gw != nil {
					for _, encl := range gw.Enclosing {
						if encl.Description == "composite literal" {
							if gd := QueryDesc_Guru(fullsrcfilepath, srcin, umisc.Str(encl.Start)); gd != nil && gd.Type != nil {
								tname = gd.Type.Type
							}
							break
						}
					}
				}
			}
			if tname = strings.TrimLeft(tname, "*"); len(tname) > 0 {
				dochash = tname + "." + dochash
			}
			ggd.DocUrl += dochash
			ggd.ImpN = ggd.Pkg
			if len(tname) > 0 || ggd.Name != ggd.ImpN {
				if ggd.ImpN != "" {
					ggd.ImpN = "_" + ggd.ImpN + "_ · "
				}
				ggd.ImpN += strings.TrimLeft(tname+"."+ggd.Name, ".")
			}
		}
	}
	if cmdout == "gogetdoc: no documentation found" {
		return nil
	}
	if ggd.ErrMsgs = cmderr; err != nil {
		ggd.Err = err.Error()
	}
	return &ggd
}
