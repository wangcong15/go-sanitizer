package main

import (
	"flag"
	"log"
)

var (
	p string
)

func init() {
	flag.StringVar(&p, "p", "", "set a golang package to recommend assertions")
}

func main() {
	flag.Parse()
	if p == "" {
		log.Println("Package path should not be empty. Use -p to set.")
		return
	}
	log.Printf("Package Path Set: %v\n", p)
	// TODO
}
