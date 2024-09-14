//go:build ignore

package main

import "fmt"

func main() {
	var a int = 1
	var b string

	if a > 1 {
		b = "hoge"
	} else if a == 1 {
		b = "fuga"
	} else {
		b = "boo"
	}

	fmt.Println(b) // fuga

	// if文のスコープとなる変数が宣言可能
	if s := "abc"; a > 0 {
		fmt.Println(s) // abc
	}

}
