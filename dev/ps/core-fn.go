package udevps

type CoreFn struct {
	ModuleName []string `json:"moduleName"` // ["Control","Monad","Eff"]
	ModulePath string   `json:"modulePath"` // /home/_/c/ps/gonad-test/src/Mini/DataType.purs   ,   /home/_/c/ps/gonad-test/bower_components/purescript-eff/src/Control/Monad/Eff.purs
	Imports    []struct {
		Ann        *CoreAnn `json:"annotation"`
		ModuleName []string `json:"moduleName"`
	} `json:"imports"`
}

func (me *CoreFn) prep() {
	for _, imp := range me.Imports {
		imp.Ann.prep()
	}
}
