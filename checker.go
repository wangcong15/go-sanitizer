package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func Reflect(node ast.Node) {
	value := reflect.ValueOf(node)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Println(fmt.Sprintf("method[%d]%s and type is %v", i, typ.Method(i).Name, typ.Method(i).Type))
	}
}

// CWE-777: Regular Expression without Anchors
func C777(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var val1 map[string]int = make(map[string]int)
	var expr string
	var location int
	var weak_id int = 777
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.CallExpr); ok {
					if ret3, ok := ret2.Fun.(*ast.SelectorExpr); ok {
						if ret4, ok := ret3.X.(*ast.Ident); ok && ret4.Name == "regexp" && ret3.Sel.Name == "MatchString" {
							ret5 := ret2.Args[1]
							if ret6, ok := ret5.(*ast.Ident); ok {
								val1[ret6.Name] = 1
							}
						}
					}
				}
				// C3
				if ret2, ok := n2.(*ast.CallExpr); ok {
					if ret3, ok := ret2.Fun.(*ast.SelectorExpr); ok {
						if ret4, ok := ret3.X.(*ast.Ident); ok && ret4.Name == "path" && ret3.Sel.Name == "Join" {
							ret5 := ret2.Args
							for _, arg := range ret5 {
								if ret6, ok := arg.(*ast.Ident); ok {
									if val1[ret6.Name] == 1 {
										location = fset.Position(ret2.Lparen).Line
										expr = "goassert.AssertStrNotIn(\"..\", " + ret6.Name + ")"
										// NEW ASSERTION
										result = append(result, assertion{file_path, location, expr, weak_id})
										break
									}
								}
							}
						}
					}
				}
				return true
			})
		}
		return true
	})
	return
}

// CWE-478: Missing Default Case in Switch Statement
func C478(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-839: Numeric Range Comparison Without Minimum Check
func C839(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-486: Comparison of Classes by Name
func C486(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-1077: Floating Point Comparison with Incorrect Operator
func C1077(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-785: Use of Path Manipulation Function without Maximum-sized Buffer
func C785(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-466: Return of Pointer Value Outside of Expected Range
func C466(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-822: Untrusted Pointer Dereference
func C822(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-823: Use of Out-of-range Pointer Offset
func C823(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-824: Access of Uninitialized Pointer
func C824(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-128: Wrap-around Error
func C128(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-190: Integer Overflow or Wraparound
func C190(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}

// CWE-191: Integer Underflow (Wrap or Wraparound)
func C191(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	return
}
