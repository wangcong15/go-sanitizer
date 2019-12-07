package checkers

import (
	"go/ast"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars466 struct {
	x1      *ast.FuncDecl
	x2      *ast.ValueSpec
	x3      *ast.AssignStmt
	ptrVals map[string]int
	valList []string

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check466(c *code.Code) {
	lv := localVars466{
		funcName: "AssertNNil",
		bugType:  "466",
	}
	lv.ptrVals = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check466x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check466x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars466) {
	lv.x1 = x1
	x2List := Inspect(x1, &ast.ValueSpec{})
	for _, x2 := range x2List {
		check466x2(c, x2.(*ast.ValueSpec), lv)
	}
	x3List := Inspect(x1, &ast.AssignStmt{})
	for _, x3 := range x3List {
		check466x3(c, x3.(*ast.AssignStmt), lv)
	}
}

func check466x2(c *code.Code, x2 *ast.ValueSpec, lv *localVars466) {
	lv.x2 = x2
	if _, ok := x2.Type.(*ast.StarExpr); ok {
		for _, v := range x2.Names {
			lv.ptrVals[getExpr(v)] = 1
		}
	}
}

func check466x3(c *code.Code, x3 *ast.AssignStmt, lv *localVars466) {
	lv.x3 = x3
	lv.valList = []string{}
	for _, v := range x3.Lhs {
		lv.valList = append(lv.valList, getExpr(v))
	}
	for i, v := range x3.Rhs {
		if _, ok := v.(*ast.CallExpr); ok && isNormalName(lv.valList[i]) && lv.ptrVals[lv.valList[i]] == 1 {
			lv.params = lv.valList[i]
			lv.lineNo = c.Fset.Position(x3.TokPos).Line + 1
			genAssert466(c, lv)
		}
	}
}

func genAssert466(c *code.Code, lv *localVars466) {
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lv.lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
