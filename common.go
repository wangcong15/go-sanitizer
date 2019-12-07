package main

import (
	"fmt"
	"reflect"
	"strings"
)

// stress in terminals
func notice(concatWelcome string) {
	fmt.Println(strings.Repeat("-", len(concatWelcome)))
	fmt.Println(concatWelcome)
	fmt.Println(strings.Repeat("-", len(concatWelcome)))
}

// Call : map string to function
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := m[name]; !ok {
		return
	}
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
