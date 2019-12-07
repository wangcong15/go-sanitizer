package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars777 struct {
	x1     *ast.FuncDecl
	x2     *ast.CallExpr
	x3     *ast.SelectorExpr
	x4     *ast.Ident
	x5     *ast.Ident
	x6     *ast.Ident
	x7     []ast.Expr
	x8     *ast.Ident
	valMap map[string]int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check777(c *code.Code) {
	lv := localVars777{
		funcName: "AssertStrNotIn",
		bugType:  "777",
	}
	lv.valMap = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check777x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check777x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars777) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.CallExpr{})
	for _, x2 := range x2List {
		check777x2(c, x2.(*ast.CallExpr), lv)
	}
}

func check777x2(c *code.Code, x2 *ast.CallExpr, lv *localVars777) {
	lv.x2 = x2
	x3 := x2.Fun
	if _, ok := x3.(*ast.SelectorExpr); ok {
		check777x3(c, x3.(*ast.SelectorExpr), lv)
	}
}

func check777x3(c *code.Code, x3 *ast.SelectorExpr, lv *localVars777) {
	lv.x3 = x3
	if x3.Sel.Name == "MatchString" {
		x4 := x3.X
		if _, ok := x4.(*ast.Ident); ok {
			check777x4(c, x4.(*ast.Ident), lv)
		}
	}
	if x3.Sel.Name == "Join" {
		x6 := x3.X
		if _, ok := x6.(*ast.Ident); ok {
			check777x6(c, x6.(*ast.Ident), lv)
		}
	}
}

func check777x4(c *code.Code, x4 *ast.Ident, lv *localVars777) {
	lv.x4 = x4
	if x4.Name == "regexp" {
		x5 := lv.x2.Args[1]
		if _, ok := x5.(*ast.Ident); ok {
			check777x5(c, x5.(*ast.Ident), lv)
		}
	}
}

func check777x5(c *code.Code, x5 *ast.Ident, lv *localVars777) {
	lv.x5 = x5
	lv.valMap[x5.Name] = 1
}

func check777x6(c *code.Code, x6 *ast.Ident, lv *localVars777) {
	lv.x6 = x6
	if x6.Name == "path" {
		x7 := lv.x2.Args
		check777x7(c, x7, lv)
	}
}

func check777x7(c *code.Code, x7 []ast.Expr, lv *localVars777) {
	lv.x7 = x7
	for _, x8 := range x7 {
		if _, ok := x8.(*ast.Ident); ok {
			check777x8(c, x8.(*ast.Ident), lv)
		}
	}
}

func check777x8(c *code.Code, x8 *ast.Ident, lv *localVars777) {
	lv.x8 = x8
	if lv.valMap[x8.Name] == 1 {
		lv.lineNo = c.Fset.Position(lv.x2.Lparen).Line
		lv.params = "\"..\", " + lv.x8.Name
		genAssert777(c, lv)
	}
}

func genAssert777(c *code.Code, lv *localVars777) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
