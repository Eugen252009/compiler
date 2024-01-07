package parser

import (
	"testing"

	"github.com/eugen252009/compiler/lexer"
)

func TestParsing(t *testing.T) {
	input := `int main(){return 0;}`
	tokens := lexer.New(input)
	parsedTokens := Parse(tokens)
	for _, val := range parsedTokens.Function {
		if val.ReturnType != "int" {
			t.Fatalf("ReturnType wrong: wanted %s, got %s", lexer.TOKEN_INT, val.ReturnType)
		}
		if val.FunctionName != "main" {
			t.Fatalf("FunctionName wrong: wanted %s, got %s", lexer.TOKEN_INT, val.FunctionName)
		}
		for _, val2 := range val.Functions {
			if val2.Name != "return" {
				t.Fatalf("FunctionName wrong: wanted %s, got %s", lexer.TOKEN_INT, val.FunctionName)
			}
			if val2.Args[0] != "0" {
				t.Fatalf("FunctionName wrong: wanted %s, got %s", lexer.TOKEN_INT, val.FunctionName)
			}
		}
	}

}
