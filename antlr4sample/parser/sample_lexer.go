// Code generated from c://git//src//github.com//lercher//gotools//antlr4sample//Sample.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SampleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SampleLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func samplelexerLexerInit() {
	staticData := &SampleLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Sample'", "", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "SAMPLE", "INT", "DOT", "WS",
	}
	staticData.RuleNames = []string{
		"SAMPLE", "INT", "DIG", "DOT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 33, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 21,
		8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 0, 7, 3, 9, 4, 1, 0, 3, 1, 0, 49, 57, 1,
		0, 48, 57, 3, 0, 9, 10, 12, 13, 32, 32, 32, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 18,
		1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11,
		12, 5, 83, 0, 0, 12, 13, 5, 97, 0, 0, 13, 14, 5, 109, 0, 0, 14, 15, 5,
		112, 0, 0, 15, 16, 5, 108, 0, 0, 16, 17, 5, 101, 0, 0, 17, 2, 1, 0, 0,
		0, 18, 22, 7, 0, 0, 0, 19, 21, 3, 5, 2, 0, 20, 19, 1, 0, 0, 0, 21, 24,
		1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 4, 1, 0, 0, 0,
		24, 22, 1, 0, 0, 0, 25, 26, 7, 1, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 46,
		0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 7, 2, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32,
		6, 4, 0, 0, 32, 10, 1, 0, 0, 0, 2, 0, 22, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SampleLexerInit initializes any static state used to implement SampleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSampleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SampleLexerInit() {
	staticData := &SampleLexerLexerStaticData
	staticData.once.Do(samplelexerLexerInit)
}

// NewSampleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSampleLexer(input antlr.CharStream) *SampleLexer {
	SampleLexerInit()
	l := new(SampleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SampleLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
