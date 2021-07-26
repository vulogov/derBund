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
	if l.VM.MustIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("ENTERING Namespace: %v", c.GetName().GetText())
		l.VM.GetNS(c.GetName().GetText())
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "NS", Value: c.GetName().GetText()})
		}
	}
}

func (l *bundExecListener) ExitNs(c *parser.NsContext) {
	if l.VM.MustIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("EXITING Namespace: %v", c.GetName().GetText())
		l.VM.EndNS()
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "exitNS", Value: c.GetName().GetText()})
		}
	}
}

func (l *bundExecListener) EnterBlock(c *parser.BlockContext) {
	if l.VM.MustIgnore() {
		return
	}
	blockname := uuid.New().String()
	if !l.VM.InLambda() {
		log.Debugf("ENTERING Block")
		l.VM.GetNS(blockname)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "BLOCK", Value: blockname})
		}
	}
}

func (l *bundExecListener) ExitBlock(c *parser.BlockContext) {
	if l.VM.MustIgnore() {
		return
	}
	if !l.VM.InLambda() {
		if l.VM.Current != nil {
			log.Debugf("EXITING Block. Stack size: %v", l.VM.Current.Len())
		} else {
			log.Debugf("EXITING Block. No current stack")
		}
		l.VM.EndNS()
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "exitBLOCK", Value: nil})
		}
	}
}

func (l *bundExecListener) EnterTrue_term(c *parser.True_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("bool")
	if err != nil {
		log.Errorf("BUND type 'bool' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText())
	if !l.VM.InLambda() {
		log.Debug("Value: TRUE")
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterFalse_term(c *parser.False_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("bool")
	if err != nil {
		log.Errorf("BUND type 'bool' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText())
	if !l.VM.InLambda() {
		log.Debug("Value: FALSE")
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterString_term(c *parser.String_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("str")
	if err != nil {
		log.Errorf("BUND type 'str' not defined: %v", err)
		return
	}
	s := c.GetValue().GetText()
	sz := len(s) - 1
	val := eh.FromString(s[1:sz])
	if !l.VM.InLambda() {
		log.Debugf("String Value: %v", c.GetValue().GetText())
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterInteger(c *parser.IntegerContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("int")
	if err != nil {
		log.Errorf("BUND type 'int' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText())
	if !l.VM.InLambda() {
		log.Debugf("64-bit Integer Value: %v", c.GetValue().GetText())
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterUinteger(c *parser.UintegerContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("uint")
	if err != nil {
		log.Errorf("BUND type 'uint' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText()[1:])
	if !l.VM.InLambda() {
		log.Debugf("64-bit Unsigned Integer Value: %v", c.GetValue().GetText())
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterFloat(c *parser.FloatContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("flt")
	if err != nil {
		log.Errorf("BUND type 'flt' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText())
	if !l.VM.InLambda() {
		log.Debugf("64-bit Float Value: %v", c.GetValue().GetText())
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterComplex_term(c *parser.Complex_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("cpx")
	if err != nil {
		log.Errorf("BUND type 'cpx' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText())
	if !l.VM.InLambda() {
		log.Debugf("128-bit Complex Value: %v", c.GetValue().GetText())
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterUfloat(c *parser.UfloatContext) {
	if l.VM.CheckIgnore() {
		return
	}
	eh, err := vm.GetType("uflt")
	if err != nil {
		log.Errorf("BUND type 'uflt' not defined: %v", err)
		return
	}
	val := eh.FromString(c.GetValue().GetText()[1:])
	if !l.VM.InLambda() {
		log.Debugf("64-bit Unsigned Float Value: %v", c.GetValue().GetText()[1:])
		l.VM.Put(val)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(val)
		}
	}
}

func (l *bundExecListener) EnterBegin(c *parser.BeginContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("STACK: pushing to BEGIN")
		l.VM.Mode = false
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "MODE", Value: false})
		}
	}
}

func (l *bundExecListener) EnterEnd(c *parser.EndContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("STACK: pushing to END")
		l.VM.Mode = true
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "MODE", Value: true})
		}
	}
}

func (l *bundExecListener) EnterCall_term(c *parser.Call_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("CALLING: %v", c.GetValue().GetText())
		l.VM.Exec(c.GetValue().GetText())
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "CALL", Value: c.GetValue().GetText()})
		}
	}
}

func (l *bundExecListener) EnterCmd_term(c *parser.Cmd_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("OPERATOR: %v", c.GetValue().GetText())
		l.VM.Op(c.GetValue().GetText())
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "OP", Value: c.GetValue().GetText()})
		}
	}
}

func (l *bundExecListener) EnterDrop(c *parser.DropContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
		log.Debugf("STACK: Drop")
		if l.VM.IsStack() {
			if l.VM.Current.Len() == 0 {
				log.Warn("Attempt to Drop value from an empty stack")
			} else {
				l.VM.Take()
			}
		} else {
			log.Error("Attempt to Drop value with empty context")
		}
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "DROP", Value: nil})
		}
	}
}

func (l *bundExecListener) EnterDatablock(c *parser.DatablockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	blockname := uuid.New().String()
	if !l.VM.InLambda() {
		log.Debugf("ENTERING Data Block")
		l.VM.GetNS(blockname)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "DBLOCK", Value: blockname})
		}
	}
}

func (l *bundExecListener) ExitDatablock(c *parser.DatablockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
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
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "exitDBLOCK", Value: nil})
		}
	}
}

func (l *bundExecListener) EnterFloatblock(c *parser.FloatblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	blockname := uuid.New().String()
	if !l.VM.InLambda() {
		log.Debugf("ENTERING Float Block")
		l.VM.GetNS(blockname)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "FBLOCK", Value: blockname})
		}
	}
}

func (l *bundExecListener) ExitFloatblock(c *parser.FloatblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
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
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "exitFBLOCK", Value: nil})
		}
	}
}

func (l *bundExecListener) EnterIntblock(c *parser.IntblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	blockname := uuid.New().String()
	if !l.VM.InLambda() {
		log.Debugf("ENTERING Int Block")
		l.VM.GetNS(blockname)
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "IBLOCK", Value: blockname})
		}
	}
}

func (l *bundExecListener) ExitIntblock(c *parser.IntblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.InLambda() {
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
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "exitIBLOCK", Value: nil})
		}
	}
}

func (l *bundExecListener) EnterUintblock(c *parser.UintblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	log.Debugf("ENTERING Unsigned Int Block")
	blockname := uuid.New().String()
	l.VM.GetNS(blockname)
}

func (l *bundExecListener) ExitUintblock(c *parser.UintblockContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if l.VM.Current != nil {
		log.Debugf("EXITING Unsigned Int Block. Stack size: %v", l.VM.Current.Len())
		res := new(vm.Elem)
		res.Type = "uiblock"
		res.Value = l.VM.Current
		l.VM.EndNS()
		if l.VM.IsStack() {
			l.VM.Put(res)
		}
	} else {
		log.Debugf("EXITING Unsigned Int Block. No current stack")
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
				l.VM.Ignore()
				log.Debugf("True Block will not be executed: %v", l.VM.CheckIgnore())
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

func (l *bundExecListener) EnterReturn_term(c *parser.Return_termContext) {
	if l.VM.CheckIgnore() {
		return
	}
	cmd := c.GetValue().GetText()
	if !l.VM.InLambda() {
		log.Debugf("STACK: Return")
		if l.VM.NSStack.Len() < 1 {
			log.Errorf("Namespace stack is too shallow for RETURN operation: %v", l.VM.NSStack.Len())
			return
		}
		nsr := l.VM.NSStack.Back().(*vm.NS)
		log.Debugf("Return push to %v", nsr.Name)
		var e *vm.Elem
		switch cmd {
		case "$":
			e = l.VM.Take()
		case "$$":
			e = l.VM.Get()
		default:
			log.Errorf("I do not know how to do this RETURN operation: %v", cmd)
		}
		if l.VM.Mode {
			nsr.Stack.PushBack(e)
		} else {
			nsr.Stack.PushFront(e)
		}
	} else {
		ls := l.VM.CurrentLambda()
		if ls != nil {
			ls.PushBack(&vm.Elem{Type: "RETURN", Value: cmd})
		}
	}
}

func (l *bundExecListener) EnterLambda(c *parser.LambdaContext) {
	if l.VM.CheckIgnore() {
		return
	}
	log.Debugf("LAMBDA(start): %v", c.GetName().GetText())
	if !l.VM.IsStack() {
		log.Errorf("Attempt of defining Lambda function with empty context")
		return
	}
	l.VM.CurrentNS.GetLambda(c.GetName().GetText())
	l.VM.CurrentNS.InLambda(c.GetName().GetText())
}

func (l *bundExecListener) ExitLambda(c *parser.LambdaContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.IsStack() {
		log.Errorf("Attempt to close Lambda function with empty context")
		return
	}

	ls := l.VM.CurrentLambda()
	l.VM.CurrentNS.CloseLambda()
	log.Debugf("LAMBDA(fin): %v, size: %v", c.GetName().GetText(), ls.Len())

}

func (l *bundExecListener) EnterAlambda(c *parser.AlambdaContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.IsStack() {
		log.Errorf("Attempt of defining Lambda function with empty context")
		return
	}
	lambdaname := uuid.New().String()
	log.Debugf("ALAMBDA(start): %v", lambdaname)
	l.VM.CurrentNS.GetLambda(lambdaname)
	l.VM.CurrentNS.InLambda(lambdaname)
}

func (l *bundExecListener) ExitAlambda(c *parser.AlambdaContext) {
	if l.VM.CheckIgnore() {
		return
	}
	if !l.VM.IsStack() {
		log.Errorf("Attempt to close Lambda function with empty context")
		return
	}
	name := l.VM.CurrentNS.NameOfCurrentLambda()
	if name == "" {
		log.Errorf("Attempt to close Lambda function with no Lambda function")
		return
	}
	res := new(vm.Elem)
	res.Type = "CALL"
	res.Value = name
	ls := l.VM.CurrentLambda()
	l.VM.CurrentNS.CloseLambda()
	if l.VM.IsStack() {
		l.VM.Put(res)
	}
	log.Debugf("ALAMBDA(fin): %v, size: %v", name, ls.Len())
}
