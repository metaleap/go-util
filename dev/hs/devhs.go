package udevhs

import (
	"github.com/metaleap/go-util/run"
)

var (
	StackVersion string

	Has_hindent         bool
	Has_brittany        bool
	Has_stylish_haskell bool
	Has_hlint           bool
	Has_ghcmod          bool
	Has_pointfree       bool
	Has_pointful        bool
	Has_hsimport        bool
	Has_htrefact        bool
	Has_htdaemon        bool
	Has_hare            bool
	Has_hasktags        bool
	Has_hothasktags     bool
	Has_lushtags        bool
	Has_deadcodedetect  bool
	Has_hsautofix       bool
	Has_hscabal         bool
	Has_hsclearimports  bool
	Has_hsdev           bool
	Has_hshayoo         bool
	Has_hsinspect       bool
	Has_intero          bool
	Has_doctest         bool
	Has_hoogle          bool
	Has_apply_refact    bool

	StackArgs      = []string{"--dump-logs", "--no-time-in-log", "--no-install-ghc", "--skip-ghc-check", "--skip-msys", "--no-terminal", "--color", "never", "--jobs", "8", "--verbosity", "info"}
	StackArgsBuild = []string{"--copy-bins", "--no-haddock", "--no-open", "--no-haddock-internal", "--no-haddock-deps", "--no-keep-going", "--no-test", "--no-rerun-tests", "--no-bench", "--no-run-benchmarks", "--no-cabal-verbose", "--no-split-objs"}
)

func HasHsDevEnv() bool {
	var cmdout string
	var err error

	if len(StackVersion) > 0 {
		return true
	}
	if cmdout, err = urun.CmdExec("stack", "--numeric-version", "--no-terminal", "--color", "never"); err == nil && len(cmdout) > 0 {
		StackVersion = cmdout

		urun.CmdsTryStart(map[string]*urun.CmdTry{
			"ghc-mod":             {Ran: &Has_ghcmod, Args: []string{"--version"}},
			"ghc-hare":            {Ran: &Has_hare, Args: []string{"--version"}},
			"hsimport":            {Ran: &Has_hsimport, Args: []string{"--version"}},
			"hasktags":            {Ran: &Has_hasktags, Args: []string{"--help"}},
			"lushtags":            {Ran: &Has_lushtags, Args: []string{"--help"}},
			"hothasktags":         {Ran: &Has_hothasktags, Args: []string{"--help"}},
			"dead-code-detection": {Ran: &Has_deadcodedetect, Args: []string{"--version"}},
			"pointfree":           {Ran: &Has_pointfree},
			"pointful":            {Ran: &Has_pointful},
			"refactor":            {Ran: &Has_apply_refact},
			"hoogle":              {Ran: &Has_hoogle, Args: []string{"--version"}},
			"hlint":               {Ran: &Has_hlint, Args: []string{"--version"}},
			"doctest":             {Ran: &Has_doctest, Args: []string{"--version"}},
			"intero":              {Ran: &Has_intero, Args: []string{"--version"}},
			"hindent":             {Ran: &Has_hindent, Args: []string{"--version"}},
			"brittany":            {Ran: &Has_brittany, Args: []string{"--version"}},
			"stylish-haskell":     {Ran: &Has_stylish_haskell, Args: []string{"--version"}},
			"ht-refact":           {Ran: &Has_htrefact},
			"ht-daemon":           {Ran: &Has_htdaemon, Args: []string{"hows'it hangin holmes"}},
		})

	}
	return (len(StackVersion) > 0)
}
