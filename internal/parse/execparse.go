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

func ParserExec(name string, code string) *vm.VM {
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
		return nil
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
	if errorListener.errors > 0 {
		log.Errorf("Errors detected: %v", errorListener.errors)
		return nil
	} else {
		log.Debug("No errors detected")
	}
	return listener.VM
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

func (l *bundExecListener) EnterFloat(c *parser.FloatContext) {
	log.Debugf("64-bit Float Value: %v", c.GetValue().GetText())
	eh, err := vm.GetType("flt")
	if err != nil {
		log.Errorf("BUND type 'flt' not defined: %v", err)
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

func (l *bundExecListener) EnterDatablock(c *parser.DatablockContext) {
	log.Debugf("ENTERING Data Block")
	blockname := uuid.New().String()
	l.VM.GetNS(blockname)
}

func (l *bundExecListener) ExitDatablock(c *parser.DatablockContext) {
	if l.VM.Current != nil {
		log.Debugf("EXITING Data Block. Stack size: %v", l.VM.Current.Len())
		res := new(vm.Elem)
		res.Type = "dblock"
		res.Value = l.VM.Current
		l.VM.EndNS()
		if l.VM.IsStack() {
			l.VM.Put(res)
		}
	} else {
		log.Debugf("EXITING Data Block. No current stack")
		l.VM.EndNS()
	}
}

func (l *bundExecListener) EnterFloatblock(c *parser.FloatblockContext) {
	log.Debugf("ENTERING Float Block")
	blockname := uuid.New().String()
	l.VM.GetNS(blockname)
}

func (l *bundExecListener) ExitFloatblock(c *parser.FloatblockContext) {
	if l.VM.Current != nil {
		log.Debugf("EXITING Float Block. Stack size: %v", l.VM.Current.Len())
		res := new(vm.Elem)
		res.Type = "fblock"
		res.Value = l.VM.Current
		l.VM.EndNS()
		if l.VM.IsStack() {
			l.VM.Put(res)
		}
	} else {
		log.Debugf("EXITING Float Block. No current stack")
		l.VM.EndNS()
	}
}

func (l *bundExecListener) EnterIntblock(c *parser.IntblockContext) {
	log.Debugf("ENTERING Int Block")
	blockname := uuid.New().String()
	l.VM.GetNS(blockname)
}

func (l *bundExecListener) ExitIntblock(c *parser.IntblockContext) {
	if l.VM.Current != nil {
		log.Debugf("EXITING Int Block. Stack size: %v", l.VM.Current.Len())
		res := new(vm.Elem)
		res.Type = "iblock"
		res.Value = l.VM.Current
		l.VM.EndNS()
		if l.VM.IsStack() {
			l.VM.Put(res)
		}
	} else {
		log.Debugf("EXITING Int Block. No current stack")
		l.VM.EndNS()
	}
}

func (l *bundExecListener) EnterTrueblock(c *parser.TrueblockContext) {
	log.Debugf("ENTERING True Block")
	if l.VM.CanGet() {
		e := l.VM.Get()
		if e.Type == "bool" {
			if e.Value.(bool) == true {
				l.VM.NotIgnore()
				blockname := uuid.New().String()
				l.VM.GetNS(blockname)
			} else {
				log.Debugf("True Block will not be executed")
				l.VM.Ignore()
			}
		} else {
			l.VM.Ignore()
		}
	}
}

func (l *bundExecListener) ExitTrueblock(c *parser.TrueblockContext) {
	log.Debugf("EXITING True Block")
	if !l.VM.MustIgnore() {
		if l.VM.CanGet() {
			l.VM.EndNS()
		}
	}
}

func (l *bundExecListener) EnterFalseblock(c *parser.FalseblockContext) {
	log.Debugf("ENTERING False Block")
	if l.VM.CanGet() {
		e := l.VM.Get()
		if e.Type == "bool" {
			if e.Value.(bool) == false {
				l.VM.NotIgnore()
				blockname := uuid.New().String()
				l.VM.GetNS(blockname)
			} else {
				log.Debugf("False Block will not be executed")
				l.VM.Ignore()
			}
		} else {
			l.VM.Ignore()
		}
	}
}

func (l *bundExecListener) ExitFalseblock(c *parser.FalseblockContext) {
	log.Debugf("EXITING False Block")
	if !l.VM.MustIgnore() {
		if l.VM.CanGet() {
			l.VM.EndNS()
		}
	}
}
