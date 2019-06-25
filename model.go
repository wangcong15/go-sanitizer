package main

type assertion struct {
	file_path     string
	line_no       int
	expression    string
	weakness_type int
}

type assertionSlice []assertion

func (a assertionSlice) Len() int {
	return len(a)
}

func (a assertionSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a assertionSlice) Less(i, j int) bool {
	if a[i].file_path == a[j].file_path {
		return a[i].line_no > a[j].line_no
	} else {
		return a[i].file_path > a[j].file_path
	}
}
