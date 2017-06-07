package main 

import (
	"fmt"

	"github.com/8byt/gox/parser"
	//"github.com/8byt/gox/scanner"
	"github.com/8byt/gox/scanner"
	"github.com/8byt/gox/token"
)

func main() {
	fmt.Println("Hello World")
	testExpr(`1 + 2`)
	testExpr(`int[...] {1, 2, 3}`)
	testExpr(`func() { return <|div>12</div>; }`)

	parseAst("var tokens = [...]string {\n 2: \"sting\"}")

	parseAst("var tagYo = <|div>{help}</div>")
}

func testExpr(strExpr string) {
	expr, err := parser.ParseExpr(strExpr)
	fmt.Printf("err: %v value: %v\n", err, expr)
}

func parseAst(strExpr string)  {
	// src is the input that we want to tokenize.
	src := []byte(strExpr)

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}