package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"fmt"
	"strings"
)

func Parse(tokens lexer.Tokenized) (asttree ast.AST) {
	for tok := tokens.NextToken(); true; tok = tokens.NextToken() {
		switch tok.Type {
		case lexer.TOKEN_INT:
			token := tokens.NextToken()
			if token.Type == lexer.TOKEN_STRING_LITERAL {
				_ = parseFunction(tokens)
				// fmt.Println(function)
				return
			}
			continue
		case lexer.TOKEN_END:
			fmt.Println("EOF", tok.Type)
			// panic(fmt.Sprintf(
			// 	"Unsupported Token %q",
			// 	tokens.Token[i].Type,
			// ))
		}
	}
	return ast.AST{}
}

func parseFunction(token lexer.Tokenized) (fcall ast.FUNCTION) {
	fcall = ast.FUNCTION{ReturnType: token.Token[token.Index-1].Value}
	tok := token.LastToken
	// fmt.Println("line 34: tok", tok)
	if tok.Type != lexer.TOKEN_STRING_LITERAL {
		fmt.Printf(
			"Error on [%d:%d], wanted %q, got %q\n",
			tok.Location.Line,
			tok.Location.Column,
			lexer.TOKEN_CCURLY,
			tok.Type,
		)
		panic(
			fmt.Sprintf(
				"error while parsing on line[%d:%d]",
				tok.Location.Line,
				tok.Location.Column,
			),
		)
	}
	fcall.FunctionName = tok.Value
	tok = token.NextToken()
	if tok.Type != lexer.TOKEN_OPAREN {
		fmt.Printf(
			"Error on [%d:%d], wanted %q, got %q\n",
			tok.Location.Line,
			tok.Location.Column,
			lexer.TOKEN_CCURLY,
			tok.Type,
		)
		panic(
			fmt.Sprintf(
				"error while parsing on line[%d:%d]",
				tok.Location.Line,
				tok.Location.Column,
			),
		)
	}
	tok = token.NextToken()
	if tok.Type != lexer.TOKEN_CPAREN {
		fmt.Printf(
			"Error on [%d:%d], wanted %q, got %q\n",
			tok.Location.Line,
			tok.Location.Column,
			lexer.TOKEN_CCURLY,
			tok.Type,
		)
		panic(
			fmt.Sprintf(
				"error while parsing on line[%d:%d]",
				tok.Location.Line,
				tok.Location.Column,
			),
		)
	}
	tok = token.NextToken()
	if tok.Type != lexer.TOKEN_OCURLY {
		fmt.Printf(
			"Error on [%d:%d], wanted %q, got %q\n",
			tok.Location.Line,
			tok.Location.Column,
			lexer.TOKEN_CCURLY,
			tok.Type,
		)
		panic(
			fmt.Sprintf(
				"error while parsing on line[%d:%d]",
				tok.Location.Line,
				tok.Location.Column,
			),
		)
	}
	fcall.Functions, token = parseFunctionCalls(token)
	if tok.Type != lexer.TOKEN_CCURLY {
		// fmt.Printf(
		// 	"Error on [%d:%d], wanted %q, got %q\n",
		// 	tok.Location.Line,
		// 	tok.Location.Column,
		// 	lexer.TOKEN_CCURLY,
		// 	tok.Type,
		// )
		// panic(
		// 	fmt.Sprintf(
		// 		"error while parsing on line[%d:%d]",
		// 		tok.Location.Line,
		// 		tok.Location.Column,
		// 	),
		// )
	}
	// fmt.Printf("Name: %s, Returntype: %s, Functions:%v\n", fcall.FunctionName, fcall.ReturnType, fcall.Functions)
	return
}
func parseFunctionCalls(token lexer.Tokenized) ([]ast.FUNCTIONCALL, lexer.Tokenized) {
	tiktok := make([]ast.FUNCTIONCALL, 0)
	parsedfunc := ast.FUNCTIONCALL{}
	tok := token.NextToken()
	if tok.Type != lexer.TOKEN_FUNCTIONCALL {
		fmt.Printf(
			"Error on [%d:%d], wanted %q, got %q\n",
			tok.Location.Line,
			tok.Location.Column,
			lexer.TOKEN_FUNCTIONCALL,
			tok.Type,
		)
		panic(
			fmt.Sprintf(
				"error while parsing on line[%d:%d]",
				tok.Location.Line,
				tok.Location.Column,
			),
		)

	}
	parsedfunc.Name = tok.Value
	name := tok.Value
	parseFuncCall(name, token)
	// if tok.Type != lexer.TOKEN_OPAREN {
	// 	fmt.Printf(
	// 		"Error on [%d:%d], wanted %q, got %q\n",
	// 		tok.Location.Line,
	// 		tok.Location.Column,
	// 		lexer.TOKEN_CPAREN,
	// 		tok.Type,
	// 	)
	// 	panic(
	// 		fmt.Sprintf(
	// 			"error while parsing on line[%d:%d]",
	// 			tok.Location.Line,
	// 			tok.Location.Column,
	// 		),
	// 	)
	//
	// }
	// tok = token.NextToken()
	// fmt.Println("FUnction:", tok)
	// if tok.Type != lexer.TOKEN_STRING_LITERAL {
	// 	fmt.Printf(
	// 		"Error on [%d:%d], wanted %q, got %q\n",
	// 		tok.Location.Line,
	// 		tok.Location.Column,
	// 		lexer.TOKEN_CPAREN,
	// 		tok.Type,
	// 	)
	// 	panic(
	// 		fmt.Sprintf(
	// 			"error while parsing on line[%d:%d]",
	// 			tok.Location.Line,
	// 			tok.Location.Column,
	// 		),
	// 	)
	//
	// }
	// fmt.Println(parsedfunc)
	// switch tok.Value {
	// case "printf":
	//
	// 	break
	// default:
	// 	panic("Function is not Implemented yet.")
	// }
	// tiktok = append(tiktok, parsedfunc)
	return tiktok, token
}
func parseFuncCall(name string, token lexer.Tokenized) {
	switch name {
	case "printf":
		Printf(token)
		return
	default:
		panic("Function is not Implemented yet.")
	}
}
func Printf(token lexer.Tokenized) int {
	// fmt.Println("Called Printf")
	for tok := token.NextToken(); tok.Type != lexer.TOKEN_SEMICOL; tok = token.NextToken() {
		if token.LastToken.Type == lexer.TOKEN_OPAREN {
			params, id := parseParam(token)
			token.Index = id
			Output(fmt.Sprintf("console.log(\"%s\")", strings.Join(params, "\",\"")))
		}
	}
	return token.Index
}
func parseParam(token lexer.Tokenized) (args []string, id int) {
	if token.LastToken.Type == lexer.TOKEN_OPAREN {
		for tok := token.NextToken(); tok.Type != lexer.TOKEN_CPAREN; tok = token.NextToken() {
			args = append(args, tok.Value)
		}
		id = token.Index
	} else {
		// fmt.Println("-----------------------------")
		// fmt.Println("Error parsing args")
		// fmt.Println("-----------------------------")
	}
	return
}

func Output(compiler string) {
	fmt.Println(compiler)
}
