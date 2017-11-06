package udevps

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
}

type CoreAnnotationMeta struct {
	MetaType          string   `json:"metaType"`        // IsConstructor or IsNewtype or IsTypeClassConstructor or IsForeign
	ConstructorType   string   `json:"constructorType"` // if MetaType=IsConstructor: SumType or ProductType
	ConstructorIdents []string `json:"identifiers"`     // if MetaType=IsConstructor
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

type CoreSourceSpan struct {
	Name  string `json:"name"`
	Start []int  `json:"start"`
	End   []int  `json:"end"`
}
