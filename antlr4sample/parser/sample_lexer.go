// Code generated from /home/lercher/go/src/github.com/lercher/gotools/antlr4sample/Sample.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 35, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 23, 10, 3, 12, 3, 14,
	3, 26, 11, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 2, 2, 7,
	3, 3, 5, 4, 7, 2, 9, 5, 11, 6, 3, 2, 5, 3, 2, 51, 59, 3, 2, 50, 59, 5,
	2, 11, 12, 14, 15, 34, 34, 2, 34, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 20, 3, 2, 2, 2,
	7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 14, 7, 85,
	2, 2, 14, 15, 7, 99, 2, 2, 15, 16, 7, 111, 2, 2, 16, 17, 7, 114, 2, 2,
	17, 18, 7, 110, 2, 2, 18, 19, 7, 103, 2, 2, 19, 4, 3, 2, 2, 2, 20, 24,
	9, 2, 2, 2, 21, 23, 5, 7, 4, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2,
	24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 6, 3, 2, 2, 2, 26, 24, 3, 2,
	2, 2, 27, 28, 9, 3, 2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 48, 2, 2, 30, 10,
	3, 2, 2, 2, 31, 32, 9, 4, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 8, 6, 2, 2,
	34, 12, 3, 2, 2, 2, 4, 2, 24, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'Sample'", "", "'.'",
}

var lexerSymbolicNames = []string{
	"", "SAMPLE", "INT", "DOT", "WS",
}

var lexerRuleNames = []string{
	"SAMPLE", "INT", "DIG", "DOT", "WS",
}

type SampleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSampleLexer(input antlr.CharStream) *SampleLexer {

	l := new(SampleLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sample.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SampleLexer tokens.
const (
	SampleLexerSAMPLE = 1
	SampleLexerINT    = 2
	SampleLexerDOT    = 3
	SampleLexerWS     = 4
)
