//go:build ignore

package main

import "fmt"

// int型をベースにした型宣言
type MyInt int

// 複数の宣言も可能
type (
	MyInt2 int
	MyInt3 int
)

// 宣言した型の利用
var n1 MyInt = 1
var n2 MyInt = 2

var n int = 10

func main() {
	fmt.Println(n1 + n2) // 3
	// fmt.Println(n1 + n)  // 型が異なるため演算できない
}
