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

// EnterData is called when production data is entered.
func (s *BaseBundListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseBundListener) ExitData(ctx *DataContext) {}

// EnterNs is called when production ns is entered.
func (s *BaseBundListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BaseBundListener) ExitNs(ctx *NsContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBundListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBundListener) ExitBlock(ctx *BlockContext) {}

// EnterCblock is called when production cblock is entered.
func (s *BaseBundListener) EnterCblock(ctx *CblockContext) {}

// ExitCblock is called when production cblock is exited.
func (s *BaseBundListener) ExitCblock(ctx *CblockContext) {}

// EnterDatablock is called when production datablock is entered.
func (s *BaseBundListener) EnterDatablock(ctx *DatablockContext) {}

// ExitDatablock is called when production datablock is exited.
func (s *BaseBundListener) ExitDatablock(ctx *DatablockContext) {}

// EnterFloatblock is called when production floatblock is entered.
func (s *BaseBundListener) EnterFloatblock(ctx *FloatblockContext) {}

// ExitFloatblock is called when production floatblock is exited.
func (s *BaseBundListener) ExitFloatblock(ctx *FloatblockContext) {}

// EnterIntblock is called when production intblock is entered.
func (s *BaseBundListener) EnterIntblock(ctx *IntblockContext) {}

// ExitIntblock is called when production intblock is exited.
func (s *BaseBundListener) ExitIntblock(ctx *IntblockContext) {}

// EnterUintblock is called when production uintblock is entered.
func (s *BaseBundListener) EnterUintblock(ctx *UintblockContext) {}

// ExitUintblock is called when production uintblock is exited.
func (s *BaseBundListener) ExitUintblock(ctx *UintblockContext) {}

// EnterTrueblock is called when production trueblock is entered.
func (s *BaseBundListener) EnterTrueblock(ctx *TrueblockContext) {}

// ExitTrueblock is called when production trueblock is exited.
func (s *BaseBundListener) ExitTrueblock(ctx *TrueblockContext) {}

// EnterFalseblock is called when production falseblock is entered.
func (s *BaseBundListener) EnterFalseblock(ctx *FalseblockContext) {}

// ExitFalseblock is called when production falseblock is exited.
func (s *BaseBundListener) ExitFalseblock(ctx *FalseblockContext) {}

// EnterIgnoreblock is called when production ignoreblock is entered.
func (s *BaseBundListener) EnterIgnoreblock(ctx *IgnoreblockContext) {}

// ExitIgnoreblock is called when production ignoreblock is exited.
func (s *BaseBundListener) ExitIgnoreblock(ctx *IgnoreblockContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseBundListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseBundListener) ExitLambda(ctx *LambdaContext) {}

// EnterAlambda is called when production alambda is entered.
func (s *BaseBundListener) EnterAlambda(ctx *AlambdaContext) {}

// ExitAlambda is called when production alambda is exited.
func (s *BaseBundListener) ExitAlambda(ctx *AlambdaContext) {}

// EnterLambda_cmd is called when production lambda_cmd is entered.
func (s *BaseBundListener) EnterLambda_cmd(ctx *Lambda_cmdContext) {}

// ExitLambda_cmd is called when production lambda_cmd is exited.
func (s *BaseBundListener) ExitLambda_cmd(ctx *Lambda_cmdContext) {}

// EnterTrue_term is called when production true_term is entered.
func (s *BaseBundListener) EnterTrue_term(ctx *True_termContext) {}

// ExitTrue_term is called when production true_term is exited.
func (s *BaseBundListener) ExitTrue_term(ctx *True_termContext) {}

// EnterFalse_term is called when production false_term is entered.
func (s *BaseBundListener) EnterFalse_term(ctx *False_termContext) {}

// ExitFalse_term is called when production false_term is exited.
func (s *BaseBundListener) ExitFalse_term(ctx *False_termContext) {}

// EnterString_term is called when production string_term is entered.
func (s *BaseBundListener) EnterString_term(ctx *String_termContext) {}

// ExitString_term is called when production string_term is exited.
func (s *BaseBundListener) ExitString_term(ctx *String_termContext) {}

// EnterGlob_term is called when production glob_term is entered.
func (s *BaseBundListener) EnterGlob_term(ctx *Glob_termContext) {}

// ExitGlob_term is called when production glob_term is exited.
func (s *BaseBundListener) ExitGlob_term(ctx *Glob_termContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseBundListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseBundListener) ExitInteger(ctx *IntegerContext) {}

// EnterUinteger is called when production uinteger is entered.
func (s *BaseBundListener) EnterUinteger(ctx *UintegerContext) {}

// ExitUinteger is called when production uinteger is exited.
func (s *BaseBundListener) ExitUinteger(ctx *UintegerContext) {}

// EnterBinteger is called when production binteger is entered.
func (s *BaseBundListener) EnterBinteger(ctx *BintegerContext) {}

// ExitBinteger is called when production binteger is exited.
func (s *BaseBundListener) ExitBinteger(ctx *BintegerContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseBundListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseBundListener) ExitFloat(ctx *FloatContext) {}

// EnterUfloat is called when production ufloat is entered.
func (s *BaseBundListener) EnterUfloat(ctx *UfloatContext) {}

// ExitUfloat is called when production ufloat is exited.
func (s *BaseBundListener) ExitUfloat(ctx *UfloatContext) {}

// EnterAllfloat is called when production allfloat is entered.
func (s *BaseBundListener) EnterAllfloat(ctx *AllfloatContext) {}

// ExitAllfloat is called when production allfloat is exited.
func (s *BaseBundListener) ExitAllfloat(ctx *AllfloatContext) {}

// EnterComplex_term is called when production complex_term is entered.
func (s *BaseBundListener) EnterComplex_term(ctx *Complex_termContext) {}

// ExitComplex_term is called when production complex_term is exited.
func (s *BaseBundListener) ExitComplex_term(ctx *Complex_termContext) {}

// EnterCall_term is called when production call_term is entered.
func (s *BaseBundListener) EnterCall_term(ctx *Call_termContext) {}

// ExitCall_term is called when production call_term is exited.
func (s *BaseBundListener) ExitCall_term(ctx *Call_termContext) {}

// EnterCall_sys is called when production call_sys is entered.
func (s *BaseBundListener) EnterCall_sys(ctx *Call_sysContext) {}

// ExitCall_sys is called when production call_sys is exited.
func (s *BaseBundListener) ExitCall_sys(ctx *Call_sysContext) {}

// EnterCmd_term is called when production cmd_term is entered.
func (s *BaseBundListener) EnterCmd_term(ctx *Cmd_termContext) {}

// ExitCmd_term is called when production cmd_term is exited.
func (s *BaseBundListener) ExitCmd_term(ctx *Cmd_termContext) {}

// EnterCmd_sys is called when production cmd_sys is entered.
func (s *BaseBundListener) EnterCmd_sys(ctx *Cmd_sysContext) {}

// ExitCmd_sys is called when production cmd_sys is exited.
func (s *BaseBundListener) ExitCmd_sys(ctx *Cmd_sysContext) {}

// EnterRef_call is called when production ref_call is entered.
func (s *BaseBundListener) EnterRef_call(ctx *Ref_callContext) {}

// ExitRef_call is called when production ref_call is exited.
func (s *BaseBundListener) ExitRef_call(ctx *Ref_callContext) {}

// EnterRef_cmd is called when production ref_cmd is entered.
func (s *BaseBundListener) EnterRef_cmd(ctx *Ref_cmdContext) {}

// ExitRef_cmd is called when production ref_cmd is exited.
func (s *BaseBundListener) ExitRef_cmd(ctx *Ref_cmdContext) {}

// EnterBegin is called when production begin is entered.
func (s *BaseBundListener) EnterBegin(ctx *BeginContext) {}

// ExitBegin is called when production begin is exited.
func (s *BaseBundListener) ExitBegin(ctx *BeginContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseBundListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseBundListener) ExitEnd(ctx *EndContext) {}

// EnterDrop is called when production drop is entered.
func (s *BaseBundListener) EnterDrop(ctx *DropContext) {}

// ExitDrop is called when production drop is exited.
func (s *BaseBundListener) ExitDrop(ctx *DropContext) {}

// EnterDuplicate is called when production duplicate is entered.
func (s *BaseBundListener) EnterDuplicate(ctx *DuplicateContext) {}

// ExitDuplicate is called when production duplicate is exited.
func (s *BaseBundListener) ExitDuplicate(ctx *DuplicateContext) {}

// EnterExecute_term is called when production execute_term is entered.
func (s *BaseBundListener) EnterExecute_term(ctx *Execute_termContext) {}

// ExitExecute_term is called when production execute_term is exited.
func (s *BaseBundListener) ExitExecute_term(ctx *Execute_termContext) {}

// EnterSeparate_term is called when production separate_term is entered.
func (s *BaseBundListener) EnterSeparate_term(ctx *Separate_termContext) {}

// ExitSeparate_term is called when production separate_term is exited.
func (s *BaseBundListener) ExitSeparate_term(ctx *Separate_termContext) {}

// EnterReturn_term is called when production return_term is entered.
func (s *BaseBundListener) EnterReturn_term(ctx *Return_termContext) {}

// ExitReturn_term is called when production return_term is exited.
func (s *BaseBundListener) ExitReturn_term(ctx *Return_termContext) {}
