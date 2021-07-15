// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 112, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 7, 6, 41, 10, 6, 12, 6, 14, 6, 44, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 51, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 57, 10,
	8, 3, 8, 5, 8, 60, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 65, 10, 9, 3, 10, 3,
	10, 5, 10, 69, 10, 10, 3, 10, 5, 10, 72, 10, 10, 3, 10, 3, 10, 5, 10, 76,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 82, 10, 11, 12, 11, 14, 11,
	85, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 91, 10, 12, 12, 12, 14,
	12, 94, 11, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	103, 10, 13, 12, 13, 14, 13, 106, 11, 13, 3, 14, 6, 14, 109, 10, 14, 13,
	14, 14, 14, 110, 3, 92, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 3, 2, 8, 4, 2, 67, 92, 99,
	124, 3, 2, 99, 124, 3, 2, 50, 59, 4, 2, 12, 12, 14, 15, 4, 2, 12, 12, 15,
	15, 4, 2, 11, 11, 34, 34, 2, 120, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 36, 3, 2, 2,
	2, 11, 38, 3, 2, 2, 2, 13, 50, 3, 2, 2, 2, 15, 59, 3, 2, 2, 2, 17, 64,
	3, 2, 2, 2, 19, 66, 3, 2, 2, 2, 21, 77, 3, 2, 2, 2, 23, 86, 3, 2, 2, 2,
	25, 98, 3, 2, 2, 2, 27, 108, 3, 2, 2, 2, 29, 30, 7, 93, 2, 2, 30, 4, 3,
	2, 2, 2, 31, 32, 7, 60, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 61, 2, 2, 34,
	35, 7, 61, 2, 2, 35, 8, 3, 2, 2, 2, 36, 37, 7, 49, 2, 2, 37, 10, 3, 2,
	2, 2, 38, 42, 5, 15, 8, 2, 39, 41, 5, 17, 9, 2, 40, 39, 3, 2, 2, 2, 41,
	44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 12, 3, 2, 2,
	2, 44, 42, 3, 2, 2, 2, 45, 51, 5, 27, 14, 2, 46, 51, 5, 21, 11, 2, 47,
	51, 5, 19, 10, 2, 48, 51, 5, 25, 13, 2, 49, 51, 5, 23, 12, 2, 50, 45, 3,
	2, 2, 2, 50, 46, 3, 2, 2, 2, 50, 47, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50,
	49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 8, 7, 2, 2, 53, 14, 3, 2, 2,
	2, 54, 57, 9, 2, 2, 2, 55, 57, 5, 9, 5, 2, 56, 54, 3, 2, 2, 2, 56, 55,
	3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 60, 9, 3, 2, 2, 59, 56, 3, 2, 2, 2,
	59, 58, 3, 2, 2, 2, 60, 16, 3, 2, 2, 2, 61, 65, 5, 15, 8, 2, 62, 65, 9,
	4, 2, 2, 63, 65, 5, 9, 5, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64,
	63, 3, 2, 2, 2, 65, 18, 3, 2, 2, 2, 66, 68, 7, 94, 2, 2, 67, 69, 5, 27,
	14, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 75, 3, 2, 2, 2, 70,
	72, 7, 15, 2, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3, 2,
	2, 2, 73, 76, 7, 12, 2, 2, 74, 76, 4, 14, 15, 2, 75, 71, 3, 2, 2, 2, 75,
	74, 3, 2, 2, 2, 76, 20, 3, 2, 2, 2, 77, 78, 7, 37, 2, 2, 78, 79, 7, 37,
	2, 2, 79, 83, 3, 2, 2, 2, 80, 82, 10, 5, 2, 2, 81, 80, 3, 2, 2, 2, 82,
	85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 22, 3, 2, 2,
	2, 85, 83, 3, 2, 2, 2, 86, 87, 7, 49, 2, 2, 87, 88, 7, 44, 2, 2, 88, 92,
	3, 2, 2, 2, 89, 91, 11, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2,
	92, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3,
	2, 2, 2, 95, 96, 7, 44, 2, 2, 96, 97, 7, 49, 2, 2, 97, 24, 3, 2, 2, 2,
	98, 99, 7, 49, 2, 2, 99, 100, 7, 49, 2, 2, 100, 104, 3, 2, 2, 2, 101, 103,
	10, 6, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 26, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	107, 109, 9, 7, 2, 2, 108, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 28, 3, 2, 2, 2, 15, 2, 42,
	50, 56, 59, 64, 68, 71, 75, 83, 92, 104, 110, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "':'", "';;'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "SLASH", "NAME", "SKIP_",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "SLASH", "NAME", "SKIP_", "ID_START", "ID_CONTINUE",
	"LINE_JOINING", "COMMENT", "BLOCK_COMMENT", "CCOMMENT", "SPACES",
}

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *BundLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	l := new(BundLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bund.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerT__0  = 1
	BundLexerT__1  = 2
	BundLexerT__2  = 3
	BundLexerSLASH = 4
	BundLexerNAME  = 5
	BundLexerSKIP_ = 6
)
