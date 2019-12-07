package checkers

import (
	"go/ast"
	"strings"

	"github.com/wangcong15/go-sanitizer/code"
)

type localVars191 struct {
	x1        *ast.FuncDecl
	x2        *ast.ValueSpec
	x3        *ast.Ident
	x4        *ast.BinaryExpr
	dirtyVals map[string]int

	// For Assertions
	funcName string
	params   string
	lineNo   int
	bugType  string
}

func check191(c *code.Code) {
	lv := localVars191{
		funcName: "AssertUnderflow",
		bugType:  "191",
	}
	lv.dirtyVals = make(map[string]int)
	x1List := Inspect(c.F, &ast.FuncDecl{})
	for _, x1 := range x1List {
		check191x1(c, x1.(*ast.FuncDecl), &lv)
	}
}

func check191x1(c *code.Code, x1 *ast.FuncDecl, lv *localVars191) {
	lv.x1 = x1
	for _, v := range x1.Type.Params.List {
		lv.dirtyVals[getExpr(v)] = 1
	}
	x2List := Inspect(x1, &ast.ValueSpec{})
	for _, x2 := range x2List {
		check191x2(c, x2.(*ast.ValueSpec), lv)
	}
	x4List := Inspect(x1, &ast.BinaryExpr{})
	for _, x4 := range x4List {
		check191x4(c, x4.(*ast.BinaryExpr), lv)
	}
}

func check191x2(c *code.Code, x2 *ast.ValueSpec, lv *localVars191) {
	lv.x2 = x2
	x3 := x2.Type
	if _, ok := x3.(*ast.Ident); ok {
		check191x3(c, x3.(*ast.Ident), lv)
	}
}

func check191x3(c *code.Code, x3 *ast.Ident, lv *localVars191) {
	lv.x3 = x3
	if isWrapFunc(x3.Name) && len(Inspect(lv.x2, &ast.CallExpr{})) > 0 {
		for _, v := range lv.x2.Names {
			lv.dirtyVals[v.Name] = 1
		}
	}
}

func check191x4(c *code.Code, x4 *ast.BinaryExpr, lv *localVars191) {
	lv.x4 = x4
	if x4.Op.String() == "-" {
		exp1 := getExpr(x4.X)
		exp2 := getExpr(x4.Y)
		if (lv.dirtyVals[exp1] == 1 || lv.dirtyVals[exp2] == 1) && (exp1 != "" && exp2 != "") && (!strings.Contains(exp1, "\"") && !strings.Contains(exp2, "\"")) {
			lv.params = exp1 + ", " + exp2 + ", " + exp1 + "-" + exp2
			genAssert191(c, lv)
		}
	}
}

func genAssert191(c *code.Code, lv *localVars191) {
	lineNo := c.Fset.Position(lv.x4.OpPos).Line
	newAssert := code.Assertion{
		FuncName: lv.funcName,
		Params:   lv.params,
		LineNo:   lineNo,
		BugType:  lv.bugType,
	}
	c.Asserts = append(c.Asserts, newAssert)
}
