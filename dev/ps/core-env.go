package udevps

import (
	"fmt"
)

/*
Representations of PureScript top-level declarations:
type-alias defs, data-type defs, data-type constructors,
type-classes, type-class instances, and the signatures
of top-level functions.
*/

type CoreEnv struct {
	TypeSyns   map[string]*CoreEnvTypeSyn           `json:"typeSynonyms"`
	TypeDefs   map[string]*CoreEnvTypeDef           `json:"types"`
	DataCtors  map[string]*CoreEnvTypeCtor          `json:"dataConstructors"`
	Classes    map[string]*CoreEnvClass             `json:"typeClasses"`
	ClassDicts []map[string]map[string]*CoreEnvInst `json:"typeClassDictionaries"`
	Functions  map[string]*CoreEnvName              `json:"names"`
}

func (me *CoreEnv) prep() {
	for _, ts := range me.TypeSyns {
		ts.prep()
	}
	for _, td := range me.TypeDefs {
		td.prep()
	}
	for _, tdc := range me.DataCtors {
		tdc.prep()
	}
	for _, tc := range me.Classes {
		tc.prep()
	}
	for _, tcdmap := range me.ClassDicts {
		for _, tcdsubmap := range tcdmap {
			for _, tcd := range tcdsubmap {
				tcd.prep()
			}
		}
	}
	for _, fn := range me.Functions {
		fn.prep()
	}
}

type CoreEnvClass struct {
	CoveringSets   [][]int               `json:"tcCoveringSets"`
	DeterminedArgs []int                 `json:"tcDeterminedArgs"`
	Args           []*CoreEnvClassArg    `json:"tcArgs"`
	Members        []*CoreEnvClassMember `json:"tcMembers"`
	Superclasses   []*CoreConstr         `json:"tcSuperclasses"`
	Dependencies   []struct {
		Determiners []int `json:"determiners"`
		Determined  []int `json:"determined"`
	} `json:"tcDependencies"`
}

func (me *CoreEnvClass) prep() {
	for _, tca := range me.Args {
		tca.prep()
	}
	for _, tcm := range me.Members {
		tcm.prep()
	}
	for _, tcs := range me.Superclasses {
		tcs.prep()
	}
}

type CoreEnvClassArg struct {
	Name string       `json:"tcaName"`
	Kind *CoreTagKind `json:"tcaKind"`
}

func (me *CoreEnvClassArg) prep() {
	me.Kind.prep()
}

type CoreEnvClassMember struct {
	Ident string       `json:"tcmIdent"`
	Type  *CoreTagType `json:"tcmType"`
}

func (me *CoreEnvClassMember) prep() {
	me.Type.prep()
}

type CoreEnvInst struct {
	Chain []string `json:"tcdChain"`
	Index int      `json:"tcdIndex"`
	Value string   `json:"tcdValue"`
	Path  []struct {
		Class string `json:"tcdpClass"`
		Int   int    `json:"tcdpInt"`
	} `json:"tcdPath"`
	ClassName     string         `json:"tcdClassName"`
	InstanceTypes []*CoreTagType `json:"tcdInstanceTypes"`
	Dependencies  []*CoreConstr  `json:"tcdDependencies"`
}

func (me *CoreEnvInst) prep() {
	if len(me.Path) > 0 {
		panic(NotImplErr("tcdPath", me.Path[0].Class, "'typeClassDictionaries'"))
	}
	if me.Index != 0 {
		panic(NotImplErr("tcdIndex", fmt.Sprint(me.Index), "'typeClassDictionaries'"))
	}
	for _, it := range me.InstanceTypes {
		it.prep()
	}
	for _, id := range me.Dependencies {
		id.prep()
	}
}

type CoreEnvTypeSyn struct {
	Args []struct {
		Name string       `json:"tsaName"`
		Kind *CoreTagKind `json:"tsaKind"`
	} `json:"tsArgs"`
	Type *CoreTagType `json:"tsType"`
}

func (me *CoreEnvTypeSyn) prep() {
	if me.Type != nil {
		me.Type.prep()
	}
	for _, tsa := range me.Args {
		tsa.Kind.prep()
	}
}

type CoreEnvTypeCtor struct {
	Decl string       `json:"cDecl"`
	Type string       `json:"cType"`
	Ctor *CoreTagType `json:"cCtor"`
	Args []string     `json:"cArgs"` // value0, value1 ..etc.
}

func (me *CoreEnvTypeCtor) IsDeclData() bool    { return me.Decl == "data" }
func (me *CoreEnvTypeCtor) IsDeclNewtype() bool { return me.Decl == "newtype" }
func (me *CoreEnvTypeCtor) prep() {
	if !(me.IsDeclData() || me.IsDeclNewtype()) {
		panic(NotImplErr("cDecl", me.Decl, "'dataConstructors'"))
	}
	if me.Ctor != nil {
		me.Ctor.prep()
	}
}

type CoreEnvTypeDef struct {
	Kind *CoreTagKind     `json:"tKind"`
	Decl *CoreEnvTypeDecl `json:"tDecl"`
}

func (me *CoreEnvTypeDef) prep() {
	if me.Kind != nil {
		me.Kind.prep()
	}
	if me.Decl != nil {
		me.Decl.prep()
	}
}

type CoreEnvTypeDecl struct {
	TypeSynonym       bool
	ExternData        bool
	LocalTypeVariable bool
	ScopedTypeVar     bool
	DataType          *CoreEnvTypeData
}

func (me *CoreEnvTypeDecl) prep() {
	if me.LocalTypeVariable {
		panic(NotImplErr("tDecl", "LocalTypeVariable", "'types'"))
	}
	if me.ScopedTypeVar {
		panic(NotImplErr("tDecl", "ScopedTypeVar", "'types'"))
	}
	if me.DataType != nil {
		me.DataType.prep()
	}
}

type CoreEnvTypeData struct {
	Args []struct {
		Name string       `json:"dtaName"`
		Kind *CoreTagKind `json:"dtaKind"`
	} `json:"dtArgs"`
	Ctors []struct {
		Name  string         `json:"dtcName"`
		Types []*CoreTagType `json:"dtcTypes"`
	} `json:"dtCtors"`
}

func (me *CoreEnvTypeData) prep() {
	for _, tda := range me.Args {
		tda.Kind.prep()
	}
	for _, tdc := range me.Ctors {
		for _, tdct := range tdc.Types {
			tdct.prep()
		}
	}
}

type CoreEnvName struct {
	Vis  string       `json:"nVis"`
	Kind string       `json:"nKind"`
	Type *CoreTagType `json:"nType"`
}

func (me *CoreEnvName) isVisDefined() bool   { return me.Vis == "Defined" }
func (me *CoreEnvName) isVisUndefined() bool { return me.Vis == "Undefined" }
func (me *CoreEnvName) isKindPrivate() bool  { return me.Kind == "Private" }
func (me *CoreEnvName) isKindPublic() bool   { return me.Kind == "Public" }
func (me *CoreEnvName) isKindExternal() bool { return me.Kind == "External" }
func (me *CoreEnvName) prep() {
	if !(me.isVisDefined() || me.isVisUndefined()) {
		panic(NotImplErr("nVis", me.Vis, "'names'"))
	}
	if !(me.isKindPublic() || me.isKindPrivate() || me.isKindExternal()) {
		panic(NotImplErr("nKind", me.Kind, "'names'"))
	}
	if me.Type != nil {
		me.Type.prep()
	}
}
