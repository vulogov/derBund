package bund

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/conf"
	"github.com/vulogov/derBund/internal/parse"
	"github.com/vulogov/derBund/internal/signal"
)

func Eval() {
	Init()
	log.Debug("[ BUND ] tsak.Eval() is reached")
	if *conf.LexerPrint {
		parse.LexerPrint(*conf.Expr)
		return
	}
	if *conf.ParserPrint {
		parse.ParserPrint(*conf.Expr)
		return
	}
	parse.ParserExec("<eval>", *conf.Expr)
	signal.ExitRequest()
}
