//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println(f1(1, 2))
	f2("国", "日本", "アメリカ")
	var a, b = f3()
	fmt.Println(a, b)
	fmt.Println(f4())
}

// ノーマルな関数
func f1(i1 int, i2 int) int {
	return i1 + i2
}

// 可変長引数
func f2(s string, params ...string) {
	fmt.Printf("%s:", s)
	for _, p := range params {
		fmt.Printf("%s ", p)
	}
	fmt.Printf("\n")
}

// 複数の戻り値
func f3() (int, string) {
	return 100, "hoge"
}

// 戻り値に名前をつけて扱う
func f4() (i int, t int) {
	i = 5
	t = 2
	return
}
