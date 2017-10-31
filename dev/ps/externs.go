package udevps

/*
We ignore most of the stuff in externs.json now, as it's
provided in coreimp's DeclEnv for both exports & non-exports.

BUT we make use of EfExports still, as coreimp's `exports`
don't capture type synonyms.
*/

type Extern struct {
	// EfSourceSpan *coreImpSourceSpan `json:"efSourceSpan"`
	// EfVersion    string             `json:"efVersion"`
	// EfModuleName []string           `json:"efModuleName"`
	// EfDecls      []*ExternDecl       `json:"efDeclarations"`
	// EfImports    []*ExternImport     `json:"efImports"`
	EfExports []*ExternRefs `json:"efExports"`
}

// type ExternImport struct {
// 	EiModule     []string         `json:"eiModule"`
// 	EiImportType *ExternImportType `json:"eiImportType"`
// 	EiImportedAs []string         `json:"eiImportedAs"`
// }

type ExternRefs struct {
	TypeRef         []interface{}
	TypeClassRef    []interface{}
	TypeInstanceRef []interface{}
	ValueRef        []interface{}
	ValueOpRef      []interface{}
	ModuleRef       []interface{}
	ReExportRef     []interface{}
}

// type ExternImportType struct {
// 	Implicit []interface{}
// 	Explicit []*ExternRefs
// }

// type ExternDecl struct {
// 	EDClass           *ExternTypeClass
// 	EDType            *ExternType
// 	EDTypeSynonym     *ExternTypeAlias
// 	EDValue           *ExternVal
// 	EDInstance        *ExternInst
// 	EDDataConstructor map[string]interface{}
// }

// type ExternIdent struct {
// 	Ident string
// }

// type ExternVal struct {
// 	Name ExternIdent    `json:"edValueName"`
// 	Type coreImpEnvTag `json:"edValueType"`
// }

// type ExternType struct {
// 	Name     string        `json:"edTypeName"`
// 	Kind     coreImpEnvTag `json:"edTypeKind"`
// 	DeclKind interface{}   `json:"edTypeDeclarationKind"`
// }

// type ExternTypeAlias struct {
// 	Name      string         `json:"edTypeSynonymName"`
// 	Arguments []interface{}  `json:"edTypeSynonymArguments"`
// 	Type      *coreImpEnvTag `json:"edTypeSynonymType"`
// }

// type ExternInst struct {
// 	ClassName   []interface{}   `json:"edInstanceClassName"`
// 	Name        ExternIdent      `json:"edInstanceName"`
// 	Types       []coreImpEnvTag `json:"edInstanceTypes"`
// 	Constraints []CoreConstr   `json:"edInstanceConstraints"`
// 	Chain       [][]interface{} `json:"edInstanceChain"`
// 	ChainIndex  int             `json:"edInstanceChainIndex"`
// }

// type ExternTypeClass struct {
// 	Name           string          `json:"edClassName"`
// 	TypeArgs       [][]interface{} `json:"edClassTypeArguments"`
// 	FunctionalDeps []struct {
// 		Determiners []int `json:"determiners"`
// 		Determined  []int `json:"determined"`
// 	} `json:"edFunctionalDependencies"`
// 	Members     [][]interface{} `json:"edClassMembers"`
// 	Constraints []CoreConstr   `json:"edClassConstraints"`
// }
