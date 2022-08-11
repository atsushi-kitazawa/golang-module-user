package foo

import "github.com/atsushi-kitazawa/golang-module-example1/module1"

func FooFactory() Fooer {
	return module1.New()
}
