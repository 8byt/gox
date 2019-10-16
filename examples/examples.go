package main

import (
	components "github.com/8byt/gox/goxtests"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

func main() {
	vecty.SetTitle("gox lang")
	p := &components.BodyComponent{}
	vecty.RenderBody(p)
	js.Global.Get("console").Call("log", "dang")
}
