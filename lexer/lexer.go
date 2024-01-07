package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Lexer_Type struct {
	Input string
	Index int
	Line  int
	Bol   int
}

func (lexer *Lexer_Type) NextChar() rune {
	lexer.Index++
	if lexer.Index >= len(string(lexer.Input)) {
		return '$'
	}
	return rune(lexer.Input[lexer.Index])
}
func New(input string) Tokenized {
	lex := Lexer_Type{Input: input, Index: -1, Bol: 1}
	// lex.Printline()
	return Tokenize(lex)
}

func Tokenize(rawstring Lexer_Type) (tokens Tokenized) {
	tokens.Index = -1
	for char := rawstring.NextChar(); char != '$'; char = rawstring.NextChar() {
		switch char {
		case '\r':
		case '\n':
			rawstring.Bol = rawstring.Index
			rawstring.Line++
			continue
		case '\t':
		case ' ':
			continue
		case '#':
			for char := rawstring.NextChar(); char != '\n'; char = rawstring.NextChar() {
			}
			rawstring.Bol = rawstring.Index
			rawstring.Line++

			continue
		case '(':
			tokens.Token = append(tokens.Token,
				Token{
					Type:     TOKEN_OPAREN,
					Value:    "(",
					Location: getLocation(rawstring, 0),
				})
			continue
		case ')':
			tokens.Token = append(
				tokens.Token,
				Token{
					Type:     TOKEN_CPAREN,
					Value:    ")",
					Location: getLocation(rawstring, 0),
				})
			continue
		case '{':
			tokens.Token = append(
				tokens.Token,
				Token{
					Type:     TOKEN_OCURLY,
					Value:    "{",
					Location: getLocation(rawstring, 0),
				})
			continue
		case '}':
			tokens.Token = append(tokens.Token, Token{
				Type:     TOKEN_CCURLY,
				Value:    "{",
				Location: getLocation(rawstring, 0),
			})
			continue
		case '"':
			char = rawstring.NextChar()
			var word = []string{}
			for char != '"' {
				word = append(word, string(char))
				char = rawstring.NextChar()
			}
			token := strings.TrimSpace(strings.Join(word, ""))
			tokens.Token = append(tokens.Token, Token{
				Type:     TOKEN_STRING_LITERAL,
				Value:    token,
				Location: getLocation(rawstring, len(token)),
			})
			// tokens = append(tokens, Token{Type: TOKEN_OPAREN, Value: "\"", Location: TextLocation{Line: rawstring.Line, Column: rawstring.Index - rawstring.Bol}})
			continue
		case ';':
			tokens.Token = append(
				tokens.Token,
				Token{
					Type:     TOKEN_SEMICOL,
					Value:    ";",
					Location: getLocation(rawstring, 0),
				})
			continue
		}
		var word = []string{}
		id := rawstring.Index
		if unicode.IsLetter(char) {
			for unicode.IsLetter(char) {
				// fmt.Println("letter", string(char))
				word = append(word, string(char))
				char = rawstring.NextChar()
			}
			rawstring.Index--
			token := strings.TrimSpace(strings.Join(word, ""))
			// fmt.Println("token", string(token))
			if GetToken(token) == TOKEN_ILLEGAL {
				panic(
					fmt.Sprintf(
						"ILLEGAL Token on line [%d:%d]:%s\n",
						rawstring.Line,
						id-rawstring.Bol,
						token,
					))
			} else {
				tokens.Token = append(
					tokens.Token,
					Token{
						Type:     GetToken(token),
						Value:    token,
						Location: getLocation(rawstring, 0),
					})
			}
		}
		if unicode.IsDigit(char) {
			for unicode.IsDigit(char) {
				// fmt.Println("letter", string(char))
				word = append(word, string(char))
				char = rawstring.NextChar()
			}
			rawstring.Index--
			token := strings.TrimSpace(strings.Join(word, ""))
			// fmt.Println("token", string(token))
			if GetToken(token) == TOKEN_ILLEGAL {
				panic(
					fmt.Sprintf(
						"ILLEGAL Token on line [%d:%d]:%s\n",
						rawstring.Line,
						id-rawstring.Bol,
						token,
					))
			} else {
				tokens.Token = append(
					tokens.Token,
					Token{
						Type:     GetToken(token),
						Value:    token,
						Location: getLocation(rawstring, 0),
					})
			}
		}

	}
	return
}
