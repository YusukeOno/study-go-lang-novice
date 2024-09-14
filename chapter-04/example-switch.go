//go:build ignore

package main

import "fmt"

func main() {
	var a int = 1

	// if文と同じくローカル変数が定義できる
	switch i := 0; a {
	case 1:
		fmt.Println(i + 1) // 1
		fallthrough
	case 2:
		fmt.Println(i + 2) // 2
	case 3, 4:
		fmt.Println(i + 3) // 実行されない
	default:
		fmt.Println(i) // 実行されない
	}

	// このような書き方もできる
	switch {
	case a == 0:
		fmt.Println(0) // 実行されない
	case a == 1:
		fmt.Println(1) // 1
	case a == 2:
		fmt.Println(2) // 実行されない
	default:
		fmt.Println(999) // 実行されない
	}
}
