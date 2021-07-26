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

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterNs is called when entering the ns production.
	EnterNs(c *NsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDatablock is called when entering the datablock production.
	EnterDatablock(c *DatablockContext)

	// EnterFloatblock is called when entering the floatblock production.
	EnterFloatblock(c *FloatblockContext)

	// EnterIntblock is called when entering the intblock production.
	EnterIntblock(c *IntblockContext)

	// EnterUintblock is called when entering the uintblock production.
	EnterUintblock(c *UintblockContext)

	// EnterTrueblock is called when entering the trueblock production.
	EnterTrueblock(c *TrueblockContext)

	// EnterFalseblock is called when entering the falseblock production.
	EnterFalseblock(c *FalseblockContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterAlambda is called when entering the alambda production.
	EnterAlambda(c *AlambdaContext)

	// EnterLambda_cmd is called when entering the lambda_cmd production.
	EnterLambda_cmd(c *Lambda_cmdContext)

	// EnterTrue_term is called when entering the true_term production.
	EnterTrue_term(c *True_termContext)

	// EnterFalse_term is called when entering the false_term production.
	EnterFalse_term(c *False_termContext)

	// EnterString_term is called when entering the string_term production.
	EnterString_term(c *String_termContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterUinteger is called when entering the uinteger production.
	EnterUinteger(c *UintegerContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterUfloat is called when entering the ufloat production.
	EnterUfloat(c *UfloatContext)

	// EnterAllfloat is called when entering the allfloat production.
	EnterAllfloat(c *AllfloatContext)

	// EnterComplex_term is called when entering the complex_term production.
	EnterComplex_term(c *Complex_termContext)

	// EnterCall_term is called when entering the call_term production.
	EnterCall_term(c *Call_termContext)

	// EnterCall_sys is called when entering the call_sys production.
	EnterCall_sys(c *Call_sysContext)

	// EnterCmd_term is called when entering the cmd_term production.
	EnterCmd_term(c *Cmd_termContext)

	// EnterCmd_sys is called when entering the cmd_sys production.
	EnterCmd_sys(c *Cmd_sysContext)

	// EnterBegin is called when entering the begin production.
	EnterBegin(c *BeginContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterDrop is called when entering the drop production.
	EnterDrop(c *DropContext)

	// EnterDuplicate is called when entering the duplicate production.
	EnterDuplicate(c *DuplicateContext)

	// EnterExecute_term is called when entering the execute_term production.
	EnterExecute_term(c *Execute_termContext)

	// EnterReturn_term is called when entering the return_term production.
	EnterReturn_term(c *Return_termContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDatablock is called when exiting the datablock production.
	ExitDatablock(c *DatablockContext)

	// ExitFloatblock is called when exiting the floatblock production.
	ExitFloatblock(c *FloatblockContext)

	// ExitIntblock is called when exiting the intblock production.
	ExitIntblock(c *IntblockContext)

	// ExitUintblock is called when exiting the uintblock production.
	ExitUintblock(c *UintblockContext)

	// ExitTrueblock is called when exiting the trueblock production.
	ExitTrueblock(c *TrueblockContext)

	// ExitFalseblock is called when exiting the falseblock production.
	ExitFalseblock(c *FalseblockContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitAlambda is called when exiting the alambda production.
	ExitAlambda(c *AlambdaContext)

	// ExitLambda_cmd is called when exiting the lambda_cmd production.
	ExitLambda_cmd(c *Lambda_cmdContext)

	// ExitTrue_term is called when exiting the true_term production.
	ExitTrue_term(c *True_termContext)

	// ExitFalse_term is called when exiting the false_term production.
	ExitFalse_term(c *False_termContext)

	// ExitString_term is called when exiting the string_term production.
	ExitString_term(c *String_termContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitUinteger is called when exiting the uinteger production.
	ExitUinteger(c *UintegerContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitUfloat is called when exiting the ufloat production.
	ExitUfloat(c *UfloatContext)

	// ExitAllfloat is called when exiting the allfloat production.
	ExitAllfloat(c *AllfloatContext)

	// ExitComplex_term is called when exiting the complex_term production.
	ExitComplex_term(c *Complex_termContext)

	// ExitCall_term is called when exiting the call_term production.
	ExitCall_term(c *Call_termContext)

	// ExitCall_sys is called when exiting the call_sys production.
	ExitCall_sys(c *Call_sysContext)

	// ExitCmd_term is called when exiting the cmd_term production.
	ExitCmd_term(c *Cmd_termContext)

	// ExitCmd_sys is called when exiting the cmd_sys production.
	ExitCmd_sys(c *Cmd_sysContext)

	// ExitBegin is called when exiting the begin production.
	ExitBegin(c *BeginContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitDrop is called when exiting the drop production.
	ExitDrop(c *DropContext)

	// ExitDuplicate is called when exiting the duplicate production.
	ExitDuplicate(c *DuplicateContext)

	// ExitExecute_term is called when exiting the execute_term production.
	ExitExecute_term(c *Execute_termContext)

	// ExitReturn_term is called when exiting the return_term production.
	ExitReturn_term(c *Return_termContext)
}
