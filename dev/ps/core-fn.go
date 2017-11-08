package udevps

import (
	"encoding/json"
	"strings"
)

type CoreFn struct {
	CoreModuleRef
	Imports []struct {
		CoreAnnotated
		CoreModuleRef
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

type CoreFnIdent struct {
	CoreModuleRef
	Identifier string `json:"identifier"`
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
	CoreAnnotated
	Identifier string     `json:"identifier"`
	Expression CoreFnExpr `json:"expression"`
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

	prep   func()
	String func() string
}

func (me *CoreFnExpr) UnmarshalJSON(data []byte) (err error) {
	var exprtype struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(data, &exprtype); err == nil {
		me.String = func() string { return exprtype.Type }
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
	CoreAnnotated
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
	Expressions  []CoreFnExpr        `json:"caseExpressions"`
	Alternatives []CoreFnExprCaseAlt `json:"caseAlternatives"`
}

func (me *CoreFnExprCase) prep() {
	me.CoreFnExprBase.prep()
	for _, expr := range me.Expressions {
		expr.prep()
	}
	for _, alt := range me.Alternatives {
		alt.prep()
	}
}

type CoreFnExprCaseAlt struct {
	Binders     []CoreFnBinder `json:"binders"`
	IsGuarded   bool           `json:"isGuarded"`
	Expression  *CoreFnExpr    `json:"expression"`
	Expressions []struct {
		Guard      *CoreFnExpr `json:"guard"`
		Expression *CoreFnExpr `json:"expression"`
	} `json:"expressions"`
}

func (me *CoreFnExprCaseAlt) prep() {
	for _, b := range me.Binders {
		b.prep()
	}
	if me.Expression != nil {
		me.Expression.prep()
	}
	for _, expr := range me.Expressions {
		expr.Guard.prep()
		expr.Expression.prep()
	}
}

type CoreFnBinder struct {
	CoreAnnotated
	BinderType string `json:"binderType"`

	Identifier string            `json:"identifier"`
	Literal    *CoreFnExprLitVal `json:"literal"`

	CtorName CoreFnIdent `json:"constructorName"`
	CtorType CoreFnIdent `json:"typeName"`

	CtorBinders []CoreFnBinder `json:"binders"`
	Named       *CoreFnBinder  `json:"binder"`
}

func (me *CoreFnBinder) prep() {
	if me.Literal != nil {
		me.Literal.prep()
	}
	for _, cb := range me.CtorBinders {
		cb.prep()
	}
	if me.Named != nil {
		me.Named.prep()
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
	Val CoreFnExprLitVal `json:"value"`
}

func (me *CoreFnExprLit) prep() {
	me.Val.prep()
}

type CoreFnExprLitVal struct {
	Type           string                `json:"-"`
	Number         float64               `json:"-"`
	Int            int                   `json:"-"`
	Boolean        bool                  `json:"-"`
	Char           rune                  `json:"-"`
	String         string                `json:"-"`
	Array          []CoreFnExpr          `json:"-"`
	ArrayOfBinders []CoreFnBinder        `json:"-"`
	Obj            []CoreFnExprLitObjFld `json:"-"`
}

func (me *CoreFnExprLitVal) prep() {
	for _, a := range me.Array {
		a.prep()
	}
	for _, ab := range me.ArrayOfBinders {
		ab.prep()
	}
	for _, o := range me.Obj {
		o.prep()
	}
}

func (me *CoreFnExprLitVal) UnmarshalJSON(data []byte) (err error) {
	var raw struct {
		Type string `json:"literalType"`
	}
	if err = json.Unmarshal(data, &raw); err == nil {
		me.Type = raw.Type
		switch me.Type {
		case "ArrayLiteral":
			var arr struct {
				Val []CoreFnExpr `json:"value"`
			}
			if err = json.Unmarshal(data, &arr); err == nil {
				me.Array = arr.Val
			} else {
				var binders struct {
					Val []CoreFnBinder `json:"value"`
				}
				if err = json.Unmarshal(data, &binders); err == nil {
					me.ArrayOfBinders = binders.Val
				}
			}
		case "ObjectLiteral":
			var obj struct {
				Val []CoreFnExprLitObjFld `json:"value"`
			}
			if err = json.Unmarshal(data, &obj); err == nil {
				me.Obj = obj.Val
			}
		default:
			var prim struct {
				Val interface{} `json:"value"`
			}
			if err = json.Unmarshal(data, &prim); err == nil {
				switch me.Type {
				case "IntLiteral":
					me.Int = int(prim.Val.(float64))
				case "NumberLiteral":
					me.Number = prim.Val.(float64)
				case "CharLiteral":
					for _, r := range prim.Val.(string) {
						me.Char = r
						break
					}
				case "StringLiteral":
					me.String = prim.Val.(string)
				case "BooleanLiteral":
					me.Boolean = prim.Val.(bool)
				default:
					err = NotImplErr("CoreFnExprLit.Type", me.Type, string(data))
				}
			}
		}
	}
	return
}

type CoreFnExprLitObjFld struct {
	Name   string        `json:"-"`
	Val    *CoreFnExpr   `json:"-"`
	Binder *CoreFnBinder `json:"-"`
}

func (me *CoreFnExprLitObjFld) prep() {
	if me.Val != nil {
		me.Val.prep()
	} else if me.Binder != nil {
		me.Binder.prep()
	}
}

func (me *CoreFnExprLitObjFld) UnmarshalJSON(data []byte) (err error) {
	hacky := string(data)
	if strings.HasPrefix(hacky, "[\"") && strings.HasSuffix(hacky, "}]") {
		if i := strings.Index(hacky, "\",{\""); i > 0 {
			me.Name = hacky[:i][2:]
			data = []byte(hacky[:len(hacky)-1][i+2:])
			if err = json.Unmarshal(data, me.Val); err != nil {
				err = json.Unmarshal(data, &me.Binder)
			}
		}
	}
	return
}

type CoreFnExprObjUpd struct {
	CoreFnExprBase
	Expression CoreFnExpr            `json:"expression"`
	Updates    []CoreFnExprLitObjFld `json:"updates"`
}

func (me *CoreFnExprObjUpd) prep() {
	me.CoreFnExprBase.prep()
	me.Expression.prep()
	for _, upd := range me.Updates {
		upd.prep()
	}
}

type CoreFnExprVar struct {
	CoreFnExprBase
	Value CoreFnIdent `json:"value"`
}
