package parse

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/uuid"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/parser"
	"github.com/vulogov/derBund/internal/vm"
)

type bundExecListener struct {
	*parser.BaseBundListener
	VM     *vm.VM
	errors int
}

type bundExecErrorListener struct {
	antlr.ErrorListener
	errors int
}

func ParserExec(name string, code string) {
	errorListener := new(bundExecErrorListener)
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	listener := new(bundExecListener)
	listener.VM = vm.NewVM(name)
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

func (l *bundExecErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.Errorf("Syntax error line=%v, column=%v : %v", line, column, msg)
	l.errors += 1
}
func (l *bundExecErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Ambiguity Error")
	l.errors += 1
}
func (l *bundExecErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Attempting in Full Context")
	l.errors += 1
}
func (l *bundExecErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	log.Errorf("Context sensitivity error")
	l.errors += 1
}

func (l *bundExecListener) EnterNs(c *parser.NsContext) {
	log.Debugf("ENTERING Namespace: %v", c.GetName().GetText())
	l.VM.GetNS(c.GetName().GetText())
}

func (l *bundExecListener) ExitNs(c *parser.NsContext) {
	log.Debugf("EXITING Namespace: %v", c.GetName().GetText())
	l.VM.EndNS()
}

func (l *bundExecListener) EnterBlock(c *parser.BlockContext) {
	log.Debugf("ENTERING Block")
	blockname := uuid.New().String()
	l.VM.GetNS(blockname)
}

func (l *bundExecListener) ExitBlock(c *parser.BlockContext) {
	if l.VM.Current != nil {
		log.Debugf("EXITING Block. Stack size: %v", l.VM.Current.Len())
	} else {
		log.Debugf("EXITING Block. No current stack")
	}
	l.VM.EndNS()
}

func (l *bundExecListener) EnterTrue_term(c *parser.True_termContext) {
	log.Debug("Value: TRUE")
	eh, err := vm.GetType("bool")
	if err != nil {
		log.Errorf("BUND type 'bool' not defined: %v", err)
		return
	}
	l.VM.Put(eh.FromString(c.GetValue().GetText()))
}

func (l *bundExecListener) EnterFalse_term(c *parser.False_termContext) {
	log.Debug("Value: FALSE")
	eh, err := vm.GetType("bool")
	if err != nil {
		log.Errorf("BUND type 'bool' not defined: %v", err)
		return
	}
	l.VM.Put(eh.FromString(c.GetValue().GetText()))
}
