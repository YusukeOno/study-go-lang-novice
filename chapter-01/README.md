# Introduction

## Environment variable

```zsh
> go env
GO111MODULE=''
GOARCH='arm64'
GOBIN=''
GOCACHE='/Users/yusuke.ono/Library/Caches/go-build'
GOENV='/Users/yusuke.ono/Library/Application Support/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='arm64'
GOHOSTOS='darwin'
GOINSECURE=''
GOMODCACHE='/Users/yusuke.ono/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='darwin'
GOPATH='/Users/yusuke.ono/go'
GOPRIVATE=''
GOPROXY='https://proxy.golang.org,direct'
GOROOT='/opt/homebrew/Cellar/go/1.23.0/libexec'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='local'
GOTOOLDIR='/opt/homebrew/Cellar/go/1.23.0/libexec/pkg/tool/darwin_arm64'
GOVCS=''
GOVERSION='go1.23.0'
GODEBUG=''
GOTELEMETRY='local'
GOTELEMETRYDIR='/Users/yusuke.ono/Library/Application Support/go/telemetry'
GCCGO='gccgo'
GOARM64='v8.0'
AR='ar'
CC='cc'
CXX='c++'
CGO_ENABLED='1'
GOMOD='/dev/null'
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -ffile-prefix-map=/var/folders/c0/j8z9nmwj3r93swy5clqm0jb00000gn/T/go-build4186934639=/tmp/go-build -gno-record-gcc-switches -fno-common'
```

## Compile/Execute

```go
package main

import "fmt"

func main() {
        fmt.Println("Hello World")
}
```

```zsh
> go build ./hello.go
```

```zsh
> ./hello
Hello World
```

buildコマンドではなくrunコマンドを使用すると、ビルドと実行を同時に実施する。

```zsh
> go run ./hello.go
Hello World
```

## Official Services

### Go Playgroung

ブラウザ上でGo言語を実行する環境を提供する。

![https://go.dev/play/](https://go.dev/play/)

![https://goplay.tools/](https://goplay.tools/)

### A Tour of Go

このサービスは、チュートリアル形式でGo言語の一通りの文法を学ぶことができる。

![https://go-tour-jp.appspot.com/](https://go-tour-jp.appspot.com/)

学習の途中で気になったことを動かしながら進めることもできる。

以上
