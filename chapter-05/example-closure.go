//go:build ignore

package main

import "fmt"

func main() {
	var fnc = f()
	fmt.Println(fnc()) // 1
	fmt.Println(fnc()) // 2
	fmt.Println(fnc()) // 3
}

// ローカル関数を返却する関数
func f() func() int {
	l := 0

	// ローカル変数を参照する関数を返却する
	return func() int {
		l++
		return l
	}
}
