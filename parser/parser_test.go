package parser

import (
	"compiler/lexer"
	"testing"
)

func TestParsing(t *testing.T) {
	input := `int main(){return 0;}`
	tokens := lexer.New(input)
	_ = Parse(tokens)
	t.Fail()
}
