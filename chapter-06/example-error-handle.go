//go:build ignore

package main

import (
	"errors"
	"fmt"
)

func main() {
	e := test()
	fmt.Println(e.Error()) // エラーを生成
}

func test() error {
	return errors.New("エラーを生成")
}
