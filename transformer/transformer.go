package transformer

import (
	"fmt"
	"os"
	"strings"

	"github.com/8byt/gox/ast"
	"github.com/8byt/gox/parser"
	"github.com/8byt/gox/printer"
	"github.com/8byt/gox/token"
)

func rename() {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "../parser/goxtests/args_and_more.gox", nil, 0)
	if err != nil {
		fmt.Println("Can't parse file", err)
	}

	file.Name.Name = "what" // change package name

	r := &Renamer{"Foo", "Bar"}
	ast.Walk(r, file)

	printer.Fprint(os.Stdout, fs, file)
}

type Renamer struct {
	find    string
	replace string
}

func (r *Renamer) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		switch n := node.(type) {
		case *ast.FuncDecl:
			if n.Recv != nil && n.Recv.List != nil && len(n.Recv.List) > 0 {
				field := n.Recv.List[0]
				typ := field.Type.(*ast.StarExpr).X.(*ast.Ident).Name
				if typ == r.find {
					field.Names[0].Name = strings.ToLower(r.replace[0:1])
				}
			}
		case *ast.Ident:
			if n.Name == r.find {
				n.Name = r.replace
			}
		}
	}
	return r
}
