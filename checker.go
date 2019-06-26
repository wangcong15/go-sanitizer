package main

import (
	"go/ast"
    "go/token"
    "log"
)

// CWE-777: Regular Expression without Anchors
func C777(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	// demo: result = append(result, assertion{file_path, 2, "123", 777})
	ast.Inspect(f, func(n ast.Node) bool {
		// work on each function
		ret, ok := n.(*ast.FuncDecl)
		if ok {
			log.Printf("==> Analyzing function in Line.%v\n", fset.Position(ret.Pos()).Line)
			ast.Inspect(ret, func(n ast.Node) bool {
				ret2, ok2 := n.(*ast.CallExpr)
				if ok2 {
					// TODO
					if ret3, ok3 := 
					log.Println(ret2.Fun)
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
