package checkers

import (
	"go/ast"
	"go/token"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars193 struct {
	x1 *ast.FuncDecl
	x2 *ast.ForStmt
	x3 *ast.BinaryExpr
	x4 *ast.Ident
	x5 *ast.BlockStmt
	x6 *ast.IndexExpr
	x8 *ast.Ident
	x9 *ast.Ident
	x7 *ast.Ident

	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check193(c *code.Code) {
	lv := localVars193{
		funcName: "AssertLt",
		bugType:  "193",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check193x1(c, x1.(*ast.FuncDecl), &lv)
	}
}
func check193x8(c *code.Code, x8 *ast.Ident, lv *localVars193) {
	lv.x8 = x8
	x9 := lv.x6.X
	if _, ok := x9.(*ast.Ident); ok {
		check193x9(c, x9.(*ast.Ident), lv)
	}
}
func check193x9(c *code.Code, x9 *ast.Ident, lv *localVars193) {
	lv.x9 = x9
	x7 := lv.x6.Index
	if _, ok := x7.(*ast.Ident); ok {
		check193x7(c, x7.(*ast.Ident), lv)
	}
}
func check193x2(c *code.Code, x2 *ast.ForStmt, lv *localVars193) {
	lv.x2 = x2
	x3 := x2.Cond
	if _, ok := x3.(*ast.BinaryExpr); ok {
		check193x3(c, x3.(*ast.BinaryExpr), lv)
	}
}
func check193x3(c *code.Code, x3 *ast.BinaryExpr, lv *localVars193) {
	lv.x3 = x3
	if x3.Op == token.LEQ {
		x4 := x3.X
		if _, ok := x4.(*ast.Ident); ok {
			check193x4(c, x4.(*ast.Ident), lv)
		}
	}
}
func check193x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars193) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.ForStmt{})
	for _, x2 := range x2List {
		check193x2(c, x2.(*ast.ForStmt), lv)
	}
}
func check193x6(c *code.Code, x6 *ast.IndexExpr, lv *localVars193) {
	lv.x6 = x6
	x8 := lv.x3.Y
	if _, ok := x8.(*ast.Ident); ok {
		check193x8(c, x8.(*ast.Ident), lv)
	}
}
func check193x7(c *code.Code, x7 *ast.Ident, lv *localVars193) {
	lv.x7 = x7
	if x7.Name == lv.x4.Name {
		genAssert193(c, lv)
	}
}
func check193x4(c *code.Code, x4 *ast.Ident, lv *localVars193) {
	lv.x4 = x4
	x5 := lv.x2.Body
	check193x5(c, x5, lv)
}
func check193x5(c *code.Code, x5 *ast.BlockStmt, lv *localVars193) {
	lv.x5 = x5
	x6List := Inspect(x5, &ast.IndexExpr{})
	for _, x6 := range x6List {
		check193x6(c, x6.(*ast.IndexExpr), lv)
	}
}
func genAssert193(c *code.Code, lv *localVars193) {
	lineNo := c.Fset.Position(lv.x2.For).Line
	params := lv.x8.Name + ", len(" + lv.x9.Name + ")"
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   params,
		LineNo:   lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
