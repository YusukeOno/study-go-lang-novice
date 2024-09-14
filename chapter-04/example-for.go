//go:build ignore

package main

import "fmt"

func main() {
	// 基本的なfor文
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// foreach文
	for i, t := range []int{5, 6, 7} {
		fmt.Printf("i=%d,t=%d\n", i, t)
	}

	// while文
	w := 0
	for w < 2 {
		fmt.Println(w)
		w++
	}

	// 無限ループ
	for {

	}
}
