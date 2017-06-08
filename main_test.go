package main

import (
	"fmt"
	"testing"

	"github.com/8byt/gox/parser"
	"github.com/8byt/gox/scanner"
	"github.com/8byt/gox/token"
)

//	case token.ASSIGN, token.EQL, token.NEQ, token.DEFINE,
//token.LPAREN, token.LBRACE, token.COMMA, token.COLON,
//token.RETURN, token.IF, token.SWITCH, token.CASE:

func TestOtagParse(t *testing.T) {
	t.Log(doParseAst("return <a>hello world</a>"))
	t.Log(doParseAst("lol := <a>hahaha</a>"))
	t.Log(doParseAst("if <abc></abc> == <abc></abc> {}"))
	t.Log(doParseAst(`return <a attr="value">hello world</a>`))

	t.Skip()
}

func doTestExpr(strExpr string) {
	expr, err := parser.ParseExpr(strExpr)
	fmt.Printf("err: %v value: %v\n", err, expr)
}

func doParseAst(strExpr string) (result string) {
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
		result += fmt.Sprintf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	return result
}
