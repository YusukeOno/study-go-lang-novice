//go:build ignore

package main

import "fmt"

func main() {
	// 型を指定してメモリを割り当てる
	var p *int = new(int)

	fmt.Println(p)  // 0x14000102020 (環境によって値は異なる)
	fmt.Println(*p) // 0
}
