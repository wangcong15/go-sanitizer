package main

import (
	"github.com/wangcong15/go-sanitizer/argparser"
	"github.com/wangcong15/go-sanitizer/checkers"
	"github.com/wangcong15/go-sanitizer/code"
	"github.com/wangcong15/go-sanitizer/config"
)

var cfg config.Config
var cs code.CodeSlice

func main() {
	// STEP.1 parse execution params and load configurations
	argparser.Parse(&cfg)
	config.Load(&cfg)
	cfg.Validate()
	notice(cfg.GetWelcome())

	// STEP.2 load Go files
	cs = code.LoadFiles(cs, cfg.ArgPkg)
	for _, cwe := range cfg.GetCWEs() {
		for i := range cs {
			// STEP.3 call checkers to recommend assertions
			Call(checkers.CheckFuncs, cwe, &cs[i])
		}
	}
	// STEP.4 merge assertions into code
	cs.MergeAssert()
	notice(cfg.GetGoodbye())
}
