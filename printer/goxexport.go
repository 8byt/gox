package printer

import (
	"strconv"

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
			expr := ast.CallExpr{
				Fun: newSelectorExpr("vecty", "Attribute"),
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING,
						Value: strconv.Quote(attr.Lhs.Name)},
					attr.Rhs,
				},
				Ellipsis: token.NoPos, Lparen: token.NoPos, Rparen: token.NoPos,
			}

			args = append(args, &expr)
		}

		// Add the contents
		for _, expr := range gox.X {
			switch expr.(type) {

			default:
				args = append(args, expr)
			}
		}

		return &ast.CallExpr{
			Fun:      newSelectorExpr("vecty", "Tag"),
			Args:     args,
			Ellipsis: token.NoPos, Lparen: token.NoPos, Rparen: token.NoPos}
	}
}

func newSelectorExpr(x, sel string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(x),
		Sel: ast.NewIdent(sel)}
}
