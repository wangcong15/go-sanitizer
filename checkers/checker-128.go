package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars128 struct {
	x1      *ast.FuncDecl
	x2      *ast.AssignStmt
	x3      []ast.Expr
	x4      *ast.Ident
	x5      []ast.Expr
	x6      *ast.CallExpr
	x7      *ast.Ident
	x8      *ast.Ident
	valList []string
	idx     int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check128(c *code.Code) {
	lv := localVars128{
		funcName: "AssertValEq",
		bugType:  "128",
	}
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check128x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check128x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars128) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.AssignStmt{})
	for _, x2 := range x2List {
		check128x2(c, x2.(*ast.AssignStmt), lv)
	}
}

func check128x2(c *code.Code, x2 *ast.AssignStmt, lv *localVars128) {
	lv.x2 = x2
	x3 := x2.Lhs
	check128x3(c, x3, lv)
	x5 := x2.Rhs
	check128x5(c, x5, lv)
}

func check128x3(c *code.Code, x3 []ast.Expr, lv *localVars128) {
	lv.x3 = x3
	for _, x4 := range x3 {
		if _, ok := x4.(*ast.Ident); ok {
			check128x4(c, x4.(*ast.Ident), lv)
		} else {
			lv.valList = append(lv.valList, "")
		}
	}
}

func check128x4(c *code.Code, x4 *ast.Ident, lv *localVars128) {
	lv.x4 = x4
	lv.valList = append(lv.valList, x4.Name)
}

func check128x5(c *code.Code, x5 []ast.Expr, lv *localVars128) {
	lv.x5 = x5
	for idx, x6 := range x5 {
		lv.idx = idx
		if _, ok := x6.(*ast.CallExpr); ok {
			check128x6(c, x6.(*ast.CallExpr), lv)
		}
	}
}

func check128x6(c *code.Code, x6 *ast.CallExpr, lv *localVars128) {
	lv.x6 = x6
	if len(x6.Args) > 0 {
		x7 := x6.Fun
		if _, ok := x7.(*ast.Ident); ok {
			check128x7(c, x7.(*ast.Ident), lv)
		}
	}
}

func check128x7(c *code.Code, x7 *ast.Ident, lv *localVars128) {
	lv.x7 = x7
	if isWrapFunc(x7.Name) {
		x8 := lv.x6.Args[0]
		if _, ok := x8.(*ast.Ident); ok {
			check128x8(c, x8.(*ast.Ident), lv)
		}
	}
}

func check128x8(c *code.Code, x8 *ast.Ident, lv *localVars128) {
	lv.x8 = x8
	if lv.valList[lv.idx] != "" {
		genAssert128(c, lv)
	}
}

func genAssert128(c *code.Code, lv *localVars128) {
	lineNo := c.Fset.Position(lv.x2.TokPos).Line + 1
	params := lv.valList[lv.idx] + ", " + lv.x8.Name
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   params,
		LineNo:   lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
