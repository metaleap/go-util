package udevps

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

	prep       func()
	Annotation func() *CoreAnnotation
	String     func() string
}

func (me *CoreFnExpr) UnmarshalJSON(data []byte) (err error) {
	var exprtype struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(data, &exprtype); err == nil {
		setann := func(ann *CoreAnnotation) {
			me.Annotation = func() *CoreAnnotation { return ann }
		}
		switch exprtype.Type {
		case "Abs":
			var abs CoreFnExprAbs
			if err = json.Unmarshal(data, &abs); err == nil {
				me.prep, me.Abs, me.String = abs.prep, &abs, abs.String
				setann(&abs.Annotation)
			}
		case "Accessor":
			var acc CoreFnExprAcc
			if err = json.Unmarshal(data, &acc); err == nil {
				me.prep, me.Accessor, me.String = acc.prep, &acc, acc.String
				setann(&acc.Annotation)
			}
		case "App":
			var app CoreFnExprApp
			if err = json.Unmarshal(data, &app); err == nil {
				me.prep, me.App, me.String = app.prep, &app, app.String
				setann(&app.Annotation)
			}
		case "Case":
			var cªse CoreFnExprCase
			if err = json.Unmarshal(data, &cªse); err == nil {
				me.prep, me.Case, me.String = cªse.prep, &cªse, cªse.String
				setann(&cªse.Annotation)
			}
		case "Constructor":
			var ctor CoreFnExprCtor
			if err = json.Unmarshal(data, &ctor); err == nil {
				me.prep, me.Constructor, me.String = ctor.prep, &ctor, ctor.String
				setann(&ctor.Annotation)
			}
		case "Let":
			var let CoreFnExprLet
			if err = json.Unmarshal(data, &let); err == nil {
				me.prep, me.Let, me.String = let.prep, &let, let.String
				setann(&let.Annotation)
			}
		case "Literal":
			var lit CoreFnExprLit
			if err = json.Unmarshal(data, &lit); err == nil {
				me.prep, me.Literal, me.String = lit.prep, &lit, lit.String
				setann(&lit.Annotation)
			}
		case "ObjectUpdate":
			var objupd CoreFnExprObjUpd
			if err = json.Unmarshal(data, &objupd); err == nil {
				me.prep, me.ObjectUpdate, me.String = objupd.prep, &objupd, objupd.String
				setann(&objupd.Annotation)
			}
		case "Var":
			var vªr CoreFnExprVar
			if err = json.Unmarshal(data, &vªr); err == nil {
				me.prep, me.Var, me.String = vªr.prep, &vªr, vªr.String
				setann(&vªr.Annotation)
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

func (me *CoreFnExprAbs) Meta() *CoreAnnotationMeta {
	return me.Annotation.Meta
}

func (me *CoreFnExprAbs) String() string {
	return "ABS:\\" + me.Argument + "-> " + me.Body.String()
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
func (me *CoreFnExprAcc) String() string {
	return "ACC:" + me.Expression.String() + "@" + me.FieldName
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
func (me *CoreFnExprApp) String() string {
	return "ABS:" + me.Abstraction.String() + "(" + me.Argument.String() + ")"
}

type CoreFnExprCase struct {
	CoreFnExprBase
	Expressions  []CoreFnExpr        `json:"caseExpressions"`
	Alternatives []CoreFnExprCaseAlt `json:"caseAlternatives"`
}

func (me *CoreFnExprCase) prep() {
	me.CoreFnExprBase.prep()
	for i, _ := range me.Expressions {
		me.Expressions[i].prep()
	}
	for i, _ := range me.Alternatives {
		me.Alternatives[i].prep()
	}
}

func (me *CoreFnExprCase) String() string {
	s := "CASE:"
	if len(me.Expressions) > 1 {
		println(len(me.Expressions))
	}
	for i, _ := range me.Expressions {
		if i > 0 {
			s += ", "
		}
		s += me.Expressions[i].String()
	}
	s += ":OF:"
	for i, _ := range me.Alternatives {
		s += me.Alternatives[i].String()
	}
	return s
}

type CoreFnExprCaseAlt struct {
	Binders     []CoreFnBinder `json:"binders"`
	IsGuarded   bool           `json:"isGuarded"`
	Expression  *CoreFnExpr    `json:"expression"`
	Expressions []struct {
		Guard      CoreFnExpr `json:"guard"`
		Expression CoreFnExpr `json:"expression"`
	} `json:"expressions"`
}

func (me *CoreFnExprCaseAlt) prep() {
	for i, _ := range me.Binders {
		me.Binders[i].prep()
	}
	if me.Expression != nil {
		me.Expression.prep()
	}
	for i, _ := range me.Expressions {
		me.Expressions[i].Guard.prep()
		me.Expressions[i].Expression.prep()
	}
}

func (me *CoreFnExprCaseAlt) String() string {
	s := fmt.Sprintf(" ❬C:%v| ", me.IsGuarded)
	for i, _ := range me.Binders {
		s += "B:" + me.Binders[i].String() + " "
	}
	if me.Expression != nil {
		s += me.Expression.String()
	} else {
		s += "EXPRS["
		for i, _ := range me.Expressions {
			s += "X:" + me.Expressions[i].Guard.String() + " |?| " + me.Expressions[i].Expression.String() + " "
		}
		s += "]"
	}
	return s + " |C:❭"
}

type CoreFnExprCtor struct {
	CoreFnExprBase
	ConstructorName string   `json:"constructorName"`
	TypeName        string   `json:"typeName"`
	FieldNames      []string `json:"fieldNames"`
}

func (me *CoreFnExprCtor) String() string {
	s := me.TypeName + "::" + me.ConstructorName + "<"
	for i, fn := range me.FieldNames {
		if i > 0 {
			s += ", "
		}
		s += fn
	}
	return s + ">"
}

type CoreFnExprLet struct {
	CoreFnExprBase
	Binds      []CoreFnDecl `json:"binds"`
	Expression CoreFnExpr   `json:"expression"`
}

func (me *CoreFnExprLet) prep() {
	me.CoreFnExprBase.prep()
	me.Expression.prep()
	for i, _ := range me.Binds {
		me.Binds[i].prep()
	}
}
func (me *CoreFnExprLet) String() string {
	s := "LET:"
	for i, _ := range me.Binds {
		if i > 0 {
			s += ", "
		}
		s += me.Binds[i].String()
	}
	return s + ":IN:" + me.Expression.String()
}

type CoreFnExprLit struct {
	CoreFnExprBase
	Val CoreFnExprLitVal `json:"value"`
}

func (me *CoreFnExprLit) prep() {
	me.Val.prep()
}

func (me *CoreFnExprLit) Meta() *CoreAnnotationMeta {
	return me.Annotation.Meta
}

func (me *CoreFnExprLit) String() string {
	return me.Val.String()
}

type CoreFnExprLitVal struct {
	Type           string                `json:"-"`
	Number         float64               `json:"-"`
	Int            int                   `json:"-"`
	Boolean        bool                  `json:"-"`
	Char           rune                  `json:"-"`
	Str            string                `json:"-"`
	Array          []CoreFnExpr          `json:"-"`
	ArrayOfBinders []CoreFnBinder        `json:"-"`
	Obj            []CoreFnExprLitObjFld `json:"-"`
}

func (me *CoreFnExprLitVal) prep() {
	for i, _ := range me.Array {
		me.Array[i].prep()
	}
	for i, _ := range me.ArrayOfBinders {
		me.ArrayOfBinders[i].prep()
	}
	for i, _ := range me.Obj {
		me.Obj[i].prep()
	}
}

func (me *CoreFnExprLitVal) String() string {
	switch me.Type {
	case "ArrayLiteral":
		s := "La:["
		for i, _ := range me.Array {
			if i > 0 {
				s += ", "
			}
			s += me.Array[i].String()
		}
		for i, _ := range me.ArrayOfBinders {
			if i > 0 {
				s += ", "
			}
			s += me.ArrayOfBinders[i].String()
		}
		return s + "]"
	case "ObjectLiteral":
		s := "Lo:{"
		for i, _ := range me.Obj {
			if i > 0 {
				s += ", "
			}
			s += me.Obj[i].String()
		}
		return s + "}"
	case "IntLiteral":
		return fmt.Sprintf("Li:%d", me.Int)
	case "NumberLiteral":
		return fmt.Sprintf("Ln:%f", me.Number)
	case "CharLiteral":
		return fmt.Sprintf("Lc:%q", me.Char)
	case "StringLiteral":
		return fmt.Sprintf("Ls:%q", me.Str)
	case "BooleanLiteral":
		return fmt.Sprintf("Lb:%t", me.Boolean)
	default:
		panic(NotImplErr("CoreFnExprLit.Type", me.Type, *me))
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
					me.Str = prim.Val.(string)
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

func (me *CoreFnExprLitObjFld) String() (s string) {
	s = me.Name + ":"
	if me.Val != nil {
		s += me.Val.String()
	} else if me.Binder != nil {
		s += me.Binder.String()
	}
	return
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
	for i, _ := range me.Updates {
		me.Updates[i].prep()
	}
}

func (me *CoreFnExprObjUpd) String() string {
	s := "UPDOBJ:" + me.Expression.String() + "{"
	for i, _ := range me.Updates {
		if i > 0 {
			s += ", "
		}
		s += me.Updates[i].String()
	}
	return s + "}"
}

type CoreFnExprVar struct {
	CoreFnExprBase
	Value CoreFnIdent `json:"value"`
}

func (me *CoreFnExprVar) Meta() *CoreAnnotationMeta {
	return me.Annotation.Meta
}

func (me *CoreFnExprVar) String() string {
	return "V:" + me.Value.String()
}
