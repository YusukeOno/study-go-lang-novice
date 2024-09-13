# Operators and Datatypes

## 演算子

三項演算子（条件演算子）は無い。

## 基本型と複合型

Go言語のデータ型は「基本型」と「複合型」に分かれます。

基本型は単体で意味をなすデータを扱い、複合型は他の方と組み合わせて使用するようなデータを扱う。

「数値型」「文字列型」「真偽値型」が基本形で、それ以外が複合型である。

## 型宣言

型に別名をつけて新しい型を作成することを「型宣言」という。なお、宣言した型と元の型は「別物」として扱われる。

```go
// int型をベースにした型宣言
type MyInt int

// 複数の宣言も可能
type (
    MyInt2 int
    MyInt3 int
)

// 宣言した型の利用
var n1 MyInt = 1
var n2 MyInt = 2

var n int = 10

func main() {
    fmt.Println(n1 + n2) // 3
    // fmt.Println(n1 + n) // 型が異なるため演算できない
}
```

## 数値型

数値型は「整数型」「浮動小数点型」「複素数型」の3つの分類がある。

## 文字列型

「string」。メモリ上への参照をUTF-8のバイト列で扱い、ゼロ値は空文字（""）。

JavaのStringと同じくイミュータブル（状態変更不可）な性質を持つ。

## 真偽値型

「bool」。"true"と"false"のどちらかを扱い、ゼロ値は"false"。

## 構造体型

C言語の構造体と比べると、Go言語特有の機能が備わっている。

構造体は`struct`キーワードを使って宣言する。

```go
// 構造体型の宣言
type MyStruct struct {
    a string
    b, c int
}

func main() {
    var st MyStruct
    st.a = "hoge"
    st.b = 1
    st.c = 2
    fmt.Pringln(st.a, st.b + st.c) // hoge 3
}
```

次に、Go言語の構造体は別の構造体を自身に埋め込むことができる。

```go
type MyStuct struct {
    a string
    b, c int
}

type MyStruct2 struct {
    MyStruct // 構造体の埋め込み
    d int
}

func main() {
    var st2 MyStruct2
    st2.a = "hoge" // 埋め込んだ構造体のメンバアクセス
    st2.d = 10
    fmt.Pringln(st2.a, st2.d)
}
```

構造体に別の構造体を埋め込む際には、例のように「型名」のみを指定する。（これを匿名フィールドと呼ぶ）

埋め込んだ構造体のメンバは、自身のメンバの如く扱うことができるようになる。

最後に、Go言語の構造体ではメンバに対してタグづけをすることができる。また、タグづけした値はreflectパッケージの機能を使って参照することができる。

```go
import (
    "fmt"
    "reflect"
)

type MyStruct struct {
    a string `tag1:"value1" tag2:"value2"`
    b int `tag3:"value3"`
}

func main() {
    var st MyStruct
    field1 := reflect.TypeOf(st).Filed(0)
    field2 := reflect.TypeOf(st).Field(1)
    fmt.Println(field1.Tag.Get("tag1")) // value1
    fmt.Println(field2.Tag.Get("tag3")) // value3
}
```

## ポインタ型

Go言語にもポインタの概念がある。

ポインタを扱うための型が「ポインタ型」であり、変数宣言の際に型名に "*" をつけて定義する。

また、変数に "&" をつけると変数のアドレスを参照することができ、アドレスから変数の値を参照するには "*" を変数につける。

なお、Go言語では「ポインタの値を演算して別アドレスにする」ようなことはできない。

```go
func main() {
    // int型のポインタ変数
    var p *int

    // int型の変数
    n := 10

    // 変数nのアドレスを取得
    p = &n

    fmt.Println(p) // 0x116d006c (環境によって値は異なる)
    fmt.Println(*p) // 10
}
```

Go言語ではnew関数を使って動的にメモリを確保することができる。その際、確保したメモリのアドレスはポインタ型の変数に格納することができる。

なお、new関数で割り当てたメモリは方ごとのゼロ値で初期化されている。

```go
func main() {
    // 型を指定してメモリを割り当てる
    var p *int = new(int)

    fmt.Println(p) // 0x116da0bc (環境によって値は異なる)
    fmt.Println(*p) // 0
}
```

