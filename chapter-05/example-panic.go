//go:build ignore

package main

import "fmt"

func main() {
	fmt.Println("start")
	sub()
	fmt.Println("end") // 実行されない
}

func sub() {
	fmt.Println("sub start")
	panic("パニック！")
	fmt.Println("sub end") // 実行されない
}
