// Code generated from Match.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 66, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 47, 10,
	11, 13, 11, 14, 11, 48, 3, 12, 3, 12, 7, 12, 53, 10, 12, 12, 12, 14, 12,
	56, 11, 12, 3, 12, 3, 12, 3, 13, 6, 13, 61, 10, 13, 13, 13, 14, 13, 62,
	3, 13, 3, 13, 3, 54, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3, 2, 4, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2,
	2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35,
	3, 2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2,
	19, 43, 3, 2, 2, 2, 21, 46, 3, 2, 2, 2, 23, 50, 3, 2, 2, 2, 25, 60, 3,
	2, 2, 2, 27, 28, 7, 42, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30, 7, 43, 2, 2, 30,
	6, 3, 2, 2, 2, 31, 32, 7, 48, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 40, 2,
	2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 126, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38,
	7, 35, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7, 45, 2, 2, 40, 16, 3, 2, 2,
	2, 41, 42, 7, 63, 2, 2, 42, 18, 3, 2, 2, 2, 43, 44, 7, 64, 2, 2, 44, 20,
	3, 2, 2, 2, 45, 47, 9, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2,
	48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 22, 3, 2, 2, 2, 50, 54, 7,
	41, 2, 2, 51, 53, 11, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2,
	54, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54, 3,
	2, 2, 2, 57, 58, 7, 41, 2, 2, 58, 24, 3, 2, 2, 2, 59, 61, 9, 3, 2, 2, 60,
	59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2,
	2, 63, 64, 3, 2, 2, 2, 64, 65, 8, 13, 2, 2, 65, 26, 3, 2, 2, 2, 7, 2, 46,
	48, 54, 62, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'.'", "'&'", "'|'", "'!'", "'+'", "'='", "'>'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "AND", "OR", "NOT", "EXIST", "EQUAL", "CONTAINS", "KI",
	"VALUE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "AND", "OR", "NOT", "EXIST", "EQUAL", "CONTAINS",
	"KI", "VALUE", "WHITESPACE",
}

type MatchLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewMatchLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *MatchLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMatchLexer(input antlr.CharStream) *MatchLexer {
	l := new(MatchLexer)
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
	l.GrammarFileName = "Match.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MatchLexer tokens.
const (
	MatchLexerT__0       = 1
	MatchLexerT__1       = 2
	MatchLexerT__2       = 3
	MatchLexerAND        = 4
	MatchLexerOR         = 5
	MatchLexerNOT        = 6
	MatchLexerEXIST      = 7
	MatchLexerEQUAL      = 8
	MatchLexerCONTAINS   = 9
	MatchLexerKI         = 10
	MatchLexerVALUE      = 11
	MatchLexerWHITESPACE = 12
)
