package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars369 struct {
	x1        *ast.FuncDecl
	x2        *ast.ValueSpec
	x3        *ast.Ident
	x4        *ast.BinaryExpr
	x5        *ast.AssignStmt
	dirtyVals map[string]int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check369(c *code.Code) {
	lv := localVars369{
		funcName: "AssertNEq",
		bugType:  "369",
	}
	lv.dirtyVals = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check369x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check369x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars369) {
	lv.x1 = x1
	for _, v := range x1.Type.Params.List {
		lv.dirtyVals[getExpr(v)] = 1
	}
	x2List := Inspect(x1, &ast.ValueSpec{})
	for _, x2 := range x2List {
		check369x2(c, x2.(*ast.ValueSpec), lv)
	}
	x4List := Inspect(x1, &ast.BinaryExpr{})
	for _, x4 := range x4List {
		check369x4(c, x4.(*ast.BinaryExpr), lv)
	}
}

func check369x2(c *code.Code, x2 *ast.ValueSpec, lv *localVars369) {
	lv.x2 = x2
	x3 := x2.Type
	if _, ok := x3.(*ast.Ident); ok {
		check369x3(c, x3.(*ast.Ident), lv)
	}
}

func check369x3(c *code.Code, x3 *ast.Ident, lv *localVars369) {
	lv.x3 = x3
	if len(Inspect(lv.x2, &ast.CallExpr{})) > 0 {
		for _, v := range lv.x2.Names {
			lv.dirtyVals[v.Name] = 1
		}
	}
}

func check369x4(c *code.Code, x4 *ast.BinaryExpr, lv *localVars369) {
	lv.x4 = x4
	if x4.Op.String() == "/" {
		exp := getExpr(x4.Y)
		if lv.dirtyVals[exp] == 1 {
			lv.params = exp + ", 0"
			genAssert369(c, lv)
		}
	}
}

func genAssert369(c *code.Code, lv *localVars369) {
	lineNo := c.Fset.Position(lv.x4.OpPos).Line
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
