package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/golang-module-user/foo"
)

func main() {
	f := foo.FooFactory()
	s := f.Foo()
	fmt.Printf("f type is %T\n", f)
	fmt.Printf("function return is %v\n", s)
}
