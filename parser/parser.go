package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"fmt"
)

func Parse(tokens lexer.Tokenized) (asttree ast.AST) {
	for tok := tokens.NextToken(); true; tok = tokens.NextToken() {
		switch tok.Type {
		case lexer.TOKEN_INT:
			token := tokens.NextToken()
			if token.Type == lexer.TOKEN_STRING_LITERAL {
				f := parseFunction(tokens)
				asttree.Function = append(asttree.Function, f)
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
	return
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
	functioncalls, id := parseFunctionCalls(token)
	fcall.Functions = functioncalls
	token.Index = id
	return fcall
}
func parseFunctionCalls(token lexer.Tokenized) ([]ast.FUNCTIONCALL, int) {
	tiktok := make([]ast.FUNCTIONCALL, 0)
	tok := token.NextToken()
	id, funccall := parseFuncCall(tok.Value, token)
	token.Index = id + 1
	tiktok = append(tiktok, funccall)
	for token.NextToken().Type != lexer.TOKEN_CCURLY {
		// fmt.Println(token.LastToken)
		id, funccall := parseFuncCall(token.LastToken.Value, token)
		token.Index = id
		tiktok = append(tiktok, funccall)

	}
	return tiktok, token.Index
}
func parseFuncCall(name string, token lexer.Tokenized) (int, ast.FUNCTIONCALL) {
	switch name {
	case "printf":
		return Printf(token)
	case "return":
		return Return(token)
	default:
		panic(fmt.Sprintf("Function %s is not Implemented yet.", name))
	}
}
func Printf(token lexer.Tokenized) (int, ast.FUNCTIONCALL) {
	params := []string{}
	for tok := token.NextToken(); tok.Type != lexer.TOKEN_SEMICOL; tok = token.NextToken() {
		if token.LastToken.Type == lexer.TOKEN_OPAREN {
			params, id := parseParam(token)
			token.Index = id
			return token.Index, ast.FUNCTIONCALL{Name: "printf", Args: params}
		}
		if token.LastToken.Type == lexer.TOKEN_STRING_LITERAL {
			params = append(params, token.LastToken.Value)
		}
	}
	return token.Index, ast.FUNCTIONCALL{}
}
func Return(token lexer.Tokenized) (int, ast.FUNCTIONCALL) {
	params := []string{}
	token.NextToken()
	for tok := token.LastToken; tok.Type != lexer.TOKEN_SEMICOL; tok = token.NextToken() {
		params = append(params, token.LastToken.Value)
	}
	return token.Index, ast.FUNCTIONCALL{Name: "return", Args: params}
}
func parseParam(token lexer.Tokenized) (args []string, id int) {
	if token.LastToken.Type == lexer.TOKEN_OPAREN {
		for tok := token.NextToken(); tok.Type != lexer.TOKEN_CPAREN; tok = token.NextToken() {
			args = append(args, tok.Value)
		}
		id = token.Index
	} else {
		fmt.Println("-----------------------------")
		fmt.Println("Error parsing args")
		fmt.Println("-----------------------------")
	}
	return
}

func Output(compiler string) {
	fmt.Println(compiler)
}
