package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strconv"
	"strings"
)

func gofmtFile(fset *token.FileSet, f *ast.File) (string, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func AddImport(fset *token.FileSet, f *ast.File) string {
	newImport := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"github.com/wangcong15/goassert\"",
		},
	}

	hasGoassert := false
	var tempPos token.Pos
	ast.Inspect(f, func(n1 ast.Node) bool {
		ret, ok := n1.(*ast.GenDecl)
		if ok && ret.Tok == token.IMPORT {
			for _, v := range ret.Specs {
				if ret2, ok := v.(*ast.ImportSpec); ok {
					if ret3 := ret2.Path; ok && ret3.Kind == token.STRING && ret3.Value == "\"github.com/wangcong15/goassert\"" {
						hasGoassert = true
						return false
					} else {
						tempPos = ret2.Pos()
					}
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

	raw_code, err := gofmtFile(fset, f)
	if err != nil {
		panic(err)
	}
	return raw_code
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
			val1 = make(map[string]int)
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
	var variable string
	var cases []string
	var flag bool
	var expr string
	var location int
	var weak_id int = 478
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				if ret2, ok := n2.(*ast.SwitchStmt); ok {
					variable = getExpr(ret2.Tag)
					if variable == "" {
						return true
					}
					cases = []string{}
					flag = true
					for _, c := range ret2.Body.List {
						if ret3, ok := c.(*ast.CaseClause); ok {
							if len(ret3.List) == 0 {
								flag = false
								break
							}
							for _, e := range ret3.List {
								cases = append(cases, getExpr(e))
							}
						}
					}
					if flag {
						cases_to_str := cases[0]
						for idx, val := range cases {
							if idx > 0 {
								cases_to_str += ", " + val
							}
						}
						expr = "goassert.AssertIntIn(" + variable + ", []int{" + cases_to_str + "})"
						location = fset.Position(ret2.Switch).Line
						// NEW ASSERTION
						result = append(result, assertion{file_path, location, expr, weak_id})
					}
				}
				return true
			})
		}
		return true
	})
	return
}

// CWE-1077: Floating Point Comparison with Incorrect Operator
func C1077(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var exp1 string
	var exp2 string
	var expr string
	var location int
	var weak_id int = 777
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.BinaryExpr); ok {
					if checkBinaryExprOp(ret2, "==") {
						exp1 = getExpr(ret2.X)
						exp2 = getExpr(ret2.Y)
						if exp1 != "" && exp1 != "err" && exp2 != "" && exp2 != "nil" {
							if _, err := strconv.Atoi(exp2); err == nil {
								return true
							}
							expr = "goassert.AssertPresion(" + exp1 + ", " + exp2 + ")"
							location = fset.Position(ret2.OpPos).Line + 1
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

// CWE-785: Use of Path Manipulation Function without Maximum-sized Buffer
func C785(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var exp1 string
	var exp2 string
	var expr string
	var location int
	var weak_id int = 785
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.CallExpr); ok {
					if getExpr(ret2.Fun) == "copy" {
						args := ret2.Args
						exp1 = getExpr(args[0])
						exp2 = getExpr(args[1])
						if exp1 != "" && exp2 != "" {
							expr = "goassert.AssertGte(len(" + exp1 + "), len(" + exp2 + "))"
							location = fset.Position(ret2.Lparen).Line
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

// CWE-466: Return of Pointer Value Outside of Expected Range
func C466(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var ptr_dict map[string]int = make(map[string]int)
	var var_list []string
	var expr string
	var location int
	var weak_id int = 466
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2 collect pointers
				if ret5, ok := n2.(*ast.ValueSpec); ok {
					if _, ok := ret5.Type.(*ast.StarExpr); ok {
						for _, v := range ret5.Names {
							ptr_dict[getExpr(v)] = 1
						}
					}
				}

				// C3
				if ret2, ok := n2.(*ast.AssignStmt); ok {
					var_list = []string{}
					for _, v := range ret2.Lhs {
						var_list = append(var_list, getExpr(v))
					}
					for i, v := range ret2.Rhs {
						if _, ok := v.(*ast.CallExpr); ok && var_list[i] != "" && var_list[i] != "err" && var_list[i] != "e" && var_list[i] != "_" && ptr_dict[var_list[i]] == 1 {
							expr = "goassert.AssertNNil(" + var_list[i] + ")"
							ret3 := ret2.Rhs[len(ret2.Rhs)-1]
							if ret4, ok := ret3.(*ast.CallExpr); ok {
								location = fset.Position(ret4.Rparen).Line + 1
							} else {
								location = fset.Position(ret2.TokPos).Line + 1
							}
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

// CWE-824: Access of Uninitialized Pointer
func C824(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var uninit_vars map[string]int = make(map[string]int)
	var expr string
	var location int
	var weak_id int = 824

	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope = ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {

				// C2
				if ret2, ok := n2.(*ast.ValueSpec); ok && len(ret2.Values) == 0 {
					if _, ok := ret2.Type.(*ast.StarExpr); ok {
						for _, name := range ret2.Names {
							uninit_vars[getExpr(name)] = 1
						}
					}
				}
				// C3
				if ret3, ok := n2.(*ast.AssignStmt); ok {
					for _, name := range ret3.Lhs {
						temp_name := getExpr(name)
						if uninit_vars[temp_name] == 1 {
							uninit_vars[temp_name] = 2
						}
					}
				}
				// C4
				if ret3, ok := n2.(*ast.CallExpr); ok {
					for _, name := range ret3.Args {
						temp_name := getExpr(name)
						if uninit_vars[temp_name] == 1 {
							expr = "goassert.AssertNNil(" + temp_name + ")"
							location = fset.Position(ret3.Rparen).Line + 1
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

// CWE-128: Wrap-around Error
func C128(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var val_list []string
	var expr string
	var location int
	var weak_id int = 128
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.AssignStmt); ok {
					val_list = []string{}
					location = fset.Position(ret2.TokPos).Line + 1
					ret3 := ret2.Lhs
					for _, args := range ret3 {
						if ret4, ok := args.(*ast.Ident); ok {
							val_list = append(val_list, ret4.Name)
						} else {
							val_list = append(val_list, "")
						}
					}
					ret5 := ret2.Rhs
					// C3
					for idx, args := range ret5 {
						if ret6, ok := args.(*ast.CallExpr); ok {
							if ret7, ok := ret6.Fun.(*ast.Ident); ok {
								if ret7.Name == "int8" || ret7.Name == "int16" || ret7.Name == "int32" || ret7.Name == "uint8" || ret7.Name == "uint16" || ret7.Name == "uint32" {
									if len(ret6.Args) > 0 {
										if ret8, ok := ret6.Args[0].(*ast.Ident); ok {
											if val_list[idx] != "" {
												expr = "goassert.AssertValEq(" + val_list[idx] + ", " + ret8.Name + ")"
												// NEW ASSERTION
												result = append(result, assertion{file_path, location, expr, weak_id})
											}
										}
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

// CWE-190: Integer Overflow or Wraparound
func C190(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var dirty_vals map[string]int = make(map[string]int)
	var exp1 string
	var exp2 string
	var expr string
	var location int
	var weak_id int = 190
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			tys := ret.Type
			for _, pp := range tys.Params.List {
				for _, nn := range pp.Names {
					dirty_vals[getExpr(nn)] = 1
				}
			}

			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.ValueSpec); ok {
					if ret3, ok := ret2.Type.(*ast.Ident); ok {
						if isInteger(ret3.Name) && hasCallExpr(ret2) {
							for _, arg := range ret2.Names {
								dirty_vals[arg.Name] = 1
							}
						}
					}
				}
				// C3
				if ret4, ok := n2.(*ast.BinaryExpr); ok {
					if checkBinaryExprOp(n2, "+") {
						exp1 = getExpr(ret4.X)
						exp2 = getExpr(ret4.Y)
						if (dirty_vals[exp1] == 1 || dirty_vals[exp2] == 1) && !strings.Contains(exp1, "\"") && !strings.Contains(exp2, "\"") && exp1 != "" && exp2 != "" {
							location = fset.Position(ret4.OpPos).Line
							expr = "goassert.AssertOverflow(" + exp1 + ", " + exp2 + ", " + exp1 + "+" + exp2 + ")"
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

// CWE-191: Integer Underflow (Wrap or Wraparound)
func C191(fset *token.FileSet, f *ast.File, file_path string) (result assertionSlice) {
	var dirty_vals map[string]int = make(map[string]int)
	var exp1 string
	var exp2 string
	var expr string
	var location int
	var weak_id int = 191
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			tys := ret.Type
			for _, pp := range tys.Params.List {
				for _, nn := range pp.Names {
					dirty_vals[getExpr(nn)] = 1
				}
			}
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.ValueSpec); ok {
					if ret3, ok := ret2.Type.(*ast.Ident); ok {
						if isInteger(ret3.Name) && hasCallExpr(ret2) {
							for _, arg := range ret2.Names {
								dirty_vals[arg.Name] = 1
							}
						}
					}
				}
				// C3
				if ret4, ok := n2.(*ast.BinaryExpr); ok {
					if checkBinaryExprOp(n2, "-") {
						exp1 = getExpr(ret4.X)
						exp2 = getExpr(ret4.Y)
						if (dirty_vals[exp1] == 1 || dirty_vals[exp2] == 1) && !strings.Contains(exp2, "\"") && exp1 != "" && exp2 != "" {
							location = fset.Position(ret4.OpPos).Line
							expr = "goassert.AssertUnderflow(" + exp1 + ", " + exp2 + ", " + exp1 + "-" + exp2 + ")"
							// NEW ASSERTION
							result = append(result, assertion{file_path, location, expr, weak_id})
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

func isInteger(typeName string) bool {
	return typeName == "int8" || typeName == "int16" || typeName == "int32"
}

func hasCallExpr(node ast.Node) bool {
	result := false
	ast.Inspect(node, func(n1 ast.Node) bool {
		if _, ok := n1.(*ast.CallExpr); ok {
			result = true
			return false
		}
		return true
	})
	return result
}

func checkBinaryExprOp(node ast.Node, op string) bool {
	if ret, ok := node.(*ast.BinaryExpr); ok {
		if ret.Op.String() == op {
			return true
		}
	}
	return false
}

func getExpr(X ast.Node) (exp1 string) {
	if ret5, ok := X.(*ast.Ident); ok {
		exp1 = ret5.Name
	} else if ret6, ok := X.(*ast.BasicLit); ok {
		exp1 = ret6.Value
	}
	return
}
