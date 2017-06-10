package examples

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/examples/todomvc/store"
	"github.com/8byt/gox/goxtests_gen"
)

func main() {
	vecty.SetTitle("gox lang")
	p := &components.FilterButton{}
	store.Listeners.Add(p, func() {
		p.Items = store.Items
		vecty.Rerender(p)
	})
	vecty.RenderBody(p)
}
