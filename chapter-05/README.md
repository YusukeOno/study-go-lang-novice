# Function

## 関数定義

関数の定義は`func`キーワードで行う。

Go言語の関数には次の特徴がある。

- 複数の戻り値を返却できる。
- 戻り値に名前をつけて扱える

- [example-function.go](./example-function.go)

f2関数の中ではfor文の中で「_」を使用している箇所があるが、これは「ブランク識別子」と呼ばれ、「文法上は値を受け取るけど保存する必要がない」ような場合に使用する。

f3関数のように戻り値の定義をカッコで括って型を指定する。また、main関数内でf3関数を呼び出している箇所のような記述で戻り値を取得することができる。

最後に、戻り値に名前をつけて扱うにはf4関数のような記述になる。戻り値に名前をつけた場合、関数内で変数のように扱うことができ、return式に戻り値を指定する必要がなくなる。

なお、return式に値を明記した場合はそれが優先される。

## 関数型

Go言語では「関数を変数に格納する」ことができる。そのためのデータ型が「関数型」である。

- [example-function-type.go](./example-function-type.go)

Go言語の関数は第一級オブジェクトなので、引数やお戻り値にすることも可能。

## クロージャ

Go言語でも「ローカル変数を参照している関数内関数」を実現できる。

- [example-closure.go](./example-closure.go)

## defer文

defer文を使うと、それを呼び出した関数が終了する際に実行すべき処理を機打つすることができる。

リソースの開封など、確実に行いたい処理を記述するのに適している。

- [example-defer.go](./example-defer.go)

## パニック

Go言語には「パニック」というエラー処理の仕組みがある。

パニックは処理を停止して関数の呼び出し元に戻る仕組みで、発生すると順番に関数の呼び出し元に戻り、main関数まで戻るとスタックトレースを出力してプログラムを終了させる。

- [example-panic.go](./example-panic.go)

```zsh
> go run ./example-panic.go
start
sub start
panic: パニック！

goroutine 1 [running]:
main.sub()
        /Users/yusuke.ono/Develop/GitHub/study-go-lang-novice/chapter-05/example-panic.go:15 +0x64
main.main()
        /Users/yusuke.ono/Develop/GitHub/study-go-lang-novice/chapter-05/example-panic.go:9 +0x54
exit status 2
```
