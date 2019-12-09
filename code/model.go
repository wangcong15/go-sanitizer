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
	newCodeLines []string
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

type sortByLine struct { AssertSlice }

func (a sortByLine) Less(i, j int) bool {
	return a.AssertSlice[i].LineNo > a.AssertSlice[j].LineNo
}
func (p AssertSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p AssertSlice) Len() int { return len(p) }
