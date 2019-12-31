package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars131 struct {
	x1     *ast.FuncDecl
	x2     *ast.AssignStmt
	x3     *ast.Ident
	x4     *ast.CallExpr
	x5     *ast.CallExpr
	x6     *ast.Ident
	x7     *ast.Ident
	valMap map[string]int

	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check131(c *code.Code) {
	lv := localVars131{
		funcName: "AssertGt",
		bugType:  "131",
	}
	lv.valMap = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check131x1(c, x1.(*ast.FuncDecl), &lv)
	}
}
func check131x2(c *code.Code, x2 *ast.AssignStmt, lv *localVars131) {
	lv.x2 = x2
	x3 := x2.Lhs[0]
	if _, ok := x3.(*ast.Ident); ok {
		check131x3(c, x3.(*ast.Ident), lv)
	}
	x4 := x2.Rhs[0]
	if _, ok := x4.(*ast.CallExpr); ok {
		check131x4(c, x4.(*ast.CallExpr), lv)
	}
}
func check131x3(c *code.Code, x3 *ast.Ident, lv *localVars131) {
	lv.x3 = x3

}
func check131x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars131) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.AssignStmt{})
	for _, x2 := range x2List {
		check131x2(c, x2.(*ast.AssignStmt), lv)
	}
	x5List := Inspect(x1, &ast.CallExpr{})
	for _, x5 := range x5List {
		check131x5(c, x5.(*ast.CallExpr), lv)
	}
}
func check131x6(c *code.Code, x6 *ast.Ident, lv *localVars131) {
	lv.x6 = x6

}
func check131x7(c *code.Code, x7 *ast.Ident, lv *localVars131) {
	lv.x7 = x7
	if lv.valMap[lv.x7.Name] == 1 {
		genAssert131(c, lv)
	}
}
func check131x4(c *code.Code, x4 *ast.CallExpr, lv *localVars131) {
	lv.x4 = x4
	lv.valMap[lv.x3.Name] = 1
}
func check131x5(c *code.Code, x5 *ast.CallExpr, lv *localVars131) {
	lv.x5 = x5
	x6 := x5.Fun
	if _, ok := x6.(*ast.Ident); ok {
		check131x6(c, x6.(*ast.Ident), lv)
	}
	if len(x5.Args) == 2 && lv.x6.Name == "make" {
		x7 := x5.Args[1]
		if _, ok := x7.(*ast.Ident); ok {
			check131x7(c, x7.(*ast.Ident), lv)
		}
	}
}
func genAssert131(c *code.Code, lv *localVars131) {
	lineNo := c.Fset.Position(lv.x5.Lparen).Line
	params := lv.x7.Name + ", 0"
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   params,
		LineNo:   lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
