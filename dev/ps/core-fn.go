package udevps

import (
	"encoding/json"
)

type CoreFn struct {
	ModuleName []string `json:"moduleName"` // ["Control","Monad","Eff"]
	Imports    []struct {
		Annotation CoreAnnotation `json:"annotation"`
		ModuleName []string       `json:"moduleName"`
	} `json:"imports"`
	ModulePath string        `json:"modulePath"` // /home/_/c/ps/gonad-test/src/Mini/DataType.purs   ,   /home/_/c/ps/gonad-test/bower_components/purescript-eff/src/Control/Monad/Eff.purs
	Exports    []string      `json:"exports"`
	Decls      []CoreFnDecl  `json:"decls"`
	Comments   []CoreComment `json:"comments"`
	Foreign    []string      `json:"foreign"`
}

func (me *CoreFn) Prep() {
	for _, imp := range me.Imports {
		imp.Annotation.prep()
	}
	for _, decl := range me.Decls {
		decl.prep()
	}
}

type CoreFnDecl struct {
	*CoreFnDeclBind
	Binds []CoreFnDeclBind `json:"binds"`
}

func (me *CoreFnDecl) IsRecursive() bool    { return me.CoreFnDeclBind == nil }
func (me *CoreFnDecl) IsNonRecursive() bool { return me.CoreFnDeclBind != nil }

func (me *CoreFnDecl) prep() {
	if me.CoreFnDeclBind != nil {
		me.CoreFnDeclBind.prep()
	} else {
		for _, b := range me.Binds {
			b.prep()
		}
	}
}

type CoreFnDeclBind struct {
	Identifier string         `json:"identifier"`
	Annotation CoreAnnotation `json:"annotation"`
	Expression CoreFnExpr     `json:"expression"`
}

func (me *CoreFnDeclBind) prep() {
	me.Annotation.prep()
	me.Expression.prep()
}

type CoreFnExpr struct {
	Abs          *CoreFnExprAbs    `json:"-"`
	Accessor     *CoreFnExprAcc    `json:"-"`
	App          *CoreFnExprApp    `json:"-"`
	Case         *CoreFnExprCase   `json:"-"`
	Constructor  *CoreFnExprCtor   `json:"-"`
	Let          *CoreFnExprLet    `json:"-"`
	Literal      *CoreFnExprLit    `json:"-"`
	ObjectUpdate *CoreFnExprObjUpd `json:"-"`
	Var          *CoreFnExprVar    `json:"-"`

	prep func()
}

func (me *CoreFnExpr) UnmarshalJSON(data []byte) (err error) {
	var exprtype struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(data, &exprtype); err == nil {
		switch exprtype.Type {
		case "Abs":
			var abs CoreFnExprAbs
			if err = json.Unmarshal(data, &abs); err == nil {
				me.prep, me.Abs = abs.prep, &abs
			}
		case "Accessor":
			var acc CoreFnExprAcc
			if err = json.Unmarshal(data, &acc); err == nil {
				me.prep, me.Accessor = acc.prep, &acc
			}
		case "App":
			var app CoreFnExprApp
			if err = json.Unmarshal(data, &app); err == nil {
				me.prep, me.App = app.prep, &app
			}
		case "Case":
			var cªse CoreFnExprCase
			if err = json.Unmarshal(data, &cªse); err == nil {
				me.prep, me.Case = cªse.prep, &cªse
			}
		case "Constructor":
			var ctor CoreFnExprCtor
			if err = json.Unmarshal(data, &ctor); err == nil {
				me.prep, me.Constructor = ctor.prep, &ctor
			}
		case "Let":
			var let CoreFnExprLet
			if err = json.Unmarshal(data, &let); err == nil {
				me.prep, me.Let = let.prep, &let
			}
		case "Literal":
			var lit CoreFnExprLit
			if err = json.Unmarshal(data, &lit); err == nil {
				me.prep, me.Literal = lit.prep, &lit
			}
		case "ObjectUpdate":
			var objupd CoreFnExprObjUpd
			if err = json.Unmarshal(data, &objupd); err == nil {
				me.prep, me.ObjectUpdate = objupd.prep, &objupd
			}
		case "Var":
			var vªr CoreFnExprVar
			if err = json.Unmarshal(data, &vªr); err == nil {
				me.prep, me.Var = vªr.prep, &vªr
			}
		default:
			err = NotImplErr("CoreFnExpr.Type", exprtype.Type, string(data))
		}
	}
	return
}

type CoreFnExprBase struct {
	Annotation CoreAnnotation `json:"annotation"`
}

func (me *CoreFnExprBase) prep() {
	me.Annotation.prep()
}

type CoreFnExprAbs struct {
	CoreFnExprBase
	Argument string     `json:"argument"`
	Body     CoreFnExpr `json:"body"`
}

func (me *CoreFnExprAbs) prep() {
	me.CoreFnExprBase.prep()
	me.Body.prep()
}

type CoreFnExprAcc struct {
	CoreFnExprBase
	FieldName  string     `json:"fieldName"`
	Expression CoreFnExpr `json:"expression"`
}

func (me *CoreFnExprAcc) prep() {
	me.CoreFnExprBase.prep()
	me.Expression.prep()
}

type CoreFnExprApp struct {
	CoreFnExprBase
	Abstraction CoreFnExpr `json:"abstraction"`
	Argument    CoreFnExpr `json:"argument"`
}

func (me *CoreFnExprApp) prep() {
	me.CoreFnExprBase.prep()
	me.Abstraction.prep()
	me.Argument.prep()
}

type CoreFnExprCase struct {
	CoreFnExprBase
	Expressions  []CoreFnExpr  `json:"caseExpressions"`
	Alternatives []interface{} `json:"caseAlternatives"`
}

func (me *CoreFnExprCase) prep() {
	me.CoreFnExprBase.prep()
	for _, expr := range me.Expressions {
		expr.prep()
	}
}

type CoreFnExprCtor struct {
	CoreFnExprBase
	ConstructorName string   `json:"constructorName"`
	TypeName        string   `json:"typeName"`
	FieldNames      []string `json:"fieldNames"`
}

type CoreFnExprLet struct {
	CoreFnExprBase
	Binds      []CoreFnDecl `json:"binds"`
	Expression CoreFnExpr   `json:"expression"`
}

func (me *CoreFnExprLet) prep() {
	me.CoreFnExprBase.prep()
	me.Expression.prep()
	for _, bind := range me.Binds {
		bind.prep()
	}
}

type CoreFnExprLit struct {
	CoreFnExprBase
	Number  float64      `json:"-"`
	Int     int          `json:"-"`
	Boolean bool         `json:"-"`
	Char    rune         `json:"-"`
	String  string       `json:"-"`
	Array   []CoreFnExpr `json:"-"`

	Type string `json:"-"`
}

func (me *CoreFnExprLit) UnmarshalJSON(data []byte) (err error) {
	var raw struct {
		CoreFnExprBase
		Lit struct {
			Type string      `json:"literalType"`
			Val  interface{} `json:"value"`
		} `json:"value"`
	}
	if err = json.Unmarshal(data, &raw); err == nil {
		me.Type, me.CoreFnExprBase = raw.Lit.Type, raw.CoreFnExprBase
		switch me.Type {
		case "ArrayLiteral":
			var arr struct {
				Lit struct {
					Val []CoreFnExpr `json:"value"`
				} `json:"value"`
			}
			if err = json.Unmarshal(data, &arr); err == nil {
				me.Array = arr.Lit.Val
			}
		case "ObjectLiteral":
		case "IntLiteral":
			me.Int = int(raw.Lit.Val.(float64))
		case "NumberLiteral":
			me.Number = raw.Lit.Val.(float64)
		case "CharLiteral":
			for _, r := range raw.Lit.Val.(string) {
				me.Char = r
				break
			}
		case "StringLiteral":
			me.String = raw.Lit.Val.(string)
		case "BooleanLiteral":
			me.Boolean = raw.Lit.Val.(bool)
		default:
			err = NotImplErr("CoreFnExprLit.Type", me.Type, string(data))
		}
	}
	return
}

type CoreFnExprObjUpd struct {
	CoreFnExprBase
	Expression CoreFnExpr    `json:"expression"`
	Updates    []interface{} `json:"updates"`
}

func (me *CoreFnExprObjUpd) prep() {
	me.CoreFnExprBase.prep()
	me.Expression.prep()
}

type CoreFnExprVar struct {
	CoreFnExprBase
	Value struct {
		ModuleName []string `json:"moduleName"`
		Identifier string   `json:"identifier"`
	} `json:"value"`
}
