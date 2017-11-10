package udevps

import "strings"

type CoreAnnotated struct {
	Annotation CoreAnnotation `json:"annotation"`
}

type CoreAnnotation struct {
	SourceSpan *CoreSourceSpan     `json:"sourceSpan"`
	Type       *CoreTagType        `json:"type"`
	Comments   []CoreComment       `json:"comments"`
	Meta       *CoreAnnotationMeta `json:"meta"`
}

func (me *CoreAnnotation) prep() {
	if me.Type != nil {
		me.Type.prep()
	}
	if me.Meta != nil {
		me.Meta.prep()
	}
}

type CoreAnnotationMeta struct {
	MetaType          string   `json:"metaType"`        // IsConstructor or IsNewtype or IsTypeClassConstructor or IsForeign
	ConstructorType   string   `json:"constructorType"` // if MetaType=IsConstructor: SumType or ProductType
	ConstructorIdents []string `json:"identifiers"`     // if MetaType=IsConstructor
}

func (me *CoreAnnotationMeta) IsConstructor() bool      { return me.MetaType == "IsConstructor" }
func (me *CoreAnnotationMeta) IsForeign() bool          { return me.MetaType == "IsForeign" }
func (me *CoreAnnotationMeta) IsNewtype() bool          { return me.MetaType == "IsNewtype" }
func (me *CoreAnnotationMeta) IsTypeClassCtor() bool    { return me.MetaType == "IsTypeClassConstructor" }
func (me *CoreAnnotationMeta) IsCtorˇSumType() bool     { return me.ConstructorType == "SumType" }
func (me *CoreAnnotationMeta) IsCtorˇProductType() bool { return me.ConstructorType == "ProductType" }

func (me *CoreAnnotationMeta) prep() {
	if isctor := me.IsConstructor(); !(isctor || me.IsForeign() || me.IsNewtype() || me.IsTypeClassCtor()) {
		panic(NotImplErr("CoreFn Annotation.MetaType", me.MetaType, *me))
	} else if isctor && !(me.IsCtorˇSumType() || me.IsCtorˇProductType()) {
		panic(NotImplErr("CoreFn Annotation.ConstructorType", me.ConstructorType, *me))
	}
}

func (me *CoreAnnotationMeta) String() string {
	return me.MetaType + "::" + me.ConstructorType + "❭" + strings.Join(me.ConstructorIdents, "❬")
}

type CoreComment struct {
	LineComment  string
	BlockComment string
}

type CoreConstr struct {
	Class interface{}    `json:"constraintClass"`
	Args  []*CoreTagType `json:"constraintArgs"`
	Data  interface{}    `json:"constraintData"`

	Cls string
}

func (me *CoreConstr) prep() {
	if me.Cls == "" {
		switch cls := me.Class.(type) {
		case string:
			me.Cls = cls
		default:
			panic(me.Class.(string))
		}
	}
	if me.Data != nil {
		panic(NotImplErr("constraintData for 'typeClasses' or", "typeClassDictionaries", me.Data))
	}
	for _, ca := range me.Args {
		ca.prep()
	}
}

type CoreModuleRef struct {
	ModuleName []string `json:"moduleName"`
}

type CoreSourceSpan struct {
	Name  string `json:"name"`
	Start []int  `json:"start"`
	End   []int  `json:"end"`
}
