// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BundListener is a complete listener for a parse tree produced by BundParser.
type BundListener interface {
	antlr.ParseTreeListener

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterRoot_term is called when entering the root_term production.
	EnterRoot_term(c *Root_termContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterNs is called when entering the ns production.
	EnterNs(c *NsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterTrue_term is called when entering the true_term production.
	EnterTrue_term(c *True_termContext)

	// EnterFalse_term is called when entering the false_term production.
	EnterFalse_term(c *False_termContext)

	// EnterCall_term is called when entering the call_term production.
	EnterCall_term(c *Call_termContext)

	// EnterBegin is called when entering the begin production.
	EnterBegin(c *BeginContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitTrue_term is called when exiting the true_term production.
	ExitTrue_term(c *True_termContext)

	// ExitFalse_term is called when exiting the false_term production.
	ExitFalse_term(c *False_termContext)

	// ExitCall_term is called when exiting the call_term production.
	ExitCall_term(c *Call_termContext)

	// ExitBegin is called when exiting the begin production.
	ExitBegin(c *BeginContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)
}
