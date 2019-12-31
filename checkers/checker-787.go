package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars787 struct {
	x1        *ast.FuncDecl
	x2        *ast.AssignStmt
	x3        *ast.Ident
	x4        *ast.AssignStmt
	dirtyVals map[string]int

	// For Assertions
	funcName  string
	params    string
	funcName2 string
	params2   string
	lineNo    int
	bugType   string
}

func check787(c *code.Code) {
	lv := localVars787{
		funcName:  "AssertGte",
		funcName2: "AssertLt",
		bugType:   "787",
	}
	lv.dirtyVals = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check787x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check787x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars787) {
	lv.x1 = x1
	for _, v := range x1.Type.Params.List {
		lv.dirtyVals[getExpr(v)] = 1
	}
	x2List := Inspect(x1, &ast.AssignStmt{})
	for _, x2 := range x2List {
		check787x2(c, x2.(*ast.AssignStmt), lv)
	}
	x4List := Inspect(x1, &ast.AssignStmt{})
	for _, x4 := range x4List {
		check787x4(c, x4.(*ast.AssignStmt), lv)
	}
}

func check787x2(c *code.Code, x2 *ast.AssignStmt, lv *localVars787) {
	lv.x2 = x2
	x3 := x2.Rhs[0]
	if _, ok := x2.Lhs[0].(*ast.Ident); !ok {
		return
	}
	if _, ok := x3.(*ast.Ident); ok {
		check787x3(c, x3.(*ast.Ident), lv)
	}
}

func check787x3(c *code.Code, x3 *ast.Ident, lv *localVars787) {
	lv.x3 = x3
	if len(Inspect(lv.x2, &ast.CallExpr{})) > 0 {
		lv.dirtyVals[lv.x2.Lhs[0].(*ast.Ident).Name] = 1
	}
}

func check787x4(c *code.Code, x4 *ast.AssignStmt, lv *localVars787) {
	lv.x4 = x4
	if _, ok := x4.Lhs[0].(*ast.IndexExpr); !ok {
		return
	}
	x4_ := x4.Lhs[0].(*ast.IndexExpr)
	if _, ok := x4_.X.(*ast.Ident); !ok {
		return
	}
	if _, ok := x4_.Index.(*ast.Ident); !ok {
		return
	}
	lv.params = x4_.Index.(*ast.Ident).Name + ", 0"
	lv.params2 = x4_.Index.(*ast.Ident).Name + ", len(" + x4_.X.(*ast.Ident).Name + ")"
	lv.lineNo = c.Fset.Position(x4_.Lbrack).Line
	genAssert787(c, lv)
}

func genAssert787(c *code.Code, lv *localVars787) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
	newAssert2 := code.Assertion{
		FuncName: lv.funcName2,
		Params:   lv.params2,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert2)
}
