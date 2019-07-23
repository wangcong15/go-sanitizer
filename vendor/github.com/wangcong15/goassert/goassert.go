package goassert

import (
	"log"
	"reflect"
	"strings"
)

var (
	debug_mode = false
)

// AssertEq: check equality
func AssertEq(val1, val2 interface{}) {
	// checker
	if !reflect.DeepEqual(val1, val2) {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Eq assertion fails.")
	}
}

// AssertNEq: check non-equality
func AssertNEq(val1, val2 interface{}) {
	// checker
	if reflect.DeepEqual(val1, val2) {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("NEq assertion fails.")
	}
}

func AssertValEq(val1, val2 interface{}) {
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss != 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("NEq assertion fails.")
	}
}

func AssertNValEq(val1, val2 interface{}) {
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss == 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("NEq assertion fails.")
	}
}

// AssertGt: check greater than
func AssertGt(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss <= 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Gt assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertGte: check greater or equal
func AssertGte(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss < 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Gte assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertLt: check less than
func AssertLt(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss >= 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Lt assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertLte: check less or equal
func AssertLte(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss > 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Lte assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertEmpty: check empty
func AssertEmpty(val1 []interface{}) int {
	length := len(val1)
	if length > 0 {
		log.Printf("val1=%v\n", val1)
		panic("Empty assertion fails.")
	}
	return length
}

// AssertNEmpty: check none-empty
func AssertNEmpty(val1 []interface{}) int {
	length := len(val1)
	if length == 0 {
		log.Printf("val1=%v\n", val1)
		panic("NEmpty assertion fails.")
	}
	return length
}

// AssertNNil: check not nil
func AssertNNil(val1 interface{}) {
	// checker
	if val1 == nil || reflect.ValueOf(val1).IsNil() {
		panic("AssertNNil assertion fails.")
	}
}

func AssertIntIn(val1 int, val2 []int) {
	flag := 0
	for _, i := range val2 {
		if val1 == i {
			flag = 1
			break
		}
	}
	if flag == 0 {
		panic("AssertIntIn assertion fails.")
	}
}

func AssertStrIn(val1 string, val2 string) {
	if strings.Contains(val2, val1) == false {
		panic("AssertStrIn assertion fails.")
	}
}

func AssertStrNotIn(val1 string, val2 string) {
	if strings.Contains(val2, val1) == true {
		panic("AssertStrNotIn assertion fails.")
	}
}

func AssertOverflow(val1, val2, val3 interface{}) {
	v1, v2, v3 := allToFloat64(val1), allToFloat64(val2), allToFloat64(val3)
	if v1 >= 0 && v2 >= 0 && v3 < 0 || v1 <= 0 && v2 <= 0 && v3 > 0 {
		panic("AssertOverflow assertion fails")
	}
}

func AssertUnderflow(val1, val2, val3 interface{}) {
	v1, v2, v3 := allToFloat64(val1), allToFloat64(val2), allToFloat64(val3)
	if v1 <= 0 && v2 >= 0 && v3 > 0 || v1 >= 0 && v2 <= 0 && v3 < 0 {
		panic("AssertUnderflow assertion fails")
	}
}

func AssertPresion(val1, val2 interface{}) {
	switch val1.(type) {
	case float32, float64:
		v1, v2 := allToFloat64(val1), allToFloat64(val2)
		if v1-v2 < 0.0001 || v2-v1 < 0.0001 {
			panic("AssertPresion assertion fails")
		}
	}
}
