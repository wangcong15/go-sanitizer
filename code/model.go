package code

import (
	"go/ast"
	"go/token"
)

// CodeSlice : list of struct Code
type CodeSlice []Code

// Code : the struct to record assertions
type Code struct {
	FilePath string
	Fset     *token.FileSet
	F        *ast.File
	Asserts  AssertSlice
}

// AssertSlice : list of struct Assertion
type AssertSlice []Assertion

// Assertion : the struct of generated assertions
type Assertion struct {
	FuncName string
	Params   string
	LineNo   int
	BugType  string
}
