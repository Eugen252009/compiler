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
	tokens := lexer.Lexer(content)
	parsed := Parse(&tokens)
	fmt.Println(parsed)
}
func Parse(token *[]lexer.Token) bool {
	for _, val := range *token {
		fmt.Printf("Type: %s, Value: %s\n", val.Type, val.Value)
	}
	return true
}
