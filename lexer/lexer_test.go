package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `int main(){printf("Hallo Welt")}`
	tokens := []Token{
		{Type: TOKEN_INT, Value: "int"},
		{Type: TOKEN_STRING_LITERAL, Value: "main"},
		{Type: TOKEN_OPAREN, Value: "("},
		{Type: TOKEN_CPAREN, Value: ")"},
		{Type: TOKEN_OCURLY, Value: "{"},
		{Type: TOKEN_FUNCTIONCALL, Value: "printf"},
		{Type: TOKEN_OPAREN, Value: "("},
		{Type: TOKEN_STRING_LITERAL, Value: "Hallo Welt"},
		{Type: TOKEN_CPAREN, Value: ")"},
		{Type: TOKEN_CCURLY, Value: "}"},
	}
	lex := New(input)
	for _, v := range lex.Token {
		fmt.Println(v.Type)
	}
	for id, val := range lex.Token {
		if val.Type != tokens[id].Type {
			fmt.Printf("%d\tExpected Type:%q, got=%q\n", id, val, tokens[id])
			t.Fatalf("%v", val)
		}
	}
}
