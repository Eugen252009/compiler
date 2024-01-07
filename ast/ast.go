package ast

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
