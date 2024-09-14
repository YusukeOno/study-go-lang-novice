//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Open関数はファイルを開くための関数
	// 存在しないファイルを開こうとしてエラーを発生させる
	_, e := os.Open("./hoge.txt")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
}
