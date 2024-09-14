//go:build ignore

package main

import "fmt"

type Calc struct {
	a, b int
}

type MyInt int

// メソッド（Calc構造体に紐づく関数）
func (p Calc) Add() int {
	return p.a + p.b
}

// メソッド（MyInt型に紐づく関数）
func (m MyInt) Add(n int) MyInt {
	return m + MyInt(n)
}

func main() {
	p := Calc{3, 2}
	var m MyInt = 1

	// メソッドの呼び出し
	fmt.Println(p.Add())  // 5
	fmt.Println(m.Add(2)) // 3
}
