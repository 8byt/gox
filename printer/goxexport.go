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
				Fun: &ast.SelectorExpr{X: ast.NewIdent("vecty"), Sel: ast.NewIdent("Attribute")},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING,
						Value: strconv.Quote(attr.Lhs.Name)},
					attr.Rhs,
				},
			}

			args = append(args, &expr)
		}

		// Add the contents
		for _, expr := range gox.X {
			args = append(args, expr)
		}

		selector := &ast.SelectorExpr{
			X:   ast.NewIdent("vecty"),
			Sel: ast.NewIdent("Tag")}

		return &ast.CallExpr{Fun: selector, Args: args}
	}
}

func newSelectorExpr() {}
