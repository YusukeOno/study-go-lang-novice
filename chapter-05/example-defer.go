//go:build ignore

package main

import "fmt"

func main() {
	f() // processing...の後にfinish!
}

func f() {
	defer fmt.Println("finish!")
	fmt.Println("processing...")
}
