# gox
### gox is an extension of Go's syntax that let's you write HTML-style tags directly in your source code.
#### In other words, it's JSX for Go.

Write HTML-style tags directly in your GopherJS source, and have them get transpiled into [Vecty](https://github.com/gopherjs/vecty) components.

Okay take a look:
```
package main

import "github.com/gopherjs/vecty"

func main() {
	woah := <body>
		<div class="amazing">
			<h1>gox</h1>
			<span class={"you could put dynamic content here"}/>
			yeah you can do bare words too
		</div>
	</body>
	
	vecty.RenderBody(woah)
}
```
## Why?
Two big reasons:
 - It would be nice to have type safety, but I'm unwilling to write Vecty components
 - It would be nice to know how Go parsing works
 - I would like to learn Go by modifying its AST (Danny's reason)
 - I want to write frontend code, but I don't want JS (Eric's reason)

## How?
We basically vendored the Go parser/scanner/AST/etc. and just modified it until it fit our needs.

## Wot?
Here's a more complicated example portion of a `.gox` file.
```
func (p *PageView) renderItemList() *vecty.HTML {
	var items vecty.List
	for i, item := range store.Items {
		if (store.Filter == model.Active && item.Completed) || (store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, <ItemView Index={i} Item={item} />)
	}

	return <section class="main">
		<input
			id="toggle-all"
			class="toggle-all"
			type="checkbox"
			checked={store.CompletedItemCount() == len(store.Items)}
			onChange={p.onToggleAllCompleted}/>
		<label for="toggle-all">Mark all as complete</label>
		<ul class="todo-list">
			{items}
		</ul>
	</section>
}
```
from [our TodoMVC implementation](https://github.com/8byt/gox/blob/master/examples/todomvc/components/pageview.gox)

## alright, I'm convinced, get me started
Wow! Okay I don't think we thought that would happen.

For now, clone this repo, and build it.

Use `gox <directory>` to convert `.gox` files into `.go` files (they stay in the same directory)

GopherJS should take care of the rest, use [Vecty's docs](https://github.com/gopherjs/vecty) and [GopherJS's docs](https://github.com/gopherjs/gopherjs) to learn more. We use `gopherjs serve` and things magically get transpiled again.

If you want to make this process better, we'd be happy to consider your ideas/PRs.

Thanks,
[Eric](https://github.com/HALtheWise) and [Danny](https://github.com/wolfd)

## License
All modifications are MIT

Original Go code is all BSD
