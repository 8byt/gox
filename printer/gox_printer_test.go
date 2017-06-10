package printer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"path/filepath"

	"github.com/8byt/gox/parser"
	"github.com/8byt/gox/token"
)

const goxTestsDir = "../goxtests"

// TestGoxParse verifies that gox can parse
func TestGoxParse(t *testing.T) {
	list, err := ioutil.ReadDir(goxTestsDir)
	if err != nil {
		t.Fatalf("Unable to read directory: %v", err)
	}

	cfg := &Config{Mode: GoxToGo}

	for _, fi := range list {
		name := fi.Name()

		if !fi.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".gox") {
			t.Run(name, func(t *testing.T) {
				fset := token.NewFileSet()

				file, err := parser.ParseFile(fset, filepath.Join(goxTestsDir, name), nil, parser.ParseComments)
				if err != nil {
					fmt.Println("Can't parse file", err)
				}

				cfg.Fprint(os.Stdout, fset, file)

				if err != nil {
					fmt.Printf("Failed with error: %v", err)
					t.Fatalf("ParseFile(%s): %v", name, err)
				}
			})
		}
	}

}
