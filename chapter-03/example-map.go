//go:build ignore

package main

import "fmt"

func main() {
	// マップの定義
	m := map[string]int{"a": 1, "b": 2}

	// 要素への代入
	m["b"] = 10

	// 要素の参照
	fmt.Println(m["a"], m["b"]) // 1 10
}
