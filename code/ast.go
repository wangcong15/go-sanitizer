package code

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
)

// GofmtFile : generate raw code file from AST
func GofmtFile(fset *token.FileSet, f *ast.File) (string, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

// AddImport : add `goassert` package into imports
func AddImport(fset *token.FileSet, f *ast.File, importPath string) {
	newImport := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: importPath,
		},
	}
	hasGoassert := false
	var tempPos token.Pos
	ast.Inspect(f, func(n1 ast.Node) bool {
		ret, ok := n1.(*ast.GenDecl)
		if ok && ret.Tok == token.IMPORT {
			for _, v := range ret.Specs {
				if ret2, ok := v.(*ast.ImportSpec); ok {
					if ret3 := ret2.Path; ok && ret3.Kind == token.STRING && ret3.Value == importPath {
						hasGoassert = true
						return false
					}
					tempPos = ret2.Pos()
				}
			}
			if !hasGoassert {
				newImport.Path.ValuePos = tempPos
				newImport.EndPos = tempPos
				f.Imports = append(f.Imports, newImport)
				ret.Specs = append(ret.Specs, newImport)
			}
		}
		return true
	})
}

// AddAssert : add assertion statement into AST
func AddAssert(fset *token.FileSet, f *ast.File, ass Assertion) {
	newAssert := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: "goassert",
				},
				Sel: &ast.Ident{
					Name: ass.FuncName,
				},
			},
			Args: []ast.Expr{&ast.Ident{Name: ass.Params}},
		},
	}

	ast.Inspect(f, func(n1 ast.Node) bool {
		// Try to find a block statement
		if ret, ok := n1.(*ast.BlockStmt); ok {
			// filter the block
			if fset.Position(ret.Lbrace).Line <= ass.LineNo && fset.Position(ret.Rbrace).Line >= ass.LineNo {
				var idx int
				for i, v := range ret.List {
					var tempLineNo, maxLineNo int
					ast.Inspect(v, func(n ast.Node) bool {
						if ret2, ok := n.(*ast.Ident); ok {
							tempLineNo = fset.Position(ret2.NamePos).Line
							if tempLineNo > maxLineNo {
								maxLineNo = tempLineNo
							}
							return true
						}
						return true
					})
					if maxLineNo >= ass.LineNo {
						idx = i
						break
					}
					idx = i + 1
				}
				rear := append([]ast.Stmt{}, ret.List[idx:]...)
				ret.List = append(ret.List[:idx], newAssert)
				ret.List = append(ret.List, rear...)
				return false
			}
			return true
		}
		return true
	})
}
