//go:build ignore

package main

import "fmt"

func main() {
	// 関数型の変数を定義
	var fnc func(int, int) int

	// 関数の代入
	fnc = add
	fmt.Println(fnc(2, 1)) // 3

	// 関数の代入
	fnc = mul
	fmt.Println(fnc(2, 1)) // 2
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
