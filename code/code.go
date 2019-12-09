package code

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
	"log"
	"fmt"
	"sort"
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
	for i, v := range cs {
		for _, ass := range v.Asserts {
			AddAssert(v.Fset, v.F, ass)
		}
		if len(v.Asserts) > 0 {
			AddImport(v.Fset, v.F, "\"github.com/wangcong15/goassert\"")
		}
		if newCode, err := GofmtFile(v.Fset, v.F); err == nil {
			cs[i].newCodeLines = strings.Split(newCode, "\n")
			ioutil.WriteFile(v.FilePath, []byte(newCode), 0644)
		}
	}
}

// Dump : Print assertions and save metadata
func (cs CodeSlice) Dump(dumpFile string) {
	stringToWrite := ""
	lineNo := 0
	for _, v := range cs {
		log.Printf("文件 : %v", v.FilePath)
		sort.Sort(sortByLine{v.Asserts})
		lineNo = 0
		for _, assert := range v.Asserts {
			lineNo += 1
			statement := fmt.Sprintf("goassert.%v(%v)", assert.FuncName, assert.Params)
			log.Printf("CWE-%v，第%v行，断言语句为%v", assert.BugType, assert.LineNo, statement)

			lineNo = adjustLineNo(v.newCodeLines, lineNo)
			stringToWrite += fmt.Sprintf("%v\t%v\t%v\t%v\n", assert.BugType, lineNo+1, statement, v.FilePath)
		}
	}
	ioutil.WriteFile(dumpFile, []byte(stringToWrite), 0644)
}

func adjustLineNo(lines []string, startLine int) int {
	result := startLine
	for {
		if (result >= len(lines) || strings.Contains(lines[result], "goassert.")) {
			break
		}
		result += 1
	}
	if result == len(lines) {
		return startLine
	}
	return result
}