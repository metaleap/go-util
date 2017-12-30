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
	// ImpS   string `json:",omitempty"`
	ImpN string `json:",omitempty"`
	Type string `json:",omitempty"`

	Err     string `json:",omitempty"`
	ErrMsgs string `json:",omitempty"`
}

type Guru struct {
	gurujson.Describe

	IsLessThan func(*gurujson.DescribeMember, *gurujson.DescribeMember) bool `json:"-"`
}

func (me *Guru) Len() int { return len(me.Package.Members) }

func (me *Guru) Less(i int, j int) bool {
	dis, dat := me.Package.Members[i], me.Package.Members[j]
	if me.IsLessThan != nil {
		return me.IsLessThan(dis, dat)
	}
	if dis.Kind != dat.Kind {
		return dis.Kind < dat.Kind
	} else if dis.Type != dat.Type {
		return dis.Type < dat.Type
	}
	return dis.Name < dat.Name
}

func (me *Guru) Swap(i int, j int) {
	me.Package.Members[i], me.Package.Members[j] = me.Package.Members[j], me.Package.Members[i]
}

func (me *Guru) Matches(pM *gurujson.DescribeMember, lowerCaseQuery string) bool {
	return strings.Contains(strings.ToLower(pM.Kind), lowerCaseQuery) ||
		strings.Contains(strings.ToLower(pM.Type), lowerCaseQuery) ||
		strings.Contains(strings.ToLower(pM.Name), lowerCaseQuery) ||
		strings.Contains(strings.ToLower(pM.Value), lowerCaseQuery)
}

var (
	GuruScopes        string
	GuruScopeExclPkgs = map[string]bool{}
)

func queryGuru(gurucmd string, fullsrcfilepath string, srcin string, bpos1 string, bpos2 string, singlevar interface{}, multnextvar func() (interface{}, func(interface{})), guruScopes string) (allok bool, err error) {
	cmdargs := []string{"-json", gurucmd, fullsrcfilepath + ":#" + bpos1}
	if len(bpos2) > 0 {
		cmdargs[len(cmdargs)-1] = cmdargs[len(cmdargs)-1] + ",#" + bpos2
	}
	if srcin != "" {
		cmdargs = append([]string{"-modified"}, cmdargs...)
		srcin = queryModSrcIn(fullsrcfilepath, srcin)
	}
	canscope := gurucmd == "callees" || gurucmd == "callers" || gurucmd == "callstack" || gurucmd == "pointsto" || gurucmd == "peers" || gurucmd == "whicherrs"
	if guruScopes == "" {
		guruScopes = GuruScopes
	}
	if guruScopes != "" && canscope {
		cmdargs = append([]string{"-scope", guruScopes}, cmdargs...)
		for exclpkg, excl := range GuruScopeExclPkgs {
			if excl {
				cmdargs[1] = cmdargs[1] + ",-" + exclpkg
			}
		}
	}
	var jsonerr error
	cmdout, cmderr, e := urun.CmdExecStdin(srcin, "", "guru", cmdargs...)
	if len(cmderr) > 0 {
		if ustr.Has(cmderr, "is not a Go source file") {
			cmderr = ""
		} else if ustr.Has(cmderr, "guru: ") {
			cmderr = ustr.Join(uslice.StrFiltered(ustr.Split(cmderr, "\n"), func(ln string) bool { return ustr.Pref(ln, "guru: ") }), "\n")
		}
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
	} else if len(cmderr) > 0 {
		err = umisc.E(cmderr)
	}
	if err == nil && e != nil {
		err = e
	}
	return
}

func QueryDef_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.Definition {
	var gr gurujson.Definition
	if ok, _ := queryGuru("definition", fullsrcfilepath, srcin, bytepos, "", &gr, nil, ""); ok {
		return &gr
	}
	return nil
}

func QueryDesc_Guru(fullsrcfilepath string, srcin string, bytepos string) (*Guru, error) {
	var gr Guru
	if _, err := queryGuru("describe", fullsrcfilepath, srcin, bytepos, "", &gr, nil, ""); err != nil {
		return nil, err
	}
	return &gr, nil
}

func QueryImpl_Guru(fullsrcfilepath string, srcin string, bytepos string) *gurujson.Implements {
	var gr gurujson.Implements
	if ok, _ := queryGuru("implements", fullsrcfilepath, srcin, bytepos, "", &gr, nil, ""); ok {
		return &gr
	}
	return nil
}

func QueryWhat_Guru(fullsrcfilepath string, srcin string, bytepos string) (*gurujson.What, error) {
	var gr gurujson.What
	if _, err := queryGuru("what", fullsrcfilepath, srcin, bytepos, "", &gr, nil, ""); err != nil {
		return nil, err
	}
	return &gr, nil
}

func QueryRefs_Guru(fullsrcfilepath string, srcin string, bytepos string) (refs []gurujson.Ref) {
	onrefpkg := func(mayberefpkg interface{}) {
		if refpkg, _ := mayberefpkg.(*gurujson.ReferrersPackage); refpkg != nil {
			refs = append(refs, refpkg.Refs...)
		}
	}
	queryGuru("referrers", fullsrcfilepath, srcin, bytepos, "", nil, func() (interface{}, func(interface{})) {
		return &gurujson.ReferrersPackage{}, onrefpkg
	}, "")
	return
}

func QueryCallees_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gc *gurujson.Callees, err error) {
	var gr gurujson.Callees
	var ok bool
	if ok, err = queryGuru("callees", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); ok {
		gc = &gr
	}
	return
}

func QueryCallers_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gr []gurujson.Caller, err error) {
	var ok bool
	if ok, err = queryGuru("callers", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); !ok {
		gr = nil
	}
	return
}

func QueryCallstack_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gcs *gurujson.CallStack, err error) {
	var gr gurujson.CallStack
	var ok bool
	if ok, err = queryGuru("callstack", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); ok {
		gcs = &gr
	}
	return
}

func QueryWhicherrs_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gwe *gurujson.WhichErrs, err error) {
	var gr gurujson.WhichErrs
	var ok bool
	if ok, err = queryGuru("whicherrs", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); ok {
		gwe = &gr
	}
	return
}

func QueryPeers_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gp *gurujson.Peers, err error) {
	var gr gurujson.Peers
	var ok bool
	if ok, err = queryGuru("peers", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); ok {
		gp = &gr
	}
	return
}

func QueryPointsto_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string, altScopes string) (gr []gurujson.PointsTo, err error) {
	var ok bool
	if ok, err = queryGuru("pointsto", fullsrcfilepath, srcin, bytepos1, bytepos2, &gr, nil, altScopes); !ok {
		gr = nil
	}
	return
}

func QueryFreevars_Guru(fullsrcfilepath string, srcin string, bytepos1 string, bytepos2 string) (gfvs []*gurujson.FreeVar, err error) {
	ongfv := func(maybegfv interface{}) {
		if gfv, ok := maybegfv.(*gurujson.FreeVar); ok && gfv != nil {
			gfvs = append(gfvs, gfv)
		}
	}
	_, err = queryGuru("freevars", fullsrcfilepath, srcin, bytepos1, bytepos2, nil, func() (interface{}, func(interface{})) {
		return &gurujson.FreeVar{}, ongfv
	}, "")
	return
}

func QueryCmplSugg_Gocode(fullsrcfilepath string, srcin string, pos string) (cmpls []map[string]string, err error) {
	var args []string
	if len(srcin) == 0 {
		args = []string{"-in=" + fullsrcfilepath}
	}
	args = append(args, "-f=json", "autocomplete", fullsrcfilepath, pos)
	cmdout, cmderr, e := urun.CmdExecStdin(srcin, filepath.Dir(fullsrcfilepath), "gocode", args...)
	if cmdout = ustr.Trim(cmdout); len(cmdout) > 0 {
		if i := ustr.Idx(cmdout, "[{"); i > 0 {
			cmdout = cmdout[:len(cmdout)-1][i:]
		}
		err = json.Unmarshal([]byte(cmdout), &cmpls)
	} else if e != nil {
		err = e
	} else if len(cmderr) > 0 {
		err = umisc.E(cmderr)
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
	if cmdout, cmderr, err := urun.CmdExecStdin(srcin, filepath.Dir(fullsrcfilepath), "godef", args...); err != nil {
		defdecl = err.Error()
	} else if cmdout = ustr.Trim(cmdout); cmdout == "" && cmderr != "" && !(strings.HasPrefix(cmderr, "godef: ") && (strings.ContainsAny(cmderr, "found") || strings.Contains(cmderr, "error finding import path for"))) {
		defdecl = cmderr
	} else if cmdout != "" {
		defdecl = strings.TrimSpace(ustr.Join(uslice.StrWithout(uslice.StrMap(ustr.Split(cmdout, "\n"), foreachln), true, ""), "\n"))
	}
	return
}

func QueryDefLoc_Gogetdoc(fullsrcfilepath string, srcin string, bytepos string) *udev.SrcMsg {
	if ggd := Query_Gogetdoc(fullsrcfilepath, srcin, bytepos, true, false); ggd != nil && len(ggd.Pos) > 0 {
		if srcref := udev.SrcMsgFromLn(ggd.Pos); srcref != nil {
			return srcref
		}
	}
	return nil
}

func queryModSrcIn(fullsrcfilepath string, srcin string) string {
	return fullsrcfilepath + "\n" + umisc.Str(len([]byte(srcin))) + "\n" + srcin
}

func Query_Gogetdoc(fullsrcfilepath string, srcin string, bytepos string, onlyDocAndDecl bool, docFromPlainToMarkdown bool) *Gogetdoc {
	var ggd Gogetdoc
	cmdargs := []string{"-json", "-u", "-linelength", "50", "-pos", fullsrcfilepath + ":#" + bytepos}
	if len(srcin) > 0 {
		cmdargs = append(cmdargs, "-modified")
		srcin = queryModSrcIn(fullsrcfilepath, srcin)
	}
	cmdout, cmderr, err := urun.CmdExecStdin(srcin, "", "gogetdoc", cmdargs...)
	if cmdout, cmderr = ustr.Trim(cmdout), ustr.Trim(cmderr); err == nil && len(cmdout) > 0 {
		if err = json.Unmarshal([]byte(cmdout), &ggd); err == nil {
			ggd.DocUrl = ggd.ImpP + "#"
			if ispkgstd := (ggd.ImpP == "builtin"); docFromPlainToMarkdown {
				if (!ispkgstd) && PkgsByImP != nil {
					if pkg := PkgsByImP[ggd.ImpP]; pkg != nil {
						ispkgstd = pkg.Standard
					}
				}
				if ispkgstd {
					ggd.Doc = strings.Replace(ggd.Doc, "\n\t", "\n\n\t", -1)
					ggd.Doc = strings.Replace(ggd.Doc, "\n\n\n", "\n\n", -1)
				}
			}
			dochash := ggd.Name
			if !onlyDocAndDecl {
				if ustr.Pref(ggd.Decl, "func (") {
					ggd.Type = ggd.Decl[6:ustr.Idx(ggd.Decl, ")")]
				} else if Has_guru && ustr.Pref(ggd.Decl, "field ") {
					if gw, e := QueryWhat_Guru(fullsrcfilepath, srcin, bytepos); e != nil {
						err = e
					} else {
						for _, encl := range gw.Enclosing {
							if encl.Description == "composite literal" || encl.Description == "selector" {
								if gd, e := QueryDesc_Guru(fullsrcfilepath, srcin, umisc.Str(encl.Start)); gd != nil {
									if gd.Type != nil {
										ggd.Type = gd.Type.Type
									} else if encl.Description != "selector" && gd.Value != nil {
										ggd.Type = gd.Value.Type
									}
								} else if e != nil {
									err = e
								}
								break
							}
						}
					}
				}
				if err == nil {
					if ggd.Type != "" {
						_, ggd.Type = ustr.BreakOn(ggd.Type, " ")
						_, ggd.Type = ustr.BreakOnLast(ggd.Type, ".")
					}
					if ggd.Type = strings.TrimLeft(ggd.Type, "*[]"); len(ggd.Type) > 0 {
						dochash = ggd.Type + "." + dochash
					}
					ggd.DocUrl += dochash
					ggd.ImpN = ggd.Pkg
					if len(ggd.Type) > 0 || ggd.Name != ggd.ImpN {
						if ggd.ImpN != "" {
							ggd.ImpN = "_" + ggd.ImpN + "_ · "
						}
						ggd.ImpN += strings.TrimLeft(ggd.Type+"."+ggd.Name, ".")
					}
				}
			}
		}
	}
	if ggd.ErrMsgs = ustr.Trim(strings.Replace(cmderr, "gogetdoc: no documentation found", "", -1)); err != nil {
		ggd.Err = err.Error()
	}
	return &ggd
}
