package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("simplx", map[string]interface{}{
		"NewSimplx": NewSimplx,
	})
}

func NewSimplx() *js.Object {
	return js.MakeWrapper(&SplxObj{})
}
