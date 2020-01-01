package checkers

import (
	"go/ast"
	"reflect"
	"strings"
)

// Inspect : search AST nodes
func Inspect(f ast.Node, t interface{}) []ast.Node {
	nodes := []ast.Node{}
	ast.Inspect(f, func(n ast.Node) bool {
		if reflect.TypeOf(t) == reflect.TypeOf(n) {
			nodes = append(nodes, n)
		}
		return true
	})
	return nodes
}

func isWrapFunc(funcName string) bool {
	return (funcName == "int8" || funcName == "int16" || funcName == "int32" || funcName == "uint8" || funcName == "uint16" || funcName == "uint32")
}

func getExpr(X ast.Node) (exp1 string) {
	if x1, ok := X.(*ast.Ident); ok {
		exp1 = x1.Name
	} else if x2, ok := X.(*ast.BasicLit); ok {
		exp1 = x2.Value
	}
	return
}

func isNormalName(name string) bool {
	return (name != "nil" && name != "" && name != "err" && name != "e" && name != "_" && !strings.Contains(name, "\""))
}

func hasBool(names ...string) bool {
	for _, n := range names {
		if n == "true" || n == "false" {
			return true
		}
	}
	return false
}