//go:build ignore

package main

import "fmt"

func main() {
	// チャネルの生成
	ch := make(chan int)

	// ゴルーチンの生成
	go send(ch)

	// チャネルによる値の受信
	c := <-ch
	fmt.Printf("[値を受信:%d]\n", c)
}

func send(ch chan int) {
	fmt.Println("[値を送信:1]")

	// チャネルによる値の送信
	ch <- 1
}
