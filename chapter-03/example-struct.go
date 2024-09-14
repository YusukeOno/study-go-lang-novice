//go:build ignore

package main

import "fmt"

// 構造体型の宣言
type MyStruct struct {
	a    string
	b, c int
}

func main() {
	var st MyStruct
	st.a = "hoge"
	st.b = 1
	st.c = 2
	fmt.Println(st.a, st.b+st.c) // hoge 3
}
