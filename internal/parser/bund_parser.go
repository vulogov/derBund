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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 244,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 7, 2, 64, 10, 2, 12, 2, 14,
	2, 67, 11, 2, 3, 3, 3, 3, 5, 3, 71, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 97, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 108, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 7, 6, 114, 10, 6, 12, 6, 14, 6, 117, 11, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 7, 7, 123, 10, 7, 12, 7, 14, 7, 126, 11, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 7, 8, 132, 10, 8, 12, 8, 14, 8, 135, 11, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 7, 9, 141, 10, 9, 12, 9, 14, 9, 144, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	7, 10, 150, 10, 10, 12, 10, 14, 10, 153, 11, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 7, 11, 159, 10, 11, 12, 11, 14, 11, 162, 11, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 7, 12, 168, 10, 12, 12, 12, 14, 12, 171, 11, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 7, 13, 177, 10, 13, 12, 13, 14, 13, 180, 11, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 188, 10, 14, 12, 14, 14, 14, 191, 11,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 199, 10, 15, 12, 15,
	14, 15, 202, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 2, 2, 32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 2,
	3, 4, 2, 30, 30, 32, 32, 2, 256, 2, 65, 3, 2, 2, 2, 4, 70, 3, 2, 2, 2,
	6, 96, 3, 2, 2, 2, 8, 107, 3, 2, 2, 2, 10, 109, 3, 2, 2, 2, 12, 120, 3,
	2, 2, 2, 14, 129, 3, 2, 2, 2, 16, 138, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2,
	20, 156, 3, 2, 2, 2, 22, 165, 3, 2, 2, 2, 24, 174, 3, 2, 2, 2, 26, 183,
	3, 2, 2, 2, 28, 194, 3, 2, 2, 2, 30, 205, 3, 2, 2, 2, 32, 207, 3, 2, 2,
	2, 34, 209, 3, 2, 2, 2, 36, 211, 3, 2, 2, 2, 38, 213, 3, 2, 2, 2, 40, 215,
	3, 2, 2, 2, 42, 217, 3, 2, 2, 2, 44, 219, 3, 2, 2, 2, 46, 224, 3, 2, 2,
	2, 48, 226, 3, 2, 2, 2, 50, 231, 3, 2, 2, 2, 52, 233, 3, 2, 2, 2, 54, 235,
	3, 2, 2, 2, 56, 237, 3, 2, 2, 2, 58, 239, 3, 2, 2, 2, 60, 241, 3, 2, 2,
	2, 62, 64, 5, 4, 3, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 3, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	68, 71, 5, 10, 6, 2, 69, 71, 5, 12, 7, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3,
	2, 2, 2, 71, 5, 3, 2, 2, 2, 72, 97, 5, 10, 6, 2, 73, 97, 5, 12, 7, 2, 74,
	97, 5, 14, 8, 2, 75, 97, 5, 16, 9, 2, 76, 97, 5, 18, 10, 2, 77, 97, 5,
	20, 11, 2, 78, 97, 5, 22, 12, 2, 79, 97, 5, 24, 13, 2, 80, 97, 5, 30, 16,
	2, 81, 97, 5, 32, 17, 2, 82, 97, 5, 34, 18, 2, 83, 97, 5, 36, 19, 2, 84,
	97, 5, 38, 20, 2, 85, 97, 5, 40, 21, 2, 86, 97, 5, 42, 22, 2, 87, 97, 5,
	44, 23, 2, 88, 97, 5, 46, 24, 2, 89, 97, 5, 48, 25, 2, 90, 97, 5, 50, 26,
	2, 91, 97, 5, 52, 27, 2, 92, 97, 5, 58, 30, 2, 93, 97, 5, 60, 31, 2, 94,
	97, 5, 54, 28, 2, 95, 97, 5, 56, 29, 2, 96, 72, 3, 2, 2, 2, 96, 73, 3,
	2, 2, 2, 96, 74, 3, 2, 2, 2, 96, 75, 3, 2, 2, 2, 96, 76, 3, 2, 2, 2, 96,
	77, 3, 2, 2, 2, 96, 78, 3, 2, 2, 2, 96, 79, 3, 2, 2, 2, 96, 80, 3, 2, 2,
	2, 96, 81, 3, 2, 2, 2, 96, 82, 3, 2, 2, 2, 96, 83, 3, 2, 2, 2, 96, 84,
	3, 2, 2, 2, 96, 85, 3, 2, 2, 2, 96, 86, 3, 2, 2, 2, 96, 87, 3, 2, 2, 2,
	96, 88, 3, 2, 2, 2, 96, 89, 3, 2, 2, 2, 96, 90, 3, 2, 2, 2, 96, 91, 3,
	2, 2, 2, 96, 92, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96,
	95, 3, 2, 2, 2, 97, 7, 3, 2, 2, 2, 98, 108, 5, 30, 16, 2, 99, 108, 5, 32,
	17, 2, 100, 108, 5, 34, 18, 2, 101, 108, 5, 36, 19, 2, 102, 108, 5, 40,
	21, 2, 103, 108, 5, 42, 22, 2, 104, 108, 5, 44, 23, 2, 105, 108, 5, 46,
	24, 2, 106, 108, 5, 48, 25, 2, 107, 98, 3, 2, 2, 2, 107, 99, 3, 2, 2, 2,
	107, 100, 3, 2, 2, 2, 107, 101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 107,
	103, 3, 2, 2, 2, 107, 104, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106,
	3, 2, 2, 2, 108, 9, 3, 2, 2, 2, 109, 110, 7, 3, 2, 2, 110, 111, 7, 30,
	2, 2, 111, 115, 7, 25, 2, 2, 112, 114, 5, 6, 4, 2, 113, 112, 3, 2, 2, 2,
	114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116,
	118, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 7, 4, 2, 2, 119, 11, 3,
	2, 2, 2, 120, 124, 7, 5, 2, 2, 121, 123, 5, 6, 4, 2, 122, 121, 3, 2, 2,
	2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125,
	127, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7, 6, 2, 2, 128, 13, 3,
	2, 2, 2, 129, 133, 7, 7, 2, 2, 130, 132, 5, 8, 5, 2, 131, 130, 3, 2, 2,
	2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134,
	136, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 137, 7, 6, 2, 2, 137, 15, 3,
	2, 2, 2, 138, 142, 7, 8, 2, 2, 139, 141, 5, 40, 21, 2, 140, 139, 3, 2,
	2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2,
	143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 6, 2, 2, 146,
	17, 3, 2, 2, 2, 147, 151, 7, 9, 2, 2, 148, 150, 5, 36, 19, 2, 149, 148,
	3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2,
	2, 2, 152, 154, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 7, 6, 2, 2,
	155, 19, 3, 2, 2, 2, 156, 160, 7, 10, 2, 2, 157, 159, 5, 38, 20, 2, 158,
	157, 3, 2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 163, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 6,
	2, 2, 164, 21, 3, 2, 2, 2, 165, 169, 7, 11, 2, 2, 166, 168, 5, 6, 4, 2,
	167, 166, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169,
	170, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173,
	7, 6, 2, 2, 173, 23, 3, 2, 2, 2, 174, 178, 7, 12, 2, 2, 175, 177, 5, 6,
	4, 2, 176, 175, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181,
	182, 7, 6, 2, 2, 182, 25, 3, 2, 2, 2, 183, 184, 7, 3, 2, 2, 184, 185, 7,
	30, 2, 2, 185, 189, 7, 13, 2, 2, 186, 188, 5, 6, 4, 2, 187, 186, 3, 2,
	2, 2, 188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2,
	190, 192, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 192, 193, 7, 14, 2, 2, 193,
	27, 3, 2, 2, 2, 194, 195, 7, 15, 2, 2, 195, 196, 7, 31, 2, 2, 196, 200,
	7, 16, 2, 2, 197, 199, 5, 6, 4, 2, 198, 197, 3, 2, 2, 2, 199, 202, 3, 2,
	2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 3, 2, 2, 2,
	202, 200, 3, 2, 2, 2, 203, 204, 7, 14, 2, 2, 204, 29, 3, 2, 2, 2, 205,
	206, 7, 18, 2, 2, 206, 31, 3, 2, 2, 2, 207, 208, 7, 19, 2, 2, 208, 33,
	3, 2, 2, 2, 209, 210, 7, 24, 2, 2, 210, 35, 3, 2, 2, 2, 211, 212, 7, 20,
	2, 2, 212, 37, 3, 2, 2, 2, 213, 214, 7, 21, 2, 2, 214, 39, 3, 2, 2, 2,
	215, 216, 7, 23, 2, 2, 216, 41, 3, 2, 2, 2, 217, 218, 7, 30, 2, 2, 218,
	43, 3, 2, 2, 2, 219, 220, 7, 30, 2, 2, 220, 221, 7, 17, 2, 2, 221, 222,
	9, 2, 2, 2, 222, 223, 7, 6, 2, 2, 223, 45, 3, 2, 2, 2, 224, 225, 7, 31,
	2, 2, 225, 47, 3, 2, 2, 2, 226, 227, 7, 31, 2, 2, 227, 228, 7, 17, 2, 2,
	228, 229, 9, 2, 2, 2, 229, 230, 7, 6, 2, 2, 230, 49, 3, 2, 2, 2, 231, 232,
	7, 25, 2, 2, 232, 51, 3, 2, 2, 2, 233, 234, 7, 26, 2, 2, 234, 53, 3, 2,
	2, 2, 235, 236, 7, 28, 2, 2, 236, 55, 3, 2, 2, 2, 237, 238, 7, 29, 2, 2,
	238, 57, 3, 2, 2, 2, 239, 240, 7, 33, 2, 2, 240, 59, 3, 2, 2, 2, 241, 242,
	7, 34, 2, 2, 242, 61, 3, 2, 2, 2, 16, 65, 70, 96, 107, 115, 124, 133, 142,
	151, 160, 169, 178, 189, 200,
}
var literalNames = []string{
	"", "'['", "';;'", "'('", "')'", "'(*'", "'(float'", "'(int'", "'(uint'",
	"'(true'", "'(false'", "']'", "'.'", "'[['", "']]'", "'.('", "", "", "",
	"", "", "", "", "':'", "';'", "'/'", "','", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE",
	"FALSE", "INTEGER", "UINTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING",
	"TOBEGIN", "TOEND", "SLASH", "DROP", "DUPLICATE", "NAME", "CMD", "SYS",
	"EXECUTE", "RETURN", "COMMENT", "BLOCK_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"expressions", "root_term", "term", "data", "ns", "block", "datablock",
	"floatblock", "intblock", "uintblock", "trueblock", "falseblock", "lambda",
	"lambda_cmd", "true_term", "false_term", "string_term", "integer", "uinteger",
	"float", "call_term", "call_sys", "cmd_term", "cmd_sys", "begin", "end",
	"drop", "duplicate", "execute_term", "return_term",
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
	BundParserEOF             = antlr.TokenEOF
	BundParserT__0            = 1
	BundParserT__1            = 2
	BundParserT__2            = 3
	BundParserT__3            = 4
	BundParserT__4            = 5
	BundParserT__5            = 6
	BundParserT__6            = 7
	BundParserT__7            = 8
	BundParserT__8            = 9
	BundParserT__9            = 10
	BundParserT__10           = 11
	BundParserT__11           = 12
	BundParserT__12           = 13
	BundParserT__13           = 14
	BundParserT__14           = 15
	BundParserTRUE            = 16
	BundParserFALSE           = 17
	BundParserINTEGER         = 18
	BundParserUINTEGER        = 19
	BundParserDECIMAL_INTEGER = 20
	BundParserFLOAT_NUMBER    = 21
	BundParserSTRING          = 22
	BundParserTOBEGIN         = 23
	BundParserTOEND           = 24
	BundParserSLASH           = 25
	BundParserDROP            = 26
	BundParserDUPLICATE       = 27
	BundParserNAME            = 28
	BundParserCMD             = 29
	BundParserSYS             = 30
	BundParserEXECUTE         = 31
	BundParserRETURN          = 32
	BundParserCOMMENT         = 33
	BundParserBLOCK_COMMENT   = 34
	BundParserWS              = 35
	BundParserSHEBANG         = 36
)

// BundParser rules.
const (
	BundParserRULE_expressions  = 0
	BundParserRULE_root_term    = 1
	BundParserRULE_term         = 2
	BundParserRULE_data         = 3
	BundParserRULE_ns           = 4
	BundParserRULE_block        = 5
	BundParserRULE_datablock    = 6
	BundParserRULE_floatblock   = 7
	BundParserRULE_intblock     = 8
	BundParserRULE_uintblock    = 9
	BundParserRULE_trueblock    = 10
	BundParserRULE_falseblock   = 11
	BundParserRULE_lambda       = 12
	BundParserRULE_lambda_cmd   = 13
	BundParserRULE_true_term    = 14
	BundParserRULE_false_term   = 15
	BundParserRULE_string_term  = 16
	BundParserRULE_integer      = 17
	BundParserRULE_uinteger     = 18
	BundParserRULE_float        = 19
	BundParserRULE_call_term    = 20
	BundParserRULE_call_sys     = 21
	BundParserRULE_cmd_term     = 22
	BundParserRULE_cmd_sys      = 23
	BundParserRULE_begin        = 24
	BundParserRULE_end          = 25
	BundParserRULE_drop         = 26
	BundParserRULE_duplicate    = 27
	BundParserRULE_execute_term = 28
	BundParserRULE_return_term  = 29
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__2 {
		{
			p.SetState(60)
			p.Root_term()
		}

		p.SetState(65)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(66)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(67)
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

func (s *TermContext) Datablock() IDatablockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatablockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatablockContext)
}

func (s *TermContext) Floatblock() IFloatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatblockContext)
}

func (s *TermContext) Intblock() IIntblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntblockContext)
}

func (s *TermContext) Uintblock() IUintblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintblockContext)
}

func (s *TermContext) Trueblock() ITrueblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrueblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrueblockContext)
}

func (s *TermContext) Falseblock() IFalseblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalseblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalseblockContext)
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

func (s *TermContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Uinteger() IUintegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintegerContext)
}

func (s *TermContext) Float() IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *TermContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *TermContext) Call_sys() ICall_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_sysContext)
}

func (s *TermContext) Cmd_term() ICmd_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_termContext)
}

func (s *TermContext) Cmd_sys() ICmd_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_sysContext)
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

func (s *TermContext) Execute_term() IExecute_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecute_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecute_termContext)
}

func (s *TermContext) Return_term() IReturn_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_termContext)
}

func (s *TermContext) Drop() IDropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDropContext)
}

func (s *TermContext) Duplicate() IDuplicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDuplicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDuplicateContext)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(70)
			p.Ns()
		}

	case 2:
		{
			p.SetState(71)
			p.Block()
		}

	case 3:
		{
			p.SetState(72)
			p.Datablock()
		}

	case 4:
		{
			p.SetState(73)
			p.Floatblock()
		}

	case 5:
		{
			p.SetState(74)
			p.Intblock()
		}

	case 6:
		{
			p.SetState(75)
			p.Uintblock()
		}

	case 7:
		{
			p.SetState(76)
			p.Trueblock()
		}

	case 8:
		{
			p.SetState(77)
			p.Falseblock()
		}

	case 9:
		{
			p.SetState(78)
			p.True_term()
		}

	case 10:
		{
			p.SetState(79)
			p.False_term()
		}

	case 11:
		{
			p.SetState(80)
			p.String_term()
		}

	case 12:
		{
			p.SetState(81)
			p.Integer()
		}

	case 13:
		{
			p.SetState(82)
			p.Uinteger()
		}

	case 14:
		{
			p.SetState(83)
			p.Float()
		}

	case 15:
		{
			p.SetState(84)
			p.Call_term()
		}

	case 16:
		{
			p.SetState(85)
			p.Call_sys()
		}

	case 17:
		{
			p.SetState(86)
			p.Cmd_term()
		}

	case 18:
		{
			p.SetState(87)
			p.Cmd_sys()
		}

	case 19:
		{
			p.SetState(88)
			p.Begin()
		}

	case 20:
		{
			p.SetState(89)
			p.End()
		}

	case 21:
		{
			p.SetState(90)
			p.Execute_term()
		}

	case 22:
		{
			p.SetState(91)
			p.Return_term()
		}

	case 23:
		{
			p.SetState(92)
			p.Drop()
		}

	case 24:
		{
			p.SetState(93)
			p.Duplicate()
		}

	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *DataContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *DataContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *DataContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *DataContext) Float() IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *DataContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *DataContext) Call_sys() ICall_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_sysContext)
}

func (s *DataContext) Cmd_term() ICmd_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_termContext)
}

func (s *DataContext) Cmd_sys() ICmd_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_sysContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *BundParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_data)

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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(96)
			p.True_term()
		}

	case 2:
		{
			p.SetState(97)
			p.False_term()
		}

	case 3:
		{
			p.SetState(98)
			p.String_term()
		}

	case 4:
		{
			p.SetState(99)
			p.Integer()
		}

	case 5:
		{
			p.SetState(100)
			p.Float()
		}

	case 6:
		{
			p.SetState(101)
			p.Call_term()
		}

	case 7:
		{
			p.SetState(102)
			p.Call_sys()
		}

	case 8:
		{
			p.SetState(103)
			p.Cmd_term()
		}

	case 9:
		{
			p.SetState(104)
			p.Cmd_sys()
		}

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
	p.EnterRule(localctx, 8, BundParserRULE_ns)
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
		p.SetState(107)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(108)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(109)
		p.Match(BundParserTOBEGIN)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(110)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
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
	p.EnterRule(localctx, 10, BundParserRULE_block)
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
		p.SetState(118)
		p.Match(BundParserT__2)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(119)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IDatablockContext is an interface to support dynamic dispatch.
type IDatablockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetBody returns the body rule context list.
	GetBody() []IDataContext

	// SetBody sets the body rule context list.
	SetBody([]IDataContext)

	// IsDatablockContext differentiates from other interfaces.
	IsDatablockContext()
}

type DatablockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_data  IDataContext
	body   []IDataContext
}

func NewEmptyDatablockContext() *DatablockContext {
	var p = new(DatablockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_datablock
	return p
}

func (*DatablockContext) IsDatablockContext() {}

func NewDatablockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatablockContext {
	var p = new(DatablockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_datablock

	return p
}

func (s *DatablockContext) GetParser() antlr.Parser { return s.parser }

func (s *DatablockContext) Get_data() IDataContext { return s._data }

func (s *DatablockContext) Set_data(v IDataContext) { s._data = v }

func (s *DatablockContext) GetBody() []IDataContext { return s.body }

func (s *DatablockContext) SetBody(v []IDataContext) { s.body = v }

func (s *DatablockContext) AllData() []IDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataContext)(nil)).Elem())
	var tst = make([]IDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataContext)
		}
	}

	return tst
}

func (s *DatablockContext) Data(i int) IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *DatablockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatablockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatablockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDatablock(s)
	}
}

func (s *DatablockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDatablock(s)
	}
}

func (p *BundParser) Datablock() (localctx IDatablockContext) {
	localctx = NewDatablockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_datablock)
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
		p.SetState(127)
		p.Match(BundParserT__4)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserNAME)|(1<<BundParserCMD))) != 0 {
		{
			p.SetState(128)

			var _x = p.Data()

			localctx.(*DatablockContext)._data = _x
		}
		localctx.(*DatablockContext).body = append(localctx.(*DatablockContext).body, localctx.(*DatablockContext)._data)

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IFloatblockContext is an interface to support dynamic dispatch.
type IFloatblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_float returns the _float rule contexts.
	Get_float() IFloatContext

	// Set_float sets the _float rule contexts.
	Set_float(IFloatContext)

	// GetBody returns the body rule context list.
	GetBody() []IFloatContext

	// SetBody sets the body rule context list.
	SetBody([]IFloatContext)

	// IsFloatblockContext differentiates from other interfaces.
	IsFloatblockContext()
}

type FloatblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_float IFloatContext
	body   []IFloatContext
}

func NewEmptyFloatblockContext() *FloatblockContext {
	var p = new(FloatblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_floatblock
	return p
}

func (*FloatblockContext) IsFloatblockContext() {}

func NewFloatblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatblockContext {
	var p = new(FloatblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_floatblock

	return p
}

func (s *FloatblockContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatblockContext) Get_float() IFloatContext { return s._float }

func (s *FloatblockContext) Set_float(v IFloatContext) { s._float = v }

func (s *FloatblockContext) GetBody() []IFloatContext { return s.body }

func (s *FloatblockContext) SetBody(v []IFloatContext) { s.body = v }

func (s *FloatblockContext) AllFloat() []IFloatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFloatContext)(nil)).Elem())
	var tst = make([]IFloatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFloatContext)
		}
	}

	return tst
}

func (s *FloatblockContext) Float(i int) IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *FloatblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloatblock(s)
	}
}

func (s *FloatblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloatblock(s)
	}
}

func (p *BundParser) Floatblock() (localctx IFloatblockContext) {
	localctx = NewFloatblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_floatblock)
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
		p.SetState(136)
		p.Match(BundParserT__5)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserFLOAT_NUMBER {
		{
			p.SetState(137)

			var _x = p.Float()

			localctx.(*FloatblockContext)._float = _x
		}
		localctx.(*FloatblockContext).body = append(localctx.(*FloatblockContext).body, localctx.(*FloatblockContext)._float)

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IIntblockContext is an interface to support dynamic dispatch.
type IIntblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_integer returns the _integer rule contexts.
	Get_integer() IIntegerContext

	// Set_integer sets the _integer rule contexts.
	Set_integer(IIntegerContext)

	// GetBody returns the body rule context list.
	GetBody() []IIntegerContext

	// SetBody sets the body rule context list.
	SetBody([]IIntegerContext)

	// IsIntblockContext differentiates from other interfaces.
	IsIntblockContext()
}

type IntblockContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_integer IIntegerContext
	body     []IIntegerContext
}

func NewEmptyIntblockContext() *IntblockContext {
	var p = new(IntblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_intblock
	return p
}

func (*IntblockContext) IsIntblockContext() {}

func NewIntblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntblockContext {
	var p = new(IntblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_intblock

	return p
}

func (s *IntblockContext) GetParser() antlr.Parser { return s.parser }

func (s *IntblockContext) Get_integer() IIntegerContext { return s._integer }

func (s *IntblockContext) Set_integer(v IIntegerContext) { s._integer = v }

func (s *IntblockContext) GetBody() []IIntegerContext { return s.body }

func (s *IntblockContext) SetBody(v []IIntegerContext) { s.body = v }

func (s *IntblockContext) AllInteger() []IIntegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerContext)(nil)).Elem())
	var tst = make([]IIntegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerContext)
		}
	}

	return tst
}

func (s *IntblockContext) Integer(i int) IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *IntblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterIntblock(s)
	}
}

func (s *IntblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitIntblock(s)
	}
}

func (p *BundParser) Intblock() (localctx IIntblockContext) {
	localctx = NewIntblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_intblock)
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
		p.SetState(145)
		p.Match(BundParserT__6)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserINTEGER {
		{
			p.SetState(146)

			var _x = p.Integer()

			localctx.(*IntblockContext)._integer = _x
		}
		localctx.(*IntblockContext).body = append(localctx.(*IntblockContext).body, localctx.(*IntblockContext)._integer)

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IUintblockContext is an interface to support dynamic dispatch.
type IUintblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_uinteger returns the _uinteger rule contexts.
	Get_uinteger() IUintegerContext

	// Set_uinteger sets the _uinteger rule contexts.
	Set_uinteger(IUintegerContext)

	// GetBody returns the body rule context list.
	GetBody() []IUintegerContext

	// SetBody sets the body rule context list.
	SetBody([]IUintegerContext)

	// IsUintblockContext differentiates from other interfaces.
	IsUintblockContext()
}

type UintblockContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_uinteger IUintegerContext
	body      []IUintegerContext
}

func NewEmptyUintblockContext() *UintblockContext {
	var p = new(UintblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_uintblock
	return p
}

func (*UintblockContext) IsUintblockContext() {}

func NewUintblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintblockContext {
	var p = new(UintblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_uintblock

	return p
}

func (s *UintblockContext) GetParser() antlr.Parser { return s.parser }

func (s *UintblockContext) Get_uinteger() IUintegerContext { return s._uinteger }

func (s *UintblockContext) Set_uinteger(v IUintegerContext) { s._uinteger = v }

func (s *UintblockContext) GetBody() []IUintegerContext { return s.body }

func (s *UintblockContext) SetBody(v []IUintegerContext) { s.body = v }

func (s *UintblockContext) AllUinteger() []IUintegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUintegerContext)(nil)).Elem())
	var tst = make([]IUintegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUintegerContext)
		}
	}

	return tst
}

func (s *UintblockContext) Uinteger(i int) IUintegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUintegerContext)
}

func (s *UintblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterUintblock(s)
	}
}

func (s *UintblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitUintblock(s)
	}
}

func (p *BundParser) Uintblock() (localctx IUintblockContext) {
	localctx = NewUintblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_uintblock)
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
		p.SetState(154)
		p.Match(BundParserT__7)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserUINTEGER {
		{
			p.SetState(155)

			var _x = p.Uinteger()

			localctx.(*UintblockContext)._uinteger = _x
		}
		localctx.(*UintblockContext).body = append(localctx.(*UintblockContext).body, localctx.(*UintblockContext)._uinteger)

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ITrueblockContext is an interface to support dynamic dispatch.
type ITrueblockContext interface {
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

	// IsTrueblockContext differentiates from other interfaces.
	IsTrueblockContext()
}

type TrueblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyTrueblockContext() *TrueblockContext {
	var p = new(TrueblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_trueblock
	return p
}

func (*TrueblockContext) IsTrueblockContext() {}

func NewTrueblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueblockContext {
	var p = new(TrueblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_trueblock

	return p
}

func (s *TrueblockContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueblockContext) Get_term() ITermContext { return s._term }

func (s *TrueblockContext) Set_term(v ITermContext) { s._term = v }

func (s *TrueblockContext) GetBody() []ITermContext { return s.body }

func (s *TrueblockContext) SetBody(v []ITermContext) { s.body = v }

func (s *TrueblockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TrueblockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TrueblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTrueblock(s)
	}
}

func (s *TrueblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTrueblock(s)
	}
}

func (p *BundParser) Trueblock() (localctx ITrueblockContext) {
	localctx = NewTrueblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BundParserRULE_trueblock)
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
		p.SetState(163)
		p.Match(BundParserT__8)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(164)

			var _x = p.Term()

			localctx.(*TrueblockContext)._term = _x
		}
		localctx.(*TrueblockContext).body = append(localctx.(*TrueblockContext).body, localctx.(*TrueblockContext)._term)

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IFalseblockContext is an interface to support dynamic dispatch.
type IFalseblockContext interface {
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

	// IsFalseblockContext differentiates from other interfaces.
	IsFalseblockContext()
}

type FalseblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyFalseblockContext() *FalseblockContext {
	var p = new(FalseblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_falseblock
	return p
}

func (*FalseblockContext) IsFalseblockContext() {}

func NewFalseblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FalseblockContext {
	var p = new(FalseblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_falseblock

	return p
}

func (s *FalseblockContext) GetParser() antlr.Parser { return s.parser }

func (s *FalseblockContext) Get_term() ITermContext { return s._term }

func (s *FalseblockContext) Set_term(v ITermContext) { s._term = v }

func (s *FalseblockContext) GetBody() []ITermContext { return s.body }

func (s *FalseblockContext) SetBody(v []ITermContext) { s.body = v }

func (s *FalseblockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FalseblockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FalseblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FalseblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFalseblock(s)
	}
}

func (s *FalseblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFalseblock(s)
	}
}

func (p *BundParser) Falseblock() (localctx IFalseblockContext) {
	localctx = NewFalseblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BundParserRULE_falseblock)
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
		p.SetState(172)
		p.Match(BundParserT__9)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(173)

			var _x = p.Term()

			localctx.(*FalseblockContext)._term = _x
		}
		localctx.(*FalseblockContext).body = append(localctx.(*FalseblockContext).body, localctx.(*FalseblockContext)._term)

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
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

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) GetName() antlr.Token { return s.name }

func (s *LambdaContext) SetName(v antlr.Token) { s.name = v }

func (s *LambdaContext) Get_term() ITermContext { return s._term }

func (s *LambdaContext) Set_term(v ITermContext) { s._term = v }

func (s *LambdaContext) GetBody() []ITermContext { return s.body }

func (s *LambdaContext) SetBody(v []ITermContext) { s.body = v }

func (s *LambdaContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *LambdaContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *LambdaContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *BundParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BundParserRULE_lambda)
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
		p.SetState(181)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(182)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(183)
		p.Match(BundParserT__10)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(184)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(BundParserT__11)
	}

	return localctx
}

// ILambda_cmdContext is an interface to support dynamic dispatch.
type ILambda_cmdContext interface {
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

	// IsLambda_cmdContext differentiates from other interfaces.
	IsLambda_cmdContext()
}

type Lambda_cmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambda_cmdContext() *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda_cmd
	return p
}

func (*Lambda_cmdContext) IsLambda_cmdContext() {}

func NewLambda_cmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda_cmd

	return p
}

func (s *Lambda_cmdContext) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_cmdContext) GetName() antlr.Token { return s.name }

func (s *Lambda_cmdContext) SetName(v antlr.Token) { s.name = v }

func (s *Lambda_cmdContext) Get_term() ITermContext { return s._term }

func (s *Lambda_cmdContext) Set_term(v ITermContext) { s._term = v }

func (s *Lambda_cmdContext) GetBody() []ITermContext { return s.body }

func (s *Lambda_cmdContext) SetBody(v []ITermContext) { s.body = v }

func (s *Lambda_cmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Lambda_cmdContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Lambda_cmdContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Lambda_cmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_cmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_cmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda_cmd(s)
	}
}

func (s *Lambda_cmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda_cmd(s)
	}
}

func (p *BundParser) Lambda_cmd() (localctx ILambda_cmdContext) {
	localctx = NewLambda_cmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BundParserRULE_lambda_cmd)
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
		p.SetState(192)
		p.Match(BundParserT__12)
	}
	{
		p.SetState(193)

		var _m = p.Match(BundParserCMD)

		localctx.(*Lambda_cmdContext).name = _m
	}
	{
		p.SetState(194)
		p.Match(BundParserT__13)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(BundParserT__0-1))|(1<<(BundParserT__2-1))|(1<<(BundParserT__4-1))|(1<<(BundParserT__5-1))|(1<<(BundParserT__6-1))|(1<<(BundParserT__7-1))|(1<<(BundParserT__8-1))|(1<<(BundParserT__9-1))|(1<<(BundParserTRUE-1))|(1<<(BundParserFALSE-1))|(1<<(BundParserINTEGER-1))|(1<<(BundParserUINTEGER-1))|(1<<(BundParserFLOAT_NUMBER-1))|(1<<(BundParserSTRING-1))|(1<<(BundParserTOBEGIN-1))|(1<<(BundParserTOEND-1))|(1<<(BundParserDROP-1))|(1<<(BundParserDUPLICATE-1))|(1<<(BundParserNAME-1))|(1<<(BundParserCMD-1))|(1<<(BundParserEXECUTE-1))|(1<<(BundParserRETURN-1)))) != 0 {
		{
			p.SetState(195)

			var _x = p.Term()

			localctx.(*Lambda_cmdContext)._term = _x
		}
		localctx.(*Lambda_cmdContext).body = append(localctx.(*Lambda_cmdContext).body, localctx.(*Lambda_cmdContext)._term)

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(BundParserT__11)
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
	p.EnterRule(localctx, 28, BundParserRULE_true_term)

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
		p.SetState(203)

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
	p.EnterRule(localctx, 30, BundParserRULE_false_term)

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
		p.SetState(205)

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
	p.EnterRule(localctx, 32, BundParserRULE_string_term)

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
		p.SetState(207)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).value = _m
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) GetValue() antlr.Token { return s.value }

func (s *IntegerContext) SetValue(v antlr.Token) { s.value = v }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserINTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *BundParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BundParserRULE_integer)

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
		p.SetState(209)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*IntegerContext).value = _m
	}

	return localctx
}

// IUintegerContext is an interface to support dynamic dispatch.
type IUintegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsUintegerContext differentiates from other interfaces.
	IsUintegerContext()
}

type UintegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyUintegerContext() *UintegerContext {
	var p = new(UintegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_uinteger
	return p
}

func (*UintegerContext) IsUintegerContext() {}

func NewUintegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintegerContext {
	var p = new(UintegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_uinteger

	return p
}

func (s *UintegerContext) GetParser() antlr.Parser { return s.parser }

func (s *UintegerContext) GetValue() antlr.Token { return s.value }

func (s *UintegerContext) SetValue(v antlr.Token) { s.value = v }

func (s *UintegerContext) UINTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserUINTEGER, 0)
}

func (s *UintegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterUinteger(s)
	}
}

func (s *UintegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitUinteger(s)
	}
}

func (p *BundParser) Uinteger() (localctx IUintegerContext) {
	localctx = NewUintegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BundParserRULE_uinteger)

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
		p.SetState(211)

		var _m = p.Match(BundParserUINTEGER)

		localctx.(*UintegerContext).value = _m
	}

	return localctx
}

// IFloatContext is an interface to support dynamic dispatch.
type IFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFloatContext differentiates from other interfaces.
	IsFloatContext()
}

type FloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFloatContext() *FloatContext {
	var p = new(FloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_float
	return p
}

func (*FloatContext) IsFloatContext() {}

func NewFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatContext {
	var p = new(FloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_float

	return p
}

func (s *FloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatContext) GetValue() antlr.Token { return s.value }

func (s *FloatContext) SetValue(v antlr.Token) { s.value = v }

func (s *FloatContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (p *BundParser) Float() (localctx IFloatContext) {
	localctx = NewFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BundParserRULE_float)

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
		p.SetState(213)

		var _m = p.Match(BundParserFLOAT_NUMBER)

		localctx.(*FloatContext).value = _m
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
	p.EnterRule(localctx, 40, BundParserRULE_call_term)

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
		p.SetState(215)

		var _m = p.Match(BundParserNAME)

		localctx.(*Call_termContext).value = _m
	}

	return localctx
}

// ICall_sysContext is an interface to support dynamic dispatch.
type ICall_sysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// GetSyscmd returns the syscmd token.
	GetSyscmd() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// SetSyscmd sets the syscmd token.
	SetSyscmd(antlr.Token)

	// IsCall_sysContext differentiates from other interfaces.
	IsCall_sysContext()
}

type Call_sysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
	syscmd antlr.Token
}

func NewEmptyCall_sysContext() *Call_sysContext {
	var p = new(Call_sysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_call_sys
	return p
}

func (*Call_sysContext) IsCall_sysContext() {}

func NewCall_sysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_sysContext {
	var p = new(Call_sysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_call_sys

	return p
}

func (s *Call_sysContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_sysContext) GetValue() antlr.Token { return s.value }

func (s *Call_sysContext) GetSyscmd() antlr.Token { return s.syscmd }

func (s *Call_sysContext) SetValue(v antlr.Token) { s.value = v }

func (s *Call_sysContext) SetSyscmd(v antlr.Token) { s.syscmd = v }

func (s *Call_sysContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(BundParserNAME)
}

func (s *Call_sysContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(BundParserNAME, i)
}

func (s *Call_sysContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Call_sysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_sysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_sysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCall_sys(s)
	}
}

func (s *Call_sysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCall_sys(s)
	}
}

func (p *BundParser) Call_sys() (localctx ICall_sysContext) {
	localctx = NewCall_sysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BundParserRULE_call_sys)
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
		p.SetState(217)

		var _m = p.Match(BundParserNAME)

		localctx.(*Call_sysContext).value = _m
	}
	{
		p.SetState(218)
		p.Match(BundParserT__14)
	}
	{
		p.SetState(219)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Call_sysContext).syscmd = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserNAME || _la == BundParserSYS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Call_sysContext).syscmd = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(220)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ICmd_termContext is an interface to support dynamic dispatch.
type ICmd_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsCmd_termContext differentiates from other interfaces.
	IsCmd_termContext()
}

type Cmd_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyCmd_termContext() *Cmd_termContext {
	var p = new(Cmd_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cmd_term
	return p
}

func (*Cmd_termContext) IsCmd_termContext() {}

func NewCmd_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_termContext {
	var p = new(Cmd_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cmd_term

	return p
}

func (s *Cmd_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_termContext) GetValue() antlr.Token { return s.value }

func (s *Cmd_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Cmd_termContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Cmd_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCmd_term(s)
	}
}

func (s *Cmd_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCmd_term(s)
	}
}

func (p *BundParser) Cmd_term() (localctx ICmd_termContext) {
	localctx = NewCmd_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BundParserRULE_cmd_term)

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
		p.SetState(222)

		var _m = p.Match(BundParserCMD)

		localctx.(*Cmd_termContext).value = _m
	}

	return localctx
}

// ICmd_sysContext is an interface to support dynamic dispatch.
type ICmd_sysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// GetSyscmd returns the syscmd token.
	GetSyscmd() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// SetSyscmd sets the syscmd token.
	SetSyscmd(antlr.Token)

	// IsCmd_sysContext differentiates from other interfaces.
	IsCmd_sysContext()
}

type Cmd_sysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
	syscmd antlr.Token
}

func NewEmptyCmd_sysContext() *Cmd_sysContext {
	var p = new(Cmd_sysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cmd_sys
	return p
}

func (*Cmd_sysContext) IsCmd_sysContext() {}

func NewCmd_sysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_sysContext {
	var p = new(Cmd_sysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cmd_sys

	return p
}

func (s *Cmd_sysContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_sysContext) GetValue() antlr.Token { return s.value }

func (s *Cmd_sysContext) GetSyscmd() antlr.Token { return s.syscmd }

func (s *Cmd_sysContext) SetValue(v antlr.Token) { s.value = v }

func (s *Cmd_sysContext) SetSyscmd(v antlr.Token) { s.syscmd = v }

func (s *Cmd_sysContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Cmd_sysContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Cmd_sysContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Cmd_sysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_sysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_sysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCmd_sys(s)
	}
}

func (s *Cmd_sysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCmd_sys(s)
	}
}

func (p *BundParser) Cmd_sys() (localctx ICmd_sysContext) {
	localctx = NewCmd_sysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BundParserRULE_cmd_sys)
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
		p.SetState(224)

		var _m = p.Match(BundParserCMD)

		localctx.(*Cmd_sysContext).value = _m
	}
	{
		p.SetState(225)
		p.Match(BundParserT__14)
	}
	{
		p.SetState(226)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Cmd_sysContext).syscmd = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserNAME || _la == BundParserSYS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Cmd_sysContext).syscmd = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(227)
		p.Match(BundParserT__3)
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
	p.EnterRule(localctx, 48, BundParserRULE_begin)

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
		p.SetState(229)

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
	p.EnterRule(localctx, 50, BundParserRULE_end)

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
		p.SetState(231)

		var _m = p.Match(BundParserTOEND)

		localctx.(*EndContext).value = _m
	}

	return localctx
}

// IDropContext is an interface to support dynamic dispatch.
type IDropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDropContext differentiates from other interfaces.
	IsDropContext()
}

type DropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDropContext() *DropContext {
	var p = new(DropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_drop
	return p
}

func (*DropContext) IsDropContext() {}

func NewDropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropContext {
	var p = new(DropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_drop

	return p
}

func (s *DropContext) GetParser() antlr.Parser { return s.parser }

func (s *DropContext) GetValue() antlr.Token { return s.value }

func (s *DropContext) SetValue(v antlr.Token) { s.value = v }

func (s *DropContext) DROP() antlr.TerminalNode {
	return s.GetToken(BundParserDROP, 0)
}

func (s *DropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDrop(s)
	}
}

func (s *DropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDrop(s)
	}
}

func (p *BundParser) Drop() (localctx IDropContext) {
	localctx = NewDropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BundParserRULE_drop)

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
		p.SetState(233)

		var _m = p.Match(BundParserDROP)

		localctx.(*DropContext).value = _m
	}

	return localctx
}

// IDuplicateContext is an interface to support dynamic dispatch.
type IDuplicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDuplicateContext differentiates from other interfaces.
	IsDuplicateContext()
}

type DuplicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDuplicateContext() *DuplicateContext {
	var p = new(DuplicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_duplicate
	return p
}

func (*DuplicateContext) IsDuplicateContext() {}

func NewDuplicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DuplicateContext {
	var p = new(DuplicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_duplicate

	return p
}

func (s *DuplicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DuplicateContext) GetValue() antlr.Token { return s.value }

func (s *DuplicateContext) SetValue(v antlr.Token) { s.value = v }

func (s *DuplicateContext) DUPLICATE() antlr.TerminalNode {
	return s.GetToken(BundParserDUPLICATE, 0)
}

func (s *DuplicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DuplicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DuplicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDuplicate(s)
	}
}

func (s *DuplicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDuplicate(s)
	}
}

func (p *BundParser) Duplicate() (localctx IDuplicateContext) {
	localctx = NewDuplicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BundParserRULE_duplicate)

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
		p.SetState(235)

		var _m = p.Match(BundParserDUPLICATE)

		localctx.(*DuplicateContext).value = _m
	}

	return localctx
}

// IExecute_termContext is an interface to support dynamic dispatch.
type IExecute_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsExecute_termContext differentiates from other interfaces.
	IsExecute_termContext()
}

type Execute_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyExecute_termContext() *Execute_termContext {
	var p = new(Execute_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_execute_term
	return p
}

func (*Execute_termContext) IsExecute_termContext() {}

func NewExecute_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Execute_termContext {
	var p = new(Execute_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_execute_term

	return p
}

func (s *Execute_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Execute_termContext) GetValue() antlr.Token { return s.value }

func (s *Execute_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Execute_termContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(BundParserEXECUTE, 0)
}

func (s *Execute_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Execute_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Execute_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExecute_term(s)
	}
}

func (s *Execute_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExecute_term(s)
	}
}

func (p *BundParser) Execute_term() (localctx IExecute_termContext) {
	localctx = NewExecute_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BundParserRULE_execute_term)

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
		p.SetState(237)

		var _m = p.Match(BundParserEXECUTE)

		localctx.(*Execute_termContext).value = _m
	}

	return localctx
}

// IReturn_termContext is an interface to support dynamic dispatch.
type IReturn_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsReturn_termContext differentiates from other interfaces.
	IsReturn_termContext()
}

type Return_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyReturn_termContext() *Return_termContext {
	var p = new(Return_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_return_term
	return p
}

func (*Return_termContext) IsReturn_termContext() {}

func NewReturn_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_termContext {
	var p = new(Return_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_return_term

	return p
}

func (s *Return_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_termContext) GetValue() antlr.Token { return s.value }

func (s *Return_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Return_termContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BundParserRETURN, 0)
}

func (s *Return_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterReturn_term(s)
	}
}

func (s *Return_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitReturn_term(s)
	}
}

func (p *BundParser) Return_term() (localctx IReturn_termContext) {
	localctx = NewReturn_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BundParserRULE_return_term)

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
		p.SetState(239)

		var _m = p.Match(BundParserRETURN)

		localctx.(*Return_termContext).value = _m
	}

	return localctx
}
