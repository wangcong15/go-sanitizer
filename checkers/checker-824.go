package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars824 struct {
	x1        *ast.FuncDecl
	x2        *ast.ValueSpec
	x3        *ast.AssignStmt
	x4        *ast.CallExpr
	varStatus map[string]int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check824(c *code.Code) {
	lv := localVars824{
		funcName: "AssertNNil",
		bugType:  "824",
	}
	lv.varStatus = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check824x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check824x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars824) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.ValueSpec{})
	for _, x2 := range x2List {
		check824x2(c, x2.(*ast.ValueSpec), lv)
	}
	x3List := Inspect(x1, &ast.AssignStmt{})
	for _, x3 := range x3List {
		check824x3(c, x3.(*ast.AssignStmt), lv)
	}
	x4List := Inspect(x1, &ast.CallExpr{})
	for _, x4 := range x4List {
		check824x4(c, x4.(*ast.CallExpr), lv)
	}
}

func check824x2(c *code.Code, x2 *ast.ValueSpec, lv *localVars824) {
	lv.x2 = x2
	if _, ok := x2.Type.(*ast.StarExpr); ok {
		for _, name := range x2.Names {
			lv.varStatus[getExpr(name)] = 1
		}
	}
}

func check824x3(c *code.Code, x3 *ast.AssignStmt, lv *localVars824) {
	lv.x3 = x3
	for _, name := range x3.Lhs {
		tmpName := getExpr(name)
		if lv.varStatus[tmpName] == 1 {
			lv.varStatus[tmpName] = 2
		}
	}
}

func check824x4(c *code.Code, x4 *ast.CallExpr, lv *localVars824) {
	lv.x4 = x4
	for _, name := range x4.Args {
		tmpName := getExpr(name)
		if lv.varStatus[tmpName] == 1 {
			lv.lineNo = c.Fset.Position(x4.Rparen).Line
			lv.params = tmpName
			genAssert824(c, lv)
		}
	}
}

func genAssert824(c *code.Code, lv *localVars824) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
