package printer

import (
	"strconv"

	"unicode"

	"strings"

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
			expr := newCallExpr(
				newSelectorExpr("vecty", "Attribute"),
				[]ast.Expr{
					&ast.BasicLit{Kind: token.STRING,
						Value: strconv.Quote(attr.Lhs.Name)},
					attr.Rhs,
				},
			)

			args = append(args, expr)
		}

		// Add the contents
		for _, expr := range gox.X {
			switch expr := expr.(type) {
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
