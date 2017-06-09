package parser

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

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

		if !fi.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".src") {
			t.Run(name, func(t *testing.T) {
				_, err := ParseFile(token.NewFileSet(), filepath.Join(goxtestsfolder, name), nil, DeclarationErrors)
				if err != nil {
					t.Fatalf("ParseFile(%s): %v", name, err)
				}
			})
		}
	}

}
