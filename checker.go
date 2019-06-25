package main

func C777(raw_code string) {
	var result []assertion
	result = append(result, assertion{"a",1,"b",2})
	checker_chan <- result
}