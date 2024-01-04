package lexer

import (
	"strings"
	"unicode"
)

func Lexer(code []byte) []Token {
	var Tokens = make([]Token, 0)
	for i := 0; i < len(code); i++ {
		token := make([]string, 0)
		if unicode.IsLetter(rune(code[i])) {
			for unicode.IsLetter(rune(code[i])) {
				token = append(token, string(code[i]))
				i++
			}
			i--
			word := strings.TrimSpace(strings.Join(token, ""))
			Tokens = append(Tokens, IsToken(word, word))
		} else if unicode.IsDigit(rune(code[i])) {
			for unicode.IsDigit(rune(code[i])) {
				token = append(token, string(code[i]))
				i++
			}
			i--
			word := strings.TrimSpace(strings.Join(token, ""))
			Tokens = append(Tokens, Token{TOKEN_INT, word})
		} else if unicode.IsPunct(rune(code[i])) {
			word := strings.TrimSpace(string(code[i]))
			Tokens = append(Tokens, IsToken(word, word))
		}
		// else if unicode.IsSpace(rune(code[i])) {
		// }
		clear(token)
	}
	return Tokens
}
