//go:build ignore

package main

import "fmt"

type MyStruct struct {
	Calc
}

type Calc struct {
	a, b int
}

func (p Calc) Add() int {
	return p.a + p.b
}

func main() {
	var s MyStruct
	s.a = 5
	s.b = 4

	// 別の型のメソッド呼び出し
	fmt.Println(s.Add()) // 9
}
