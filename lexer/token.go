package lexer

const (
	LITERAL_INT       TokenType = "int"
	TOKEN_FUNCTION    TokenType = "function"
	TOKEN_RETURN      TokenType = "return"
	TOKEN_SEMICOL     TokenType = ";"
	TOKEN_OPAREN      TokenType = "("
	TOKEN_CPAREN      TokenType = ")"
	TOKEN_OCURLY      TokenType = "{"
	TOKEN_CCURLY      TokenType = "}"
	TOKEN_EMPTY       TokenType = ""
	TOKEN_INT_LITERAL TokenType = "int"
)

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}
