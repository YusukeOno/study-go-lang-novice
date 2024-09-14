//go:build ignore

package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	a string `tag1:"value1" tag2:"value2"`
	b int    `tag3:"value3"`
}

func main() {
	var st MyStruct
	field1 := reflect.TypeOf(st).Field(0)
	field2 := reflect.TypeOf(st).Field(1)
	fmt.Println(field1.Tag.Get("tag1")) // value1
	fmt.Println(field2.Tag.Get("tag3")) // value3
}
