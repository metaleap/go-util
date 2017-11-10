package udevps

import (
	"fmt"
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
	for i, _ := range me.Imports {
		me.Imports[i].Annotation.prep()
	}
	for i, _ := range me.Decls {
		me.Decls[i].prep()
	}
}

func (me *CoreFn) RemoveAt(i int) {
	me.Decls = append(me.Decls[:i], me.Decls[i+1:]...)
}

type CoreFnBinder struct {
	CoreAnnotated
	BinderType string `json:"binderType"`

	Identifier string            `json:"identifier"`
	Literal    *CoreFnExprLitVal `json:"literal"`

	CtorName *CoreFnIdent `json:"constructorName"`
	CtorType *CoreFnIdent `json:"typeName"`

	CtorBinders []CoreFnBinder `json:"binders"`
	Named       *CoreFnBinder  `json:"binder"`
}

func (me *CoreFnBinder) prep() {
	if me.Literal != nil {
		me.Literal.prep()
	}
	for i, _ := range me.CtorBinders {
		me.CtorBinders[i].prep()
	}
	if me.Named != nil {
		me.Named.prep()
	}
}

func (me *CoreFnBinder) String() (s string) {
	s = fmt.Sprintf("❬B:%s`%s`", me.BinderType, me.Identifier)
	if me.Literal != nil {
		s += fmt.Sprintf(" L:%s", me.Literal.String())
	}
	if me.CtorName != nil {
		s += fmt.Sprintf(" Cn:%s", me.CtorName.String())
	}
	if me.CtorType != nil {
		s += fmt.Sprintf(" Ct:%s", me.CtorType.String())
	}
	if len(me.CtorBinders) > 0 {
		s += " Cb:["
		for i, _ := range me.CtorBinders {
			s += ", " + me.CtorBinders[i].String()
		}
		s += "]"
	}
	if me.Named != nil {
		s += fmt.Sprintf(" N:%s", me.Named.String())
	}
	s += " B:" + me.BinderType + "❭"
	return
}

type CoreFnDecl struct {
	*CoreFnDeclBind                  // directly upon decode may be nil or non-nil, after prep() ALWAYS nil
	Binds           []CoreFnDeclBind `json:"binds"`
}

func (me *CoreFnDecl) IsRecursive() bool    { return me.CoreFnDeclBind == nil }
func (me *CoreFnDecl) IsNonRecursive() bool { return me.CoreFnDeclBind != nil }

func (me *CoreFnDecl) prep() {
	if me.CoreFnDeclBind != nil {
		me.Binds = []CoreFnDeclBind{*me.CoreFnDeclBind}
		me.CoreFnDeclBind = nil
	}
	for i, _ := range me.Binds {
		me.Binds[i].prep()
	}
}

func (me *CoreFnDecl) String() (s string) {
	if me.CoreFnDeclBind != nil {
		s = me.CoreFnDeclBind.String()
	} else {
		s = "["
		for i, _ := range me.Binds {
			s += ", " + me.Binds[i].String()
		}
		s += "]"
	}
	return
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

func (me *CoreFnDeclBind) String() string {
	return fmt.Sprintf("%s{ %s }", me.Identifier, me.Expression.String())
}

type CoreFnIdent struct {
	CoreModuleRef
	Identifier string `json:"identifier"`
}

func (me *CoreFnIdent) String() string {
	return strings.Join(append(me.ModuleName, me.Identifier), ".")
}
