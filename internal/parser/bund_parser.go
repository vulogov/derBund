// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 77, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 7,
	2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 3, 3, 5, 3, 33, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 7, 5, 49, 10, 5, 12, 5, 14, 5, 52, 11, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 7, 6, 58, 10, 6, 12, 6, 14, 6, 61, 11, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 2, 2, 76, 2, 27,
	3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10,
	55, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2, 14, 66, 3, 2, 2, 2, 16, 68, 3, 2, 2,
	2, 18, 70, 3, 2, 2, 2, 20, 72, 3, 2, 2, 2, 22, 74, 3, 2, 2, 2, 24, 26,
	5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2,
	27, 28, 3, 2, 2, 2, 28, 3, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 33, 5, 8,
	5, 2, 31, 33, 5, 10, 6, 2, 32, 30, 3, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33,
	5, 3, 2, 2, 2, 34, 43, 5, 8, 5, 2, 35, 43, 5, 10, 6, 2, 36, 43, 5, 12,
	7, 2, 37, 43, 5, 14, 8, 2, 38, 43, 5, 16, 9, 2, 39, 43, 5, 18, 10, 2, 40,
	43, 5, 20, 11, 2, 41, 43, 5, 22, 12, 2, 42, 34, 3, 2, 2, 2, 42, 35, 3,
	2, 2, 2, 42, 36, 3, 2, 2, 2, 42, 37, 3, 2, 2, 2, 42, 38, 3, 2, 2, 2, 42,
	39, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 7, 3, 2, 2,
	2, 44, 45, 7, 3, 2, 2, 45, 46, 7, 13, 2, 2, 46, 50, 7, 10, 2, 2, 47, 49,
	5, 6, 4, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2,
	50, 51, 3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7,
	4, 2, 2, 54, 9, 3, 2, 2, 2, 55, 59, 7, 5, 2, 2, 56, 58, 5, 6, 4, 2, 57,
	56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 6, 2, 2, 63, 11,
	3, 2, 2, 2, 64, 65, 7, 7, 2, 2, 65, 13, 3, 2, 2, 2, 66, 67, 7, 8, 2, 2,
	67, 15, 3, 2, 2, 2, 68, 69, 7, 9, 2, 2, 69, 17, 3, 2, 2, 2, 70, 71, 7,
	13, 2, 2, 71, 19, 3, 2, 2, 2, 72, 73, 7, 10, 2, 2, 73, 21, 3, 2, 2, 2,
	74, 75, 7, 11, 2, 2, 75, 23, 3, 2, 2, 2, 7, 27, 32, 42, 50, 59,
}
var literalNames = []string{
	"", "'['", "';;'", "'('", "')'", "", "", "", "':'", "';'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "TRUE", "FALSE", "STRING", "TOBEGIN", "TOEND", "SLASH",
	"NAME", "SKIP_",
}

var ruleNames = []string{
	"expressions", "root_term", "term", "ns", "block", "true_term", "false_term",
	"string_term", "call_term", "begin", "end",
}

type BundParser struct {
	*antlr.BaseParser
}

// NewBundParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BundParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundParser(input antlr.TokenStream) *BundParser {
	this := new(BundParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bund.g4"

	return this
}

// BundParser tokens.
const (
	BundParserEOF     = antlr.TokenEOF
	BundParserT__0    = 1
	BundParserT__1    = 2
	BundParserT__2    = 3
	BundParserT__3    = 4
	BundParserTRUE    = 5
	BundParserFALSE   = 6
	BundParserSTRING  = 7
	BundParserTOBEGIN = 8
	BundParserTOEND   = 9
	BundParserSLASH   = 10
	BundParserNAME    = 11
	BundParserSKIP_   = 12
)

// BundParser rules.
const (
	BundParserRULE_expressions = 0
	BundParserRULE_root_term   = 1
	BundParserRULE_term        = 2
	BundParserRULE_ns          = 3
	BundParserRULE_block       = 4
	BundParserRULE_true_term   = 5
	BundParserRULE_false_term  = 6
	BundParserRULE_string_term = 7
	BundParserRULE_call_term   = 8
	BundParserRULE_begin       = 9
	BundParserRULE_end         = 10
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllRoot_term() []IRoot_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRoot_termContext)(nil)).Elem())
	var tst = make([]IRoot_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRoot_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Root_term(i int) IRoot_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRoot_termContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *BundParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BundParserRULE_expressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__2 {
		{
			p.SetState(22)
			p.Root_term()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRoot_termContext is an interface to support dynamic dispatch.
type IRoot_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_termContext differentiates from other interfaces.
	IsRoot_termContext()
}

type Root_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_termContext() *Root_termContext {
	var p = new(Root_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_root_term
	return p
}

func (*Root_termContext) IsRoot_termContext() {}

func NewRoot_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_termContext {
	var p = new(Root_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_root_term

	return p
}

func (s *Root_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Root_termContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Root_termContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Root_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterRoot_term(s)
	}
}

func (s *Root_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitRoot_term(s)
	}
}

func (p *BundParser) Root_term() (localctx IRoot_termContext) {
	localctx = NewRoot_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BundParserRULE_root_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(28)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(29)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *TermContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TermContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *TermContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *TermContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *TermContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *TermContext) Begin() IBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginContext)
}

func (s *TermContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BundParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(32)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(33)
			p.Block()
		}

	case BundParserTRUE:
		{
			p.SetState(34)
			p.True_term()
		}

	case BundParserFALSE:
		{
			p.SetState(35)
			p.False_term()
		}

	case BundParserSTRING:
		{
			p.SetState(36)
			p.String_term()
		}

	case BundParserNAME:
		{
			p.SetState(37)
			p.Call_term()
		}

	case BundParserTOBEGIN:
		{
			p.SetState(38)
			p.Begin()
		}

	case BundParserTOEND:
		{
			p.SetState(39)
			p.End()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INsContext is an interface to support dynamic dispatch.
type INsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsNsContext differentiates from other interfaces.
	IsNsContext()
}

type NsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyNsContext() *NsContext {
	var p = new(NsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ns
	return p
}

func (*NsContext) IsNsContext() {}

func NewNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsContext {
	var p = new(NsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ns

	return p
}

func (s *NsContext) GetParser() antlr.Parser { return s.parser }

func (s *NsContext) GetName() antlr.Token { return s.name }

func (s *NsContext) SetName(v antlr.Token) { s.name = v }

func (s *NsContext) Get_term() ITermContext { return s._term }

func (s *NsContext) Set_term(v ITermContext) { s._term = v }

func (s *NsContext) GetBody() []ITermContext { return s.body }

func (s *NsContext) SetBody(v []ITermContext) { s.body = v }

func (s *NsContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *NsContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *NsContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *NsContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *NsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterNs(s)
	}
}

func (s *NsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitNs(s)
	}
}

func (p *BundParser) Ns() (localctx INsContext) {
	localctx = NewNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_ns)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(43)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(44)
		p.Match(BundParserTOBEGIN)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserSTRING)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserNAME))) != 0 {
		{
			p.SetState(45)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(BundParserT__1)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_term() ITermContext { return s._term }

func (s *BlockContext) Set_term(v ITermContext) { s._term = v }

func (s *BlockContext) GetBody() []ITermContext { return s.body }

func (s *BlockContext) SetBody(v []ITermContext) { s.body = v }

func (s *BlockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BlockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BundParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(BundParserT__2)
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserSTRING)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserNAME))) != 0 {
		{
			p.SetState(54)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ITrue_termContext is an interface to support dynamic dispatch.
type ITrue_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsTrue_termContext differentiates from other interfaces.
	IsTrue_termContext()
}

type True_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyTrue_termContext() *True_termContext {
	var p = new(True_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_true_term
	return p
}

func (*True_termContext) IsTrue_termContext() {}

func NewTrue_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *True_termContext {
	var p = new(True_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_true_term

	return p
}

func (s *True_termContext) GetParser() antlr.Parser { return s.parser }

func (s *True_termContext) GetValue() antlr.Token { return s.value }

func (s *True_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *True_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BundParserTRUE, 0)
}

func (s *True_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *True_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *True_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTrue_term(s)
	}
}

func (s *True_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTrue_term(s)
	}
}

func (p *BundParser) True_term() (localctx ITrue_termContext) {
	localctx = NewTrue_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_true_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)

		var _m = p.Match(BundParserTRUE)

		localctx.(*True_termContext).value = _m
	}

	return localctx
}

// IFalse_termContext is an interface to support dynamic dispatch.
type IFalse_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFalse_termContext differentiates from other interfaces.
	IsFalse_termContext()
}

type False_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFalse_termContext() *False_termContext {
	var p = new(False_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_false_term
	return p
}

func (*False_termContext) IsFalse_termContext() {}

func NewFalse_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *False_termContext {
	var p = new(False_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_false_term

	return p
}

func (s *False_termContext) GetParser() antlr.Parser { return s.parser }

func (s *False_termContext) GetValue() antlr.Token { return s.value }

func (s *False_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *False_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BundParserFALSE, 0)
}

func (s *False_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *False_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *False_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFalse_term(s)
	}
}

func (s *False_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFalse_term(s)
	}
}

func (p *BundParser) False_term() (localctx IFalse_termContext) {
	localctx = NewFalse_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_false_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)

		var _m = p.Match(BundParserFALSE)

		localctx.(*False_termContext).value = _m
	}

	return localctx
}

// IString_termContext is an interface to support dynamic dispatch.
type IString_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsString_termContext differentiates from other interfaces.
	IsString_termContext()
}

type String_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyString_termContext() *String_termContext {
	var p = new(String_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_string_term
	return p
}

func (*String_termContext) IsString_termContext() {}

func NewString_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_termContext {
	var p = new(String_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_string_term

	return p
}

func (s *String_termContext) GetParser() antlr.Parser { return s.parser }

func (s *String_termContext) GetValue() antlr.Token { return s.value }

func (s *String_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *String_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(BundParserSTRING, 0)
}

func (s *String_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterString_term(s)
	}
}

func (s *String_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitString_term(s)
	}
}

func (p *BundParser) String_term() (localctx IString_termContext) {
	localctx = NewString_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_string_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).value = _m
	}

	return localctx
}

// ICall_termContext is an interface to support dynamic dispatch.
type ICall_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsCall_termContext differentiates from other interfaces.
	IsCall_termContext()
}

type Call_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyCall_termContext() *Call_termContext {
	var p = new(Call_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_call_term
	return p
}

func (*Call_termContext) IsCall_termContext() {}

func NewCall_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_termContext {
	var p = new(Call_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_call_term

	return p
}

func (s *Call_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_termContext) GetValue() antlr.Token { return s.value }

func (s *Call_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Call_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Call_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCall_term(s)
	}
}

func (s *Call_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCall_term(s)
	}
}

func (p *BundParser) Call_term() (localctx ICall_termContext) {
	localctx = NewCall_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_call_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)

		var _m = p.Match(BundParserNAME)

		localctx.(*Call_termContext).value = _m
	}

	return localctx
}

// IBeginContext is an interface to support dynamic dispatch.
type IBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsBeginContext differentiates from other interfaces.
	IsBeginContext()
}

type BeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyBeginContext() *BeginContext {
	var p = new(BeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_begin
	return p
}

func (*BeginContext) IsBeginContext() {}

func NewBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginContext {
	var p = new(BeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_begin

	return p
}

func (s *BeginContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginContext) GetValue() antlr.Token { return s.value }

func (s *BeginContext) SetValue(v antlr.Token) { s.value = v }

func (s *BeginContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *BeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBegin(s)
	}
}

func (s *BeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBegin(s)
	}
}

func (p *BundParser) Begin() (localctx IBeginContext) {
	localctx = NewBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_begin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)

		var _m = p.Match(BundParserTOBEGIN)

		localctx.(*BeginContext).value = _m
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) GetValue() antlr.Token { return s.value }

func (s *EndContext) SetValue(v antlr.Token) { s.value = v }

func (s *EndContext) TOEND() antlr.TerminalNode {
	return s.GetToken(BundParserTOEND, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *BundParser) End() (localctx IEndContext) {
	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BundParserRULE_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)

		var _m = p.Match(BundParserTOEND)

		localctx.(*EndContext).value = _m
	}

	return localctx
}
