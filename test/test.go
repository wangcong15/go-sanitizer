package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/wangcong15/go-sanitizer/checkers"
	"github.com/wangcong15/go-sanitizer/code"
)

func main() {
	// testAST()
	testChecker()
}

func testChecker() {
	src := `package main
	import "log"
	func main() {
		i := 1
		hello := 2
		world := 3
	}`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", src, parser.ParseComments)
	ast.Print(fset, f)
	log.Println(checkers.Inspect(f, &ast.FuncDecl{}))
}

func testAST() {
	src := `package main
	import "log"
	func main() {
		i := 1
		hello := 2
		world := 3
	}`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	// ast.Print(fset, f)

	// ADD IMPORT
	log.Println(code.GofmtFile(fset, f))
	log.Println("------")
	code.AddImport(fset, f, "\"github.com/wangcong15/goassert\"")
	// log.Println(gofmtFile(fset, f))

	// ADD ASSERT
	// log.Println(gofmtFile(fset, f))
	// log.Println("------")
	code.AddAssert(fset, f, code.Assertion{"AssertNNil", "hello, world", 5, "777"})
	code.AddAssert(fset, f, code.Assertion{"AssertNNil", "hello", 6, "777"})
	log.Println(code.GofmtFile(fset, f))
}
