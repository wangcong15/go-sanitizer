package main

import (
	"flag"
	"log"
	"os"
	"io/ioutil"
	"strings"
	"path"
)

var (
	p string
	rec_chan chan []assertion
	checker_chan chan []assertion
)

type assertion struct {
	file_path string
	line_no int
	expression string
	weakness_type int
}

func init() {
	flag.StringVar(&p, "p", "", "set a golang package to recommend assertions")
	rec_chan = make(chan []assertion)
	checker_chan = make(chan []assertion)
}

func main() {
	var asserts []assertion
	var chan_counter int
	var file_path string
	// Flag Parse
	flag.Parse()
	if p == "" {
		log.Println("Package path should not be empty. Use -p to set.")
		return
	}
	log.Printf("Package path is set: %v\n", p)

	// Read File List
	if s, err := os.Stat(p); err != nil || !s.IsDir() {
		log.Println("Package path does not exist")
		return
	}
	files, _ := ioutil.ReadDir(p)
	chan_counter = 0
	for _, f := range files {
        if strings.HasSuffix(f.Name(), ".go") && !f.IsDir() {
        	file_path = path.Join(p, f.Name())
        	chan_counter += 1
        	// handle source files in concurrent mode
        	go rec(file_path)
        }
	}
	for i := 0; i < chan_counter; i++ {
		val, ok := <- rec_chan
		if ok {
			asserts = append(asserts, val...)
		}
	}
	log.Println(asserts)
}

func rec(file_path string) {
	var result []assertion
	var raw_code string

	log.Println(file_path)
	raw_code = file_path

	// checkers in concurrent mode
	go C777(raw_code)

	for i := 0; i < 1; i++ {
		val, ok := <- checker_chan
		if ok {
			result = append(result, val...)
		}
	}

	rec_chan <- result
}