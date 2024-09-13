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
