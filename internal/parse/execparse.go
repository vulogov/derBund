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
	code   *string
	errors int
}

func ParserExec(name string, code string) {
	errorListener := new(bundExecErrorListener)
	errorListener.code = &code
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
	var val *vm.Elem
	if l.VM.Current != nil {
		log.Debugf("EXITING Block. Stack size: %v", l.VM.Current.Len())
	} else {
		log.Debugf("EXITING Block. No current stack")
	}
	val = nil
	if l.VM.CanGet() {
		val = l.VM.Get()
	}
	l.VM.EndNS()
	if l.VM.IsStack() {
		if val != nil {
			log.Debug("Pushing value from block to upper stack")
			l.VM.Put(val)
		}
	}
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

func (l *bundExecListener) EnterString_term(c *parser.String_termContext) {
	log.Debugf("String Value: %v", c.GetValue().GetText())
	eh, err := vm.GetType("str")
	if err != nil {
		log.Errorf("BUND type 'str' not defined: %v", err)
		return
	}
	s := c.GetValue().GetText()
	sz := len(s) - 1
	l.VM.Put(eh.FromString(s[1:sz]))
}

func (l *bundExecListener) EnterInteger(c *parser.IntegerContext) {
	log.Debugf("64-bit Integer Value: %v", c.GetValue().GetText())
	eh, err := vm.GetType("int")
	if err != nil {
		log.Errorf("BUND type 'int' not defined: %v", err)
		return
	}
	l.VM.Put(eh.FromString(c.GetValue().GetText()))
}

func (l *bundExecListener) EnterBegin(c *parser.BeginContext) {
	log.Debugf("STACK: pushing to BEGIN")
	l.VM.Mode = false
}

func (l *bundExecListener) EnterEnd(c *parser.EndContext) {
	log.Debugf("STACK: pushing to END")
	l.VM.Mode = true
}

func (l *bundExecListener) EnterCall_term(c *parser.Call_termContext) {
	log.Debugf("CALLING: %v", c.GetValue().GetText())
	l.VM.Exec(c.GetValue().GetText())
}

func (l *bundExecListener) EnterDrop(c *parser.DropContext) {
	log.Debugf("STACK: Drop")
	if l.VM.Current.Len() == 0 {
		log.Warn("Attempt to Drop value from an empty stack")
	} else {
		l.VM.Take()
	}
}
