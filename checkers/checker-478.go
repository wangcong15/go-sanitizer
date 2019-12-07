package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars478 struct {
	x1         *ast.FuncDecl
	x2         *ast.SwitchStmt
	x3         *ast.CaseClause
	varName    string
	caseStr    string
	hasDefault bool

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check478(c *code.Code) {
	lv := localVars478{
		funcName: "AssertIntIn",
		bugType:  "478",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check478x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check478x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars478) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.SwitchStmt{})
	for _, x2 := range x2List {
		check478x2(c, x2.(*ast.SwitchStmt), lv)
	}
}

func check478x2(c *code.Code, x2 *ast.SwitchStmt, lv *localVars478) {
	lv.x2 = x2
	lv.varName = getExpr(x2.Tag)
	lv.caseStr = ""
	lv.hasDefault = false
	if lv.varName == "" {
		return
	}
	x3List := x2.Body.List
	for _, x3 := range x3List {
		if _, ok := x3.(*ast.CaseClause); ok {
			check478x3(c, x3.(*ast.CaseClause), lv)
		}
	}
	if !lv.hasDefault {
		lv.lineNo = c.Fset.Position(x2.Switch).Line
		lv.params = lv.varName + lv.caseStr
		genAssert478(c, lv)
	}
}

func check478x3(c *code.Code, x3 *ast.CaseClause, lv *localVars478) {
	lv.x3 = x3
	if !lv.hasDefault {
		if len(x3.List) == 0 {
			lv.hasDefault = true
		} else {
			for _, v := range x3.List {
				if exp1 := getExpr(v); exp1 != "" {
					lv.caseStr += ", " + exp1
				}
			}
		}
	}
}

func genAssert478(c *code.Code, lv *localVars478) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
