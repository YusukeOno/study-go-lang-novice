//go:build ignore

package main

import "fmt"

func main() {
	// int型のポインタ変数
	var p *int

	// int型の変数
	n := 10

	// 変数nのアドレスを取得
	p = &n

	fmt.Println(p)  // 0x14000104020 (環境によって値は異なる)
	fmt.Println(*p) // 10
}
