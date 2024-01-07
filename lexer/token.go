package lexer

type TokenType string
type Token struct {
	Type     TokenType
	Value    string
	Location TextLocation
}
type Tokenized struct {
	LastToken Token
	Token     []Token
	Index     int
}
type TextLocation struct {
	Line   int
	Column int
}

func (t *Tokenized) NextToken() Token {
	t.Index++
	if t.Index > len(t.Token) {
		t.LastToken = Token{Type: TOKEN_END}
		return t.LastToken
	}
	t.LastToken = t.Token[t.Index]
	return t.Token[t.Index]
}

var TokenTypes = make([]TokenType, 0)

const (
	TOKEN_INT            TokenType = "TOKEN_INT"
	TOKEN_SEMICOL        TokenType = "TOKEN_SEMICOL"
	TOKEN_OPAREN         TokenType = "TOKEN_OPAREN"
	TOKEN_CPAREN         TokenType = "TOKEN_CPAREN"
	TOKEN_OCURLY         TokenType = "TOKEN_OCURLY"
	TOKEN_CCURLY         TokenType = "TOKEN_CCURLY"
	TOKEN_RETURN         TokenType = "TOKEN_RETURN"
	TOKEN_ILLEGAL        TokenType = "TOKEN_ILLEGAL"
	TOKEN_FUNCTIONCALL   TokenType = "TOKEN_FUNCTIONCALL"
	TOKEN_STRING_LITERAL TokenType = "TOKEN_STRING_LITERAL"
	TOKEN_END            TokenType = "TOKEN_END"
)

func GetToken(tokenstr string) (validtoken TokenType) {
	switch tokenstr {
	case "printf":
		return TOKEN_FUNCTIONCALL
	case "return":
		return TOKEN_RETURN
	case "int":
		return TOKEN_INT
	default:
		return TOKEN_STRING_LITERAL
	}

}
func getLocation(lex Lexer_Type, offset int) TextLocation {
	return TextLocation{Line: lex.Line, Column: lex.Index - lex.Bol - offset}
}

// const (
//
//	TYPE_INT  string = "INT"
//	TYPE_Void string = "void"
//
// )
//
// type STMT_KIND string
//
// const (
//
//	STMT_FUNCALL STMT_KIND = "STMT_FUNCALL"
//	STMT_RETURN  STMT_KIND = "STMT_RETURN"
//
// )

//
// type FUNC struct {
// 	NAME string
// 	Body string
// }
// type RetStmt struct {
// 	Expr string
// }
// type Funccall struct {
// 	Name string
// 	args string
// }
// type STMT struct {
// 	Kind STMT_KIND
// }
