//go:build ignore

package main

import "fmt"

// インタフェースの宣言
type Human interface {
	hello()
	walk()
}

// 構造体とメソッドの宣言
type Taro struct {
	name string
	age  int
}

func (t Taro) hello() {
	fmt.Printf("%s%d歳です\n", t.name, t.age)
}

func (t Taro) walk() {
	fmt.Println("Waking...")
}

func (t Taro) shout() {
	fmt.Println("Hoooo!!!")
}

func main() {
	// 構造体の値をインタフェースに設定する
	t := Taro{"太郎", 20}
	var h Human = t

	// インタフェースから構造体のメソッドを呼び出す
	h.hello() // 太郎20歳です
	h.walk()  // Waking...
	// h.shout() // アクセスできない
}
