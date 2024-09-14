# Statement

## 条件分岐

if文とswitch文の2種類がある。

if文のスコープとなる変数が宣言可能。

なお、`else`の前で改行すると文法エラーになる。

- [example-if.go](./example-if.go)

次にswitch文では、「case節にbreakを書く必要がない」のが特徴である。

- [example-switch.go](./example-switch.go)

breakキーワードも存在し、これはswitch文そのものを抜けるような場合に使用する。

## 繰り返し

繰り返しはfor文のみが存在する。他の言語でいうwhile文やforeach文は存在しないが、それらと同様の処理もfor文で実装する。

- [example-for.go](./example-for.go)

ループを抜けるには`break`、スキップするには`continue`を使用する。

## goto


## 例外処理

