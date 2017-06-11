package main

import (
	"fmt"

	"io/ioutil"
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
		Transpile(os.Args[1])
	} else {
		Transpile("goxtests")
	}
}

func Transpile(directory string) {
	list, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Unable to read directory: %v", err)
	}

	cfg := &printer.Config{Mode: printer.GoxToGo | printer.RawFormat}

	fset := token.NewFileSet()

	for _, fi := range list {
		name := fi.Name()

		if !fi.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".gox") {
			fset.AddFile(filepath.Join(directory, name), -1, int(fi.Size()))

			file, err := parser.ParseFile(fset, filepath.Join(directory, name), nil, parser.ParseComments)
			if err != nil {
				fmt.Println("Can't parse file", err)
			}

			//cfg.Fprint(os.Stdout, fset, file)
			of, err := os.Create(filepath.Join(directory, name[:len(name)-1])) // lol
			cfg.Fprint(of, fset, file)

			if err != nil {
				fmt.Printf("Failed with error: %v", err)
				log.Fatalf("ParseFile(%s): %v", name, err)
			}
		}
	}

}
