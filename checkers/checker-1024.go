package checkers

import (
	"go/ast"
	"strconv"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars1024 struct {
	x1 *ast.FuncDecl
	x2 *ast.BinaryExpr

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check1024(c *code.Code) {
	lv := localVars1024{
		funcName: "AssertStrNotIn",
		bugType:  "1024",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check1024x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check1024x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars1024) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.BinaryExpr{})
	for _, x2 := range x2List {
		check1024x2(c, x2.(*ast.BinaryExpr), lv)
	}
}

func check1024x2(c *code.Code, x2 *ast.BinaryExpr, lv *localVars1024) {
	lv.x2 = x2
	if x2.Op.String() == "==" {
		exp1 := getExpr(x2.X)
		exp2 := getExpr(x2.Y)
		if _, err := strconv.Atoi(exp2); err == nil || hasBool(exp1, exp2) {
			return
		} else if isNormalName(exp1) && isNormalName(exp2) {
			lv.lineNo = c.Fset.Position(x2.OpPos).Line
			lv.params = "\"float\", reflect.TypeOf(" + exp1 + ").String()"
			genAssert1024(c, lv)
		}
	}
}

func genAssert1024(c *code.Code, lv *localVars1024) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
