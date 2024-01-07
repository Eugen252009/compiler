package main

import (
	"fmt"
	"os"

	"github.com/eugen252009/compiler/ast"
	"github.com/eugen252009/compiler/lexer"
	"github.com/eugen252009/compiler/parser"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("No File specified or too Many files specified")
		return
	}
	var file = os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	lex := lexer.New(string(content))
	parse := parser.Parse(lex)
	ast.ToJavaScript(parse)

	// fmt.Println("Parsed ", parse)
}
