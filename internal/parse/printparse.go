package parse

import (
	// "strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/parser"
)

type bundListener struct {
	*parser.BaseBundListener
	errors int
}

type bundErrorListener struct {
	antlr.ErrorListener
	errors int
}

func ParserPrint(code string) {
	log.Debugf("Code passed on parser print: %v", code)
	errorListener := new(bundErrorListener)
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	listener := new(bundListener)
	if errorListener.errors > 0 {
		log.Errorf("%v lexer errors detected.", errorListener.errors)
		return
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
	if errorListener.errors > 0 {
		log.Errorf("Errors detected: %v", errorListener.errors)
	} else {
		log.Debug("No errors detected")
	}
}

func (l *bundErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.Errorf("Syntax error line=%v, column=%v : %v", line, column, msg)
	l.errors += 1
}
func (l *bundErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Ambiguity Error")
	l.errors += 1
}
func (l *bundErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Attempting in Full Context")
	l.errors += 1
}
func (l *bundErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	log.Errorf("Context sensitivity error")
	l.errors += 1
}

func (l *bundListener) EnterNs(c *parser.NsContext) {
	log.Infof("ENTERING Namespace: %v", c.GetName().GetText())
}

func (l *bundListener) ExitNs(c *parser.NsContext) {
	log.Infof("EXITING Namespace: %v", c.GetName().GetText())
}

func (l *bundListener) EnterBlock(c *parser.BlockContext) {
	log.Infof("ENTERING Block")
}

func (l *bundListener) ExitBlock(c *parser.BlockContext) {
	log.Infof("EXITING Block")
}

func (l *bundListener) EnterTrue_term(c *parser.True_termContext) {
	log.Infof("Value: TRUE")
}

func (l *bundListener) EnterFalse_term(c *parser.False_termContext) {
	log.Infof("Value: FALSE")
}

func (l *bundListener) EnterString_term(c *parser.String_termContext) {
	log.Infof("String Value: %v", c.GetValue().GetText())
}

func (l *bundListener) EnterBegin(c *parser.BeginContext) {
	log.Infof("STACK: pushing to BEGIN")
}

func (l *bundListener) EnterEnd(c *parser.EndContext) {
	log.Infof("STACK: pushing to END")
}

func (l *bundListener) EnterCall_term(c *parser.Call_termContext) {
	log.Infof("CALLING: %v", c.GetValue().GetText())
}
