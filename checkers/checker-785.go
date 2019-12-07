package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars785 struct {
	x1 *ast.FuncDecl
	x2 *ast.CallExpr

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check785(c *code.Code) {
	lv := localVars785{
		funcName: "AssertGte",
		bugType:  "785",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check785x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check785x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars785) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.CallExpr{})
	for _, x2 := range x2List {
		check785x2(c, x2.(*ast.CallExpr), lv)
	}
}

func check785x2(c *code.Code, x2 *ast.CallExpr, lv *localVars785) {
	lv.x2 = x2
	if getExpr(x2.Fun) == "copy" {
		args := x2.Args
		exp1 := getExpr(args[0])
		exp2 := getExpr(args[1])
		if exp1 != "" && exp2 != "" {
			lv.lineNo = c.Fset.Position(x2.Lparen).Line
			lv.params = "len(" + exp1 + "), len(" + exp2 + "))"
			genAssert785(c, lv)
		}
	}
}

func genAssert785(c *code.Code, lv *localVars785) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
