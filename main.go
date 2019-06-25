package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	p        string
	rec_chan chan assertionSlice
)

func init() {
	flag.StringVar(&p, "p", "", "set a golang package to recommend assertions")
	rec_chan = make(chan assertionSlice)
}

func main() {
	var asserts assertionSlice
	var chan_counter int
	var file_path string
	// Flag Parse
	flag.Parse()
	if p == "" {
		log.Println("==> Package path should not be empty. Use -p to set.")
		return
	}
	log.Printf("==> Package path is set: %v\n", p)

	// Read File List
	if s, err := os.Stat(p); err != nil || !s.IsDir() {
		log.Println("==> Package path does not exist")
		return
	}
	files, _ := ioutil.ReadDir(p)
	chan_counter = 0
	log.Println("==> Source file list: ")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") && !f.IsDir() {
			file_path = path.Join(p, f.Name())
			chan_counter += 1
			// handle source files in concurrent mode
			go rec(file_path)
		}
	}
	for i := 0; i < chan_counter; i++ {
		val, ok := <-rec_chan
		if ok {
			asserts = append(asserts, val...)
		}
	}
	sort.Sort(asserts)
	log.Println("==> Assertion list: ")
	for j := range asserts {
		log.Println(asserts[j])
	}
	insert(asserts)
}

// assertion recommender
func rec(file_path string) {
	var result assertionSlice
	var raw_code string

	log.Println(file_path)
	raw_code = ""

	// checkers in concurrent mode
	result = append(result, C777(raw_code, file_path)...)
	result = append(result, C478(raw_code, file_path)...)
	result = append(result, C839(raw_code, file_path)...)
	result = append(result, C486(raw_code, file_path)...)
	result = append(result, C1077(raw_code, file_path)...)
	result = append(result, C785(raw_code, file_path)...)
	result = append(result, C466(raw_code, file_path)...)
	result = append(result, C822(raw_code, file_path)...)
	result = append(result, C823(raw_code, file_path)...)
	result = append(result, C824(raw_code, file_path)...)
	result = append(result, C128(raw_code, file_path)...)
	result = append(result, C190(raw_code, file_path)...)
	result = append(result, C191(raw_code, file_path)...)

	rec_chan <- result
}

// Golang does not support transformation from ast to source code.
// Thus we insert the assertion with file R/W.
func insert(asserts assertionSlice) {
	for _, val := range asserts {
		if b, err := ioutil.ReadFile(val.file_path); err == nil {
			raw_code := string(b)
			code_arr := strings.Split(raw_code, "\n")
			if !strings.Contains(raw_code, "goassert") {
				for i := range code_arr {
					if strings.HasPrefix(code_arr[i], "package ") {
						code_arr[i] += "\nimport \"github.com/wangcong15/goassert\""
					}
				}
			}
			code_arr[val.line_no-2] += "\n"
			for j := 0; ; j++ {
				if j >= len(code_arr[val.line_no]) {
					break
				}
				if code_arr[val.line_no][j] == '\t' {
					code_arr[val.line_no-2] += "\t"
				} else {
					break
				}
			}
			code_arr[val.line_no-2] += val.expression
			new_code := strings.Join(code_arr, "\n")
			if ioutil.WriteFile(val.file_path, []byte(new_code), 0644) != nil {
				log.Printf("==> Error in writing %v\n", val.file_path)
			}
		}
	}
}
