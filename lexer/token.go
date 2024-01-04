package lexer

const (
	TOKEN_FUNCTION TokenType = "main"
	TOKEN_RETURN   TokenType = "return"
	TOKEN_SEMICOL  TokenType = ";"
	TOKEN_OPAREN   TokenType = "("
	TOKEN_CPAREN   TokenType = ")"
	TOKEN_OCURLY   TokenType = "{"
	TOKEN_CCURLY   TokenType = "}"
	TOKEN_EMPTY    TokenType = "NONE"
	TOKEN_INT      TokenType = "int"
	TOKEN_Char     TokenType = "char"
)

const (
	Literal_Int   LiteralType = "_int"
	Literal_EMPTY TokenType   = "NONE"
)

type TokenType string
type Token struct {
	Type  TokenType
	Value string
}

type LiteralType string
type Literal struct {
	Literal LiteralType
	Value   string
}

func IsToken(name string, val string) Token {
	switch name {
	case string(TOKEN_FUNCTION):
		return Token{TOKEN_FUNCTION, val}
	case string(TOKEN_RETURN):
		return Token{TOKEN_RETURN, ""}
	case string(TOKEN_SEMICOL):
		return Token{TOKEN_SEMICOL, ""}
	case string(TOKEN_OPAREN):
		return Token{TOKEN_OPAREN, ""}
	case string(TOKEN_CPAREN):
		return Token{TOKEN_CPAREN, ""}
	case string(TOKEN_OCURLY):
		return Token{TOKEN_OCURLY, ""}
	case string(TOKEN_CCURLY):
		return Token{TOKEN_CCURLY, ""}
	case string(TOKEN_INT):
		return Token{TOKEN_INT, ""}
	default:
		return Token{TOKEN_EMPTY, string(Literal_EMPTY)}
	}
}
func IsLiteral(name string, val string) Literal {
	switch name {
	case string(Literal_Int):
		return Literal{LiteralType(Literal_Int), val}
	default:
		return Literal{LiteralType(Literal_EMPTY), string(Literal_Int)}

	}
}
