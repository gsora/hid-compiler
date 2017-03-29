package main

import (
	"fmt"
	"github.com/gsora/hid-compiler/compiler"
)

var s1 = "this is a string"
var s2 = "this is a [LCTRL]>string too"
var s3 = "th[LSHIFT]:is is a string [LCTRL]>again"

func main() {

	l := compiler.Compile(s3)
	fmt.Println(l)
}
