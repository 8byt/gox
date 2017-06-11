package printer

import (
	"strconv"
	"strings"

	"unicode"

	"github.com/8byt/gox/ast"
	"github.com/8byt/gox/token"
)

// Map html-style to actual js event names
var eventMap = map[string]string{
	"onAbort":          "abort",
	"onCancel":         "cancel",
	"onCanPlay":        "canplay",
	"onCanPlaythrough": "canplaythrough",
	"onChange":         "change",
	"onClick":          "click",
	"onCueChange":      "cuechange",
	"onDblClick":       "dblclick",
	"onDurationChange": "durationchange",
	"onEmptied":        "emptied",
	"onEnded":          "ended",
	"onInput":          "input",
	"onInvalid":        "invalid",
	"onKeyDown":        "keydown",
	"onKeyPress":       "keypress",
	"onKeyUp":          "keyup",
	"onLoadedData":     "loadeddata",
	"onLoadedMetadata": "loadedmetadata",
	"onLoadStart":      "loadstart",
	"onMouseDown":      "mousedown",
	"onMouseEnter":     "mouseenter",
	"onMouseleave":     "mouseleave",
	"onMouseMove":      "mousemove",
	"onMouseOut":       "mouseout",
	"onMouseOver":      "mouseover",
	"onMouseUp":        "mouseup",
	"onMouseWheel":     "mousewheel",
	"onPause":          "pause",
	"onPlay":           "play",
	"onPlaying":        "playing",
	"onProgress":       "progress",
	"onRateChange":     "ratechange",
	"onReset":          "reset",
	"onSeeked":         "seeked",
	"onSeeking":        "seeking",
	"onSelect":         "select",
	"onShow":           "show",
	"onStalled":        "stalled",
	"onSubmit":         "submit",
	"onSuspend":        "suspend",
	"onTimeUpdate":     "timeupdate",
	"onToggle":         "toggle",
	"onVolumeChange":   "volumechange",
	"onWaiting":        "waiting",
}

func goxToVecty(gox *ast.GoxExpr) ast.Expr {
	isComponent := unicode.IsUpper(rune(gox.TagName.Name[0]))

	if isComponent {
		return newComponent(gox)
	} else {
		args := []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(gox.TagName.Name),
			}}

		// Add the attributes
		args = append(args, mapProps(gox.Attrs)...)

		// Add the contents
		for _, expr := range gox.X {
			switch expr := expr.(type) {
			// TODO figure out what's a better thing to do here
			// do we want to error on compile or figure out what to do based on context?
			// (I think the latter)
			// Fallback to regular behavior, don't wrap this yet
			//case *ast.GoExpr:
			//	e := newCallExpr(
			//		newSelectorExpr("vecty", "Text"),
			//		[]ast.Expr{expr},
			//	)
			//	args = append(args, e)

			case *ast.BareWordsExpr:
				if len(strings.TrimSpace(expr.Value)) == 0 {
					continue
				}
				e := newCallExpr(
					newSelectorExpr("vecty", "Text"),
					[]ast.Expr{expr},
				)
				args = append(args, e)
			default:
				args = append(args, expr)
			}
		}

		return newCallExpr(
			newSelectorExpr("vecty", "Tag"),
			args,
		)
	}
}

func newSelectorExpr(x, sel string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(x),
		Sel: ast.NewIdent(sel)}
}

func newCallExpr(fun ast.Expr, args []ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:      fun,
		Args:     args,
		Ellipsis: token.NoPos, Lparen: token.NoPos, Rparen: token.NoPos}
}

func newComponent(gox *ast.GoxExpr) *ast.UnaryExpr {
	var args []ast.Expr
	for _, attr := range gox.Attrs {
		if attr.Rhs == nil { // default to true like JSX
			attr.Rhs = ast.NewIdent("true")
		}
		expr := &ast.KeyValueExpr{
			Key:   ast.NewIdent(attr.Lhs.Name),
			Colon: token.NoPos,
			Value: attr.Rhs,
		}

		args = append(args, expr)
	}

	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    token.AND,
		X: &ast.CompositeLit{
			Type:   ast.NewIdent(gox.TagName.Name),
			Lbrace: token.NoPos,
			Elts:   args,
			Rbrace: token.NoPos,
		},
	}
}

func mapProps(goxAttrs []*ast.GoxAttrStmt) []ast.Expr {
	var mapped = []ast.Expr{}
	for _, attr := range goxAttrs {
		// set default of Rhs to true if none provided
		if attr.Rhs == nil { // default to true like JSX
			attr.Rhs = ast.NewIdent("true")
		}

		var expr ast.Expr

		// if prop is an event listener (e.g. "onClick")
		if _, ok := eventMap[attr.Lhs.Name]; ok {
			expr = newEventListener(attr)
		} else {
			// if prop is a normal attribute
			expr = newCallExpr(
				newSelectorExpr("vecty", "Attribute"),
				[]ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote(attr.Lhs.Name)},
					attr.Rhs,
				},
			)
		}

		mapped = append(mapped, expr)
	}

	return mapped
}

func newEventListener(goxAttr *ast.GoxAttrStmt) ast.Expr {
	return &ast.UnaryExpr{
		OpPos: token.NoPos,
		Op:    token.AND,
		X: &ast.CompositeLit{
			Type:   newSelectorExpr("vecty", "EventListener"),
			Lbrace: token.NoPos,
			Elts: []ast.Expr{
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Name"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote(eventMap[goxAttr.Lhs.Name]),
					},
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("Listener"),
					Value: goxAttr.Rhs,
				},
			},
			Rbrace: token.NoPos,
		},
	}
}
