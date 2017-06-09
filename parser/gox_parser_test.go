package parser

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/8byt/gox/ast"
	"github.com/8byt/gox/token"
)

const goxtestsfolder = "goxtests"

// TestGoxParse verifies that Gox can parse
func TestGoxParse(t *testing.T) {
	list, err := ioutil.ReadDir(goxtestsfolder)
	if err != nil {
		t.Fatalf("Unable to read directory: %v", err)
	}

	for _, fi := range list {
		name := fi.Name()

		if !fi.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".gox") {
			t.Run(name, func(t *testing.T) {
				fset := token.NewFileSet()
				f, err := ParseFile(fset, filepath.Join(goxtestsfolder, name), nil, DeclarationErrors|Trace)

				if err != nil {
					t.Fatalf("ParseFile(%s): %v", name, err)
				}

				buf := bytes.NewBufferString("\n")
				ast.Fprint(buf, fset, f, ast.NotNilFilter)
				t.Log(buf.String())

			})
		}
	}

}
