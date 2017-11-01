package udevps

type CoreTag struct {
	Tag      string      `json:"tag"`
	Contents interface{} `json:"contents"`
}

func (_ *CoreTag) ident2qname(identtuple []interface{}) (qname string) {
	for _, m := range identtuple[0].([]interface{}) {
		qname += (m.(string) + ".")
	}
	switch x := identtuple[1].(type) {
	case map[string]string:
		qname += x["Ident"]
	case map[string]interface{}:
		qname += x["Ident"].(string)
	default:
		qname += x.(string)
	}
	return
}

func (_ *CoreTag) tagFrom(tc map[string]interface{}) CoreTag {
	return CoreTag{Tag: tc["tag"].(string), Contents: tc["contents"]}
}

type CoreTagKind struct {
	CoreTag

	Num   int          `json:"n,omitempty"`
	Text  string       `json:"t,omitempty"`
	Kind0 *CoreTagKind `json:"k0,omitempty"`
	Kind1 *CoreTagKind `json:"k1,omitempty"`
}

func (me *CoreTagKind) IsRow() bool       { return me.Tag == "Row" }
func (me *CoreTagKind) IsKUnknown() bool  { return me.Tag == "KUnknown" }
func (me *CoreTagKind) IsFunKind() bool   { return me.Tag == "FunKind" }
func (me *CoreTagKind) IsNamedKind() bool { return me.Tag == "NamedKind" }
func (me *CoreTagKind) new(tc map[string]interface{}) *CoreTagKind {
	return &CoreTagKind{CoreTag: me.tagFrom(tc), Num: -1}
}
func (me *CoreTagKind) prep() {
	if me != nil {
		//	no type assertions, arr-len checks or nil checks anywhere here: the panic signals us that the coreimp format has changed
		me.Num = -1
		if me.IsKUnknown() {
			me.Num = int(me.Contents.(float64))
		} else if me.IsRow() {
			me.Kind0 = me.new(me.Contents.(map[string]interface{}))
			me.Kind0.prep()
		} else if me.IsFunKind() {
			items := me.Contents.([]interface{})
			me.Kind0 = me.new(items[0].(map[string]interface{}))
			me.Kind0.prep()
			me.Kind1 = me.new(items[1].(map[string]interface{}))
			me.Kind1.prep()
		} else if me.IsNamedKind() {
			me.Text = me.ident2qname(me.Contents.([]interface{}))
		} else {
			panic(NotImplErr("tagged-kind", me.Tag, me.Contents))
		}
	}
}

type CoreTagType struct {
	CoreTag

	Num    int          `json:"n,omitempty"`
	Skolem int          `json:"s,omitempty"`
	Text   string       `json:"t,omitempty"`
	Type0  *CoreTagType `json:"t0,omitempty"`
	Type1  *CoreTagType `json:"t1,omitempty"`
	Constr *CoreConstr  `json:"c,omitempty"`
}

// func (me *CoreTagType) IsTypeWildcard() bool        { return me.Tag == "TypeWildcard" }
// func (me *CoreTagType) IsTypeOp() bool              { return me.Tag == "TypeOp" }
// func (me *CoreTagType) IsProxyType() bool           { return me.Tag == "ProxyType" }
// func (me *CoreTagType) IsKindedType() bool          { return me.Tag == "KindedType" }
// func (me *CoreTagType) IsPrettyPrintFunction() bool { return me.Tag == "PrettyPrintFunction" }
// func (me *CoreTagType) IsPrettyPrintObject() bool   { return me.Tag == "PrettyPrintObject" }
// func (me *CoreTagType) IsPrettyPrintForAll() bool   { return me.Tag == "PrettyPrintForAll" }
// func (me *CoreTagType) IsBinaryNoParensType() bool  { return me.Tag == "BinaryNoParensType" }
// func (me *CoreTagType) IsParensInType() bool        { return me.Tag == "ParensInType" }
// func (me *CoreTagType) IsTUnknown() bool            { return me.Tag == "TUnknown" }
func (me *CoreTagType) IsTypeLevelString() bool { return me.Tag == "TypeLevelString" }
func (me *CoreTagType) IsTypeVar() bool         { return me.Tag == "TypeVar" }
func (me *CoreTagType) IsTypeConstructor() bool { return me.Tag == "TypeConstructor" }
func (me *CoreTagType) IsSkolem() bool          { return me.Tag == "Skolem" }
func (me *CoreTagType) IsREmpty() bool          { return me.Tag == "REmpty" }
func (me *CoreTagType) IsRCons() bool           { return me.Tag == "RCons" }
func (me *CoreTagType) IsTypeApp() bool         { return me.Tag == "TypeApp" }
func (me *CoreTagType) IsForAll() bool          { return me.Tag == "ForAll" }
func (me *CoreTagType) IsConstrainedType() bool { return me.Tag == "ConstrainedType" }
func (me *CoreTagType) new(tc map[string]interface{}) *CoreTagType {
	return &CoreTagType{CoreTag: me.tagFrom(tc), Num: -1, Skolem: -1}
}
func (me *CoreTagType) prep() {
	//	no type assertions, arr-len checks or nil checks anywhere here: the panic signals us that the coreimp format has changed
	me.Skolem, me.Num = -1, -1
	if me.IsTypeVar() {
		me.Text = me.Contents.(string)
	} else if me.IsForAll() {
		tuple := me.Contents.([]interface{})
		me.Text = tuple[0].(string)
		me.Type0 = me.new(tuple[1].(map[string]interface{}))
		me.Type0.prep()
		if tuple[2] != nil {
			me.Skolem = int(tuple[2].(float64))
		}
	} else if me.IsTypeApp() {
		items := me.Contents.([]interface{})
		me.Type0 = me.new(items[0].(map[string]interface{}))
		me.Type0.prep()
		me.Type1 = me.new(items[1].(map[string]interface{}))
		me.Type1.prep()
	} else if me.IsTypeConstructor() {
		me.Text = me.ident2qname(me.Contents.([]interface{}))
	} else if me.IsConstrainedType() {
		tuple := me.Contents.([]interface{}) // eg [{constrstuff} , {type}]
		me.Type0 = me.new(tuple[1].(map[string]interface{}))
		me.Type0.prep()
		constr := tuple[0].(map[string]interface{})
		me.Constr = &CoreConstr{Data: constr["constraintData"], Class: constr["constraintClass"], Cls: me.ident2qname(constr["constraintClass"].([]interface{}))}
		for _, ca := range constr["constraintArgs"].([]interface{}) {
			carg := me.new(ca.(map[string]interface{}))
			me.Constr.Args = append(me.Constr.Args, carg)
		}
		me.Constr.prep()
	} else if me.IsSkolem() {
		tuple := me.Contents.([]interface{})
		me.Text = tuple[0].(string)
		me.Num = int(tuple[1].(float64))
		me.Skolem = int(tuple[2].(float64))
	} else if me.IsRCons() {
		tuple := me.Contents.([]interface{})
		me.Text = tuple[0].(string)
		me.Type0 = me.new(tuple[1].(map[string]interface{}))
		me.Type0.prep()
		me.Type1 = me.new(tuple[2].(map[string]interface{}))
		me.Type1.prep()
	} else if me.IsREmpty() && me.Contents == nil {
		// nothing to do
	} else if me.IsTypeLevelString() {
		me.Text = me.Contents.(string)
		// let any of these panic below if they ever start occurring, so we can handle them then:
		// } else if me.IsTypeWildcard() {
		// } else if me.IsTypeOp() {
		// } else if me.IsProxyType() {
		// } else if me.IsKindedType() {
		// } else if me.IsPrettyPrintFunction() {
		// } else if me.IsPrettyPrintObject() {
		// } else if me.IsPrettyPrintForAll() {
		// } else if me.IsBinaryNoParensType() {
		// } else if me.IsParensInType() {
		// } else if me.IsTUnknown() {
	} else {
		panic(NotImplErr("tagged-type", me.Tag, me.Contents))
	}
}
