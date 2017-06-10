package main

import (
	"github.com/gopherjs/vecty"
)

func getHTML2() *vecty.HTML {
	return vecty.Tag(
		"span",
		&MyComponent{Parameter1: "Hello World"},
	)
}

type MyComponent struct {
	vecty.Core
	Parameter1 string
}

func (c *MyComponent) Render() *vecty.HTML {
	return vecty.Tag(
		"div",
		vecty.Text(c.Parameter1),
	)
}

func main() {}
