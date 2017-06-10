package printer

import (
	"strconv"
	"strings"

	"unicode"

	"github.com/8byt/gox/ast"
	"github.com/8byt/gox/token"
)

func goxToVecty(gox *ast.GoxExpr) ast.Expr {
	isComponent := unicode.IsUpper(rune(gox.TagName.Name[0]))

	if isComponent {
		return ast.NewIdent("COMPONENT")
	} else {
		args := []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(gox.TagName.Name),
			}}

		// Add the attributes
		for _, attr := range gox.Attrs {
			actualRhs := attr.Rhs
			if attr.Rhs == nil { // default to true like JSX
				actualRhs = ast.NewIdent("true")
			}
			expr := newCallExpr(
				newSelectorExpr("vecty", "Attribute"),
				[]ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote(attr.Lhs.Name)},
					actualRhs,
				},
			)

			args = append(args, expr)
		}

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
