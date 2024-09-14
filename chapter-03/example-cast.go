//go:build ignore

package main

import "fmt"

func main() {
	var a int = 1
	var b int16 = 2
	var c int = a + int(b)

	fmt.Println(c) // 3
}
