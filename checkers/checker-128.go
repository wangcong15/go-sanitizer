package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars128 struct {
	x1      *ast.FuncDecl
	x2      *ast.AssignStmt
	x3      string
	x4      *ast.CallExpr
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
	x3 := x2.Lhs[0]
	if _, ok := x3.(*ast.Ident); !ok {
		return
	}
	lv.x3 = x3.(*ast.Ident).Name
	x4 := x2.Rhs[0]
	if _, ok := x4.(*ast.CallExpr); ok {
		check128x4(c, x4.(*ast.CallExpr), lv)
	}
}

func check128x4(c *code.Code, x4 *ast.CallExpr, lv *localVars128) {
	lv.x4 = x4
	if len(x4.Args) > 0 {
		x5 := x4.Fun
		if _, ok := x5.(*ast.Ident); !ok {
			return
		}
		if isWrapFunc(x5.(*ast.Ident).Name) {
			x6 := x4.Args[0]
			if _, ok := x6.(*ast.Ident); ok {
				lv.params = x6.(*ast.Ident).Name + ", " + lv.x3
				lv.lineNo = c.Fset.Position(lv.x2.TokPos).Line + 1
				genAssert128(c, lv)
			}
		}
	}
}

func genAssert128(c *code.Code, lv *localVars128) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
