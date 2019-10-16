package main

import (
	"fmt"

	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/8byt/gox/parser"
	"github.com/8byt/gox/printer"
	"github.com/8byt/gox/token"
)

func main() {
	if len(os.Args) > 1 {
		for _, dir := range os.Args[1:] {
			Transpile(dir)
		}
	} else {
		Transpile("goxtests")
	}
}

type GoxTranspiler struct {
	cfg  *printer.Config
	fset *token.FileSet
}

func (g *GoxTranspiler) TranspileFile(path string, f os.FileInfo, err error) error {
	name := f.Name()

	if !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".gox") {
		fmt.Printf("Transpiling %s\n", path)
		g.fset.AddFile(filepath.Join(path, name), -1, int(f.Size()))

		file, err := parser.ParseFile(g.fset, path, nil, parser.ParseComments)
		if err != nil {
			fmt.Println("Can't parse file", err)
			return err
		}

		//cfg.Fprint(os.Stdout, fset, file)
		of, err := os.Create(path[:len(path)-1]) // lol
		g.cfg.Fprint(of, g.fset, file)

		if err != nil {
			fmt.Printf("Failed with error: %v", err)
			log.Fatalf("ParseFile(%s): %v", name, err)
			return err
		}
	}
	return nil
}

func Transpile(directory string) {
	goxT := &GoxTranspiler{
		cfg:  &printer.Config{Mode: printer.GoxToGo | printer.RawFormat},
		fset: token.NewFileSet(),
	}

	filepath.Walk(directory, goxT.TranspileFile)

}
