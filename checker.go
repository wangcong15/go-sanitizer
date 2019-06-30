package main

import (
	"go/ast"
	"go/token"
)

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
					cases = []string{}
					flag = true
					for _, c := range ret2.Body.List {
						if ret3, ok := c.(*ast.CaseClause); ok {
							for _, e := range ret3.List {
								new_case_exp := getExpr(e)
								if new_case_exp == "default" {
									flag = false
									break
								} else {
									cases = append(cases, new_case_exp)
								}
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
						expr = "goassert.AssertPresion(" + exp1 + ", " + exp2 + ")"
						location = fset.Position(ret2.OpPos).Line
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
	var var_list []string
	var expr string
	var location int
	var weak_id int = 466
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
			ast.Inspect(ret, func(n2 ast.Node) bool {
				// C2
				if ret2, ok := n2.(*ast.AssignStmt); ok {
					var_list = []string{}
					for _, v := range ret2.Lhs {
						var_list = append(var_list, getExpr(v))
					}
					for i, v := range ret2.Rhs {
						if _, ok := v.(*ast.CallExpr); ok {
							expr = "goassert.AssertNNil(" + var_list[i] + ")"
							location = fset.Position(ret2.TokPos).Line
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
					for _, name := range ret2.Names {
						uninit_vars[getExpr(name)] = 1
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
							expr = "goassert.AssertNValEq(" + temp_name + ", nil)"
							location = fset.Position(ret3.Lparen).Line
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
						}
					}
					ret5 := ret2.Rhs
					// C3
					for idx, args := range ret5 {
						if ret6, ok := args.(*ast.CallExpr); ok {
							if ret7, ok := ret6.Fun.(*ast.Ident); ok {
								if ret7.Name == "int8" || ret7.Name == "int16" || ret7.Name == "int32" {
									if ret8, ok := ret6.Args[0].(*ast.Ident); ok {
										expr = "goassert.AssertValEq(" + val_list[idx] + ", " + ret8.Name + ")"
										// NEW ASSERTION
										result = append(result, assertion{file_path, location, expr, weak_id})
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
						if dirty_vals[exp1] == 1 || dirty_vals[exp2] == 1 {
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
	var weak_id int = 190
	ast.Inspect(f, func(n1 ast.Node) bool {
		// C1: scope=ast.FuncDecl
		if ret, ok := n1.(*ast.FuncDecl); ok {
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
						if dirty_vals[exp1] == 1 || dirty_vals[exp2] == 1 {
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
