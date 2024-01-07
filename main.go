package main

import (
	"compiler/lexer"
	"compiler/parser"
	"fmt"
	"os"
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
	fmt.Println("Parsed ", parse)
}
