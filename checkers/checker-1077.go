package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars1077 struct {
	x1 *ast.FuncDecl
	x2 *ast.BinaryExpr

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check1077(c *code.Code) {
	lv := localVars1077{
		funcName: "AssertPresion",
		bugType:  "1077",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check1077x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check1077x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars1077) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.BinaryExpr{})
	for _, x2 := range x2List {
		check1077x2(c, x2.(*ast.BinaryExpr), lv)
	}
}

func check1077x2(c *code.Code, x2 *ast.BinaryExpr, lv *localVars1077) {
	lv.x2 = x2
	if x2.Op.String() == "==" {
		exp1 := getExpr(x2.X)
		exp2 := getExpr(x2.Y)
		if isNormalName(exp1) && isNormalName(exp2) {
			lv.lineNo = c.Fset.Position(x2.OpPos).Line
			lv.params = exp1 + ", " + exp2
			genAssert1077(c, lv)
		}
	}
}

func genAssert1077(c *code.Code, lv *localVars1077) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
