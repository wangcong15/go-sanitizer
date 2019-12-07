package code

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
)

// LoadFiles : load Golang file
func LoadFiles(cs CodeSlice, pkg string) CodeSlice {
	files, _ := ioutil.ReadDir(pkg)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") && !f.IsDir() {
			filePath := path.Join(pkg, f.Name())
			b, _ := ioutil.ReadFile(filePath)
			rawCode := string(b)
			fset := token.NewFileSet()
			f, _ := parser.ParseFile(fset, "", rawCode, parser.ParseComments)
			cs = append(cs, Code{
				FilePath: filePath,
				Fset:     fset,
				F:        f,
				Asserts:  AssertSlice{},
			})
		}
	}
	return cs
}

// MergeAssert : Merge assertions into source files
func (cs CodeSlice) MergeAssert() {
	for _, v := range cs {
		for _, ass := range v.Asserts {
			AddAssert(v.Fset, v.F, ass)
		}
		if len(v.Asserts) > 0 {
			AddImport(v.Fset, v.F, "\"github.com/wangcong15/goassert\"")
		}
		if newCode, err := GofmtFile(v.Fset, v.F); err == nil {
			ioutil.WriteFile(v.FilePath, []byte(newCode), 0644)
		}
	}
}
