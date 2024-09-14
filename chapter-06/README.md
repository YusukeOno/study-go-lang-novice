# Method

## メソッド

Go言語のメソッドは「特定の型に関連づけられた関数」である。

一般的なオブジェクト指向言語での意味合いとは異なる。

- [example-method.go](./example-method.go)

メソッドと関連づいている型のことを「レシーバ」と呼び、レシーバの型が対象となる。また、レシーバにはポインタ型を指定することもできる。

次に、メソッドは構造体の匿名フィールドを使って、違う型のメソッドを使用することもできる。これによりオブジェクト指向の「継承」に似たような実装が可能になる。

- [example-method2.go](./example-method2.go)

## インタフェース

インタフェースはメソッドの集まりを定義するためのもの。インタフェースを用いることでポリモーフィズムを実現することができる。

- [example-interface.go](./example-interface.go)

## エラーインタフェース

Go言語にはエラー処理のための仕組みとしてエラーインタフェースが用意されている。

```go
type error interface {
    Error() string
}
```

エラーインタフェースは関数の戻り値として使用し、Errorメソッドでエラー内容を確認することができる。

多くの組み込み関数はエラー時にエラーインタフェースを返却するように実装されており、もちろん自作の関数で利用することもできる。

- [example-error-interface.go](./example-error-interface.go)
- [example-error-handle.go](./example-error-handle.go)
