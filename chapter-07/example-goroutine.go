//go:build ignore

package main

import (
	"fmt"
	"time"
)

// グローバル変数
var n = 0

func main() {
	fmt.Println("main実行")
	fmt.Println(&n)

	// ゴルーチンの生成
	go sub()

	// 3ミリ秒スリープする
	time.Sleep(time.Millisecond * 3)
}

func sub() {
	fmt.Println("sub実行")
	fmt.Println(&n)

	// 10ミリ秒スリープする
	time.Sleep(time.Millisecond * 10)

	fmt.Println("sub実行2") // 先にmainゴルーチンが終了するためこの行は実行されない
}
