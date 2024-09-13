//go:build ignore

package main

import "fmt"

func main() {
	// 配列の定義
	var a [3]int
	b := [2]string{"hoge", "fuga"}

	// 初期値を指定すると「...」表記で長さを省略できる
	c := [...]int{1, 2}

	// インデックスを指定して値を割り当てることも可能
	d := [3]int{1: 2, 2: 10}

	fmt.Println(a[0])             // 0
	fmt.Println(b[1])             // fuga
	fmt.Println(d[0], d[1], d[2]) // 0 2 10

	// データ長はlen関数で取得可能
	fmt.Println(len(a)) // 3
	fmt.Println(len(c)) // 2
}
