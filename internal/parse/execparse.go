package parse

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/derBund/internal/parser"
)

type bundExecListener struct {
	*parser.BaseBundListener
	inBlock bool
	inPair  bool
}

func ParserExec(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	listener := new(bundExecListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
}
