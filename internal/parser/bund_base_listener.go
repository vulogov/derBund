// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBundListener is a complete listener for a parse tree produced by BundParser.
type BaseBundListener struct{}

var _ BundListener = &BaseBundListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBundListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBundListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBundListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBundListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseBundListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseBundListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterRoot_term is called when production root_term is entered.
func (s *BaseBundListener) EnterRoot_term(ctx *Root_termContext) {}

// ExitRoot_term is called when production root_term is exited.
func (s *BaseBundListener) ExitRoot_term(ctx *Root_termContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBundListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBundListener) ExitTerm(ctx *TermContext) {}

// EnterNs is called when production ns is entered.
func (s *BaseBundListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BaseBundListener) ExitNs(ctx *NsContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBundListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBundListener) ExitBlock(ctx *BlockContext) {}

// EnterTrue_term is called when production true_term is entered.
func (s *BaseBundListener) EnterTrue_term(ctx *True_termContext) {}

// ExitTrue_term is called when production true_term is exited.
func (s *BaseBundListener) ExitTrue_term(ctx *True_termContext) {}

// EnterFalse_term is called when production false_term is entered.
func (s *BaseBundListener) EnterFalse_term(ctx *False_termContext) {}

// ExitFalse_term is called when production false_term is exited.
func (s *BaseBundListener) ExitFalse_term(ctx *False_termContext) {}

// EnterCall_term is called when production call_term is entered.
func (s *BaseBundListener) EnterCall_term(ctx *Call_termContext) {}

// ExitCall_term is called when production call_term is exited.
func (s *BaseBundListener) ExitCall_term(ctx *Call_termContext) {}

// EnterBegin is called when production begin is entered.
func (s *BaseBundListener) EnterBegin(ctx *BeginContext) {}

// ExitBegin is called when production begin is exited.
func (s *BaseBundListener) ExitBegin(ctx *BeginContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseBundListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseBundListener) ExitEnd(ctx *EndContext) {}
