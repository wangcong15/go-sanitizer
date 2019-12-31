package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars127 struct {
	x1        *ast.FuncDecl
	x2        *ast.AssignStmt
	x3        *ast.SliceExpr
	dirtyVals map[string]int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check127(c *code.Code) {
	lv := localVars127{
		funcName: "AssertGte",
		bugType:  "127",
	}
	lv.dirtyVals = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check127x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check127x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars127) {
	lv.x1 = x1
	for _, v := range x1.Type.Params.List {
		lv.dirtyVals[getExpr(v)] = 1
	}
	x2List := Inspect(x1, &ast.AssignStmt{})
	for _, x2 := range x2List {
		check127x2(c, x2.(*ast.AssignStmt), lv)
	}
	x3List := Inspect(x1, &ast.SliceExpr{})
	for _, x3 := range x3List {
		check127x3(c, x3.(*ast.SliceExpr), lv)
	}
}

func check127x2(c *code.Code, x2 *ast.AssignStmt, lv *localVars127) {
	lv.x2 = x2
	if _, ok := x2.Lhs[0].(*ast.Ident); !ok {
		return
	}
	x3 := x2.Rhs[0]
	if _, ok := x3.(*ast.CallExpr); ok {
		lv.dirtyVals[x2.Lhs[0].(*ast.Ident).Name] = 1
	}
}

func check127x3(c *code.Code, x3 *ast.SliceExpr, lv *localVars127) {
	lv.x3 = x3
	if _, ok := x3.Low.(*ast.Ident); !ok {
		return
	}
	indexName := x3.Low.(*ast.Ident).Name
	if lv.dirtyVals[indexName] == 1 {
		lv.params = indexName + ", 0"
		lv.lineNo = c.Fset.Position(x3.Lbrack).Line
		genAssert127(c, lv)
	}
}

func genAssert127(c *code.Code, lv *localVars127) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
