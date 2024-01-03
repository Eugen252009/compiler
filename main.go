package main

import (
	"compiler/lexer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No File specified")
		return
	}
	var file = os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	lexer.Lexer(content)
}
