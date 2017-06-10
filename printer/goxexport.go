package printer

import (
	"fmt"
	"strings"

	"github.com/8byt/gox/ast"
)

func goxToVecty(gox *ast.GoxExpr) string {
	args := []string{gox.TagName.Name}

	for _, attr := range gox.Attrs{
		var str string
		switch attr.Rhs.(type){
		case ast.GoExpr:
			str = 


		}
		str := fmt.Sprintf()
		args = append(args, )
	}

	return fmt.Sprintf(`vecty.Tag("%s")`, strings.Join(args, ", "))
}
