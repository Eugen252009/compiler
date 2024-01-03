package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

var Tokens = make([]Token, 0)

func Lexer(code []byte) {
	fmt.Println(string(code))
	for i := 0; i < len(code); i++ {
		if unicode.IsLetter(rune(code[i])) {
			token := make([]string, 0)
			for unicode.IsLetter(rune(code[i])) {
				token = append(token, string(code[i]))
				i++
			}
			i--
			// word := strings.TrimSpace(strings.Join(token, ""))
			Tokens = append(Tokens, checkForTokens(strings.Join(token, "")))
			clear(token)
		} else if unicode.IsDigit(rune(code[i])) {
			token := make([]string, 0)
			for unicode.IsDigit(rune(code[i])) {
				token = append(token, string(code[i]))
				i++
			}
			i--
			fmt.Println("Found ", token)
			word := strings.TrimSpace(strings.Join(token, ""))
			Tokens = append(Tokens, Token{TOKEN_INT_LITERAL, word})
			clear(token)
		} else if ';' == rune(code[i]) {
			Tokens = append(Tokens, Token{TOKEN_SEMICOL, ""})
			fmt.Println("Found ", string(TOKEN_SEMICOL))
			continue
		} else if unicode.IsSpace(rune(code[i])) {
			continue
		}

	}
	for _, val := range Tokens {
		switch val.Type {
		case TOKEN_INT_LITERAL:
			fmt.Printf(val.Value)
		case TOKEN_SEMICOL:
			fmt.Println(";")
		default:
			fmt.Printf(string(val.Type))
		}
		fmt.Printf(" ")
		// fmt.Println(Tokens)
	}
}

func checkForTokens(name string) Token {
	switch name {
	case string(TOKEN_RETURN):
		return Token{TOKEN_RETURN, ""}
	default:
		return Token{TOKEN_EMPTY, ""}
	}
}
