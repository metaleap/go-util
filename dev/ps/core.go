package udevps

type CoreAnn struct {
	SourceSpan *CoreSourceSpan `json:"sourceSpan"`
	Type       *CoreTagType    `json:"type"`
	Comments   []*CoreComment  `json:"comments"`
	Meta       struct {
		MetaType   string   `json:"metaType"`        // IsConstructor or IsNewtype or IsTypeClassConstructor or IsForeign
		CtorType   string   `json:"constructorType"` // if MetaType=IsConstructor: SumType or ProductType
		CtorIdents []string `json:"identifiers"`     // if MetaType=IsConstructor
	} `json:"meta"`
}

func (me *CoreAnn) prep() {
	if me.Type != nil {
		me.Type.prep()
	}
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

type CoreDecl struct {
	BindType string        `json:"bindType"`
	Ident    string        `json:"identifier"`
	Ann      *CoreAnn      `json:"annotation"`
	Expr     *CoreDeclExpr `json:"expression"`
}

func (me *CoreDecl) prep() {
	if me.Ann != nil {
		me.Ann.prep()
	}
	if me.Expr != nil {
		me.Expr.prep()
	}
}

type CoreDeclExpr struct {
	Ann        *CoreAnn `json:"annotation"`
	ExprTag    string   `json:"type"`            // Var or Literal or Abs or App or Let or Constructor (or Accessor or ObjectUpdate or Case)
	CtorName   string   `json:"constructorName"` // if ExprTag=Constructor
	CtorType   string   `json:"typeName"`        // if ExprTag=Constructor
	CtorFields []string `json:"fieldNames"`      // if ExprTag=Constructor
}

func (me *CoreDeclExpr) prep() {
	if me.Ann != nil {
		me.Ann.prep()
	}
}

type CoreSourceSpan struct {
	Name  string `json:"name"`
	Start []int  `json:"start"`
	End   []int  `json:"end"`
}
