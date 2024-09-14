//go:build ignore

package main

import "fmt"

func main() {
	// スライスの定義
	var a []int // nilが設定される
	b := []string{"hoge", "fuga"}
	c := []int{1, 2}

	fmt.Println(a)    // []
	fmt.Println(b[0]) // hoge
	fmt.Println(c[1]) // 2
}
