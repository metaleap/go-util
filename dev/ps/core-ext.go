package udevps

/*
We ignore most of the stuff in externs.json now:

BUT we make use of EfExports still, it covers more than
coreimp (type aliases) / corefn (exported datas with unexported ctors)
*/

type CoreExt struct {
	// SourceSpan *coreImpSourceSpan `json:"efSourceSpan"`
	// Version    string             `json:"efVersion"`
	// ModuleName []string           `json:"efModuleName"`
	// Decls      []*CoreExtDecl     `json:"efDeclarations"`
	// Imports    []*CoreExtImport   `json:"efImports"`
	Exports []*CoreExtRefs `json:"efExports"`
}

// type CoreExtImport struct {
// 	Module     []string           `json:"eiModule"`
// 	ImportType *CoreExtImportType `json:"eiImportType"`
// 	ImportedAs []string           `json:"eiImportedAs"`
// }

type CoreExtRefs struct {
	TypeRef         []interface{}
	TypeClassRef    []interface{}
	TypeInstanceRef []interface{}
	ValueRef        []interface{}
	// ValueOpRef  []interface{}
	// ModuleRef   []interface{}
	// ReExportRef []interface{}
}

// type CoreExtImportType struct {
// 	Implicit []interface{}
// 	Explicit []*CoreExtRefs
// }

// type CoreExtDecl struct {
// 	EDClass           *CoreExtTypeClass
// 	EDType            *CoreExtType
// 	EDTypeSynonym     *CoreExtTypeSyn
// 	EDValue           *CoreExtVal
// 	EDInstance        *CoreExtInst
// 	EDDataConstructor map[string]interface{}
// }

// type CoreExtIdent struct {
// 	Ident string
// }

// type CoreExtVal struct {
// 	Name CoreExtIdent  `json:"edValueName"`
// 	Type coreImpEnvTag `json:"edValueType"`
// }

// type CoreExtType struct {
// 	Name     string        `json:"edTypeName"`
// 	Kind     coreImpEnvTag `json:"edTypeKind"`
// 	DeclKind interface{}   `json:"edTypeDeclarationKind"`
// }

// type CoreExtTypeSyn struct {
// 	Name      string         `json:"edTypeSynonymName"`
// 	Arguments []interface{}  `json:"edTypeSynonymArguments"`
// 	Type      *coreImpEnvTag `json:"edTypeSynonymType"`
// }

// type CoreExtInst struct {
// 	ClassName   []interface{}   `json:"edInstanceClassName"`
// 	Name        CoreExtIdent    `json:"edInstanceName"`
// 	Types       []coreImpEnvTag `json:"edInstanceTypes"`
// 	Constraints []CoreConstr    `json:"edInstanceConstraints"`
// 	Chain       [][]interface{} `json:"edInstanceChain"`
// 	ChainIndex  int             `json:"edInstanceChainIndex"`
// }

// type CoreExtTypeClass struct {
// 	Name           string          `json:"edClassName"`
// 	TypeArgs       [][]interface{} `json:"edClassTypeArguments"`
// 	FunctionalDeps []struct {
// 		Determiners []int `json:"determiners"`
// 		Determined  []int `json:"determined"`
// 	} `json:"edFunctionalDependencies"`
// 	Members     [][]interface{} `json:"edClassMembers"`
// 	Constraints []CoreConstr    `json:"edClassConstraints"`
// }
