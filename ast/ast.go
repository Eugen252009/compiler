package ast

import (
	"fmt"
	"strings"
)

type AST struct {
	Function []FUNCTION
}

type FUNCTIONCALL struct {
	Name string
	Args []string
}

type FUNCTION struct {
	FunctionName string
	ReturnType   string
	Functions    []FUNCTIONCALL
}

func ToJavaScript(ast AST) {
	result := ""
	for _, val := range ast.Function {
		// fmt.Println(id, val)
		funcresult := printJavaScriptFunction(val)
		result += fmt.Sprintf("function %s(){\n%s}\n%s();", val.FunctionName, funcresult, val.FunctionName)
	}
	fmt.Println(result)
}

func printJavaScriptFunction(function FUNCTION) (result string) {
	for _, val := range function.Functions {
		switch val.Name {
		case "printf":
			result += jsprint(val.Args)
			break
		}
	}

	// fmt.Printf("function %s(){%s}", function.FunctionName, printJavaScriptFunctionCall())
	return
}

func printJavaScriptFunctionCall(call FUNCTIONCALL) string {
	return ""
}
func jsprint(args []string) string {
	return fmt.Sprintf("console.log(\"%s\");\n", strings.Join(args, "\",\""))
}
